package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	finished := false
	Path := os.Getenv("PATH")
	Paths := strings.Split(Path, ";")
	// Wait for user input
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
		if command == "exit" {
			finished = true
			code, _ := strconv.ParseInt(splitcomm[1], 10, 64)
			os.Exit(int(code))
		} else if command == "echo" {
			fmt.Println(strings.Join(splitcomm[1:], " "))
		} else if command == "type" {
			com := (splitcomm[1])
			if com == "exit" || com == "echo" || com == "type"{
				fmt.Println(com + " is a shell builtin")
			}else{
				for _ , val := range Paths{
					files ,_ := ioutil.ReadDir(val)
					for _ , f := range files{
						if com == f.Name(){
							fmt.Println(com + "is" + val + com)
						}else{
							fmt.Println(com + ": not found")
						}
					}

				}
				fmt.Println(com + ": not found")
			}
		} else {
			fmt.Println(command + ": command not found")
		}
	}
}
