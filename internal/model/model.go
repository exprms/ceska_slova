package model

type File struct {
	Vocabulary map[string]Entry `yaml:"vocabulary"`
}

type Entry struct {
	Tags         []string      `yaml:"tags,omitempty"`
	Translations []Translation `yaml:"translations"`
}

type Translation struct {
	Word     string `yaml:"word"`
	POS      string `yaml:"pos"`

	Gender   string `yaml:"gender,omitempty"`
	Plural   string `yaml:"plural,omitempty"`

	Aspect   string `yaml:"aspect,omitempty"`
	Pair     string `yaml:"pair,omitempty"`

	Examples []string `yaml:"examples,omitempty"`
}