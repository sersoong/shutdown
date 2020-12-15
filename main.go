package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("shutdown", "-s", "-t", "0")
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
