# go-procfile-util [![CircleCI](https://circleci.com/gh/josegonzalez/go-procfile-util.svg?style=svg)](https://circleci.com/gh/josegonzalez/go-procfile-util)

A tool for interacting with Procfiles.

## Installation

Install it using the "go get" command:

    go get github.com/josegonzalez/go-procfile-util

## What is a Procfile

A Procfile is a file that was [promoted by Heroku](https://blog.heroku.com/the_new_heroku_1_process_model_procfile) for their platform as an easy way to specify one or more distinct processes to run within Heroku. This format has since been picked up by various other tools and platforms.

The `procfile-util` tool expects a Procfile to be defined as one or more lines containing one of:

- a comment (preceeded by a `#` symbol)
- a process-type/command combination (with optional trailing whitespace or trailing comment)
  - when there is a trailing comment, the `#` symbol _must_ be preceeded by one or more `whitespace` characters.
- a blank line (with optional trailing whitespace)

Comments and blank lines are ignored, while process-type/command combinations look like the following:

```
<process type>: <command>
```

The syntax is defined as follows:

- `<process type>` – a valid DNS Label Name as per [RFC 1123](https://tools.ietf.org/html/rfc1123), a process type is a name for your command, such as `web`, `worker`, `urgentworker`, `clock`, etc.
- `<command>` – a command used to launch the process, such as `rake jobs:work`

This syntax differs common interpretations of validd `<process type>` values in that we define the process type as a DNS Label name, versus the regex `[A-Za-z0-9_]+`. The reason for this is that processes defined within Procfiles are commonly used in DNS entries. Rather than have a second level of platform-specific validation in place, this project implicitly defines the format for the process-type.

Given the above, a valid process type can be generalized to the following rules (as taken from the [Kubernetes documentation](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-label-names)):

- contain at most 63 characters
- contain only lowercase alphanumeric characters or '-'
- start with an alphanumeric character
- end with an alphanumeric character

Additionally, should a Procfile contain a line that does not match one of the three patterns mentioned above, the entire Procfile is considered invalid, and will not be processed. This is to avoid issues where a Procfile may contain merge conflicts or other improper content, thus resulting in unwanted runtime behavior for applications.

Finally, process types within a Procfile may not overlap and must be unique. Rather than assuming that the last or first specified is correct, `procfile-util` will fail to parse the Procfile with the relevant error.

## Usage

All commands take a `-P` or `--procfile` flag to specify an alternative `Procfile` path. If not specified, `procfile-util` will attempt to read the `Procfile` from the current directory.

### check

> check that the specified procfile is valid

```shell
procfile-util check
```

### delete

> delete a process type from a procfile

This command does not retain comments or extra newline characters. Specifying both the `write-path` and `stdout` flags will result in an error.

```shell
# delete the web process and write the file
procfile-util delete --process web

# delete the web process and write output to other.Procfile
procfile-util delete --process web --write-path other.Procfile

# delete the web process and write output to stdout
procfile-util delete --process web --stdout
```

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

# specify a default-port to use (default: 5000)
procfile-util expand --default-port 3000

# expand variables with getenv used for variable expansion
# may use any variable available when executing procfile-util
procfile-util expand --allow-getenv

# expand variables with a parsed .env file used for variable expansion
procfile-util expand --env-file .env

# combines getenv and .env file parsing to provide variable expansion
procfile-util expand --allow-getenv --env-file .env

# specify the default-port when performing variable expansion
procfile-util expand --allow-getenv --env-file .env --default-port 3000
```

### export

> export the application to another process management format

Due to argument parsing limitations, the `--location` flag is currently required.

In addition, not all formats support all arguments, and not all arguments have examples below.

```shell
# export systemd init files to the `tmp` directory
# support formats include: [launchd, runit, systemd, systemd-user, sysv, upstart]
# the default format is: systemd
procfile-util export --format systemd --location tmp

# override the app name
procfile-util export --location tmp --app node-js-app

# set the group and user used to launch processes
procfile-util export --location tmp --group root --user root

# set a working directory path for the process
procfile-util export --location tmp --working-directory /root
```

### list

> list all process types in a procfile

```shell
procfile-util list
```

### set

> set the command for a process type in a procfile

This command does not retain comments or extra newline characters. Specifying both the `write-path` and `stdout` flags will result in an error.

```shell
# set the web process and write the file
procfile-util set --process web --command "python app.py -p $PORT"

# set the web process and write output to other.Procfile
procfile-util set --process web --command "python app.py -p $PORT" --write-path other.Procfile

# set the web process and write output to stdout
procfile-util set --process web --command "python app.py -p $PORT" --stdout
```

### show

> show the command for a specific process type

```shell
# shows the command for the web process
procfile-util show -p web

# shows the command for the web process
# specify a default-port to use (default: 5000)
procfile-util show -p web --default-port 3000

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

# shows the command for the web process
# specify the default-port when performing variable expansion
procfile-util show web --allow-getenv --env-file .env --default-port 3000
```
