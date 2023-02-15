package generator

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"text/template"
)

func RenderTemplate(name, text string, dot any) (io.Reader, error) {
	tmpl, err := template.New(name).Parse(text)
	if err != nil {
		return nil, fmt.Errorf("parse %s error: %w", name, err)
	}

	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, dot); err != nil {
		return nil, fmt.Errorf("execute %s error: %w", name, err)
	}

	return buf, nil
}

func WriteFile(path string, reader io.Reader) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open %s error: %w", path, err)
	}
	defer file.Close()

	if _, err = io.Copy(file, reader); err != nil {
		return fmt.Errorf("write %s error: %w", path, err)
	}

	return nil
}
