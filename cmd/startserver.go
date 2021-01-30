package cmd

import (
	"muramasa/internal/infrastructure"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startServerCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	ginServer := infrastructure.NewServer(
		viper.GetInt("server.port"),
		infrastructure.DebugMode,
	)

	routerLogger := infrastructure.NewLogger("development",
		viper.GetString("log.level"),
		viper.GetString("log.format"),
	)

	db := infrastructure.NewDatabaseConnection()

	infrastructure.NewRouterWithLogger(*routerLogger, db.Conn).SetRoutes(ginServer.Router)

	ginServer.Start()
}
