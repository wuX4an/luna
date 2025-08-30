// error.go
package error

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"
	ColorGray    = "\033[90m"
	ColorBold    = "\033[1m"
)

func PrettyLuaError(mainPath string, err error, modName string) string {
	msg := err.Error()

	// --- Caso 0: Si ya es un error formateado por nosotros, no reprocesar
	if strings.Contains(msg, ">>") && strings.Contains(msg, "^") {
		return msg
	}

	// --- Caso 0.5: Error "module not found" - Formatear bonito
	if strings.Contains(msg, "module") && strings.Contains(msg, "not found") {
		return renderModuleNotFoundError(msg, mainPath)
	}

	// --- Caso 1: Error de sintaxis directo de Lua (formato <string>)
	if strings.HasPrefix(msg, "<string>") && strings.Contains(msg, "line:") {
		lineNum, colNum, detail := parseSyntaxError(msg)

		// Buscar el archivo donde realmente ocurrió el error - CON modName
		errorFile := findActualErrorFile(mainPath, msg, modName)
		if errorFile == "" {
			errorFile = mainPath
		}

		return renderErrorWithSource(errorFile, lineNum, colNum, detail)
	}

	// --- Caso 1.5: Error de sintaxis en archivo principal (formato "src/main.lua line:4(column:5)")
	if strings.Contains(msg, "line:") && strings.Contains(msg, "column:") && strings.Contains(msg, "near") {
		// Ejemplo: "src/main.lua line:4(column:5) near 'local': syntax error"
		lineNum, colNum, detail := parseSyntaxError(msg)

		// Extraer el nombre del archivo del mensaje
		errorFile := mainPath // por defecto
		if spaceIndex := strings.Index(msg, " line:"); spaceIndex != -1 {
			errorFile = msg[:spaceIndex]
			// Verificar si el archivo existe
			if _, err := os.Stat(errorFile); err != nil {
				errorFile = mainPath // usar mainPath si el archivo no existe
			}
		}

		return renderErrorWithSource(errorFile, lineNum, colNum, detail)
	}

	// --- Caso 2: Error propagado
	if strings.Contains(msg, ".lua:") && strings.Contains(msg, "error:") {
		if strings.Contains(msg, ">>") && strings.Contains(msg, "^") {
			return msg
		}

		lines := strings.Split(msg, "\n")
		for _, line := range lines {
			if strings.Contains(line, ".lua:") && strings.Contains(line, "error:") && !strings.Contains(line, ">>") {
				parts := strings.SplitN(line, ":", 3)
				if len(parts) >= 3 {
					file := parts[0]
					if lineNum, err := strconv.Atoi(parts[1]); err == nil {
						detail := strings.TrimSpace(parts[2])
						return renderErrorWithSource(file, lineNum, 1, detail)
					}
				}
			}
		}
	}

	// --- Caso 3: Error ya formateado
	if strings.Contains(msg, "-->") && strings.Contains(msg, "|") {
		return msg
	}

	// --- Caso 4: Otros errores
	return ColorRed + ColorBold + "error: " + ColorReset + cleanErrorMessage(msg)
}

func renderModuleNotFoundError(msg, mainPath string) string {
	// Extraer información del error de módulo no encontrado
	lines := strings.Split(msg, "\n")
	if len(lines) == 0 {
		return msg
	}

	// La primera línea contiene: "src/foo/foo.lua:3: module hsdasd not found:"
	firstLine := lines[0]
	parts := strings.SplitN(firstLine, ":", 3)
	if len(parts) < 3 {
		return msg
	}

	file := parts[0]
	lineNum, _ := strconv.Atoi(parts[1])
	detail := strings.TrimSpace(parts[2])

	// Limpiar el detalle (remover los dos puntos extras)
	detail = strings.TrimSuffix(detail, ":")

	// Renderizar con el formato bonito
	return renderErrorWithSource(file, lineNum, 1, detail)
}

func parseSyntaxError(msg string) (int, int, string) {
	lineNum := 1
	colNum := 1
	detail := "syntax error"

	// Extraer número de línea - manejar ambos formatos
	if lineStart := strings.Index(msg, "line:"); lineStart != -1 {
		lineRest := msg[lineStart+5:]

		// Buscar el fin del número (puede ser '(' o espacio)
		lineEnd := strings.IndexAny(lineRest, "( ")
		if lineEnd != -1 {
			lineStr := lineRest[:lineEnd]
			if num, err := strconv.Atoi(lineStr); err == nil {
				lineNum = num
			}
		} else {
			// Si no hay delimitador, intentar convertir todo
			if num, err := strconv.Atoi(lineRest); err == nil {
				lineNum = num
			}
		}
	}

	// Extraer número de columna - manejar ambos formatos
	if colStart := strings.Index(msg, "column:"); colStart != -1 {
		colRest := msg[colStart+7:]

		// Buscar el fin del número (puede ser ')' o espacio)
		colEnd := strings.IndexAny(colRest, ") ")
		if colEnd != -1 {
			colStr := colRest[:colEnd]
			if num, err := strconv.Atoi(colStr); err == nil {
				colNum = num
			}
		} else {
			// Si no hay delimitador, intentar convertir todo
			if num, err := strconv.Atoi(colRest); err == nil {
				colNum = num
			}
		}
	}

	// Extraer mensaje detallado
	if nearStart := strings.Index(msg, "near"); nearStart != -1 {
		detail = strings.TrimSpace(msg[nearStart:])
		// Limpiar si hay ":" adicional
		if colon := strings.Index(detail, ":"); colon != -1 {
			detail = strings.TrimSpace(detail[colon+1:])
		}
	} else if lastColon := strings.LastIndex(msg, ":"); lastColon != -1 {
		// Para el formato del archivo principal
		detailPart := msg[lastColon+1:]
		if paren := strings.Index(detailPart, "("); paren != -1 {
			detail = strings.TrimSpace(detailPart[:paren])
		} else {
			detail = strings.TrimSpace(detailPart)
		}
	}

	return lineNum, colNum, detail
}

