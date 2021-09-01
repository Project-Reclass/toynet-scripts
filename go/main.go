package main

import (
	"log"
	"os/exec"
)

func main() {
	wget := exec.Command("wget", "https://raw.githubusercontent.com/Project-Reclass/toynet-react/master/docker-compose.yml")
	log.Printf("Starting Toynet....")
	wgetErr := wget.Run()
	if wgetErr != nil {
		log.Printf("Toynet Container failed to run with error: %v", wgetErr)
	}

	toynet := exec.Command("docker-compose", "-f", "docker-compose.yml", "up", "-d", "--build")
	log.Printf("Starting Toynet....")
	toynetErr := toynet.Run()
	if toynetErr != nil {
		log.Printf("Toynet Container failed to run with error: %v", toynetErr)
	}
}
