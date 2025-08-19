package build

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"luna"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
)

type LunaConfig struct {
	Package struct {
		Name    string `toml:"name"`
		Version string `toml:"version"`
	} `toml:"package"`
	Build struct {
		Entry  string `toml:"entry"`
		Source string `toml:"source"`
		Target string `toml:"target"`
		Output string `toml:"output"`
	} `toml:"build"`
}

// parseTarget analiza el valor de --target o usa valores por defecto.
func parseTarget(buildTarget string) (string, string, error) {
	if buildTarget != "" {
		parts := strings.Split(buildTarget, "/")
		if len(parts) != 2 {
			return "", "", errors.New("invalid target format. Use --target=os/arch (e.g. linux/amd64)")
		}
		return parts[0], parts[1], nil
	}
	// Valores por defecto
	return "linux", "amd64", nil
}

// loadRuntime obtiene el binario de runtime embebido.
func loadRuntime(buildOS, buildArch string) ([]byte, error) {
	key := fmt.Sprintf("build/runtimes/runtime_%s_%s", buildOS, buildArch)
	runtimeBinary, err := luna.EmbeddedRuntimes.ReadFile(key)
	if err != nil {
		return nil, fmt.Errorf("no embedded runtime for target %s/%s: %w", buildOS, buildArch, err)
	}
	return runtimeBinary, nil
}

// runBuildPipeline ejecuta todo el proceso de build.
func runBuildPipeline(file, buildOS, buildArch string, runtimeBinary []byte, buildOutput string) error {
	fmt.Printf("Building %q for target=%s/%s\n", file, buildOS, buildArch)

	// 1. Crear archivo tar.gz en la ubicación de salida de la configuración
	tmpTar := filepath.Join(buildOutput, "bundle.tar.gz")

	// Asegurarnos de que el directorio exista
	if err := os.MkdirAll(buildOutput, 0o755); err != nil {
		return fmt.Errorf("no se pudo crear el directorio de salida %s: %w", buildOutput, err)
	}

	if err := createTarGz(file, tmpTar); err != nil {
		return fmt.Errorf("failed to create tar.gz: %w", err)
	}

	// 2. Definir nombre de salida final del ejecutable
	finalExe, err := resolveOutputPath(buildOS, buildArch, buildOutput, file)
	if err != nil {
		return err
	}

	// 3. Crear ejecutable final
	if err := buildExecutable(finalExe, runtimeBinary, tmpTar, buildOS); err != nil {
		return err
	}

	fmt.Printf("Built executable: %s\n", finalExe)
	return nil
}

func resolveOutputPath(buildOS, buildArch, buildOutput, entryFile string) (string, error) {
	// Carpeta base de salida
	baseDir := buildOutput
	if baseDir == "" {
		baseDir = "./app"
	}

	// Crear ruta final: ./app/<os>/<arch>/
	outDir := filepath.Join(baseDir, buildOS, buildArch)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return "", fmt.Errorf("failed to create output directory %q: %w", outDir, err)
	}

	// Nombre del ejecutable basado en entryFile
	exeName := "app"
	if buildOS == "windows" {
		exeName += ".exe"
	}

	finalExe := filepath.Join(outDir, exeName)
	return finalExe, nil
}

func loadLunaToml(buildDir string) (*LunaConfig, error) {
	tomlPath := filepath.Join(buildDir, "Luna.toml")
	if _, err := os.Stat(tomlPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("no Luna.toml found in %s", buildDir)
	}

	tree, err := toml.LoadFile(tomlPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Luna.toml: %w", err)
	}

	var cfg LunaConfig
	if err := tree.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Luna.toml: %w", err)
	}

	return &cfg, nil
}

// buildExecutable construye el ejecutable final con runtime + bundle + marker.
func buildExecutable(finalExe string, runtimeBinary []byte, tmpTar, buildOS string) error {
	outf, err := os.Create(finalExe)
	if err != nil {
		return err
	}
	defer outf.Close()

	// Escribir runtime
	runtimeSize, err := outf.Write(runtimeBinary)
	if err != nil {
		return err
	}

	// Escribir bundle
	tarFile, err := os.Open(tmpTar)
	if err != nil {
		return err
	}
	defer tarFile.Close()
	if _, err := io.Copy(outf, tarFile); err != nil {
		return err
	}

	// Escribir marker
	if _, err := outf.Write([]byte(marker)); err != nil {
		return err
	}

	// Escribir offset
	offsetBuf := make([]byte, 8)
	binary.LittleEndian.PutUint64(offsetBuf, uint64(runtimeSize))
	if _, err := outf.Write(offsetBuf); err != nil {
		return err
	}

	// Hacerlo ejecutable (no Windows)
	if buildOS != "windows" {
		if err := os.Chmod(finalExe, 0755); err != nil {
			return fmt.Errorf("failed to set executable permissions: %w", err)
		}
	}
	return nil
}
