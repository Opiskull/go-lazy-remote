package main

// Command basic structure
type Command struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Route       string            `json:"route"`
	Parameters  []Parameter       `json:"parameters,omitempty"`
	Command     string            `json:"command"`
	Type        string            `json:"type"`
	Output      string            `json:"output"`
	Values      map[string]string `json:"-"`
	Result      interface{}       `json:"-"`
}

// Parameter is used to describe a parameter
type Parameter struct {
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description"`
	Mandatory   bool   `json:"mandatory"`
}

// NewCommand creates a new Instance of a command
func NewCommand(title, command string) *Command {
	return &Command{
		Title:   title,
		Command: command,
		Values:  make(map[string]string),
	}
}
