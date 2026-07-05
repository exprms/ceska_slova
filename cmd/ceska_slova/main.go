package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/exprms/ceska_slova/internal/loader"
	"github.com/exprms/ceska_slova/internal/renderer"
)

func main() {
	// 1. CLI flag definieren
	format := flag.String("format", "cli", "output format: cli | tts")

	flag.Parse()

	// 2. Daten laden
	files, err := filepath.Glob("data/*.yaml")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	// 3. Dateien verarbeiten
	for _, file := range files {
		f, err := loader.Load(file)
		if err != nil {
			fmt.Println("error:", file, err)
			continue
		}

		// 4. Renderer auswählen
		switch *format {
		case "cli":
			renderer.PrintCLI(f)

		case "tts":
			renderer.PrintTTS(f)

		case "markdown":
			renderer.PrintMarkdown(f)

		default:
			fmt.Println("unknown format:", *format)
			os.Exit(1)
		}
	}
}