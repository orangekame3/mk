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
	"io"
	"log" // Add this line to import the "net/url" package
	"net/http"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var inputFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mk",
	Short: "mk is a CLI tool for executing make commands interactively.",
	Long:  `mk is a CLI tool for executing make commands interactively.`,
	Run: func(cmd *cobra.Command, args []string) {
		tempFile := ""
		if inputFile == "" {
			inputFile = "Makefile"
		} else if _, err := os.Stat(inputFile); os.IsNotExist(err) {
			// Download the file from URL
			tempFile = "mk_temp_makefile"
			if err := downloadFile(inputFile, tempFile); err != nil {
				log.Fatalf("Failed to download Makefile: %v", err)
			}
			inputFile = tempFile
			defer func() {
				if err := os.Remove(tempFile); err != nil {
					log.Printf("Failed to remove temporary file: %v", err)
				}
			}()
		}

		absPath, err := filepath.Abs(inputFile)
		if err != nil {
			log.Fatalf("Failed to get absolute path for %s: %v", inputFile, err)
		}

		inputFileDir := filepath.Dir(inputFile)
		commands, err := parseMakefile(absPath)
		if err != nil {
			log.Fatalf("Failed to parse Makefile: %v", err)
		}

		if inputFileDir == "" {
			inputFileDir, err = os.Getwd()
			if err != nil {
				log.Fatalf("Failed to get current directory: %v", err)
			}
		}

		p := tea.NewProgram(initialModel(commands, inputFileDir))
		m, err := p.Run()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		if m, ok := m.(model); ok && m.selected != "" {
			fmt.Printf("Running command: make -f %s %s in directory: %s\n", inputFile, m.selected, filepath.Dir(inputFile))
			if err := runCommand(filepath.Dir(inputFile), "make", "-f", inputFile, m.selected); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to run command: %v\n", err)
			}
		}
	},
}

// Execute executes the interactive make command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Failed to execute: %v", err)
	}
}

// init initializes the root command.
func init() {
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Specify an input file other than Makefile (URL is also supported)")
}

// downloadFile downloads a file from the given URL and saves it to the specified path.
func downloadFile(url, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
