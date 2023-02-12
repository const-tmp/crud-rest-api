package common

type (
	Ref struct {
		Ref string `yaml:"$ref"`
	}

	OpenAPI3 struct {
		OpenAPI string `yaml:"openapi"`

		Info struct {
			Title       string `yaml:"title"`
			Description string `yaml:"description"`
			Version     string `yaml:"version"`
		} `yaml:"info"`

		Servers []struct {
			URL         string `yaml:"url"`
			Description string `yaml:"description"`
		} `yaml:"servers"`

		Paths map[string]map[string]Ref `yaml:"paths"`

		Components struct {
			Parameters Ref `yaml:"parameters"`
			Responses  Ref `yaml:"responses"`
			Schemas    Ref `yaml:"schemas"`
		} `yaml:"components"`
	}
)
