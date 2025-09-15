package stages

import (
	"context"
	"fmt"
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

		if err := copyAndCleanFile(srcPath, dstPath); err != nil {
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

// copyAndCleanFile copia un archivo de src a dst y limpia bloques de test y require("std:test")
func copyAndCleanFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", src, err)
	}

	lines := strings.Split(string(data), "\n")
	var cleaned []string
	skipBlock, funcLevel := false, 0

	for _, line := range lines {
		trim := strings.TrimSpace(line)
		if trim == "" {
			continue
		}
		if skipBlock {
			if strings.Contains(trim, "function") {
				funcLevel++
			}
			if strings.Contains(trim, "end") {
				funcLevel--
				if funcLevel <= 0 {
					skipBlock = false
				}
			}
			continue
		}
		if strings.HasPrefix(trim, "table.insert(__TESTS__") {
			skipBlock = true
			funcLevel = 1
			continue
		}
		if strings.Contains(trim, `require("std:test")`) || strings.Contains(trim, "__TESTS__ or {}") || strings.HasPrefix(trim, "--") {
			continue
		}
		cleaned = append(cleaned, trim)
	}

	// Unir todo en una sola línea
	oneLine := strings.Join(cleaned, " ")
	oneLine = strings.Join(strings.Fields(oneLine), " ") // eliminar espacios dobles

	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return fmt.Errorf("failed to create dir %s: %w", filepath.Dir(dst), err)
	}
	return os.WriteFile(dst, []byte(oneLine), 0o644)
}
