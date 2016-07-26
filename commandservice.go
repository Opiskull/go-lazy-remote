package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/labstack/echo"
)

// CommandService is used to register commands
type CommandService struct {
	executor *CommandExecutor
	commands []*Command
	group    *echo.Group
}

// NewCommandService creates a new instance
func NewCommandService(commandsDirectory string, group *echo.Group) *CommandService {
	var service = &CommandService{
		executor: &CommandExecutor{},
		group:    group,
	}
	service.commands = LoadCommands(commandsDirectory)
	service.registerCommands()
	return service
}

// LoadCommands creates an array of commands
func LoadCommands(commandsDirectory string) []*Command {
	var commands = []*Command{}
	files, _ := filepath.Glob(filepath.Join(commandsDirectory, "*.json"))
	for _, f := range files {
		var command = &Command{}
		file, e := ioutil.ReadFile(f)
		if e != nil {
			panic(e)
		}
		err := json.NewDecoder(bytes.NewReader(file)).Decode(&command)
		if err != nil {
			panic(err)
		}
		commands = append(commands, command)
	}
	return commands
}

// RegisterCommands registers all provided commands
func (c CommandService) registerCommands() {
	log.Println("Register Commands")
	for _, command := range c.commands {
		c.registerCommand(command)
	}
	log.Println("List all commands with '/api/commands'")
	c.group.Get("/commands/info", routes.ListCommandInfos(c.commands))
	c.group.Get("/commands", routes.ListCommands(c.commands))
}

// RegisterCommand registers handlers for commands
func (c CommandService) registerCommand(command *Command) {
	log.Printf("Register Route %v", "/api"+command.Route)
	c.group.Post(command.Route, routes.ExecuteCommand(command, c.executor))
	c.group.Get(command.Route+"/info", routes.CommandInfo(command))
}
