# Tasks

One reference page per `procfile-util` subcommand. Every command also accepts the [global flags](../getting-started.md#global-flags) for selecting the Procfile, delimiter, default port, and strict mode.

## Read-only

- [`check`](check.md) - validate that a Procfile parses.
- [`list`](list.md) - print every process type, one per line.
- [`exists`](exists.md) - exit `0` if a process type exists, `1` otherwise.
- [`show`](show.md) - print the command for a single process, with variables expanded.
- [`expand`](expand.md) - print the whole Procfile with variables expanded.

## Mutating

- [`set`](set.md) - add or replace a process type's command.
- [`delete`](delete.md) - remove a process type.

## Format conversion

- [`export`](export.md) - render a Procfile as systemd, launchd, runit, sysv, or upstart unit files.
