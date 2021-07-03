package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "hail",
	Short: "hail is a cross platform script management tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd is called!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error while UserHomeDir: ", err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigFile("hail.yaml")
	}
	// read environment variable that matches.
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using Config File : ", viper.ConfigFileUsed())
	}
}
