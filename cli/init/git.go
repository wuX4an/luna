package _init

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func GitInit(dir string) {
	// Si ya existe un .git/, no hacer nada
	if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
		return
	}

	if _, err := exec.LookPath("git"); err != nil {
		fmt.Println("ℹ️ Git not found, skipping repository initialization")
		return
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
