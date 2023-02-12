package resource

import (
	"errors"
	"fmt"
	"github.com/nullc4t/crud-rest-api/pkg/common"
	"github.com/nullc4t/crud-rest-api/pkg/generator"
	"github.com/nullc4t/crud-rest-api/pkg/generator/openapi/paths"
	"gopkg.in/yaml.v3"
	"os"
)

func Generate(src string, resources []generator.ResourceTemplate) error {
	spec, err := ReadSpec(src + "openapi.yaml")
	if err != nil {
		return err
	}

	for _, resource := range resources {
		if _, err = os.Open(src + "schemas/" + resource.Schema + ".yaml"); errors.Is(err, os.ErrNotExist) {
			return err
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

func GenerateResource(resource generator.ResourceTemplate, path string) error {
	units := []generator.RenderUnit{
		{paths.Get, path + "/get.yaml"},
		{paths.Post, path + "/post.yaml"},
		{paths.GetByID, path + "/{id}/get.yaml"},
		{paths.Put, path + "/{id}/put.yaml"},
		{paths.Patch, path + "/{id}/patch.yaml"},
		{paths.Delete, path + "/{id}/delete.yaml"},
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

func AddPaths(d common.OpenAPI3, resourceName string) {
	path := "./paths/" + resourceName

	d.Paths["/"+resourceName] = map[string]common.Ref{
		"get":  {Ref: path + "/get.yaml"},
		"post": {Ref: path + "/post.yaml"},
	}

	d.Paths["/"+resourceName+"/{id}"] = map[string]common.Ref{
		"get":    {Ref: path + "/{id}/get.yaml"},
		"put":    {Ref: path + "/{id}/put.yaml"},
		"patch":  {Ref: path + "/{id}/patch.yaml"},
		"delete": {Ref: path + "/{id}/delete.yaml"},
	}
}
