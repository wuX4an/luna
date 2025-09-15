package _init

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// -------------------------
// Embedding templates
// -------------------------

//go:embed templates/*
var templatesFS embed.FS

var force bool // <- nuevo flag global

func copyTemplate(template, destDir string) error {
	return fs.WalkDir(templatesFS, "templates/"+template, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel("templates/"+template, path)
		if err != nil {
			return err
		}
		outPath := filepath.Join(destDir, rel)

		// --- Caso especial: si ya existe Luna.toml o src/, manejar según --force ---
		if rel == "Luna.toml" {
			if _, err := os.Stat(outPath); err == nil && !force {
				fmt.Println("⚠️  Found Luna.toml, skipping (use --force to overwrite)")
				return nil
			}
		}
		if rel == "src" {
			if _, err := os.Stat(outPath); err == nil && !force {
				fmt.Println("⚠️  Found src/, skipping files inside (use --force to overwrite)")
				return nil
			}
		}

		if d.IsDir() {
			return os.MkdirAll(outPath, 0o755)
		}

		data, err := templatesFS.ReadFile(path)
		if err != nil {
			return err
		}

		return os.WriteFile(outPath, data, 0o644)
	})
}

// -------------------------
// Command
// -------------------------

var template string

var InitCmd = &cobra.Command{
	Use:     "init [DIR]",
	Short:   "Initialize a new project",
	Example: "luna init\nluna init hello-world",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}

		fmt.Printf("Initializing project in %s\n", dir)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("could not create directory: %w", err)
		}

		absDir, err := filepath.Abs(dir)
		if err != nil {
			return fmt.Errorf("could not determine absolute path: %w", err)
		}

		packageName := strings.TrimSpace(filepath.Base(absDir))

		// Copiar template
		if err := copyTemplate(template, dir); err != nil {
			return fmt.Errorf("could not copy template: %w", err)
		}

		// Reemplazo dinámico de placeholders en Luna.toml
		configPath := filepath.Join(dir, "Luna.toml")
		if _, err := os.Stat(configPath); err == nil {
			data, err := os.ReadFile(configPath)
			if err != nil {
				return fmt.Errorf("could not read Luna.toml: %w", err)
			}

			content := string(data)
			// Nombre del paquete
			content = strings.ReplaceAll(content, "{{PACKAGE_NAME}}", packageName)

			// Target dinámico según el host (solo para template default)
			if template == "default" {
				content = strings.ReplaceAll(content, "{{GOOS}}", runtime.GOOS)
				content = strings.ReplaceAll(content, "{{GOARCH}}", runtime.GOARCH)
			}

			if err := os.WriteFile(configPath, []byte(content), 0o644); err != nil {
				return fmt.Errorf("could not write Luna.toml: %w", err)
			}
		} else {
			fmt.Println("⚠️  No Luna.toml found in template (skipped placeholder replacement)")
		}

		fmt.Println("✅ Project initialized successfully!")
		return nil
	},
}

func init() {
	InitCmd.Flags().StringVarP(&template, "template", "t", "default", "Project template (default, wasm)")
	InitCmd.Flags().BoolVarP(&force, "force", "f", false, "Force overwrite existing files")
}
