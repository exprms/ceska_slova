package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/exprms/ceska_slova/internal/loader"
	"github.com/exprms/ceska_slova/internal/renderer"
)

func main() {
	files, err := filepath.Glob("data/*.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		f, err := loader.Load(file)
		if err != nil {
			fmt.Println("error:", file, err)
			continue
		}

		renderer.Print(f)
	}
}