package cmd

import (
	"os"
	"testing"
)

func createTempMakefile(content string) (string, error) {
	file, err := os.CreateTemp("", "Makefile")
	if err != nil {
		return "", err
	}

	if _, err := file.WriteString(content); err != nil {
		file.Close()
		return "", err
	}

	if err := file.Close(); err != nil {
		return "", err
	}

	return file.Name(), nil
}

func TestParseMakefile(t *testing.T) {
	t.Run("side description", func(t *testing.T) {
		os.Setenv("MK_DESC_POSITION", "side")
		defer os.Unsetenv("MK_DESC_POSITION")

		content := `build: ## Build the project
	echo "Building the project"
clean: ## Clean the project
	echo "Cleaning the project"
		`
		filename, err := createTempMakefile(content)
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(filename)

		commands, err := parseMakefile(filename)
		if err != nil {
			t.Fatalf("parseMakefile failed: %v", err)
		}

		expected := []MakeCommand{
			{Name: "build", Description: "Build the project", Command: "build: ## Build the project"},
			{Name: "clean", Description: "Clean the project", Command: "clean: ## Clean the project"},
		}

		if len(commands) != len(expected) {
			t.Fatalf("Expected %d commands, got %d", len(expected), len(commands))
		}

		for i, cmd := range commands {
			if cmd.Name != expected[i].Name || cmd.Description != expected[i].Description || cmd.Command != expected[i].Command {
				t.Errorf("Expected %v, got %v", expected[i], cmd)
			}
		}
	})

	t.Run("above description", func(t *testing.T) {
		os.Unsetenv("MK_DESC_POSITION")

		content := `## Build the project
build:
	echo "Building the project"
## Clean the project
clean:
	echo "Cleaning the project"
		`
		filename, err := createTempMakefile(content)
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(filename)

		commands, err := parseMakefile(filename)
		if err != nil {
			t.Fatalf("parseMakefile failed: %v", err)
		}

		expected := []MakeCommand{
			{Name: "build", Description: "Build the project", Command: "build:"},
			{Name: "clean", Description: "Clean the project", Command: "clean:"},
		}

		if len(commands) != len(expected) {
			t.Fatalf("Expected %d commands, got %d", len(expected), len(commands))
		}

		for i, cmd := range commands {
			if cmd.Name != expected[i].Name || cmd.Description != expected[i].Description || cmd.Command != expected[i].Command {
				t.Errorf("Expected %v, got %v", expected[i], cmd)
			}
		}
	})
}
