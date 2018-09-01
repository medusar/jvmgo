package main

import (
	"fmt"
)

func main() {
	cmd := parseCmd()
	fmt.Printf("cmd: %v \n", cmd)
	if cmd.versionFlag {
		fmt.Println("version is 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
