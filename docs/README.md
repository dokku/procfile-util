# Documentation

Complete documentation for `procfile-util`, a tool for parsing, validating, and exporting Heroku-style Procfiles.

## Getting started

- [Getting started](getting-started.md) - what a Procfile is, why use procfile-util, installation, and your first Procfile

## Reference

- [Procfile format](procfile-format.md) - syntax, comments, and strict mode
- [Tasks](tasks/README.md) - reference for every subcommand

## Guides

- [Variable expansion](variable-expansion.md) - how `show` and `expand` resolve `$PORT`, `--env-file`, and `--allow-getenv`
- [Process managers](process-managers.md) - turning a Procfile into systemd, runit, launchd, sysv, or upstart units
- [Dokku integration](dokku-integration.md) - how procfile-util fits into a Dokku deploy
- [CI usage](ci-usage.md) - linting Procfiles and gating CI steps on process types
