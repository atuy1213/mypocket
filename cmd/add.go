/*
Copyright © 2023 NAME HERE atuy1213@gmail.com
*/
package cmd

import (
	"context"
	"errors"
	"log"

	"github.com/atuy1213/mypocket/pkg/interceptor"
	"github.com/atuy1213/mypocket/pkg/pocket"
	"github.com/atuy1213/mypocket/pkg/usecase"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add to your pocket",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New(ERROR_MESSAGE)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		URL := args[0]

		pocket := pocket.NewPocketClient()
		authInterceptor := interceptor.NewAuthInterceptor()
		executor := usecase.NewAddExecutor(pocket, authInterceptor)

		ctx := context.Background()
		if err := executor.Add(ctx, URL); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const ERROR_MESSAGE = `
requires URL parameter.
e.g.) pocket add "https://example.com"`
