package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LoggerConfig(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().String("log.level", "info", "one of debug, info, warn, error or fatal")
	rootCmd.PersistentFlags().String("log.format", "text", "one of text or json")
	rootCmd.PersistentFlags().Bool("log.line", false, "enable filename and line in logs")

	viper.BindPFlags(rootCmd.PersistentFlags())
}

func ServerConfig(cmd *cobra.Command) {
	cmd.Flags().String("server.host", "127.0.0.1", "host on which the server should listen")
	cmd.Flags().Int("server.port", 8000, "port on which the server should listen")
	cmd.Flags().Bool("server.debug", false, "debug mode for the server")
	cmd.Flags().String("server.allowedOrigins", "*", "allowed origins for the server")
	cmd.Flags().String("jwt.salt", "", "used to sign the JWTs")
	viper.BindPFlags(cmd.Flags())
}
