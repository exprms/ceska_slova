package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/exprms/ceska_slova/internal/loader"
)

func main() {
	dir := "data"

	files, err := filepath.Glob(filepath.Join(dir, "*.yaml"))
	if err != nil {
		fmt.Println("error reading directory:", err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Println("no yaml files found")
		return
	}

	hasError := false

	for _, file := range files {
		fmt.Println("checking:", file)

		f, err := loader.Load(file)
		if err != nil {
			fmt.Printf("❌ %s: invalid yaml: %v\n", file, err)
			hasError = true
			continue
		}

		// nur Strukturcheck (Phase 1 minimal)
		if f.Vocabulary == nil {
			fmt.Printf("❌ %s: missing 'vocabulary'\n", file)
			hasError = true
			continue
		}

		fmt.Printf("✔ %s OK (%d entries)\n", file, len(f.Vocabulary))
	}

	if hasError {
		os.Exit(1)
	}
}