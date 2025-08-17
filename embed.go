package luna

import (
	"embed"
	"io/fs"
	"path/filepath"
)

// -------------------------
// Embedding files
// -------------------------

//go:embed build/runtimes/*
var EmbeddedRuntimes embed.FS

//go:embed docs*
var embeddedDocs embed.FS

// -------------------------
// Access docs
// -------------------------

// ReadDocFile devuelve el contenido de un archivo de documentación
func ReadDocFile(path string) (string, error) {
	data, err := embeddedDocs.ReadFile(filepath.Join("docs", path))
	return string(data), err
}

// ListDocs recorre recursivamente docs y devuelve todos los paths de archivos
func ListDocs() ([]string, error) {
	var files []string
	err := fs.WalkDir(embeddedDocs, "docs", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
