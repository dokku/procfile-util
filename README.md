# go-procfile-util

A tool for interacting with Procfiles.

## Installation

Install it using the "go get" command:

    go get github.com/josegonzalez/go-procfile-util

## Usage

All commands take a `-P` or `--procfile` flag to specify an alternative `Procfile` path. If not specified, `procfile-util` will attempt to read the `Procfile` from the current directory.

### exists

> check if a process type exists

```shell
# returns 0 if the web process type exists
procfile-util exists -p web

# returns 1 if the non-existent process type does not exist
procfile-util exists -p non-existent
```

### expand

> expands a procfile against a specific environment

```shell
# expand variables with no implicit env vars
# will result in empty string for variable replacements
procfile-util expand

# expand variables with getenv used for variable expansion
# may use any variable available when executing procfile-util
procfile-util expand --allow-getenv

# expand variables with a parsed .env file used for variable expansion
procfile-util expand --env-file .env

# combines getenv and .env file parsing to provide variable expansion
procfile-util expand --allow-getenv --env-file .env
```
### list

> list all process types in a procfile

```shell
procfile-util list
```

### show

> show the command for a specific process type

```shell
# shows the command for the web process
procfile-util show -p web

# shows the command for the web process
# expand variables with getenv used for variable expansion
# may use any variable available when executing procfile-util
procfile-util show -p web --allow-getenv

# shows the command for the web process
# expand variables with a parsed .env file used for variable expansion
procfile-util show -p web --env-file .env

# shows the command for the web process
# combines getenv and .env file parsing to provide variable expansion
procfile-util show -p web --allow-getenv --env-file .env
```
