package openapi

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func FilterDefaultSchemas(dir string) ([]string, error) {
	var schemas []string

	exclude := map[string]struct{}{
		"Any.yaml":       {},
		"AnyNull.yaml":   {},
		"BaseModel.yaml": {},
		"Error.yaml":     {},
		"_index.yaml":    {},
	}

	if err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if _, ok := exclude[d.Name()]; d.IsDir() || ok {
			return nil
		}
		schemas = append(schemas, path)
		fmt.Println(path, d, err)
		return err
	}); err != nil {
		return nil, err
	}
	return schemas, nil
}
