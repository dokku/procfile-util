# Dokku integration

If you deploy to [Dokku](https://dokku.com), you are already using `procfile-util` indirectly. This page explains where it shows up and when it is worth invoking yourself.

## How Dokku uses procfile-util

When you push an application to Dokku, the build/release flow needs to know which processes the app declares. Dokku shells out to `procfile-util` to:

- **Validate** the committed `Procfile` and refuse to release a build whose Procfile cannot be parsed.
- **Enumerate** process types so commands like `dokku ps:scale` can offer the right names.
- **Read** an individual process command when the scheduler needs to construct a container `CMD`.

In other words, procfile-util is the parser of record for the format. Anything you can express in a Dokku-deployed Procfile, procfile-util can read - and anything procfile-util rejects, Dokku rejects too.

## When to run it yourself

Dokku runs procfile-util in the build container, so any error you would see there can also be reproduced locally. Practical situations:

### Pre-flight before `git push dokku`

Catch parse errors before the push, instead of waiting for the build to fail:

```bash
procfile-util check --strict
```

`--strict` matches the constraint Dokku itself applies to process type names; see [Procfile format](procfile-format.md#strict-mode).

### Debugging "process type not found"

When `dokku ps:scale my-app worker=2` complains the process type does not exist, list what your Procfile actually defines:

```bash
procfile-util list
```

If `worker` is missing from that output, fix the Procfile in source control and redeploy.

### Inspecting the command Dokku will run

To see exactly what `dokku ps:scale` would launch for a given process - with `$PORT` substituted - run:

```bash
procfile-util show --process-type web --allow-getenv --default-port 5000
```

`--allow-getenv` lets you preview substitution against environment variables you have set locally. The actual values used in production come from `dokku config`, not your shell.

### Generating units for a non-Dokku host

If you eventually need to run the same Procfile on a host that does not have Dokku, [`export`](tasks/export.md) generates systemd, launchd, runit, sysv, or upstart units. See [Process managers](process-managers.md) for the supported formats.

## A worked example

A small Procfile in a Rails app:

```
web: bundle exec puma -p $PORT
worker: bundle exec sidekiq -e $RAILS_ENV
release: bundle exec rake db:migrate
```

Local validation:

```bash
$ procfile-util check --strict
valid procfile detected web, worker, release

$ procfile-util list
web
worker
release

$ RAILS_ENV=production procfile-util show --process-type worker --allow-getenv
bundle exec sidekiq -e production
```

After that passes, `git push dokku master` will not fail at the Procfile-parse step. Anything else (build, release, deploy) is downstream of procfile-util.

## See also

- [`check`](tasks/check.md) - the lint command Dokku runs internally.
- [`list`](tasks/list.md) - process-type enumeration.
- [`show`](tasks/show.md) and [Variable expansion](variable-expansion.md) - previewing the actual command.
- [Procfile format](procfile-format.md) - the grammar Dokku enforces via this tool.
