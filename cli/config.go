package cli

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type TaskDef struct {
	Script  string   `toml:"script"`
	Desc    string   `toml:"desc"`
	Depends []string `toml:"depends"`
}

type LunaConfig struct {
	Package struct {
		Name    string `toml:"name"`
		Version string `toml:"version"`
	} `toml:"package"`

	Build struct {
		Source string `toml:"source"`
		Entry  string `toml:"entry"`
		Target string `toml:"target"`
		Output string `toml:"output"`
	} `toml:"build"`

	Tasks struct {
		Source string             `toml:"source"`
		Defs   map[string]TaskDef `toml:"-"` // lo llenamos manualmente
	} `toml:"tasks"`
}

func LoadConfig(path string) (*LunaConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read %s: %w", path, err)
	}

	// parse crudo
	var raw map[string]any
	if err := toml.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("invalid %s: %w", path, err)
	}

	// parse directo
	var conf LunaConfig
	if err := toml.Unmarshal(data, &conf); err != nil {
		return nil, fmt.Errorf("invalid %s: %w", path, err)
	}

	// extraer subtasks
	if tasksRaw, ok := raw["tasks"].(map[string]any); ok {
		conf.Tasks.Defs = make(map[string]TaskDef)
		for k, v := range tasksRaw {
			if k == "source" {
				continue
			}
			if sub, ok := v.(map[string]any); ok {
				td := TaskDef{}
				if s, ok := sub["script"].(string); ok {
					td.Script = s
				}
				if d, ok := sub["desc"].(string); ok {
					td.Desc = d
				}
				if dep, ok := sub["depends"].([]any); ok {
					for _, e := range dep {
						if s, ok := e.(string); ok {
							td.Depends = append(td.Depends, s)
						}
					}
				}
				conf.Tasks.Defs[k] = td
			}
		}
	}

	return &conf, nil
}
