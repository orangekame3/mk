package cmd

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Task represents a single task in the Taskfile.
type Task struct {
	Desc string   `yaml:"desc"`
	Cmds []string `yaml:"cmds"`
}

// Taskfile represents the structure of the Taskfile.yml.
type Taskfile struct {
	Version string          `yaml:"version"`
	Tasks   map[string]Task `yaml:"tasks"`
}

// parseTaskfile parses the Taskfile and returns a list of MakeCommand.
func parseTaskfile(filepath string) ([]MakeCommand, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read Taskfile: %v", err)
	}

	var taskfile Taskfile
	err = yaml.Unmarshal(data, &taskfile)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Taskfile: %v", err)
	}

	var commands []MakeCommand
	for name, task := range taskfile.Tasks {
		commands = append(commands, MakeCommand{
			Name:        name,
			Description: task.Desc,
			Command:     name, // We use the task name as the command to be executed.
		})
	}

	return commands, nil
}
