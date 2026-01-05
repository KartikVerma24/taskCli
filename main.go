package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/shlex"

	"github.com/KartikVerma24/taskCli/cli"
	jsonfs "github.com/KartikVerma24/taskCli/database/jsonFS"
	"github.com/KartikVerma24/taskCli/service"
)

func main() {
	fmt.Println("Welcome to the taskCLI")
	fmt.Println("Type 'help' to see commands, 'exit' to quit")

	fPath := `D:\Coding\test.json`

	// taskRepo := inmemory.NewTaskInMemRepo()
	jsonRepo, err := jsonfs.NewTaskJsonFSRepo(fPath)
	if err != nil {
		fmt.Println("error in loading file : ", err)
	}
	taskSvc := service.NewTaskService(jsonRepo)

	scanner := bufio.NewScanner(os.Stdin) // this creates an object which will read whatever is input by user

	for {
		fmt.Print("task> ")

		if !scanner.Scan() { // for ctrl + c
			fmt.Println("\nGoodbye!")
			break
		}

		line := strings.TrimSpace(scanner.Text())
		// What scanner.Text() gives you => Exactly what the user typed before pressing Enter.
		if line == "" { // in case user just enter without any command
			continue
		}

		if line == "exit" { // in case of "exit" break the flow to end the program
			fmt.Println("Goodbye!")
			break
		}

		if line == "clear" {
			fmt.Println("\033[H\033[2J")
			continue
		}

		args, inputParseErr := shlex.Split(line)
		if inputParseErr != nil {
			fmt.Println("Error :", inputParseErr)
			continue
		}
		cliErr := cli.RunCommands(args, *taskSvc)
		if cliErr != nil {
			fmt.Println("Error :", cliErr)
		}
	}
}
