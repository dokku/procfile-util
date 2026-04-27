# export

Render a Procfile as init scripts that a host's process supervisor can run. Use this when you need the Procfile-described processes to keep running on a host that does not have Dokku, Heroku, or Foreman driving them.

The conceptual material - which supervisors are supported and what each flag is for - lives in [Process managers](../process-managers.md). This page is the bare flag reference.

## Synopsis

```bash
procfile-util export --format <format> --location <dir> [flags]
```

`--location` is required. Without it the command exits `1` with `no output location specified`.

## Required flags

| Flag | Default | Description |
| ---- | ------- | ----------- |
| `--format` | `systemd` | Target supervisor. One of `systemd`, `systemd-user`, `sysv`, `upstart`, `runit`, `launchd`. |
| `--location` | (required) | Directory where generated unit files are written. |

## Identity flags

| Flag | Default | Description |
| ---- | ------- | ----------- |
| `--app` | `app` | Logical application name. Used in unit filenames and labels. |
| `--user` | `--app` | OS user to run the command as. |
| `--group` | `--app` | OS group to run the command as. |
| `--home` | `/home/<current user>` | Value used for `$HOME` inside spawned processes. |
| `--working-directory-path` | `/` | Working directory for the process. |
| `--description` | empty | Free-form description embedded in the unit. |
| `--log-path` | `/var/log` | Directory the supervisor writes stdout/stderr logs to. |

## Process behavior flags

| Flag | Default | Description |
| ---- | ------- | ----------- |
| `--formation` | `all=1` | Comma-separated `process=count` pairs that decide how many replicas to emit. |
| `--prestart` | empty | Command run before each start/restart. A non-zero exit aborts the start. |
| `--timeout` | `5` | Seconds to wait for graceful shutdown before `SIGKILL`. |
| `--nice` | empty | `nice` value applied to the process. |
| `-e, --env-file` | empty | Dotenv file whose values are baked into the generated units. |
| `--run` | `/var/run/<app>` | Directory used for pid files. |

## Resource limit flags

Each maps to a `setrlimit` resource. Unset flags leave the limit at the system default.

| Flag | `setrlimit` |
| ---- | ----------- |
| `--limit-coredump` | `RLIMIT_CORE` |
| `--limit-cputime` | `RLIMIT_CPU` |
| `--limit-data` | `RLIMIT_DATA` |
| `--limit-file-size` | `RLIMIT_FSIZE` |
| `--limit-locked-memory` | `RLIMIT_MEMLOCK` |
| `--limit-open-files` | `RLIMIT_NOFILE` |
| `--limit-user-processes` | `RLIMIT_NPROC` |
| `--limit-physical-memory` | `RLIMIT_RSS` |
| `--limit-stack-size` | `RLIMIT_STACK` |

See [Process managers](../process-managers.md#resource-limits-ulimit) for guidance on each.

## Examples

Generate one of every process as a systemd unit:

```bash
mkdir -p ./units
procfile-util export --format systemd --location ./units
```

Generate two `web` replicas and one `worker`, run as the `payments` user, with environment baked in:

```bash
procfile-util export \
  --format systemd \
  --location ./units \
  --app payments \
  --user payments --group payments \
  --working-directory-path /srv/payments \
  --formation web=2,worker=1 \
  --env-file .env.production
```

Generate launchd plists for a macOS box:

```bash
mkdir -p ~/Library/LaunchAgents/payments
procfile-util export \
  --format launchd \
  --location ~/Library/LaunchAgents/payments \
  --app payments
```

## See also

- [Process managers](../process-managers.md) - per-format guidance and a worked deployment example.
