package stages

import (
	"context"
	"fmt"
	"luna/src/config"
	"os"
	"path/filepath"
	"strings"
)

type PrepareStage struct {
	File      string
	OutputDir string
	Config    *config.LunaConfig
}

func NewPrepareStage(file, outputDir string) *PrepareStage {
	return &PrepareStage{
		File:      file,
		OutputDir: outputDir,
	}
}

func (p *PrepareStage) Name() string { return "PrepareStage" }

func (p *PrepareStage) Run(ctx context.Context) error {
	fmt.Println("\n🚀 [1/4] Build Stage: Prepare")
	fmt.Println(strings.Repeat("─", 40))

	// Validar archivo
	if _, err := os.Stat(p.File); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", p.File)
	}

	// Cargar configuración
	cfg, err := config.LoadConfig("Luna.toml")
	if err != nil {
		return err
	}
	p.Config = cfg

	// Crear carpeta de staging de módulos
	modulesDir := filepath.Join(p.OutputDir, "modules")
	os.RemoveAll(modulesDir)
	if err := os.MkdirAll(modulesDir, 0o755); err != nil {
		return fmt.Errorf("cannot create modules dir %s: %w", modulesDir, err)
	}

	fmt.Printf("Modules staged at: %s\n", modulesDir)
	fmt.Println("Status: ✅ Done")

	return nil
}
