package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"syscall"
)

var rootCmd = &cobra.Command{
	Use: "gook",
}

/*
TODO execute http server in own process
1. execute in separate process, save process-ID in .pid-file
2. implement command to stop process by reading process-ID in .pid-file
3. move cron-logic into own process
*/
var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "Launch or stop a HTTP server",
	Long:  `server is for exposing gathered health-data via a HTTP-API.`,
}

var cmdServerStart = &cobra.Command{
	Use:   "start",
	Short: "Launch a HTTP server",
	Long:  `Launch a HTTP server that exposes gathered health-data via a HTTP-API.`,
	Run: func(cmd *cobra.Command, args []string) {
		serverPath := "./http/server.go"
		process := exec.Command("go", "run", serverPath)
		process.SysProcAttr = &syscall.SysProcAttr{
			// On Unix, set to make the child process independent
			Setpgid: true,
		}

		if err := process.Start(); err != nil {
			fmt.Println("Error starting server process:", err)
			return
		}

		file, fileCreationError := os.Create("http_server.pid")
		if fileCreationError != nil {
			fmt.Println("Error creating PID file:", fileCreationError)
			return
		}

		_, fileWriteError := fmt.Fprintf(file, "%d", process.Process.Pid)
		if fileWriteError != nil {
			fmt.Println("Error writing PID file:", fileWriteError)
			return
		}

		_ = file.Close()

		fmt.Println("Server started with PID:", process.Process.Pid)
	},
}

func Execute() {
	rootCmd.AddCommand(cmdServer)
	cmdServer.AddCommand(cmdServerStart)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
