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
	"bufio"
	"os"
	"regexp"
	"strings"
)

// MakeCommand represents a make command.
type MakeCommand struct {
	Name        string
	Description string
	Command     string
}

// parseMakefile parses a Makefile and returns a slice of MakeCommand.
func parseMakefile(filename string) ([]MakeCommand, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var commands []MakeCommand
	scanner := bufio.NewScanner(file)

	descPosition := os.Getenv("MK_DESC_POSITION")
	var cmdRegexp *regexp.Regexp
	if descPosition == "side" {
		cmdRegexp = regexp.MustCompile(`^([a-zA-Z0-9\-_]+):.*?##\s*(.*)$`)
	} else {
		cmdRegexp = regexp.MustCompile(`^##\s*(.*)\n([a-zA-Z0-9\-_]+):.*$`)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if descPosition == "side" {
			matches := cmdRegexp.FindStringSubmatch(line)
			if len(matches) == 3 {
				command := MakeCommand{
					Name:        matches[1],
					Description: matches[2],
					Command:     line,
				}
				commands = append(commands, command)
			}
		} else {
			if strings.HasPrefix(line, "##") {
				description := strings.TrimSpace(strings.TrimPrefix(line, "##"))
				if scanner.Scan() {
					line = scanner.Text()
					matches := cmdRegexp.FindStringSubmatch("## " + description + "\n" + line)
					if len(matches) == 3 {
						command := MakeCommand{
							Name:        matches[2],
							Description: matches[1],
							Command:     line,
						}
						commands = append(commands, command)
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return commands, nil
}
