package main

import (
	"fmt"
	"net/http"

	"github.com/opiskull/go-jsonapi"
)

// Routes is used to store all routes
type Routes struct{}

var routes = Routes{}

func setParameters(command *Command, request *http.Request) error {
	command.Values = make(map[string]string)
	for _, param := range command.Parameters {
		var value = request.URL.Query().Get(param.Name)
		if len(value) == 0 {
			value = param.Value
		}
		if param.Mandatory && len(value) == 0 {
			return fmt.Errorf("Parameter '%v' is missing", param.Name)
		}
		command.Values[param.Name] = value
	}
	return nil
}

// CommandInfo displays the info for the command
func (c Routes) CommandInfo(command *Command) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonapi.JSON.Write(w, command)
	}
}

// ExecuteCommand the command with the commandexecutor
func (c Routes) ExecuteCommand(command *Command, executor *CommandExecutor) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var com = *command
		setParameters(&com, r)
		if err := executor.Execute(&com); err != nil {
			jsonapi.JSON.Error(w, jsonapi.NewAppError(err))
		} else {
			jsonapi.JSON.Write(w, &com.Result)
		}
	}
}

// ListCommandInfos lists all commands
func (c Routes) ListCommandInfos(commands []*Command) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonapi.JSON.Write(w, commands)
	}
}

// ListCommands is for listing all loaded commands
func (c Routes) ListCommands(commands []*Command) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var list = []string{}
		for _, command := range commands {
			list = append(list, command.Route)
		}
		jsonapi.JSON.Write(w, list)
	}
}
