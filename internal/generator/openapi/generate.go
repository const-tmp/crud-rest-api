package openapi

import (
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
		Schema   string
		Resource string
		Tag      string
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

	for _, resource := range resources {
		if _, err = os.Open(src + "schemas/" + resource.Schema + ".yaml"); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("%s: %w", src+"schemas/"+resource.Schema+".yaml", err)
		}

		resourcePath := src + "paths/" + resource.Resource

		if err = os.MkdirAll(resourcePath+"/{id}", 0755); err != nil {
			return err
		}

		if err = GenerateResource(resource, resourcePath); err != nil {
			return err
		}

		AddPaths(*spec, resource.Resource)
	}

	file, err := os.OpenFile(src+"openapi.yaml", os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if err = yaml.NewEncoder(file).Encode(spec); err != nil {
		return err
	}

	return nil
}

func ReadSpec(path string) (*common.OpenAPI3, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	spec := new(common.OpenAPI3)

	if err = yaml.Unmarshal(data, spec); err != nil {
		return nil, err
	}

	return spec, nil
}

func GenerateResource(resource TemplateData, path string) error {
	units := []RenderUnit{
		{Get, path + "/get.yaml"},
		{Post, path + "/post.yaml"},
		{GetByID, path + "/{id}/get.yaml"},
		{Put, path + "/{id}/put.yaml"},
		{Patch, path + "/{id}/patch.yaml"},
		{Delete, path + "/{id}/delete.yaml"},
	}
	for _, unit := range units {
		if reader, err := unit.Template(resource); err != nil {
			return fmt.Errorf("template %s execution error: %w", unit.Path, err)
		} else {
			if err = generator.WriteFile(unit.Path, reader); err != nil {
				return fmt.Errorf("write %s error: %w", unit.Path, err)
			}
		}
	}

	return nil
}

func AddPaths(spec common.OpenAPI3, resourceName string) {
	resourcePath := "./paths/" + resourceName

	spec.Paths["/"+resourceName] = map[string]common.Ref{
		"get":  {Ref: resourcePath + "/get.yaml"},
		"post": {Ref: resourcePath + "/post.yaml"},
	}

	spec.Paths["/"+resourceName+"/{id}"] = map[string]common.Ref{
		"get":    {Ref: resourcePath + "/{id}/get.yaml"},
		"put":    {Ref: resourcePath + "/{id}/put.yaml"},
		"patch":  {Ref: resourcePath + "/{id}/patch.yaml"},
		"delete": {Ref: resourcePath + "/{id}/delete.yaml"},
	}
}
