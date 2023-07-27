# depctl

[![build](https://github.com/twelvelabs/depctl/actions/workflows/build.yml/badge.svg)](https://github.com/twelvelabs/depctl/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/twelvelabs/depctl/branch/main/graph/badge.svg)](https://codecov.io/gh/twelvelabs/depctl)

Install your project's development dependencies with a single command ✨

## Installation

Choose one of the following:

- Download and manually install the latest [release](https://github.com/twelvelabs/depctl/releases/latest)
- Install with [Homebrew](https://brew.sh/) 🍺

  ```bash
  brew install twelvelabs/tap/depctl
  ```

- Install from source

  ```bash
  go install github.com/twelvelabs/depctl@latest
  ```

## Usage

Create a [`.dependencies.yml`](.dependencies.yml) file
in the current working directory, then run:

```bash
# Installs the "default" group of dependencies.
# For each entry:
# - Looks for the named executable in $PATH.
# - If not found, runs the `install` command.
depctl up

# Installs the "foo" group of dependencies.
depctl up foo

# View a markdown list of all your dependencies.
# Can be copied/piped to a README file for documentation.
depctl list
```

## Development

```bash
# Ensures all required dependencies are installed
# and bootstraps the project for local development.
make setup

# Run tests.
make test

# Run the app.
make run

# Show help.
make
```
