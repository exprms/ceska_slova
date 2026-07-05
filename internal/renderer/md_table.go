package renderer

import (
	"fmt"
	"strings"

	"github.com/exprms/ceska_slova/internal/model"
)

func PrintTable(file *model.File) {
	// Header
	fmt.Println("")
	fmt.Println("| Deutsch | Übersetzung | Grammatik / Notes | Beispiele |")
	fmt.Println("|--------|------------|-------------------|----------|")

	for german, entry := range file.Vocabulary {
		for _, t := range entry.Translations {

			// Grammatik / Notes bauen
			var meta []string

			if t.POS != "" {
				meta = append(meta, t.POS)
			}
			if t.Aspect != "" {
				meta = append(meta, t.Aspect)
			}
			if t.Gender != "" {
				meta = append(meta, t.Gender)
			}

			if len(t.Notes) > 0 {
				meta = append(meta, strings.Join(t.Notes, "; "))
			}

			metaStr := strings.Join(meta, ", ")

			// Beispiele zusammenfassen
			examples := ""
			if len(t.Examples) > 0 {
				examples = strings.Join(t.Examples, " / ")
			}

			// Zeile
			fmt.Printf(
				"| %s | %s | %s | %s |\n",
				german,
				t.Word,
				metaStr,
				examples,
			)
		}
	}
}