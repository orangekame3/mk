// Package cmd is a root command.
package cmd

import (
	"testing"
)

func TestRootCmd(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		actual   string
	}{
		{name: "Use", expected: "mk", actual: rootCmd.Use},
		{name: "Short", expected: "mk is a CLI tool for executing make commands interactively.", actual: rootCmd.Short},
		{name: "Long", expected: "mk is a CLI tool for executing make commands interactively.", actual: rootCmd.Long},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected != tt.actual {
				t.Errorf("Expected '%s', got '%s'", tt.expected, tt.actual)
			}
		})
	}
}
