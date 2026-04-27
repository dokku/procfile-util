# procfile-util

A tool for parsing, validating, and exporting Heroku-style Procfiles.

## Installation

Install on Linux or macOS with the install script:

```bash
curl -fsSL https://raw.githubusercontent.com/dokku/procfile-util/master/install.sh | sh
```

Or via apt (after configuring the [dokku/dokku PackageCloud repository](https://packagecloud.io/dokku/dokku)):

```bash
sudo apt-get install procfile-util
```

Or via `go install`:

```bash
go install github.com/dokku/procfile-util@latest
```

See the [Getting started](docs/getting-started.md#installation) guide for binary downloads and building from source.

## Usage

Validate a Procfile in the current directory:

```bash
procfile-util check
```

List every process type:

```bash
procfile-util list
```

Show the command for the `web` process with `$PORT` substituted:

```bash
procfile-util show --process-type web --default-port 3000
```

See the [tasks reference](docs/tasks/README.md) for every subcommand and flag.

## Documentation

- [Getting started](docs/getting-started.md) - what a Procfile is, why use procfile-util, installation, and your first Procfile
- [Procfile format](docs/procfile-format.md) - syntax, comments, and strict mode
- [Tasks](docs/tasks/README.md) - one page per subcommand
- [Variable expansion](docs/variable-expansion.md) - how `show` and `expand` resolve `$PORT`, `--env-file`, and `--allow-getenv`
- [Process managers](docs/process-managers.md) - turning a Procfile into systemd, runit, launchd, sysv, or upstart units
- [Dokku integration](docs/dokku-integration.md) - how procfile-util fits into a Dokku deploy
- [CI usage](docs/ci-usage.md) - linting Procfiles and gating CI steps on process types

## License

[MIT](LICENSE)
