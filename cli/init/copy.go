package _init

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func CopyTemplate(template, destDir string) error {
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
