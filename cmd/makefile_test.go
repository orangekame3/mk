package cmd

import (
	"os"
	"testing"
)

func TestParseMakefile(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected []MakeCommand
	}{
		{
			name: "Valid Makefile",
			content: `build: # Build the project
test: # Run tests
deploy: # Deploy the project`,
			expected: []MakeCommand{
				{Name: "build", Description: "Build the project", Command: "build: # Build the project"},
				{Name: "test", Description: "Run tests", Command: "test: # Run tests"},
				{Name: "deploy", Description: "Deploy the project", Command: "deploy: # Deploy the project"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpfile, err := os.CreateTemp("", "example.Makefile")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write([]byte(tt.content)); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			commands, err := parseMakefile(tmpfile.Name())
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if len(commands) != len(tt.expected) {
				t.Fatalf("Expected %d commands, got %d", len(tt.expected), len(commands))
			}

			for i, cmd := range commands {
				if cmd != tt.expected[i] {
					t.Errorf("Expected %v, got %v", tt.expected[i], cmd)
				}
			}
		})
	}
}
