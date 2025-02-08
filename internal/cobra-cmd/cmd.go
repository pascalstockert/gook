package cobra_cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gook",
}

var cmdCron = &cobra.Command{
	Use:   "cron",
	Short: "Manage cron jobs",
	Long:  `cron lets you add and remove jobs.`,
}

var cmdCronAdd = &cobra.Command{
	Use:   "add",
	Short: "Add cron job",
	Long:  `Add a cron job.`,
	Run:   CronAdd,
}

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
		serverPath := "./http"
		process := exec.Command("go", "run", serverPath)
		process.Env = os.Environ()

		if err := process.Start(); err != nil {
			fmt.Println("Error starting server process:", err)
			return
		}

		fmt.Println("Server started")
	},
}

var cmdServerStop = &cobra.Command{
	Use:   "stop",
	Short: "Stop the HTTP server",
	Long:  `Stops the HTTP server gracefully.`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(os.Getenv("HTTP_SERVER_PID_FILE"))
		if err != nil {
			fmt.Println("Could not find a running HTTP server. Are you sure you started one?")
			fmt.Println(err)
			return
		}

		pid, err := strconv.Atoi(string(data))
		if err != nil {
			fmt.Println("Error converting PID file:", err)
			return
		}

		process, err := os.FindProcess(pid)
		if err != nil {
			fmt.Println("Error finding process:", err)
			return
		}

		if err := process.Signal(syscall.SIGTERM); err != nil {
			fmt.Println("Error sending SIGTERM:", err)
			return
		}

		_ = os.Remove(os.Getenv("HTTP_SERVER_PID_FILE"))

		fmt.Println("Server stopped")
	},
}

func Execute() {
	rootCmd.AddCommand(cmdCron)
	cmdCron.AddCommand(cmdCronAdd)
	rootCmd.AddCommand(cmdServer)
	cmdServer.AddCommand(cmdServerStart)
	cmdServer.AddCommand(cmdServerStop)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
