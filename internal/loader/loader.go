package loader

import (
	"os"

	"gopkg.in/yaml.v3"
)

type File struct {
	Vocabulary map[string]any `yaml:"vocabulary"`
}

func Load(path string) (*File, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var f File
	err = yaml.Unmarshal(data, &f)
	if err != nil {
		return nil, err
	}

	return &f, nil
}