package cmd

import (
	"example-api/config"
	"example-api/log"
	"example-api/server"
	"fmt"

	"github.com/spf13/cobra"
)

func initFlags()  {
	root.Flags().String("logLevel", "info", "how much log things")
	root.Flags().String("addr", "0.0.0.0", "In which address listen for incoming requests")
	root.Flags().Int("port", 8080, "In which port of address listen for incoming requests")
}

var root = &cobra.Command{
	Use:   "example-api",
	Short: "example api",
	Long:  "this application acts a swagger-documented example api that provides CRUD endpoints to interact to in-memory database",
	Run: func(cmd *cobra.Command, args []string) {

		err := config.Init(cmd.Flags())
		if err != nil {
			err = fmt.Errorf("unable to initialize configuration: %w", err)
			log.Logger.Fatalln(err)
		}

		log.Init()
		server.Init()

		err = server.Run()
		if err != nil {
			err = fmt.Errorf("error when serving: %w", err)
			log.Logger.Fatalln(err)
		}

	},
}

func Execute() {

	initFlags()
	root.Execute()

}
