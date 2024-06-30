// Package cmd is a root command.
/*
Copyright © 2024 Takafumi Miyanaga <orangekame3.dev@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

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
	data, err := os.ReadFile(filepath)
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
			Command:     name,
		})
	}

	return commands, nil
}
