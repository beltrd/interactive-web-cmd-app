package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	// mt.Println("ideog")
	r := mux.NewRouter()
	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")

	// Routes Handles / Endpoints
	r.HandleFunc("/api/command/{cmd}", getCommand).Methods("GET")
	r.PathPrefix("/home").Handler(http.StripPrefix("/home", http.FileServer(http.Dir(dir))))
	// message
	fmt.Println("Server starting...")
	// to start the server
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getCommand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(runCommand(w, params["cmd"]))
}

func runCommand(w http.ResponseWriter, commandStr string) string {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()
	stdoutStderr, err := cmd.CombinedOutput()

	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())

	outStr1, errStr1 := string(stdoutStderr), string(err.Error())

	if outStr != "" {
		return outStr
	}
	if errStr != "" {
		return errStr
	}
	if outStr1 != "" {
		return outStr1
	}
	if errStr1 != "" {
		return errStr1
	}
	return ""
}
