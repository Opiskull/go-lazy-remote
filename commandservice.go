package main

import (
	"log"

	"github.com/labstack/echo"
)

// CommandService is used to register commands
type CommandService struct {
	executor *CommandExecutor
	commands []*Command
	group    *echo.Group
}

// NewCommandService creates a new instance
func NewCommandService(commands []*Command, group *echo.Group) *CommandService {
	var service = &CommandService{
		executor: &CommandExecutor{},
		commands: commands,
		group:    group,
	}
	service.registerCommands()
	return service
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
