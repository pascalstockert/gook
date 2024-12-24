package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-webhook/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "gook",
}

var cmdServe = &cobra.Command{
	Use:   "serve [port]",
	Short: "Serve API for querying statuses",
	Long: `serve is for exposing gathered health-data via a HTTP-API.
Optionally a port can be passed which defaults to 4321.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var port string
		if len(args) == 0 {
			port = "4321"
		} else {
			port = args[0]
		}

		fmt.Printf("Serving API on port %s\n", port)
		/*
			TODO execute http server in own process, not goroutine
			1. execute in separate process, save process-ID in .pid-file
			2. implement command to stop process by reading process-ID in .pid-file
		*/
		go http.Execute(port)
	},
}

func Execute() {
	rootCmd.AddCommand(cmdServe)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
