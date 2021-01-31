package cmd

import (
	"fmt"
	"log"
	"muramasa/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "muramasa",
	Short: "Welcome to the beginning of nothingness.",
	Run: func(cmd *cobra.Command, args []string) {
		run()
		fmt.Println("Welcome to the beginning of nothingness.")
	},
}

func init() {
	rootCmd.AddCommand(startServerCmd)

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Configuration file not found.")
	}

	config.ServerConfig(rootCmd)
	config.LoggerConfig(rootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal()
	}
}
