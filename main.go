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
	stdOut, stdErr, err := crcos.RunWithDefaultLocale(path, "is-active", "network-manager")

	if err != nil {
		fmt.Printf("%v : %s", err, stdErr)
	} else {

		if strings.TrimSpace(stdOut) != "active" {
			fmt.Printf("NetworkManager.service is not running")
		}
		fmt.Println("NetworkManager.service is already running")
	}

}