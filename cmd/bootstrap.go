/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/nullc4t/crud-rest-api/internal/generator/bootstrap"
	"os"

	"github.com/spf13/cobra"
)

// bootstrapCmd represents the bootstrap command
var bootstrapCmd = &cobra.Command{
	Use:     "bootstrap [path]",
	Aliases: []string{"boot", "bs"},
	Short:   "Bootstrap OpenAPI 3 files",
	Long:    `Bootstrap OpenAPI 3 files in specified path or CWD if omitted`,
	Example: `bootstrap
bootstrap .
bootstrap desired/path
bootstrap /home/user/desired/path
`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) == 0 {
			path = "."
		} else {
			path = args[0]
		}
		fmt.Println("Bootstrap path:", path)

		if err := bootstrap.Mkdir(path); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := bootstrap.Files(path); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bootstrapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bootstrapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
