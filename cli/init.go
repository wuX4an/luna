package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
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

		// Crear directorio base si no existe
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return fmt.Errorf("could not create directory: %w", err)
			}
		} else {
			// Si existe, revisamos que no esté lleno de cosas
			entries, err := os.ReadDir(dir)
			if err != nil {
				return fmt.Errorf("could not read directory: %w", err)
			}
			if len(entries) > 0 {
				fmt.Println("⚠️ Directory is not empty, init will only add missing files.")
			}
		}

		// Determinar nombre del paquete a partir del nombre de la carpeta
		absDir, err := filepath.Abs(dir)
		if err != nil {
			return fmt.Errorf("could not determine absolute path: %w", err)
		}
		packageName := filepath.Base(absDir)
		packageName = strings.TrimSpace(packageName)

		// Config path
		configPath := filepath.Join(dir, "Luna.toml")
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			defaultConfig := fmt.Sprintf(`[package]
name = "%s"
version = "0.1.0"

[build]
source = "src/"
entry = "main.lua"
target = "linux/amd64"
output = "dist"
`, packageName)

			if err := os.WriteFile(configPath, []byte(defaultConfig), 0o644); err != nil {
				return fmt.Errorf("could not write Luna.toml: %w", err)
			}
			fmt.Println("✅ Created Luna.toml")
		} else {
			fmt.Println("ℹ️  Luna.toml already exists, skipping.")
		}

		// Crear carpeta src si no existe
		srcDir := filepath.Join(dir, "src")
		if err := os.MkdirAll(srcDir, 0o755); err != nil {
			return fmt.Errorf("could not create src directory: %w", err)
		}

		// Crear main.lua si no existe
		mainPath := filepath.Join(srcDir, "main.lua")
		if _, err := os.Stat(mainPath); os.IsNotExist(err) {
			defaultMain := `print("Hello from Luna!")`
			if err := os.WriteFile(mainPath, []byte(defaultMain), 0o644); err != nil {
				return fmt.Errorf("could not write main.lua: %w", err)
			}
			fmt.Println("✅ Created src/main.lua")
		} else {
			fmt.Println("ℹ️  src/main.lua already exists, skipping.")
		}

		fmt.Println("Project initialized successfully!")
		return nil
	},
}

func init() {
	Cmd.AddCommand(initCmd)
}
