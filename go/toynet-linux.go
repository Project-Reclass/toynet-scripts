package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	resp, err := http.Get("https://raw.githubusercontent.com/Project-Reclass/toynet-react/master/docker-compose.yml")
	log.Printf("Pulling Toynet config from Github....")
	if err != nil {
		log.Printf("Toynet Container failed to run with error: %v", err)
	}
	defer resp.Body.Close()

	toynet := exec.Command("docker-compose", "-f", "docker-compose.yml", "up", "-d", "--build")
	log.Printf("Starting Toynet....")
	toynetErr := toynet.Run()
	if toynetErr != nil {
		log.Printf("Toynet Container failed to run with error: %v", toynetErr)
	} else {
		fmt.Println("Toynet is running at http://localhost:3000")
	}
}
