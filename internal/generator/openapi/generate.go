package openapi

import (
	"embed"
	"errors"
	"fmt"
	"github.com/nullc4t/crud-rest-api/pkg/common"
	"github.com/nullc4t/crud-rest-api/pkg/generator"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type (
	TemplateData struct {
		Schema   string `yaml:"schema"`
		Resource string `yaml:"resource"`
		Tag      string `yaml:"tag"`
	}

	RenderUnit struct {
		Template func(dot TemplateData) (io.Reader, error)
		Path     string
	}
)

func Generate(src string, resources []TemplateData) error {
	spec, err := ReadSpec(src + "openapi.yaml")
	if err != nil {
		return err
	}

	index, err := ReadSchemaIndex(src + "schemas/_index.yaml")
	if err != nil {
		return err
	}

	for _, resource := range resources {
		schemaPath := src + "schemas/" + resource.Schema + ".yaml"
		if _, err = os.Open(schemaPath); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("%s: %w", schemaPath, err)
		}

		index[resource.Schema] = common.Ref{Ref: fmt.Sprintf("./%s.yaml", resource.Schema)}

		resourcePath := src + "paths/" + resource.Resource

		if err = os.MkdirAll(resourcePath+"/{id}", 0755); err != nil {
			return err
		}

		if err = GenerateResource(resource, resourcePath); err != nil {
			return err
		}

		AddPaths(spec, resource.Resource)
	}

	if err = WriteYAMLFile(src+"openapi.yaml", spec); err != nil {
		return err
	}

	if err = WriteYAMLFile(src+"schemas/_index.yaml", index); err != nil {
		return err
	}

	return nil
}

func WriteYAMLFile(path string, obj any) error {
	file, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open %s error: %w", path, err)
	}
	defer file.Close()

	if err = yaml.NewEncoder(file).Encode(obj); err != nil {
		return fmt.Errorf("encode YAML to %s error: %w", path, err)
	}

	return nil
}

func ReadSpec(path string) (*common.OpenAPI3, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s error: %w", path, err)
	}

	spec := new(common.OpenAPI3)

	if err = yaml.Unmarshal(data, spec); err != nil {
		return nil, fmt.Errorf("unmarshal YAML %s error: %w", path, err)
	}

	return spec, nil
}

func ReadSchemaIndex(path string) (map[string]common.Ref, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s error: %w", path, err)
	}

	index := make(map[string]common.Ref)

	if err = yaml.Unmarshal(data, index); err != nil {
		return nil, fmt.Errorf("unmarshal YAML %s error: %w", path, err)
	}

	return index, nil
}

func ReadTemplateData(path string) ([]TemplateData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s error: %w", path, err)
	}

	var td []TemplateData

	if err = yaml.Unmarshal(data, &td); err != nil {
		return nil, fmt.Errorf("unmarshal YAML %s error: %w", path, err)
	}

	return td, nil
}

//go:embed routes/*
var f embed.FS

func GenerateResource(resource TemplateData, path string) error {
	tmplPrefix := "routes"
	routes := []string{
		"/get.yaml",
		"/post.yaml",
		"/{id}/get.yaml",
		"/{id}/put.yaml",
		"/{id}/patch.yaml",
		"/{id}/delete.yaml",
	}

	for _, route := range routes {
		data, err := f.ReadFile(tmplPrefix + route)
		if err != nil {
			return fmt.Errorf("read %s error: %w", path+route, err)
		}

		reader, err := generator.RenderTemplate(route, string(data), resource)
		if err != nil {
			return fmt.Errorf("render %s error: %w", path+route, err)
		}

		if err = generator.WriteFile(path+route, reader); err != nil {
			return fmt.Errorf("write %s error: %w", path+route, err)
		}
	}

	return nil
}

func AddPaths(spec *common.OpenAPI3, resourceName string) {
	resourcePath := "./paths/" + resourceName

	if spec.Paths == nil {
		spec.Paths = make(map[string]map[string]common.Ref)
	}

	if path, ok := spec.Paths["/"+resourceName]; ok {
		path["get"] = common.Ref{Ref: resourcePath + "/get.yaml"}
		path["post"] = common.Ref{Ref: resourcePath + "/post.yaml"}
	} else {
		spec.Paths["/"+resourceName] = map[string]common.Ref{
			"get":  {Ref: resourcePath + "/get.yaml"},
			"post": {Ref: resourcePath + "/post.yaml"},
		}
	}

	if path, ok := spec.Paths["/"+resourceName+"/{id}"]; ok {
		path["get"] = common.Ref{Ref: resourcePath + "/{id}/get.yaml"}
		path["put"] = common.Ref{Ref: resourcePath + "/{id}/put.yaml"}
		path["patch"] = common.Ref{Ref: resourcePath + "/{id}/patch.yaml"}
		path["delete"] = common.Ref{Ref: resourcePath + "/{id}/delete.yaml"}
	} else {
		spec.Paths["/"+resourceName+"/{id}"] = map[string]common.Ref{
			"get":    {Ref: resourcePath + "/{id}/get.yaml"},
			"put":    {Ref: resourcePath + "/{id}/put.yaml"},
			"patch":  {Ref: resourcePath + "/{id}/patch.yaml"},
			"delete": {Ref: resourcePath + "/{id}/delete.yaml"},
		}
	}
}
