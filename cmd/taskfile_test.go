package cmd

import (
	"os"
	"testing"
)

func TestParseTaskfile(t *testing.T) {
	// Setup: create a temporary Taskfile
	testData := `
version: '3'

tasks:
  lint:
    desc: Lint the code
    cmds:
      - npx mega-linter-runner --flavor go

  fmt:
    desc: Format the code
    cmds:
      - go fmt ./...

  test:
    desc: Run the tests
    cmds:
      - go clean -testcache
      - go test -v ./...
      - echo -e "\033[32mOK\033[0m" || echo -e "\033[31mERROR\033[0m"
`
	tempFile, err := os.CreateTemp("", "Taskfile.yml")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.Write([]byte(testData))
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()

	// Execute the function under test
	commands, err := parseTaskfile(tempFile.Name())
	if err != nil {
		t.Fatalf("parseTaskfile failed: %v", err)
	}

	// Verify the results
	expectedCommands := []MakeCommand{
		{Name: "lint", Description: "Lint the code", Command: "lint"},
		{Name: "fmt", Description: "Format the code", Command: "fmt"},
		{Name: "test", Description: "Run the tests", Command: "test"},
	}

	if len(commands) != len(expectedCommands) {
		t.Fatalf("Expected %d commands, but got %d", len(expectedCommands), len(commands))
	}

	for i, cmd := range commands {
		if cmd.Name != expectedCommands[i].Name || cmd.Description != expectedCommands[i].Description || cmd.Command != expectedCommands[i].Command {
			t.Errorf("Command %d: expected %+v, but got %+v", i, expectedCommands[i], cmd)
		}
	}
}
