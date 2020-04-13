package main

import (
	"log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	)

var rootCmd = &cobra.Command{
	Use: "columbia",
	Short: "The columbia microservice, part of The QuaerereProject",
}

func init() {

	envAppEnv := "app_env"

	viper.SetDefault(envAppEnv, "development")
	viper.SetDefault("debug_mode", false)

	viper.SetEnvPrefix("columbia")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}