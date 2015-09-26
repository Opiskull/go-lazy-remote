package main

import (
	"fmt"
	"net/http"
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

// ShowInfo displays the info for the command
func (c Routes) ShowInfo(command *Command) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(writeJSON(command))
	}
}

// Execute the command with the commandexecutor
func (c Routes) Execute(command *Command, executor *CommandExecutor) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var com = *command
		setParameters(&com, r)
		if err := executor.Execute(&com); err != nil {
			w.Write(writeJSONError(err))
		} else {
			w.Write(writeJSON(&com.Result))
		}
	}
}

// ListCommands lists all commands
func (c Routes) ListCommands(commands []*Command) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(writeJSON(commands))
	}
}
