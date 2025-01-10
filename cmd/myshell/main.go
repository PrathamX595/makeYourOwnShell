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
	found := false
	Path := os.Getenv("PATH")
	cmdArr := [...]string{
		"exit",
		"echo",
		"type",
		"pwd",
		"cd",
	}
	Paths := strings.Split(Path, ":")

	for !finished {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		trim := strings.Trim(input, "\r\n")

		// Custom parser that preserves spaces inside quotes
		var splitcomm []string
		var tokenBuilder strings.Builder
		inQuotes := false
		for i := 0; i < len(trim); i++ {
			ch := trim[i]
			if ch == '\'' {
				inQuotes = !inQuotes
				continue
			}
			if ch == ' ' && !inQuotes {
				if tokenBuilder.Len() > 0 {
					splitcomm = append(splitcomm, tokenBuilder.String())
					tokenBuilder.Reset()
				}
				continue
			}
			tokenBuilder.WriteByte(ch)
		}
		if tokenBuilder.Len() > 0 {
			splitcomm = append(splitcomm, tokenBuilder.String())
		}

		if len(splitcomm) == 0 {
			continue
		}
		command := splitcomm[0]
		switch command {
		case "exit":
			finished = true
			code, _ := strconv.ParseInt(splitcomm[1], 10, 64)
			os.Exit(int(code))
		case "echo":
			args := splitcomm[1:]
			if len(args) == 0 {
				fmt.Fprintln(os.Stdout)
			} else {
				for i := 0; i < len(args)-1; i++ {
					fmt.Fprintf(os.Stdout, "%s ", args[i])
				}
				fmt.Fprintln(os.Stdout, args[len(args)-1])
			}
		case "type":
			com := splitcomm[1]
			for _, val := range cmdArr {
				if com == val {
					fmt.Println(com + " is a shell builtin")
					found = true
					break
				}
			}
			if !found {
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
			if len(splitcomm) < 2 {
				continue
			}
			if splitcomm[1] == "~" {
				err := os.Chdir(os.Getenv("HOME"))
				if err != nil {
					fmt.Printf("%s: %s: No such file or directory\n", command, splitcomm[1])
				}
			} else {
				err := os.Chdir(splitcomm[1])
				if err != nil {
					fmt.Printf("%s: %s: No such file or directory\n", command, splitcomm[1])
				}
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
		found = false
	}
}