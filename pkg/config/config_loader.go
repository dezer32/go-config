package config

import (
	"github.com/go-yaml/yaml"
	"os"
)

func ConfigLoader(v any, files ...string) (err error) {
	for _, file := range files {
		var f *os.File
		f, err = os.Open(file)
		if err != nil {
			return
		}
		defer f.Close()

		err = yaml.NewDecoder(f).Decode(v)
		if err != nil {
			return
		}
	}

	return
}
