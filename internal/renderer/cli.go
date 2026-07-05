package renderer

import (
	"fmt"

	"github.com/exprms/ceska_slova/internal/model"
)

func Print(file *model.File) {
	for german, entry := range file.Vocabulary {
		fmt.Println(german)

		for _, t := range entry.Translations {
			line := fmt.Sprintf("  - %s (%s", t.Word, t.POS)

			if t.Gender != "" {
				line += ", " + t.Gender
			}

			if t.Aspect != "" {
				line += ", " + t.Aspect
			}

			line += ")"

			fmt.Println(line)

			for _, ex := range t.Examples {
				fmt.Println("    - " + ex)
			}
		}

		fmt.Println()
	}
}