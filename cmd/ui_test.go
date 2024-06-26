package cmd

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestInitialModel(t *testing.T) {
	tests := []struct {
		name     string
		commands []MakeCommand
	}{
		{
			name: "Single Command",
			commands: []MakeCommand{
				{Name: "build", Description: "Build the project", Command: "build: # Build the project"},
			},
		},
		{
			name: "Multiple Commands",
			commands: []MakeCommand{
				{Name: "build", Description: "Build the project", Command: "build: # Build the project"},
				{Name: "test", Description: "Run tests", Command: "test: # Run tests"},
				{Name: "deploy", Description: "Deploy the project", Command: "deploy: # Deploy the project"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := initialModel(tt.commands)

			if m.list.Title != "mk" {
				t.Errorf("Expected title 'mk', got '%s'", m.list.Title)
			}

			if len(m.list.Items()) != len(tt.commands) {
				t.Fatalf("Expected %d items, got %d", len(tt.commands), len(m.list.Items()))
			}

			for i, listItem := range m.list.Items() {
				item, ok := listItem.(item)
				if !ok {
					t.Fatalf("Expected item of type 'item', got '%T'", listItem)
				}
				cmd := tt.commands[i]
				if item.title != cmd.Name || item.description != cmd.Description {
					t.Errorf("Expected %v, got %v", cmd, item)
				}
			}
		})
	}
}

func TestRunCommand(t *testing.T) {
	tests := []struct {
		name     string
		cmdName  string
		cmdArgs  []string
		expected string
	}{
		{
			name:     "Successful Command",
			cmdName:  "echo",
			cmdArgs:  []string{"Hello, World!"},
			expected: "Hello, World!\n",
		},
		{
			name:     "Failing Command",
			cmdName:  "nonexistentcommand",
			cmdArgs:  []string{"arg"},
			expected: "", 
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			cmd := exec.Command(tt.cmdName, tt.cmdArgs...)
			cmd.Stdout = &out
			cmd.Stderr = &out

			err := cmd.Run()
			if tt.name == "Failing Command" && err == nil {
				t.Fatalf("Expected error for failing command, but got none")
			} else if tt.name != "Failing Command" && err != nil {
				t.Fatalf("Error running command: %v", err)
			}

			actual := out.String()
			if actual != tt.expected {
				t.Errorf("Expected output:\n%s\nBut got:\n%s", tt.expected, actual)
			}
		})
	}
}
