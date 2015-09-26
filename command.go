package main

// Command basic structure
type Command struct {
	Parameters  []Parameter       `json:"parameters"`
	Values      map[string]string `json:"-"`
	Title       string            `json:"title"`
	Command     string            `json:"command"`
	Route       string            `json:"route"`
	Result      interface{}       `json:"-"`
	Description string            `json:"description"`
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
