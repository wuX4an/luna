package build

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// createTarGz crea un archivo .tar.gz a partir de un directorio o archivo de origen.
// Maneja tanto archivos sueltos como directorios, evitando incluir el archivo de salida.
func createTarGz(src, dest string) error {
	src = filepath.Clean(src)
	destClean := filepath.Clean(dest)

	// Crear el archivo .tar.gz de destino
	f, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("error al crear el archivo de destino %s: %w", dest, err)
	}
	defer f.Close()

	// Preparar compresión gzip y formato tar
	gzw := gzip.NewWriter(f)
	defer gzw.Close()

	tw := tar.NewWriter(gzw)
	defer tw.Close()

	// Detectar si el origen es directorio o archivo
	info, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("error al acceder a %s: %w", src, err)
	}

	var baseDir string
	if info.IsDir() {
		baseDir = src
	} else {
		// Si es un archivo, tomar su carpeta padre como base
		baseDir = filepath.Dir(src)
	}

	// Recorrer archivos
	return filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error al acceder a %s: %w", file, err)
		}

		if fi.IsDir() {
			return nil
		}

		// Excluir el archivo de destino si está dentro de src
		if filepath.Clean(file) == destClean {
			fmt.Printf("Saltando el archivo de destino: %s\n", file)
			return nil
		}

		// Excluir bundle.tar.gz si aparece en el origen
		if fi.Name() == "bundle.tar.gz" {
			fmt.Printf("Saltando archivo problemático: %s\n", file)
			return nil
		}

		// Obtener ruta relativa para guardar en el tar
		relPath, err := filepath.Rel(baseDir, file)
		if err != nil {
			return fmt.Errorf("error al obtener la ruta relativa para %s: %w", file, err)
		}
		relPath = filepath.ToSlash(relPath)

		// Si src era un archivo individual, asegurar que se ponga como "main.lua"
		if !info.IsDir() && file == src {
			relPath = filepath.Base(src)
		}

		// Crear encabezado tar
		hdr := &tar.Header{
			Name:    relPath,
			Size:    fi.Size(),
			Mode:    int64(fi.Mode().Perm()),
			ModTime: fi.ModTime(),
			Format:  tar.FormatPAX,
		}

		if err := tw.WriteHeader(hdr); err != nil {
			return fmt.Errorf("error al escribir encabezado tar para %s (tamaño esperado: %d bytes): %w", file, hdr.Size, err)
		}

		// Abrir archivo y copiar contenido
		fh, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("error al abrir el archivo %s: %w", file, err)
		}
		defer fh.Close()

		n, err := io.Copy(tw, fh)
		if err != nil {
			return fmt.Errorf("error al copiar contenido de %s al tar: %w", file, err)
		}
		if n != hdr.Size {
			return fmt.Errorf(
				"error de tamaño de archivo para %s: se copiaron %d bytes, pero el encabezado esperaba %d bytes. Esto podría indicar que el archivo fue modificado durante la operación.",
				file, n, hdr.Size,
			)
		}

		return nil
	})
}
