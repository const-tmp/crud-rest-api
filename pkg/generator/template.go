package generator

import (
	"bytes"
	"io"
	"os"
	"text/template"
)

type (
	RenderUnit struct {
		Template func(dot ResourceTemplate) (io.Reader, error)
		Path     string
	}

	ResourceTemplate struct {
		Schema   string
		Resource string
		Tag      string
	}
)

func RenderTemplate(name, text string, dot ResourceTemplate) (io.Reader, error) {
	tmpl, err := template.New(name).Parse(text)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, dot); err != nil {
		return nil, err
	}

	return buf, nil
}

func WriteFile(path string, reader io.Reader) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = io.Copy(file, reader); err != nil {
		return err
	}

	return nil
}
