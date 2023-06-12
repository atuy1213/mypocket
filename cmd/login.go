/*
Copyright © 2023 NAME HERE atuy1213@gmail.com
*/
package cmd

import (
	"context"
	"log"

	"github.com/atuy1213/mypocket/pkg/config"
	"github.com/atuy1213/mypocket/pkg/pocket"
	"github.com/atuy1213/mypocket/pkg/usecase"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to your pocket",
	Long: `login is command to connect your pocket.
You need to register consumer key at pocket.
For more inofmation, see at https://getpocket.com/developer/apps/new`,
	Run: func(cmd *cobra.Command, args []string) {
		pocket := pocket.NewPocketClient()
		config := config.NewConfig()
		executor := usecase.NewLoginExecutor(pocket, config)

		ctx := context.Background()
		if err := executor.Login(ctx); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
