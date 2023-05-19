/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/atuy1213/mypocket/pkg/pocket"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to your pocket",
	Long: `login is command to connect your pocket.
You need to register consumer key at pocket.
For more inofmation, see at https://getpocket.com/developer/apps/new`,
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func login() {

	fmt.Print("CONSUMER KEY > ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	consumerKey := scanner.Text()

	ctx := context.Background()
	client := pocket.NewPocketClient()

	code, err := client.GetAuthCode(ctx, consumerKey)
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := client.Authorize(ctx, code); err != nil {
		log.Fatalln(err)
	}

	accessToken, err := client.GetAccessToken(ctx, consumerKey, code)
	if err != nil {
		log.Fatalln(err)
	}

	viper.Set("consumer_key", consumerKey)
	viper.Set("access_token", accessToken)
	if err := viper.WriteConfig(); err != nil {
		log.Fatalln(err)
	}
}
