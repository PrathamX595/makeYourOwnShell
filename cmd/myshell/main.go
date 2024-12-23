package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	finished := false
	
	// Wait for user input
	for !finished{
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command) + ": command not found"
		fmt.Println(command)
	}

}
