package main

import (
	"fmt"

	"github.com/labstack/echo"
)

// Routes is used to store all routes
type Routes struct{}

var routes = Routes{}

func setParameters(command *Command, c echo.Context) error {
	command.Values = make(map[string]string)
	for _, param := range command.Parameters {
		var value = c.Param(param.Name)
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
func (r Routes) CommandInfo(command *Command) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(200, &command)
	}
}

// ExecuteCommand the command with the commandexecutor
func (r Routes) ExecuteCommand(command *Command, executor *CommandExecutor) func(echo.Context) error {
	return func(c echo.Context) error {
		var com = *command
		setParameters(&com, c)
		if err := executor.Execute(&com); err != nil {
			return err
		}
		return c.JSON(200, &com.Result)
	}
}

// ListCommandInfos lists all commands
func (r Routes) ListCommandInfos(commands []*Command) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(200, commands)
	}
}

// ListCommands is for listing all loaded commands
func (r Routes) ListCommands(commands []*Command) func(echo.Context) error {
	return func(c echo.Context) error {
		var list = []string{}
		for _, command := range commands {
			list = append(list, command.Route)
		}
		return c.JSON(200, list)
	}
}
