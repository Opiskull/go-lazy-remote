package main

import (
	"log"

	"github.com/opiskull/go-jsonapi"
	"github.com/zenazn/goji/web"
)

// CommandService is used to register commands
type CommandService struct {
	executor *CommandExecutor
	commands []*Command
	Mux      *web.Mux
}

// NewCommandService creates a new instance
func NewCommandService(commands []*Command) *CommandService {
	var service = &CommandService{
		executor: &CommandExecutor{},
		commands: commands,
		Mux:      jsonapi.NewJSONSubRouterMux(),
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
	c.Mux.Get("/commands/info", routes.ListCommandInfos(c.commands))
	c.Mux.Get("/commands", routes.ListCommands(c.commands))
}

// RegisterCommand registers handlers for commands
func (c CommandService) registerCommand(command *Command) {
	log.Printf("Register Route %v", "/api"+command.Route)
	c.Mux.Handle(command.Route, routes.ExecuteCommand(command, c.executor))
	c.Mux.Get(command.Route+"/info", routes.CommandInfo(command))
}
