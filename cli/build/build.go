package build

import (
	"fmt"
	"luna/cli/build/stages"
	"luna/src/config"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// buildTarget almacena el valor de la bandera --target.
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

// buildExamples proporciona ejemplos de uso para el comando 'build'.
var buildExamples = `  luna build main.lua
  luna build .
  luna build . --target=linux/arm64`

// BuildCmd es el comando Cobra para la funcionalidad de construcción.
var BuildCmd = &cobra.Command{
	Use:           "build [FILE]",
	Short:         "Compile the script into a self-contained executable",
	Example:       buildExamples,
	Args:          cobra.MaximumNArgs(1),
	SilenceUsage:  true, // No help on errors
	SilenceErrors: true, // No duplicate errors
	RunE: func(cmd *cobra.Command, args []string) error {
		if buildList {
			fmt.Println("Supported targets:")
			for _, t := range supportedTargets {
				fmt.Println(" •", t)
			}
			return nil
		}

		// 1️⃣ Determinar archivo o directorio a empaquetar
		file := "."
		if len(args) > 0 {
			file = args[0]
		}

		var entryFile string
		info, err := os.Stat(file)
		if err != nil {
			return err
		}

		if info.IsDir() {
			configPath := filepath.Join(file, "Luna.toml")
			cfg, err := config.LoadConfig(configPath)
			if err != nil {
				return err
			}
			entryFile = cfg.Build.Source
			if !filepath.IsAbs(entryFile) {
				entryFile = filepath.Join(file, entryFile)
			}
			if buildTarget == "" && cfg.Build.Target != "" {
				buildTarget = cfg.Build.Target
			}
			if buildOutput == "" && cfg.Build.Output != "" {
				buildOutput = cfg.Build.Output
			}
		} else {
			entryFile = file
		}

		// 2️⃣ Parse target
		buildOS, buildArch, err := stages.ParseTarget(buildTarget)
		if err != nil {
			return err
		}

		// 3️⃣ Crear pipeline con stages concretos
		runtimeBinary, err := stages.LoadRuntime(buildOS, buildArch)
		if err != nil {
			return err
		}

		configPath := filepath.Join("Luna.toml")
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			return err
		}

		prepareStage := stages.NewPrepareStage(entryFile, buildOutput)

		entryMod := cfg.Build.Entry
		if strings.HasSuffix(entryMod, ".lua") {
			entryMod = entryMod[:len(entryMod)-len(".lua")]
		}
		entryMod = filepath.ToSlash(entryMod)

		rootNode := stages.ParseModule(entryMod, []string{cfg.Build.Source, "."})

		modules := stages.FlattenModules(rootNode)

		copyStage := stages.NewCopyStage(cfg, buildOutput, modules)
		bundleStage := stages.NewBundleStage(buildOutput)
		compileStage := stages.NewCompileStage(runtimeBinary, buildOutput, buildOS, buildArch)

		finalStage := stages.NewFinalStage(buildOutput+"/modules", rootNode, cfg.Build.Source)

		pipeline := stages.NewPipeline(
			prepareStage,
			copyStage,
			bundleStage,
			compileStage,
			finalStage,
		)
		// 4️⃣ Ejecutar pipeline
		return pipeline.Run(cmd.Context())
	},
}

// init se ejecuta cuando se carga el paquete 'build'.
func init() {
	BuildCmd.Flags().StringVarP(&buildTarget, "target", "t", "", "Target platform in the form os/arch (linux/amd64)")
	BuildCmd.Flags().BoolVarP(&buildList, "list", "l", false, "List available platforms")
	BuildCmd.Flags().StringVarP(&buildOutput, "output", "o", "", "Output file name (optional)")
}
