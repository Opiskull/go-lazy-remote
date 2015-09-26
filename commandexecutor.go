package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"text/template"
)

// CommandExecutor is used to execute the commands
type CommandExecutor struct {
}

func (j *CommandExecutor) buildCommandText(com *Command) (string, error) {
	t := template.Must(template.New(com.Title).Parse(com.Command))
	buf := new(bytes.Buffer)
	err := t.Execute(buf, com.Values)
	return buf.String(), err
}

func (j *CommandExecutor) parseJSON(command *Command, output []byte) error {
	if err := json.Unmarshal(output, &command.Result); err != nil {
		return fmt.Errorf("error in command-result: '%v'", string(output))
	}
	return nil
}

// Execute the command
func (j *CommandExecutor) Execute(command *Command) error {
	text, err := j.buildCommandText(command)
	if err != nil {
		return err
	}
	cmdText := text + " | ConvertTo-Json"
	log.Printf("Execute '%v'", cmdText)
	cmd := exec.Command("powershell.exe", cmdText)
	result, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error in command: '%v'", string(result))
	}
	if err := j.parseJSON(command, result); err != nil {
		return err
	}
	log.Println(command.Result)
	return nil
}
