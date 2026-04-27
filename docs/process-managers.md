# Process managers

The [`export`](tasks/export.md) command turns a Procfile into init scripts that a host's process supervisor can start, restart, and stop. This is the bridge between the platform-managed world (Heroku, Dokku, Foreman) and a plain VM where you need processes to survive a reboot.

## When to reach for export

Use `export` when:

- You are leaving a managed platform but still want to run the same Procfile-described processes.
- You want a Procfile-driven install on a non-Dokku host (a VM, an edge box, a developer machine).
- You need each process type to run as multiple replicas - `worker=3` - without writing the unit files by hand.

If you are deploying to Dokku or Heroku, you do not need `export`; the platform reads the Procfile directly. See [Dokku integration](dokku-integration.md).

## Supported supervisors

`export --format <format>` produces files for one of these targets:

| Format | When to use it |
| ------ | -------------- |
| `systemd` | Most modern Linux distributions. Generates one unit per process replica plus a `.target` that owns them. Default. |
| `systemd-user` | Same as `systemd`, but writes user units (`~/.config/systemd/user`). Use for per-user services that should not require root. |
| `sysv` | Older Linux distributions and BSDs that still use SysV-style init scripts. |
| `upstart` | Pre-systemd Ubuntu (`14.04` and earlier). Rarely needed today. |
| `runit` | Lightweight supervisor used in Void Linux, some Alpine setups, and inside some containers. |
| `launchd` | macOS. Generates `.plist` files suitable for `~/Library/LaunchAgents` or `/Library/LaunchDaemons`. |

`--location <dir>` is required for every format - it is where the generated files are written.

## Replicas: the `--formation` flag

The Procfile only declares process types; it says nothing about how many copies of each should run. `--formation` fills that in:

```bash
procfile-util export --format systemd --location ./units --formation web=2,worker=3
```

The above produces `web.1`, `web.2`, `worker.1`, `worker.2`, `worker.3` units. The default `all=1` runs a single replica of every process type.

## Identity flags

Tell the supervisor who and where to run as:

| Flag | Default | Purpose |
| ---- | ------- | ------- |
| `--app` | `app` | Logical name of the application; used in unit filenames and labels. |
| `--user` | `--app` value | OS user to run the command as. |
| `--group` | `--app` value | OS group to run as. |
| `--home` | `/home/<current user>` | Sets `$HOME` for spawned processes. |
| `--working-directory-path` | `/` | The working directory the process is started in. |
| `--log-path` | `/var/log` | Directory the supervisor writes stdout/stderr logs to. |
| `--description` | empty | Free-form description embedded in the unit (where the format supports it). |

## Lifecycle flags

| Flag | Default | Purpose |
| ---- | ------- | ------- |
| `--prestart` | empty | A command run before each start/restart. If it exits non-zero the start is aborted. Useful for config validation or warm-up checks. |
| `--timeout` | `5` | Seconds the supervisor waits for a graceful shutdown before sending `SIGKILL`. |
| `--nice` | empty | `nice` value applied to the process. |
| `--env-file` | empty | A dotenv file whose contents are written into the generated unit so the supervisor sees them at start time. |

## Resource limits (`ulimit`)

These flags translate into `ulimit` calls inside a wrapper that the generated unit runs. They map one-to-one to `setrlimit` resources:

| Flag | `setrlimit` resource | Meaning |
| ---- | -------------------- | ------- |
| `--limit-coredump` | `RLIMIT_CORE` | Largest core-dump size, in blocks. |
| `--limit-cputime` | `RLIMIT_CPU` | Maximum CPU seconds the process may consume. |
| `--limit-data` | `RLIMIT_DATA` | Maximum data segment size. |
| `--limit-file-size` | `RLIMIT_FSIZE` | Largest file the process may write. |
| `--limit-locked-memory` | `RLIMIT_MEMLOCK` | Maximum bytes the process may `mlock`. |
| `--limit-open-files` | `RLIMIT_NOFILE` | Maximum number of open files and sockets. |
| `--limit-user-processes` | `RLIMIT_NPROC` | Maximum running processes / threads. Per-user, not per-process - prefer cgroup limits where possible. |
| `--limit-physical-memory` | `RLIMIT_RSS` | Maximum resident set size, in bytes. |
| `--limit-stack-size` | `RLIMIT_STACK` | Maximum stack size, in bytes. |

Any limit you do not set is left unconstrained.

## A worked example

```
web: bundle exec puma -p $PORT
worker: bundle exec sidekiq
```

Generate two replicas of `web` and one `worker` as systemd units, run as the `payments` user:

```bash
mkdir -p ./units
procfile-util export \
  --format systemd \
  --location ./units \
  --app payments \
  --user payments --group payments \
  --working-directory-path /srv/payments \
  --formation web=2,worker=1 \
  --env-file .env.production
```

`./units` will contain unit files like `payments-web.1.service`, `payments-web.2.service`, `payments-worker.1.service`, and a `payments.target` that owns them. Copy them into `/etc/systemd/system/`, run `systemctl daemon-reload`, and `systemctl start payments.target`.

See [`export`](tasks/export.md) for the bare flag reference.
