package cmd

import (
	"github.com/rasouliali1379/movie-api/cmd/api"
	"github.com/rasouliali1379/movie-api/cmd/user"
	"github.com/spf13/cobra"
	"log"
)

var cmdList = []*cobra.Command{
	{
		Use:   "start-api",
		Short: "Start the api server",
		Long: `
			Start the api server on preferred port (default is :3000).
			Use 'movie-db start-api :<port>'
			to start api on your preferred port`,
		Run: func(cmd *cobra.Command, args []string) {

			port := ""

			if len(args) > 0 {
				port = args[0]
			}

			log.Fatalf("Error occured : %s", api.StartApi(port))
		},
	},
	{
		Use:   "login",
		Short: "Login to Panel",
		Long:  `Proceeding login you will receive a access token to manage things inside server`,
		Run: func(cmd *cobra.Command, args []string) {
			err := user.Login()
			if err != nil {
				log.Fatalf("Error occured : %s", err)
			}
		},
	},
	{
		Use:   "login",
		Short: "Login to Panel",
		Long:  `Proceeding login you will receive a access token to manage things inside server`,
		Run: func(cmd *cobra.Command, args []string) {
			err := user.Login()
			if err != nil {
				log.Fatalf("Error occured : %s", err)
			}
		},
	},
	{
		Use:   "add",
		Short: "Add user, movies and etc",
		Long: `Add anything to app database using
								movie-db add <database>`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				switch args[0] {
				case "user":
					err := user.Add()
					if err != nil {
						log.Fatalf("Error occured : %s", err)
					}
				}
			} else {
				log.Fatal("You to specify a database")
			}
		},
	},
}

var rootCmd = &cobra.Command{
	Use:     "movie-api",
	Short:   "Movie Api is a cms for movie download website",
	Long:    `Upload and download movie files , register users, add comments, like the posts and etc.`,
	Version: "1.0.0",
}

func addCommands() {
	for _, cmd := range cmdList {
		rootCmd.AddCommand(cmd)
	}
}
