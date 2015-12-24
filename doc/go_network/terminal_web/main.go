package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

const port = ":8080"

func main() {
	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/plain", handlerPlain)

	log.Println("Serving http://localhost" + port)
	if err := http.ListenAndServe(port, mainRouter); err != nil {
		panic(err)
	}
}

func handlerPlain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		cs := []string{"/bin/bash", "-c", "ls"}
		cmd := exec.Command(cs[0], cs[1:]...)
		cmd.Stdin = nil
		cmd.Stdout = w
		cmd.Stderr = w

		if err := cmd.Start(); err != nil {
			fmt.Fprintf(w, "Failed to start %v\n", cs)
			return
		}
		fmt.Println("PID:", cmd.Process.Pid)
		cmd.Wait()

	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}
