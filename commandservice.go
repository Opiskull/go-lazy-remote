package main

import (
	"log"

	"github.com/zenazn/goji"
)

// CommandService is used to register commands
type CommandService struct {
	executor *CommandExecutor
	commands []*Command
}

// NewCommandService creates a new instance
func NewCommandService(commands []*Command) *CommandService {
	var service = &CommandService{
		executor: &CommandExecutor{},
		commands: commands,
	}
	return service
}

// RegisterCommands registers all provided commands
func (c CommandService) registerCommands() {
	log.Println("Register Commands")
	for _, command := range c.commands {
		c.registerCommand(command)
	}
	log.Println("List all commands with '/api/commands'")
	goji.Get("/api/commands/info", routes.ListCommandInfos(c.commands))
	goji.Get("/api/commands", routes.ListCommands(c.commands))
}

// RegisterCommand registers handlers for commands
func (c CommandService) registerCommand(command *Command) {
	var route = "/api" + command.Route
	log.Printf("Register Route %v", route)
	goji.Handle(route, routes.ExecuteCommand(command, c.executor))
	goji.Get(route+"/info", routes.CommandInfo(command))
}

// Init and register all commands
func (c CommandService) Init() {
	c.registerCommands()
}
