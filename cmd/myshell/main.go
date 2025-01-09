package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	finished := false
	Path := os.Getenv("PATH")
	Paths := strings.Split(Path, ":")
	for !finished {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		input = strings.TrimSpace(input)
		splitcomm := strings.Split(input, " ")
		command := splitcomm[0]
		switch command {
		case "exit":
			finished = true
			code, _ := strconv.ParseInt(splitcomm[1], 10, 64)
			os.Exit(int(code))
		case "echo":
			fmt.Println(strings.Join(splitcomm[1:], " "))
		case "type":
			com := (splitcomm[1])
			if com == "exit" || com == "echo" || com == "type" || com =="pwd" || com =="cd"{
				fmt.Println(com + " is a shell builtin")
			} else {
				found := false
				for _, val := range Paths {
					exe := filepath.Join(val, com)
					file, err := os.Stat(exe)
					if err == nil && !file.IsDir() {
						fmt.Println(com + " is " + exe)
						found = true
						break
					}
				}
				if !found {
					fmt.Println(com + ": not found")
				}
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(dir)
		case "cd":
			err := os.Chdir(splitcomm[1])
			if err != nil {
				fmt.Printf("%s: %s: No such file or directory\n", command, splitcomm[1])
			}
		default:
			cmd := exec.Command(command, splitcomm[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()

			if err != nil {
				fmt.Printf("%s: command not found\n", splitcomm[0])
			}
		}
	}
}
