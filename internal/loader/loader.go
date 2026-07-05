package loader

import (
	"os"

	"github.com/exprms/ceska_slova/internal/model"
	"gopkg.in/yaml.v3"
)

func Load(path string) (*model.File, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var f model.File
	err = yaml.Unmarshal(data, &f)
	if err != nil {
		return nil, err
	}

	return &f, nil
}