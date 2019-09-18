package main

import (
	"os/exec"
	"fmt"
	"strings"
	crcos "github.com/code-ready/crc/pkg/os"
)


func main(){
	fmt.Println("check systemctl")
	// path, err := exec.LookPath("systemctl")

	// run in wsl
	path, err := exec.LookPath("service")

	if err != nil {
		fmt.Printf("lookpath error: %v", err)
	}

	// stdOut, stdErr, err := crcos.RunWithDefaultLocale(path, "is-active", "NetworkManager")
	stdOut, stdErr, err := crcos.RunWithDefaultLocale(path, "network-manager", "status")

	if err != nil {
		fmt.Printf("%v : %s", err, stdErr)
	} else {

		// if strings.TrimSpace(stdOut) != "active" {
		if strings.Contains(strings.TrimSpace(stdOut), "running") {
			fmt.Println("NetworkManager.service is already running")
		} else{
			fmt.Println("NetworkManager.service is not running")
		}
	}

}