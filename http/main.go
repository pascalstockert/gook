package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var ServerPid = os.Getenv("HTTP_SERVER_PID_FILE")
var ServerLog = os.Getenv("HTTP_SERVER_LOG_FILE")
var ServerPort = os.Getenv("HTTP_SERVER_PORT")

func jsonMiddleware(closure http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		closure(res, req)
	}
}

func healthHandler(res http.ResponseWriter, req *http.Request) {
	err := json.NewEncoder(res).Encode(`{status: ok}`)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}

func appendLog(file *os.File, content string) {
	_, fileWriteError := fmt.Fprintln(file, content)
	if fileWriteError != nil {
		fmt.Println("Error writing log file:", fileWriteError)
		return
	}
}

func main() {
	pid := os.Getpid()

	file, fileCreationError := os.Create(ServerPid)
	if fileCreationError != nil {
		fmt.Println("Error creating PID file:", fileCreationError)
		return
	}

	_, fileWriteError := fmt.Fprintf(file, "%d", pid)
	if fileWriteError != nil {
		fmt.Println("Error writing PID file:", fileWriteError)
		return
	}

	_ = file.Close()

	file, fileCreationError = os.Create(ServerLog)
	if fileCreationError != nil {
		fmt.Println("Error creating log file:", fileCreationError)
		return
	}

	appendLog(file, "Server up at port "+ServerPort)
	appendLog(file, "Using PID "+strconv.Itoa(os.Getpid()))
	http.HandleFunc("/health", jsonMiddleware(healthHandler))
	http.ListenAndServe(":"+ServerPort, nil)
}
