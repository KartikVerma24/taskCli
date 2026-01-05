package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/google/shlex"

	"github.com/KartikVerma24/taskCli/cli"
	jsonfs "github.com/KartikVerma24/taskCli/database/jsonFS"
	"github.com/KartikVerma24/taskCli/service"
	"github.com/KartikVerma24/taskCli/utils"
)

var storeDir string
var version = "taskCli v1.0.0"

func main() {
	var showVersion bool

	flag.StringVar(&storeDir, "store", "", "directory to store task data")
	flag.BoolVar(&showVersion, "version", false, "to show version of the app")
	flag.Parse()
	
	if showVersion {
		fmt.Println(version)
		return
	}
	
	storePath, err := utils.ResolveStorePath(storeDir)
	if err != nil {
		fmt.Println("Error with store path :", err)
		return
	}

	fmt.Println("Welcome to the taskCLI")
	fmt.Println("Type 'help' to see commands, 'exit' to quit")

	// taskRepo := inmemory.NewTaskInMemRepo()
	jsonRepo, err := jsonfs.NewTaskJsonFSRepo(storePath)
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
