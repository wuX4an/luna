package build // ¡Importante! Cambiado a paquete 'build'

import (
	"embed"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// embeddedRuntimes contiene los binarios de tiempo de ejecución incrustados.
//
//go:embed runtime/bin/*
var embeddedRuntimes embed.FS

var (
	// buildTarget almacena el valor de la bandera --target.
	buildTarget = ""
	buildList   = false
	buildOutput = ""
)

var supportedTargets = []string{
	"linux/amd64",
	"linux/arm64",
	"windows/amd64",
	"darwin/amd64",
	"darwin/arm64",
}

const (
	// marker es una cadena utilizada para identificar el inicio del paquete de la aplicación dentro del ejecutable.
	marker = "LUNA_BUNDLE"
)

// buildExamples proporciona ejemplos de uso para el comando 'build'.
var buildExamples = `  luna build main.lua
  luna build .
  luna build . --target=linux/arm64`

// BuildCmd es el comando Cobra para la funcionalidad de construcción.
// Se exporta (primera letra en mayúscula) para que pueda ser accedido desde otros paquetes.
var BuildCmd = &cobra.Command{
	Use:     "build [FILE]",
	Short:   "Compile the script into a self-contained executable",
	Example: buildExamples,
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if buildList {
			fmt.Println("Supported targets:")
			for _, t := range supportedTargets {
				fmt.Println(" •", t)
			}
			return nil
		}
		// Analiza el objetivo de construcción.
		var buildOS, buildArch string
		if buildTarget != "" {
			parts := strings.Split(buildTarget, "/")
			if len(parts) != 2 {
				return errors.New("invalid target format. Use --target=os/arch (e.g. linux/amd64)")
			}
			buildOS = parts[0]
			buildArch = parts[1]
		} else {
			// Valores predeterminados si no se especifica --target.
			buildOS = "linux"
			buildArch = "amd64"
		}

		// Construye la clave para el binario de tiempo de ejecución incrustado.
		key := fmt.Sprintf("runtime/bin/runtime_%s_%s", buildOS, buildArch)
		runtimeBinary, err := embeddedRuntimes.ReadFile(key)
		if err != nil {
			return fmt.Errorf("no embedded runtime for target %s/%s: %w", buildOS, buildArch, err)
		}

		// Determina el archivo o directorio a empaquetar.
		file := "."
		if len(args) > 0 {
			file = args[0]
		}

		fmt.Printf("Building %q for target=%s/%s\n", file, buildOS, buildArch)

		// 1. Crea un archivo tar.gz con el proyecto Lua.
		tmpTar := "bundle.tar.gz"
		if err := createTarGz(file, tmpTar); err != nil {
			return fmt.Errorf("failed to create tar.gz: %w", err)
		}
		// Asegúrate de limpiar el archivo temporal al finalizar.
		defer os.Remove(tmpTar)

		// 2. Crea el ejecutable final.
		var finalExe string
		if buildOutput != "" {
			finalExe = buildOutput
		} else {
			finalExe = fmt.Sprintf("main_%s_%s", buildOS, buildArch)
			if buildOS == "windows" {
				finalExe += ".exe"
			}
		}

		// Verifica que la ruta no sea un directorio
		if info, err := os.Stat(finalExe); err == nil && info.IsDir() {
			return fmt.Errorf("output path %q is a directory, not a file", finalExe)
		}

		outf, err := os.Create(finalExe)
		if err != nil {
			return err
		}
		defer outf.Close() // Asegúrate de cerrar el archivo de salida.

		// Escribe el binario de tiempo de ejecución en el ejecutable final.
		runtimeSize, err := outf.Write(runtimeBinary)
		if err != nil {
			return err
		}

		// Abre el archivo tar.gz y cópialo al ejecutable final.
		tarFile, err := os.Open(tmpTar)
		if err != nil {
			return err
		}
		defer tarFile.Close() // Asegúrate de cerrar el archivo tar.gz.

		_, err = io.Copy(outf, tarFile)
		if err != nil {
			return err
		}

		// Escribe el marcador al final del ejecutable.
		_, err = outf.Write([]byte(marker))
		if err != nil {
			return err
		}

		// Escribe el offset del inicio del paquete de la aplicación.
		offsetBuf := make([]byte, 8)
		binary.LittleEndian.PutUint64(offsetBuf, uint64(runtimeSize))
		_, err = outf.Write(offsetBuf)
		if err != nil {
			return err
		}

		fmt.Printf("Built executable: %s\n", finalExe)

		// Marca como ejecutable (solo en sistemas tipo Unix)
		if buildOS != "windows" {
			if err := os.Chmod(finalExe, 0755); err != nil {
				return fmt.Errorf("failed to set executable permissions: %w", err)
			}
		}
		return nil
	},
}

// init se ejecuta cuando se carga el paquete 'build'.
func init() {
	// Añade la bandera --target al comando BuildCmd.
	// NOTA: La adición de BuildCmd al comando principal Cmd se hará en cli.go.
	BuildCmd.Flags().StringVarP(&buildTarget, "target", "t", "", "Target platform in the form os/arch (linux/amd64)")
	BuildCmd.Flags().BoolVarP(&buildList, "list", "l", false, "List available platforms")
	BuildCmd.Flags().StringVarP(&buildOutput, "output", "o", "", "Output file name (optional)")

}
