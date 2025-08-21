package stages

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"luna"
	"luna/src/config"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// parseTarget analiza el valor de --target o usa valores por defecto.
func ParseTarget(buildTarget string) (string, string, error) {
	if buildTarget != "" {
		parts := strings.Split(buildTarget, "/")
		if len(parts) != 2 {
			return "", "", errors.New("invalid target format. Use --target=os/arch (e.g. linux/amd64)")
		}
		return parts[0], parts[1], nil
	}
	// Valores por defecto
	return runtime.GOOS, runtime.GOARCH, nil
}

// LoadRuntime obtiene el binario de runtime embebido.
func LoadRuntime(buildOS, buildArch string) ([]byte, error) {
	key := fmt.Sprintf("build/runtimes/runtime_%s_%s", buildOS, buildArch)
	runtimeBinary, err := luna.EmbeddedRuntimes.ReadFile(key)
	if err != nil {
		return nil, fmt.Errorf("no embedded runtime for target %s/%s: %w", buildOS, buildArch, err)
	}
	return runtimeBinary, nil
}

func copyModules(modulePath, dstRoot string) error {
	// quitar prefijo "src/" si existe
	relPath := strings.TrimPrefix(modulePath, "src/")
	dstPath := filepath.Join(dstRoot, relPath)

	// crear directorio padre si no existe
	if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
		return fmt.Errorf("failed to create dir %s: %w", filepath.Dir(dstPath), err)
	}

	// leer contenido
	data, err := os.ReadFile(modulePath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", modulePath, err)
	}

	// eliminar comentarios de línea que comienzan con --
	lines := strings.Split(string(data), "\n")
	var cleaned []string
	for _, line := range lines {
		trim := strings.TrimSpace(line)
		if strings.HasPrefix(trim, "--") {
			continue
		}
		cleaned = append(cleaned, line)
	}

	// escribir archivo limpio
	cleanData := strings.Join(cleaned, "\n")
	if err := os.WriteFile(dstPath, []byte(cleanData), 0o644); err != nil {
		return fmt.Errorf("failed to write %s: %w", dstPath, err)
	}

	return nil
}

func RunBuildPipeline(file, buildOS, buildArch string, runtimeBinary []byte, buildOutput string) error {
	buildDir := filepath.Dir(file)

	// 0. Cargar Luna.toml
	configPath := filepath.Join("Luna.toml")
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return err
	}

	// Validar que el archivo existe
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", file)
	}

	// Encabezado tipo CLI
	fmt.Printf("Checking %s\n", cfg.Build.Source+cfg.Build.Entry)
	fmt.Printf("Compiling %s/ → %s/\n\n", file, buildOutput)

	// 1. Crear carpeta temporal de staging
	modulesDir := filepath.Join(buildOutput, "modules")
	os.RemoveAll(modulesDir)
	if err := os.MkdirAll(modulesDir, 0o755); err != nil {
		return fmt.Errorf("cannot create modules dir %s: %w", modulesDir, err)
	}

	// 2. Obtener lista de todos los módulos y copiarlos a modulesDir
	entryMod := cfg.Build.Entry
	if strings.TrimSuffix(entryMod, ".lua") != "" {
		entryMod = entryMod[:len(entryMod)-len(".lua")]
	}
	entryMod = filepath.ToSlash(entryMod)

	rootNode := parseModule(entryMod, []string{cfg.Build.Source, "."})
	allModules := FlattenModules(rootNode, cfg.Build.Source) // Devuelve rutas tipo src/main.lua

	for _, modPath := range allModules {
		if err := copyModules(modPath, modulesDir); err != nil {
			return fmt.Errorf("failed to copy %s: %w", modPath, err)
		}
	}

	// 3. Crear archivo tar.gz temporal desde modulesDir
	tmpTar := filepath.Join(buildOutput, "bundle.tar.gz")
	os.Remove(tmpTar)
	if err := createTarGz(modulesDir, tmpTar); err != nil {
		return fmt.Errorf("failed to create tar.gz: %w", err)
	}

	// 4. Definir salida final del ejecutable
	finalExe, err := resolveOutputPath(buildOS, buildArch, buildOutput, buildDir)
	if err != nil {
		return err
	}

	// 5. Crear ejecutable final
	if err := buildExecutable(finalExe, runtimeBinary, tmpTar, buildOS); err != nil {
		return err
	}

	rootDir := cfg.Build.Source
	treeStr := DependencyTreeWithSize(rootNode, modulesDir, "", true)
	fmt.Printf("Embedded Files:\n\n%s\n%s\n", cfg.Build.Source, treeStr)

	total := TotalSize(rootNode, rootDir)
	fmt.Printf("\nSummary\n")
	fmt.Printf("Files: %s\n", HumanSize(total))
	return nil
}

func resolveOutputPath(buildOS, buildArch, buildOutput, buildDir string) (string, error) {
	configPath := filepath.Join(buildDir, "Luna.toml")
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return "", err
	}
	// Carpeta base de salida
	baseDir := buildOutput
	if baseDir == "" {
		baseDir = cfg.Package.Name
	}

	// Crear ruta final: ./app/<os>/<arch>/
	outDir := filepath.Join(baseDir, buildOS, buildArch)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return "", fmt.Errorf("failed to create output directory %q: %w", outDir, err)
	}

	// Nombre del ejecutable basado en entryFile
	exeName := cfg.Package.Name
	if buildOS == "windows" {
		exeName += ".exe"
	}

	finalExe := filepath.Join(outDir, exeName)
	return finalExe, nil
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
