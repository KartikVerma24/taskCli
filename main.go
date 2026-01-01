package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/KartikVerma24/taskCli/cli"
	inmemory "github.com/KartikVerma24/taskCli/database/inMemory"
	"github.com/KartikVerma24/taskCli/service"
)

func main() {
	fmt.Println("Welcome to the taskCLI")
	fmt.Println("Type 'help' to see commands, 'exit' to quit")

	taskRepo := inmemory.NewTaskInMemRepo()
	taskSvc := service.NewTaskService(taskRepo)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("task> ")

		if !scanner.Scan() { // for ctrl + c
			fmt.Println("\nGoodbye!")
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if line == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		args := strings.Fields(line)
		cliErr := cli.RunCommands(args, *taskSvc)
		if cliErr != nil {
			fmt.Println("Error :", cliErr)
		}
	}
}
