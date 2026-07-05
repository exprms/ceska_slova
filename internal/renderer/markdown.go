package renderer

import (
	"fmt"
	// "strings"

	"github.com/exprms/ceska_slova/internal/model"
)

func PrintMarkdown(file *model.File) {
	for german, entry := range file.Vocabulary {

		// Überschrift
		fmt.Printf("## %s\n\n", german)

		for _, t := range entry.Translations {

			// Wort + POS
			line := fmt.Sprintf("- %s (*%s*", t.Word, t.POS)

			if t.Aspect != "" {
				line += ", " + t.Aspect
			}

			if t.Gender != "" {
				line += ", " + t.Gender
			}

			line += ")"

			fmt.Println(line)

			// Beispiele eingerückt
			if len(t.Examples) > 0 {
				for _, ex := range t.Examples {
					fmt.Println("  - " + ex)
				}
			}
		}

		fmt.Println()
	}
}