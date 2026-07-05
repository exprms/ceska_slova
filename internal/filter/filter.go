package filter

import "github.com/exprms/ceska_slova/internal/model"

func MatchesTag(entry model.Entry, tags []string) bool {
	if len(tags) == 0 {
		return true
	}

	for _, t := range entry.Tags {
		for _, filter := range tags {
			if t == filter {
				return true
			}
		}
	}

	return false
}