package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"luna/std" // importa tu std
	"os"

	lua "github.com/yuin/gopher-lua"
)

const (
	marker      = "LUNA_BUNDLE"
	markerSize  = len(marker)
	offsetBytes = 8
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	f, err := os.Open(os.Args[0])
	if err != nil {
		return err
	}
	defer f.Close()

	offset, err := findTarOffset(f)
	if err != nil {
		return err
	}

	_, err = f.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	var mainLuaContent bytes.Buffer
	found := false

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if hdr.Name == "main.lua" {
			// Leer todo el contenido de main.lua en memoria
			if _, err := io.Copy(&mainLuaContent, tr); err != nil {
				return err
			}
			found = true
			break
		}
	}

	if !found {
		return errors.New("main.lua not found in bundle")
	}

	L := lua.NewState()
	defer L.Close()

	std.RegisterAll(L)

	// Ejecutar el contenido en memoria
	if err := L.DoString(mainLuaContent.String()); err != nil {
		return fmt.Errorf("error running main.lua: %w", err)
	}

	return nil
}

func findTarOffset(f *os.File) (int64, error) {
	stat, err := f.Stat()
	if err != nil {
		return 0, err
	}

	fileSize := stat.Size()
	if fileSize < int64(markerSize+offsetBytes) {
		return 0, errors.New("file too small to contain marker")
	}

	buf := make([]byte, markerSize+offsetBytes)
	_, err = f.ReadAt(buf, fileSize-int64(markerSize+offsetBytes))
	if err != nil {
		return 0, err
	}

	if string(buf[:markerSize]) != marker {
		return 0, errors.New("marker not found")
	}

	offset := int64(binary.LittleEndian.Uint64(buf[markerSize:]))
	return offset, nil
}
