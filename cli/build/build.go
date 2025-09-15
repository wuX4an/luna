package build

import (
	"context"
	"fmt"
	"io"
	"luna"
	"luna/cli/build/stages"
	"luna/src/config"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	buildTarget = ""
	buildList   = false
	buildOutput = ""
)

var supportedTargets = []string{
	"linux/amd64",
	"linux/arm64",
	"windows/amd64",
	"js/wasm",
	"darwin/amd64",
	"darwin/arm64",
}

var buildExamples = `
  luna build main.lua
  luna build .
  luna build . --target=linux/arm64
`

var BuildCmd = &cobra.Command{
	Use:           "build [FILE]",
	Short:         "Compile the script into a self-contained executable",
	Example:       buildExamples,
	Args:          cobra.MaximumNArgs(1),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if buildList {
			fmt.Println("Supported targets:")
			for _, t := range supportedTargets {
				fmt.Println(" •", t)
			}
			return nil
		}

		file := "."
		if len(args) > 0 {
			file = args[0]
		}

		entryFile, cfg, err := resolveEntry(file)
		if err != nil {
			return err
		}

		if buildTarget == "" {
			buildTarget = cfg.Build.Target
		}
		if buildOutput == "" {
			buildOutput = cfg.Build.Output
		}

		buildOS, buildArch, err := stages.ParseTarget(buildTarget)
		if err != nil {
			return err
		}

		// Target WASM tiene pipeline especial
		if buildTarget == "js/wasm" {
			return buildWASM(cmd.Context(), entryFile, cfg)
		}

		return buildNormal(cmd.Context(), entryFile, cfg, buildOS, buildArch)
	},
}

func init() {
	BuildCmd.Flags().StringVarP(&buildTarget, "target", "t", "", "Target platform in the form os/arch (linux/amd64)")
	BuildCmd.Flags().BoolVarP(&buildList, "list", "l", false, "List available platforms")
	BuildCmd.Flags().StringVarP(&buildOutput, "output", "o", "", "Output file name (optional)")
}

// resolveEntry obtiene el archivo de entrada absoluto y carga la config
func resolveEntry(file string) (string, *config.LunaConfig, error) {
	info, err := os.Stat(file)
	if err != nil {
		return "", nil, err
	}

	if info.IsDir() {
		cfgPath := filepath.Join(file, "Luna.toml")
		cfg, err := config.LoadConfig(cfgPath)
		if err != nil {
			return "", nil, err
		}
		entryFile := cfg.Build.Source
		if !filepath.IsAbs(entryFile) {
			entryFile = filepath.Join(file, entryFile)
		}
		return entryFile, cfg, nil
	}

	cfgPath := filepath.Join("Luna.toml")
	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		return "", nil, err
	}

	return file, cfg, nil
}

// buildNormal ejecuta el pipeline estándar para binarios nativos
func buildNormal(ctx context.Context, entryFile string, cfg *config.LunaConfig, buildOS, buildArch string) error {
	runtimeBinary, err := stages.LoadRuntime(buildOS, buildArch)
	if err != nil {
		return err
	}

	prepareStage := stages.NewPrepareStage(entryFile, buildOutput)

	entryMod := strings.TrimSuffix(cfg.Build.Entry, ".lua")
	entryMod = filepath.ToSlash(entryMod)
	rootNode := stages.ParseModule(entryMod, []string{cfg.Build.Source, "."})
	modules := stages.FlattenModules(rootNode)

	copyStage := stages.NewCopyStage(cfg, buildOutput, modules)
	bundleStage := stages.NewBundleStage(buildOutput)
	compileStage := stages.NewCompileStage(runtimeBinary, buildOutput, buildOS, buildArch)
	finalStage := stages.NewFinalStage(filepath.Join(buildOutput, "modules"), rootNode, cfg.Build.Source)

	pipeline := stages.NewPipeline(
		prepareStage,
		copyStage,
		bundleStage,
		compileStage,
		finalStage,
	)

	return pipeline.Run(ctx)
}

// buildWASM ejecuta pipeline custom para js/wasm
func buildWASM(ctx context.Context, entryFile string, cfg *config.LunaConfig) error {
	wasmOutput := filepath.Join(cfg.Build.Output, "js", "wasm")

	// 1️⃣ Prepare stage
	prepareStage := stages.NewPrepareStage(entryFile, cfg.Build.Output)
	if err := prepareStage.Run(ctx); err != nil {
		return err
	}

	// 2️⃣ Copy stage original: dist/modules
	os.RemoveAll(filepath.Join(wasmOutput, "_modules"))
	entryMod := strings.TrimSuffix(cfg.Build.Entry, ".lua")
	entryMod = filepath.ToSlash(entryMod)
	rootNode := stages.ParseModule(entryMod, []string{cfg.Build.Source, "."})
	modules := stages.FlattenModules(rootNode)
	copyStage := stages.NewCopyStage(cfg, cfg.Build.Output, modules)
	if err := copyStage.Run(ctx); err != nil {
		return err
	}

	// 3️⃣ Copiar modules -> dist/js/wasm/_modules
	wasmModulesDir := filepath.Join(wasmOutput, "_modules")
	if err := os.MkdirAll(wasmModulesDir, 0o755); err != nil {
		return err
	}
	modulesDir := filepath.Join(cfg.Build.Output, "modules")
	if err := filepath.Walk(modulesDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		rel, _ := filepath.Rel(modulesDir, path)
		dst := filepath.Join(wasmModulesDir, rel)
		return copyFile(path, dst)
	}); err != nil {
		return err
	}

	// 4️⃣ Runtime embebido (luna.wasm)
	runtimeBinary, err := stages.LoadRuntime("js", "wasm")
	if err != nil {
		return err
	}
	runtimeOut := filepath.Join(wasmOutput, "luna.wasm")
	if err := os.WriteFile(runtimeOut, runtimeBinary, 0o644); err != nil {
		return err
	}

	// 5️⃣ Archivos wasm.js e index.html desde embed
	wasmFiles := []string{"wasm.js", "index.html", "sw.js"}
	for _, f := range wasmFiles {
		data, err := luna.ReadWasmFile(f)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(wasmOutput, f)
		if err := os.WriteFile(dstPath, data, 0o644); err != nil {
			return err
		}
	}

	fmt.Println("✅ WASM build complete:", runtimeOut)
	return nil
}

// copyFile asegura que los directorios existen y copia el archivo
func copyFile(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}
