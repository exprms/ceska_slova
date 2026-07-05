package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/exprms/ceska_slova/internal/model"
	"github.com/exprms/ceska_slova/internal/loader"
	"github.com/exprms/ceska_slova/internal/renderer"
	"github.com/exprms/ceska_slova/internal/filter"
)

func main() {
	// 1. CLI flag definieren
	format := flag.String("format", "cli", "cli | tts | markdown")
	tag := flag.String("tag", "", "filter by tag (comma separated)")

	flag.Parse()

	var tagList []string
	if *tag != "" {
		tagList = strings.Split(*tag, ",")
	}

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

		// Filter anwenden
		filtered := make(map[string]model.Entry)

		for german, entry := range f.Vocabulary {
			if filter.MatchesTag(entry, tagList) {
				filtered[german] = entry
			}
		}

		f.Vocabulary = filtered

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