/*
Copyright © 2023 NAME HERE atuy1213@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/atuy1213/mypocket/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mypocket",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".mypocket" (without extension).
	viper.AddConfigPath(fmt.Sprintf(`%s/%s`, home, config.CONFIG_FILE_DIR))
	viper.SetConfigName(config.CONFIG_FILE_NAME)
	viper.SetConfigType(config.CONFIG_FILE_EXTENSION)

	viper.AutomaticEnv() // read in environment variables that match

	path, err := config.NewConfig().GetConfigFilePath()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if _, err := os.Create(path); err != nil {
				log.Fatalln(err)
			}
		} else {
			log.Fatalln(err)
		}
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}
