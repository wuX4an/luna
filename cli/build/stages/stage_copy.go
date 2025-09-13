package stages

import (
	"context"
	"fmt"
	"io"
	"luna/src/config"
	"os"
	"path/filepath"
	"strings"
)

type CopyStage struct {
	Config    *config.LunaConfig
	OutputDir string
	Modules   []string
}

func NewCopyStage(cfg *config.LunaConfig, outputDir string, modules []string) *CopyStage {
	return &CopyStage{
		Config:    cfg,
		OutputDir: outputDir,
		Modules:   modules,
	}
}

func (c *CopyStage) Name() string { return "CopyStage" }

func (c *CopyStage) Run(ctx context.Context) error {
	modulesDir := filepath.Join(c.OutputDir, "modules")
	fmt.Println("\n🚀 [2/4] Build Stage: Copy")
	fmt.Println(strings.Repeat("─", 40))
	fmt.Printf("Copying %d modules to: %s\n\n", len(c.Modules), modulesDir)

	var totalSize int64
	for _, relPath := range c.Modules {
		srcPath := filepath.Join(c.Config.Build.Source, relPath)
		dstPath := filepath.Join(modulesDir, relPath)

		if err := copyFile(srcPath, dstPath); err != nil {
			return fmt.Errorf("failed to copy %s: %w", relPath, err)
		}

		fi, err := os.Stat(srcPath)
		if err == nil {
			totalSize += fi.Size()
			fmt.Printf("  • %s (%s)\n", relPath, HumanSize(fi.Size()))
		}
	}

	fmt.Printf("\nStatus: ✅ Done\n")
	fmt.Printf("Total files copied: %d\n", len(c.Modules))
	fmt.Printf("Total size: %s\n", HumanSize(totalSize))

	return nil
}

// copyFile asegura que los directorios existen y copia el archivo
func copyFile(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	return out.Close()
}
