package stages

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FinalStage struct {
	ModulesDir string
	RootNode   *Node
	Source     string
}

func NewFinalStage(modulesDir string, rootNode *Node, source string) *FinalStage {
	return &FinalStage{
		ModulesDir: modulesDir,
		RootNode:   rootNode,
		Source:     source,
	}
}

func (f *FinalStage) Name() string { return "FinalStage" }

func (f *FinalStage) Run(ctx context.Context) error {
	treeStr := DependencyTreeWithSize(f.RootNode, f.ModulesDir, "", true)

	fmt.Println(strings.Repeat("─", 50))
	fmt.Printf("📦 Embedded Files:\n\n%s\n%s\n", f.Source, treeStr)

	total := TotalSize(f.RootNode, f.ModulesDir)

	var count int
	err := filepath.Walk(f.ModulesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() {
			count++
		}
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Println("\n📊 Summary")
	fmt.Println(strings.Repeat("─", 50))
	fmt.Printf("• Total files: %d\n", count)
	fmt.Printf("• Total size: %s\n", HumanSize(total))
	fmt.Printf("• Modules dir: %s\n", f.ModulesDir)

	fmt.Println(strings.Repeat("─", 50))
	fmt.Println("✅ Build complete")

	return nil
}

var (
	reRequire = regexp.MustCompile(`require\s*\(\s*["']([^"']+)["']\s*\)`)
	seen      = map[string]bool{}
)

type Node struct {
	Name     string
	Children []*Node
}

// Convierte un módulo Lua tipo "hello" o "src.hello" -> ruta de archivo real
func moduleToFile(modName string) string {
	return filepath.FromSlash(strings.ReplaceAll(modName, ".", "/")) + ".lua"
}

func ParseModule(modName string, searchDirs []string) *Node {
	if seen[modName] {
		return &Node{Name: modName} // evitar ciclos
	}
	seen[modName] = true

	var fullPath string
	found := false
	for _, dir := range searchDirs {
		try := filepath.Join(dir, moduleToFile(modName))
		if _, err := os.Stat(try); err == nil {
			fullPath = try
			found = true
			break
		}
	}
	if !found {
		// módulo externo
		return &Node{Name: modName}
	}

	data, _ := os.ReadFile(fullPath)
	content := string(data)

	// --- eliminar comentarios de línea ---
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if idx := strings.Index(line, "--"); idx >= 0 {
			lines[i] = line[:idx]
		}
	}
	content = strings.Join(lines, "\n")

	// --- eliminar comentarios multilínea ---
	reMultiLine := regexp.MustCompile(`--\[\[[\s\S]*?\]\]`)
	content = reMultiLine.ReplaceAllString(content, "")

	node := &Node{Name: modName}
	matches := reRequire.FindAllStringSubmatch(content, -1)
	for _, m := range matches {
		childName := m[1]

		// ignorar módulos "std:*"
		if strings.HasPrefix(childName, "std:") {
			continue
		}

		child := ParseModule(childName, searchDirs)
		node.Children = append(node.Children, child)
	}

	return node
}

// Convierte el nombre de módulo "foo.bar" -> "foo/bar.lua"
func formatModuleName(name string) string {
	return filepath.FromSlash(
		strings.ReplaceAll(name, ".", "/"),
	) + ".lua"
}

func DependencyTree(n *Node, prefix string, isLast bool) string {
	var sb strings.Builder
	connector := "├── "
	if isLast {
		connector = "└── "
	}
	sb.WriteString(prefix + connector + formatModuleName(n.Name))

	for i, c := range n.Children {
		last := i == len(n.Children)-1
		newPrefix := prefix
		if isLast {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		sb.WriteString("\n" + DependencyTree(c, newPrefix, last))
	}

	return sb.String()
}

func HumanSize(size int64) string {
	switch {
	case size < 1024:
		return fmt.Sprintf("%dB", size)
	case size < 1024*1024:
		return fmt.Sprintf("%.1fKB", float64(size)/1024)
	default:
		return fmt.Sprintf("%.1fMB", float64(size)/(1024*1024))
	}
}

// totalSize recorre el árbol y suma el tamaño de todos los archivos
func TotalSize(n *Node, rootDir string) int64 {
	var total int64
	var walk func(node *Node)
	walk = func(node *Node) {
		path := filepath.Join(rootDir, strings.ReplaceAll(node.Name, ".", string(os.PathSeparator))+".lua")
		if !strings.HasSuffix(path, ".lua") {
			path += ".lua"
		}
		if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
			total += fi.Size()
		}
		for _, child := range node.Children {
			walk(child)
		}
	}
	walk(n)
	return total
}

// DependencyTreeWithSize genera el árbol con tamaños de archivo en B, KB o MB
func DependencyTreeWithSize(n *Node, modulesDir, prefix string, isLast bool) string {
	var sb strings.Builder
	connector := "├── "
	if isLast {
		connector = "└── "
	}

	// Path en la carpeta modulesDir (archivos ya procesados)
	path := filepath.Join(modulesDir, strings.ReplaceAll(n.Name, ".", string(os.PathSeparator))+".lua")
	sizeStr := ""
	if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
		sizeStr = " (" + HumanSize(fi.Size()) + ")"
	}

	sb.WriteString(prefix + connector + formatModuleName(n.Name) + sizeStr)

	for i, c := range n.Children {
		last := i == len(n.Children)-1
		newPrefix := prefix
		if isLast {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		sb.WriteString("\n" + DependencyTreeWithSize(c, modulesDir, newPrefix, last))
	}

	return sb.String()
}

func FlattenModules(n *Node) []string {
	var modules []string
	var walk func(node *Node)
	walk = func(node *Node) {
		modules = append(modules, formatModuleName(node.Name))
		for _, child := range node.Children {
			walk(child)
		}
	}
	walk(n)
	return modules
}
