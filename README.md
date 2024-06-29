# mk - Interactive Task Runner for Makefile

**mk** is a command-line interface (CLI) tool designed to interactively execute `make` commands from a Makefile. It provides a user-friendly interface to select and run predefined make commands, making it easier to manage and execute build tasks.

![mk](./img/demo.gif)

## Features

- **Interactive Interface**: Browse and select from available `make` commands using arrow keys or filter by typing.
- **Documentation**: View the description of each command to understand its purpose and usage.
- **Vim like keybindings**: Use `j` and `k` to navigate, `Enter` to execute, and `q` to quit.
- **Filtering**: Quickly search for commands by typing part of the command name. check `?` for help.
- **Remote Makefile**: Load a Makefile from a remote URL and execute commands.
- **Any Local Makefile**: Load a Makefile from any directory and execute commands.

## Installation

**mk** can be installed using the following methods:

### Homebrew

```bash
brew install orangekame3/tap/mk
```

### Go Install

```bash
go install github.com/orangekame3/mk@latest
```

### Manual Installation

1. Download the latest release from the [Releases](https://github.com/orangekame3/mk/releases)
2. Extract the archive and navigate to the extracted directory.
3. Move the `mk` binary to a directory in your `PATH`.

```bash
mv mk /usr/local/bin
```

## Usage

Prepare a Makefile with predefined commands and descriptions. Each command should be documented using a comment starting with `##` to provide a description of the command.

```makefile
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help

.PHONY: test fmt
## Run tests
test:
 go test ./...

## Format source code
fmt:
 go fmt ./...

```

or you can set `MK_DESC_POSITION=side` in your `.bashrc` or `.zshrc` to show the description on the right side of the command.

```makefile
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help

.PHONY: test fmt

test: ## Run tests
 go test ./...

fmt: ## Format source code
 go fmt ./...

```

Run `mk` in your terminal to start the interactive interface. Select a command using arrow keys or filter by typing part of the command name. Press Enter to execute the selected command.

```bash
mk
```

This will display a list of available commands from the Makefile. Use the arrow keys to navigate and select a command, then press `Enter` to execute it.

## Options

**mk** supports the following options:

```bash

mk is a CLI tool for executing make commands interactively.

Usage:
  mk [flags]

Flags:
  -h, --help           help for mk
  -i, --input string   Specify an input file other than Makefile (URL is also supported)
  -v, --version        version for mk
```

### Load Makefile from a Remote URL

Use the `-i` or `--input` flag to load a Makefile from a remote URL. This allows you to execute commands from a Makefile hosted on a remote server.

```bash
mk -i https://raw.githubusercontent.com/orangekame3/mk/main/Makefile
```

> [!NOTE]
> command executed at current directory.

### Load Makefile from a Local File

Use the `-i` or `--input` flag to load a Makefile from a local file. This allows you to execute commands from a Makefile located in any directory.

```bash
mk -i /path/to/Makefile
```

> [!NOTE]
> command executed at path/to directory, and return to the original directory after the command is executed.

Contact
For questions or feedback, please contact at @orangekame3.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

## Acknowledgements

This project was inspired by [fzf-make](https://github.com/kyu08/fzf-make) by [@kyu08](https://github.com/kyu08), [make2help](https://github.com/Songmu/make2help) by [@Songmu](https://github.com/Songmu/Songmu), and [glow](https://github.com/charmbracelet/glow) Thank you for the inspiration!

## Author

👤 [**orangekame3**](https://github.com/orangekame3)
