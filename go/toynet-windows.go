package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {

	fileName := "docker-compose.yml"
	url := "https://raw.githubusercontent.com/Project-Reclass/toynet-react/main/docker-compose.yml"

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get(url)
	log.Printf("Pulling Toynet config from Github....")
	if err != nil {
		log.Printf("Toynet Container failed to run with error: %v", err)
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)
	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

	_, toynetErr := toynetFunc("docker-compose", "-f", "docker-compose.yml", "up", "-d", "--build")
	if err != nil {
		log.Printf("Toynet Container failed to run with error: %v", toynetErr)
	} else {
		fmt.Println("Toynet is running at http://localhost:3000")
	}
}

func toynetFunc(cmd string, args ...string) (string, error) {
	toynet := exec.Command(cmd, args...)
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	toynet.Stdout = mw
	toynet.Stderr = mw

	toynetErr := toynet.Run()

	return stdBuffer.String(), toynetErr
}
