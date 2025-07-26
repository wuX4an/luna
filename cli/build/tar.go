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
// Está optimizado para manejar archivos grandes de manera eficiente utilizando streaming.
//
// src: La ruta al archivo o directorio que se va a comprimir.
// dest: La ruta donde se creará el archivo .tar.gz.
func createTarGz(src, dest string) error {
	// Limpiar la ruta de origen para asegurar consistencia.
	src = filepath.Clean(src)
	// Limpiar la ruta de destino para comparaciones.
	destClean := filepath.Clean(dest)

	// Crear el archivo de destino .tar.gz.
	f, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("error al crear el archivo de destino %s: %w", dest, err)
	}
	// Asegurarse de que el archivo se cierre al finalizar la función.
	defer f.Close()

	// Crear un escritor gzip sobre el archivo de destino.
	// Esto comprimirá los datos que se le escriban.
	gzw := gzip.NewWriter(f)
	// Asegurarse de que el escritor gzip se cierre y vacíe su búfer.
	defer gzw.Close()

	// Crear un escritor tar sobre el escritor gzip.
	// Esto organizará los archivos en el formato tar antes de la compresión gzip.
	tw := tar.NewWriter(gzw)
	// Asegurarse de que el escritor tar se cierre y vacíe su búfer.
	defer tw.Close()

	// Recorrer recursivamente el directorio o archivo de origen.
	// filepath.Walk es eficiente ya que visita los archivos uno por uno,
	// sin cargar todo el árbol de directorios en memoria.
	return filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			// Si hay un error al acceder a un archivo/directorio, reportarlo.
			return fmt.Errorf("error al acceder a %s: %w", file, err)
		}

		// Ignorar directorios; solo queremos añadir archivos al tar.
		if fi.IsDir() {
			return nil
		}

		// --- NUEVA LÓGICA DE EXCLUSIÓN ---
		// Excluir el propio archivo de destino para evitar incluirse a sí mismo
		// o generar un bucle infinito si 'dest' está dentro de 'src'.
		if filepath.Clean(file) == destClean {
			fmt.Printf("Saltando el archivo de destino: %s\n", file)
			return nil
		}

		// Excluir archivos específicos que se sabe que causan problemas de concurrencia
		// o que no deben ser incluidos, como 'bundle.tar.gz' en este caso.
		if fi.Name() == "bundle.tar.gz" {
			fmt.Printf("Saltando archivo problemático: %s\n", file)
			return nil
		}
		// --- FIN DE LA NUEVA LÓGICA DE EXCLUSIÓN ---

		// Calcular la ruta relativa del archivo dentro del archivo tar.
		// Esto asegura que el archivo se almacene con su estructura de directorio relativa.
		relPath, err := filepath.Rel(src, file)
		if err != nil {
			return fmt.Errorf("error al obtener la ruta relativa para %s: %w", file, err)
		}
		// Normalizar la ruta a barras diagonales para compatibilidad con tar en diferentes sistemas operativos.
		relPath = filepath.ToSlash(relPath)

		// Crear un encabezado tar para el archivo actual.
		// El campo 'Size' es de tipo int64, lo que permite manejar archivos de tamaños muy grandes.
		// Se fuerza el formato PAX para una mejor compatibilidad con archivos grandes,
		// nombres de archivo largos y otros metadatos extendidos.
		hdr := &tar.Header{
			Name:    relPath,
			Size:    fi.Size(),
			Mode:    int64(fi.Mode().Perm()), // Permisos del archivo.
			ModTime: fi.ModTime(),            // Fecha de modificación.
			Format:  tar.FormatPAX,           // Formato PAX para soporte de archivos grandes.
		}

		// Escribir el encabezado del archivo en el archivo tar.
		// Incluye el tamaño esperado en el mensaje de error para facilitar la depuración.
		if err := tw.WriteHeader(hdr); err != nil {
			return fmt.Errorf("error al escribir el encabezado tar para %s (tamaño esperado: %d bytes): %w", file, hdr.Size, err)
		}

		// Abrir el archivo de origen para leer su contenido.
		fh, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("error al abrir el archivo %s: %w", file, err)
		}
		// Asegurarse de que el archivo de origen se cierre.
		defer fh.Close()

		// Copiar el contenido del archivo de origen al escritor tar.
		// io.Copy es crucial para el manejo de archivos grandes, ya que lee y escribe
		// el contenido en fragmentos (chunks) en lugar de cargar todo el archivo en memoria.
		// Esto minimiza el uso de memoria RAM, haciendo que la función sea muy eficiente
		// para archivos de cualquier tamaño, desde pequeños hasta terabytes.
		n, err := io.Copy(tw, fh)
		if err != nil {
			return fmt.Errorf("error al copiar el contenido de %s al tar: %w", file, err)
		}

		// Verificar si el número de bytes copiados coincide con el tamaño esperado en el encabezado.
		// El error "archive/tar: write too long" ocurre si 'n' es mayor que 'hdr.Size'.
		// Esto suele indicar que el archivo cambió entre la lectura de fi.Size()
		// y el momento en que se copió el contenido, o que el archivo estaba incompleto.
		if n != hdr.Size {
			return fmt.Errorf("error de tamaño de archivo para %s: se copiaron %d bytes, pero el encabezado esperaba %d bytes. Esto podría indicar que el archivo fue modificado durante la operación de empaquetado.", file, n, hdr.Size)
		}

		return nil
	})
}
