# mk - Interactive Task Runner for Makefiles

**mk** is a command-line interface (CLI) tool designed to interactively execute `make` commands from a Makefile. It provides a user-friendly interface to select and run predefined make commands, making it easier to manage and execute build tasks.

## Features

**Interactive Interface**: Browse and select from available `make` commands using arrow keys or filter by typing.

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

Run `mk` in your terminal to start the interactive interface. Select a command using arrow keys or filter by typing part of the command name. Press Enter to execute the selected command.

```bash
mk
```

This will display a list of available commands from the Makefile. Use the arrow keys to navigate and select a command, then press `Enter` to execute it.

Contact
For questions or feedback, please contact at @orangekame3.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

## Acknowledgements

This project was inspired by [fzf-make](https://github.com/kyu08/fzf-make) by @kyu08, [make2help](https://github.com/Songmu/make2help) @songmu, and [glow](https://github.com/charmbracelet/glow) Thank you for the inspiration!
