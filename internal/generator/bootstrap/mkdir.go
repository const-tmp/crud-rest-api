package bootstrap

import (
	"os"
	"strings"
)

func Mkdir(path string) error {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	paths := []string{
		"api/build",
		"api/src/parameters/path",
		"api/src/parameters/query",
		"api/src/paths",
		"api/src/responses",
		"api/src/schemas",
	}

	for _, p := range paths {
		if err := os.MkdirAll(path+p, 0755); err != nil {
			return err
		}
	}

	return nil
}
