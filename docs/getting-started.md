# Getting started

`procfile-util` is a CLI for working with Procfiles. It parses, validates, modifies, and exports Procfiles to formats that traditional process supervisors like `systemd` and `launchd` can run.

## Why procfile-util?

A Procfile is a plain-text manifest that declares the long-running processes an application needs - a `web` server, a background `worker`, a `clock` for scheduled jobs, and so on. Heroku introduced the format, and Dokku, Foreman, and many similar tools adopted it.

If you already use Dokku, you have probably written a Procfile without thinking much about the parser behind it. `procfile-util` is that parser, exposed as a standalone tool. Reach for it when you need to:

- Validate a Procfile in CI before pushing to Dokku or Heroku.
- Read a single process command from a script (for example, the command behind `web`).
- Render a Procfile with `$PORT` and other variables substituted, so the output matches what a deploy target would actually run.
- Translate a Procfile into systemd, launchd, runit, sysv, or upstart unit files for hosts that do not run Dokku.

## Installation

### Quick install (Linux and macOS)

```bash
curl -fsSL https://raw.githubusercontent.com/dokku/procfile-util/master/install.sh | sh
```

The installer drops `procfile-util` into `/usr/local/bin`. Override the destination by setting `INSTALL_DIR`, and pin a release with `VERSION`:

```bash
VERSION=v0.20.4 INSTALL_DIR=$HOME/.local/bin curl -fsSL https://raw.githubusercontent.com/dokku/procfile-util/master/install.sh | sh
```

### Debian and Ubuntu

`procfile-util` is published to the `dokku/dokku` PackageCloud repository. Once that repository is configured for `apt`:

```bash
sudo apt-get update
sudo apt-get install procfile-util
```

### Binary download

Download a binary from the [GitHub releases page](https://github.com/dokku/procfile-util/releases) and copy it onto your `PATH`:

```bash
curl -fsSL -o procfile-util https://github.com/dokku/procfile-util/releases/latest/download/procfile-util-linux-amd64
chmod +x procfile-util
sudo mv procfile-util /usr/local/bin/procfile-util
```

### From source

With a recent Go toolchain installed:

```bash
go install github.com/dokku/procfile-util@latest
```

Or build the repository directly:

```bash
git clone https://github.com/dokku/procfile-util.git
cd procfile-util
make build/$(uname -s | tr '[:upper:]' '[:lower:]')/procfile-util-$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/')
```

## Your first Procfile

Create a file named `Procfile` in any directory:

```
web: bundle exec puma -p $PORT
worker: bundle exec rake jobs:work
```

Validate it:

```bash
procfile-util check
```

You should see `valid procfile detected web, worker`. Now list the process types:

```bash
procfile-util list
```

And inspect what `web` would actually run, with `$PORT` resolved to the default of `5000`:

```bash
procfile-util show --process-type web
```

If your application listens on port `3000` instead, override the default:

```bash
procfile-util show --process-type web --default-port 3000
```

## Global flags

Every subcommand accepts the same four global flags. They are not repeated in each task page below:

| Flag | Default | Purpose |
| ---- | ------- | ------- |
| `-P, --procfile` | `Procfile` | Path to the Procfile to read. |
| `-D, --delimiter` | `:` | Character that separates a process type from its command. Change only when reading a non-standard format. |
| `-d, --default-port` | `5000` | Value substituted for `$PORT` during variable expansion. |
| `-S, --strict` | `false` | Apply DNS-label rules to process-type names. See [Procfile format](procfile-format.md#strict-mode). |

## What to read next

- [Procfile format](procfile-format.md) - the exact grammar `procfile-util` accepts.
- [Tasks](tasks/README.md) - one page per subcommand.
- [Variable expansion](variable-expansion.md) - how `$PORT`, env files, and getenv interact.
- [Process managers](process-managers.md) - generating systemd/launchd/runit unit files from a Procfile.
- [Dokku integration](dokku-integration.md) - how Dokku itself uses procfile-util, and when you should run it yourself.
- [CI usage](ci-usage.md) - linting Procfiles in pull requests.
