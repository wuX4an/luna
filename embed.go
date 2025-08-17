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
var EmbeddedDocs embed.FS

// -------------------------
// Access runtimes
// -------------------------

// ListRuntimes devuelve la lista de nombres de runtime embebidos
func ListRuntimes() ([]string, error) {
	entries, err := fs.ReadDir(EmbeddedRuntimes, "build/runtimes")
	if err != nil {
		return nil, err
	}

	runtimes := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		runtimes = append(runtimes, entry.Name())
	}
	return runtimes, nil
}

// ReadRuntimeFile devuelve el contenido de un archivo específico dentro de un runtime
func ReadRuntimeFile(runtime, filename string) ([]byte, error) {
	path := filepath.Join("build/runtimes", runtime, filename)
	return EmbeddedRuntimes.ReadFile(path)
}

// -------------------------
// Access docs
// -------------------------

// ReadDocFile devuelve el contenido de un archivo de documentación
func ReadDocFile(path string) (string, error) {
	data, err := EmbeddedDocs.ReadFile(filepath.Join("docs", path))
	return string(data), err
}

func MustReadDocFile(path string) string {
	s, err := ReadDocFile(path)
	if err != nil {
		panic(err)
	}
	return s
}

// ListDocs recorre recursivamente docs y devuelve todos los paths de archivos
func ListDocs() ([]string, error) {
	var files []string
	err := fs.WalkDir(EmbeddedDocs, "docs", func(path string, d fs.DirEntry, err error) error {
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
