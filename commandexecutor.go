package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"text/template"
)

// CommandExecutor is used to execute the commands
type CommandExecutor struct {
}

// StringResult is used to show the result of a string cmd
type StringResult struct {
	Result string `json:"result"`
}

// Execute the command
func (j *CommandExecutor) Execute(command *Command) error {
	cmd, err := j.buildCommand(command)
	if err != nil {
		return err
	}
	log.Printf("executing command: '%v'", cmd.Args)
	result, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error in command: '%v'", string(result))
	}
	return j.buildResult(command, result)
}

// applyTemplateValues is used to replace all placeholder with values
func (j *CommandExecutor) applyTemplateValues(com *Command) (string, error) {
	t := template.Must(template.New(com.Title).Parse(com.Command))
	buf := new(bytes.Buffer)
	err := t.Execute(buf, com.Values)
	return buf.String(), err
}

// buildCommandText is used to build the commandText
func (j *CommandExecutor) buildCommandText(com *Command) (string, error) {
	text, err := j.applyTemplateValues(com)
	if err != nil {
		return "", err
	}
	var comOut = strings.ToLower(com.Output)
	var comType = strings.ToLower(com.Type)
	if comType == "ps" && comOut == "json" {
		return text + " | ConvertTo-Json", nil
	}
	if comType == "cmd" {
		return " /C " + text, nil
	}
	return text, nil
}

// getCommandInterpreter is used to get the right interpreter for the command
func (j *CommandExecutor) getCommandInterpreter(com *Command) (string, error) {
	var comType = strings.ToLower(com.Type)
	if comType == "ps" || comType == "powershell" {
		return "powershell.exe", nil
	}
	if comType == "cmd" {
		return "cmd.exe", nil
	}
	return "", fmt.Errorf("the type '%v' is not valid", com.Type)
}

// buildCommand is used to build a exec.Cmd
func (j *CommandExecutor) buildCommand(com *Command) (*exec.Cmd, error) {
	t, err := j.buildCommandText(com)
	if err != nil {
		return nil, err
	}
	in, err := j.getCommandInterpreter(com)
	if err != nil {
		return nil, err
	}
	return exec.Command(in, t), nil
}

// buildResult is used to build a json result
func (j *CommandExecutor) buildResult(com *Command, output []byte) error {
	var comOut = strings.ToLower(com.Output)
	if comOut == "json" {
		if err := json.Unmarshal(output, &com.Result); err != nil {
			return fmt.Errorf("error in command-result: '%v'", string(output))
		}
		return nil
	}
	if comOut == "string" {
		com.Result = StringResult{string(output)}
		return nil
	}
	return fmt.Errorf("no valid output '%v'", com.Output)
}
