package std

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	lua "github.com/yuin/gopher-lua"
)

// Test que verifica que todos los módulos del std estén preloaded y disponibles con require()
func TestStdModulesAreRegistered(t *testing.T) {
	L := lua.NewState()
	defer L.Close()

	RegisterAll(L)

	// Buscar todas las carpetas de módulos en ./std/
	stdRoot := "./"
	entries, err := os.ReadDir(stdRoot)
	if err != nil {
		t.Fatalf("no se pudo leer el directorio std/: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		modName := entry.Name()

		// Comprobamos que exista el archivo "modulo.go" o "modulo/Loader"
		hasGoFiles := false
		err := filepath.Walk(filepath.Join(stdRoot, modName), func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(info.Name(), ".go") {
				hasGoFiles = true
			}
			return nil
		})
		if err != nil {
			t.Fatalf("error al buscar archivos en %s: %v", modName, err)
		}

		if !hasGoFiles {
			continue // probablemente no sea un módulo std válido
		}

		t.Run("require_"+modName, func(t *testing.T) {
			if err := L.DoString(`require("` + modName + `")`); err != nil {
				t.Errorf("El módulo std '%s' no pudo ser requerido: %v", modName, err)
			}
		})
	}
}
