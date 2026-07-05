package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/exprms/ceska_slova/internal/filter"
	"github.com/exprms/ceska_slova/internal/loader"
	"github.com/exprms/ceska_slova/internal/model"
	"github.com/exprms/ceska_slova/internal/renderer"
)

func main() {
	// CLI flags
	format := flag.String("format", "cli", "cli | tts | markdown | table")
	tag := flag.String("tag", "", "filter by tag (comma separated)")

	flag.Parse()

	var tagList []string
	if *tag != "" {
		tagList = strings.Split(*tag, ",")
	}

	// 1. merged vocabulary (WICHTIG)
	merged := model.File{
		Vocabulary: make(map[string]model.Entry),
	}

	// 2. load files
	files, err := filepath.Glob("data/*.yaml")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	// 3. merge + filter
	for _, file := range files {
		f, err := loader.Load(file)
		if err != nil {
			fmt.Println("error:", file, err)
			continue
		}

		for german, entry := range f.Vocabulary {
			if filter.MatchesTag(entry, tagList) {
				merged.Vocabulary[german] = entry
			}
		}
	}

	// 4. render
	switch *format {
	case "cli":
		renderer.PrintCLI(&merged)

	case "tts":
		renderer.PrintTTS(&merged)

	case "markdown":
		renderer.PrintMarkdown(&merged)

	case "table":
		renderer.PrintTable(&merged)

	default:
		fmt.Println("unknown format:", *format)
		os.Exit(1)
	}
}