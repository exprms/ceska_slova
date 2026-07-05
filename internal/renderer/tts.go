package renderer

import "github.com/exprms/ceska_slova/internal/model"

// Nur tschechische Sätze ausgeben
func PrintTTS(file *model.File) {
	for _, entry := range file.Vocabulary {
		for _, t := range entry.Translations {
			for _, ex := range t.Examples {
				// Nur der tschechische Satz
				println(ex)
				println("---")
				println(ex)
				println("---")
			}
		}
	}
}