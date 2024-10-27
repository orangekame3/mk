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
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	height = 20
	width  = 60
)

// item is a list item that represents a make command.
type item struct {
	title       string
	description string
	command     string
}

// Title returns the item's title.
func (i item) Title() string { return i.title }

// Description returns the item's description.
func (i item) Description() string { return i.description }

// FilterValue returns the item's title.
func (i item) FilterValue() string { return i.title }

// Model is the bubbletea model for the interactive make command.
type model struct {
	list     list.Model
	commands []MakeCommand
	selected string
}

// initialModel returns the initial model.
func initialModel(commands []MakeCommand, mode string) model {
	items := make([]list.Item, len(commands))
	for i, cmd := range commands {
		items[i] = item{title: cmd.Name, description: cmd.Description, command: cmd.Command}
	}

	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = mode

	return model{list: l, commands: commands}
}

// Init initializes the model.
func (m model) Init() tea.Cmd {
	return nil
}

// Update updates the model.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit
		case "enter":
			selectedItem := m.list.SelectedItem().(item)
			m.selected = selectedItem.title
			return m, tea.Quit
		}
	}

	return m, cmd
}

// View renders the model.
func (m model) View() string {
	return m.list.View()
}

// runCommand runs a command and returns its output.
func runCommand(dir, name string, args ...string) error {
	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %v", err)
	}
	defer func() {
		if err := os.Chdir(originalDir); err != nil {
			log.Printf("failed to change back to original directory: %v", err)
		}
	}()

	// Change to the specified directory
	if err := os.Chdir(originalDir); err != nil {
		return fmt.Errorf("failed to change directory to %s: %v", dir, err)
	}
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}
