package cmd

import (
	"fmt"
	"github.com/nullc4t/crud-rest-api/internal/generator/openapi"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate resources",
	Long:  `Generate OpenAPI3 spec from list of templates`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")

		srcPath := "api/src/"
		if !strings.HasSuffix(srcPath, "/") {
			srcPath += "/"
		}

		resources := []openapi.TemplateData{
			{Schema: "Account", Resource: "accounts", Tag: "Account"},
			{Schema: "Service", Resource: "services", Tag: "Service"},
			{Schema: "Permission", Resource: "permissions", Tag: "Permissions"},
		}

		if err := openapi.Generate(srcPath, resources); err != nil {
			log.Fatal(err)
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
