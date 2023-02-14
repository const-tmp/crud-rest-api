package cmd

import (
	"fmt"
	"github.com/nullc4t/crud-rest-api/internal/generator/openapi"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:     "generate [path] config.yaml",
	Aliases: []string{"gen", "g"},
	Short:   "Generate resources",
	Long:    `Generate OpenAPI3 spec from config.yaml`,
	Example: `generate config.yaml
generate path/to/config.yaml
generate /abs/path/to/config.yaml
generate . config.yaml
generate path/to/dir config.yaml
generate /abs/path/to/dir config.yaml
`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) == 1 {
			path = "."
		} else {
			path = args[0]
		}
		config := args[len(args)-1]
		fmt.Println("Generate path:", path)
		fmt.Println("Config path:", config)

		if !strings.HasSuffix(path, "/") {
			path += "/"
		}

		srcPath := path + "api/src/"
		if !strings.HasSuffix(srcPath, "/") {
			srcPath += "/"
		}

		resources, err := openapi.ReadTemplateData(config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := openapi.Generate(srcPath, resources); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