func findActualErrorFile(mainPath, errorMsg string, modName string) string {
	if modName == "" {
		return mainPath
	}
	mainDir := filepath.Dir(mainPath)

	// PRIMERO: Intentar usar el modName para encontrar el archivo exacto
	if modName != "" {
		// Convertir "foo.foo" -> "src/foo/foo.lua"
		// Convertir "module" -> "src/module.lua"
		modPath := strings.ReplaceAll(modName, ".", string(os.PathSeparator)) + ".lua"
		fullModPath := filepath.Join(mainDir, modPath)

		// Verificar si el archivo existe
		if _, err := os.Stat(fullModPath); err == nil {
			return fullModPath
		}

		// También buscar en la raíz de src
		simpleModPath := filepath.Join(mainDir, modName+".lua")
		if _, err := os.Stat(simpleModPath); err == nil {
			return simpleModPath
		}
	}

	// SEGUNDO: Búsqueda tradicional (como fallback)
	var luaFiles []string
	filepath.Walk(mainDir, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(path, ".lua") && path != mainPath {
			luaFiles = append(luaFiles, path)
		}
		return nil
	})

	// Si el mensaje de error menciona específicamente un archivo, usarlo
	if strings.Contains(errorMsg, ".lua:") {
		for _, file := range luaFiles {
			if strings.Contains(errorMsg, filepath.Base(file)+":") {
				return file
			}
		}
	}

	// Buscar archivos que contengan código problemático
	for _, file := range luaFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			if strings.Contains(line, "function") || strings.Contains(line, "if") ||
				strings.Contains(line, "for") || strings.Contains(line, "while") {
				return file
			}
		}
	}

	return ""
}

func renderErrorWithSource(filePath string, lineNum, colNum int, detail string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("%serror: %s (file: %s)%s", ColorRed+ColorBold, detail, filePath, ColorReset)
	}

	lines := strings.Split(string(content), "\n")

	if lineNum < 1 {
		lineNum = 1
	}
	if lineNum > len(lines) {
		lineNum = len(lines)
	}

	srcLine := lines[lineNum-1]

	if colNum < 1 {
		colNum = 1
	}
	if colNum > len(srcLine) {
		colNum = len(srcLine)
	}

	// Mostrar contexto alrededor del error
	startLine := max(1, lineNum-3)
	endLine := min(len(lines), lineNum+3)

	var output strings.Builder

	// Encabezado del error
	output.WriteString(fmt.Sprintf("%s%serror:%s %s\n", ColorRed, ColorBold, ColorReset, cleanErrorMessage(detail)))
	output.WriteString(fmt.Sprintf("%s   -->%s %s:%d:%d\n", ColorBlue, ColorReset, filePath, lineNum, colNum))

	// Mostrar las líneas de contexto
	for i := startLine; i <= endLine; i++ {
		lineContent := lines[i-1]
		lineMarker := "  "
		lineColor := ColorGray
		numberColor := ColorGray

		if i == lineNum {
			lineMarker = fmt.Sprintf("%s>>%s", ColorRed, ColorReset)
			lineColor = ColorRed + ColorBold
			numberColor = ColorRed + ColorBold
		}

		// Número de línea y contenido - CORREGIDO
		output.WriteString(fmt.Sprintf("%s %s%d%s | %s%s%s\n",
			lineMarker, numberColor, i, ColorReset, lineColor, lineContent, ColorReset))

		// Caret apuntando al error
		if i == lineNum {
			indent := strings.Repeat(" ", len(fmt.Sprintf(">> %d | ", i))+colNum-1)
			output.WriteString(fmt.Sprintf("%s%s^%s\n", indent, ColorGreen, ColorReset))
		}
	}

	return output.String()
}

// Funciones auxiliares para min/max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func cleanErrorMessage(msg string) string {
	// Limpiar mensajes de error duplicados o con formato extraño
	msg = strings.TrimSpace(msg)
	msg = strings.TrimPrefix(msg, "error: ")
	msg = strings.TrimPrefix(msg, "error:")
	return msg
}
