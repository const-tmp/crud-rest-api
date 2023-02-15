package bootstrap

import (
	"embed"
	"fmt"
	"github.com/nullc4t/crud-rest-api/pkg/generator"
	"path/filepath"
	"strings"
)

//go:embed api/* api/src/parameters/_index.yaml api/src/responses/_index.yaml api/src/schemas/_index.yaml gen.go.tmpl
var f embed.FS

func OpenAPIFiles(path string) error {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	files := []string{
		"api/src/openapi.yaml",
		"api/gen.go",
		"api/src/parameters/_index.yaml",
		"api/src/parameters/path/id.yaml",
		"api/src/parameters/query/limit.yaml",
		"api/src/parameters/query/offset.yaml",
		"api/src/parameters/query/sort.yaml",
		"api/src/responses/_index.yaml",
		"api/src/responses/OK.yaml",
		"api/src/responses/Created.yaml",
		"api/src/responses/NoContent.yaml",
		"api/src/responses/NullResponse.yaml",
		"api/src/responses/BadRequest.yaml",
		"api/src/responses/Unauthorized.yaml",
		"api/src/responses/NotFound.yaml",
		"api/src/responses/UnexpectedError.yaml",
		"api/src/schemas/_index.yaml",
		"api/src/schemas/BaseModel.yaml",
		"api/src/schemas/Any.yaml",
		"api/src/schemas/AnyNull.yaml",
		"api/src/schemas/Error.yaml",
	}

	for _, name := range files {
		if file, err := f.Open(name); err != nil {
			return fmt.Errorf("open %s error: %w", name, err)
		} else if err = generator.WriteFile(path+name, file); err != nil {
			return err
		}
	}

	return nil
}

func GenFile(path, module, apiPkg, oapiPkg, out string) error {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	data, err := f.ReadFile("gen.go.tmpl")
	if err != nil {
		return fmt.Errorf("read gen.go.tmpl error: %w", err)
	}

	if out == "" {
		out = fmt.Sprintf("pkg/%s/%s.gen.go", oapiPkg, oapiPkg)
	}
	outDir := filepath.Dir(out)

	reader, err := generator.RenderTemplate("gen.go", string(data), map[string]string{
		"module":            module,
		"api_package":       apiPkg,
		"oapi_package_name": oapiPkg,
		"oapi_out_file":     out,
		"oapi_out_dir":      outDir,
	})
	if err != nil {
		return err
	}

	if err = generator.WriteFile(path+"gen.go", reader); err != nil {
		return err
	}

	return nil
}
