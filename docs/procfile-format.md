# Procfile format

A Procfile is a plain-text file that declares one or more long-running processes for an application. The format was [introduced by Heroku](https://blog.heroku.com/the_new_heroku_1_process_model_procfile) and adopted by Dokku, Foreman, and many similar tools.

## At a glance

A Procfile is a list of `process type: command` pairs, one per line. Comments, blank lines, and trailing comments are allowed:

```
# Long-running web server. $PORT is provided by the platform.
web: bundle exec puma -p $PORT

# Background worker.
worker: bundle exec rake jobs:work  # high-priority queue

// Optional alternate comment style.
clock: bundle exec clockwork clock.rb
```

`procfile-util` reads this file, validates it, and lets every other subcommand operate on the parsed entries.

## Line types

Each line in a Procfile is one of:

- **A comment** - the line begins with `#` or `//`. The whole line is ignored.
- **A blank line** - an empty line, optionally with trailing whitespace. Ignored.
- **A process-type / command pair** - the form `<process type><delimiter> <command>`, with an optional trailing comment.

When a process line carries a trailing comment, the `#` or `//` introducing the comment **must** be preceded by at least one whitespace character. Without that whitespace, the `#` or `//` is treated as part of the command.

The delimiter defaults to `:` and can be overridden with `--delimiter`/`-D`.

## Process-type and command syntax

- `<process type>` is matched by the character class `[A-Za-z0-9_-]+`. Common names are `web`, `worker`, `urgentworker`, `clock`.
- `<command>` is the shell command that starts the process, for example `bundle exec puma -p $PORT`.

If a line does not match one of the three forms above, parsing fails and the whole file is rejected. This is intentional: rejecting an unparseable Procfile is safer than guessing past a merge conflict marker or a typo and shipping a broken release.

Process types must be unique within a single Procfile. Rather than silently picking the first or last entry for a duplicate, `procfile-util` reports the conflict and exits.

## Strict mode

Strict mode is enabled by passing `--strict`/`-S` to any subcommand. It tightens the rules for `<process type>` to match a [DNS label](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-label-names):

- At most 63 characters.
- Lowercase alphanumeric characters or `-` only.
- Starts and ends with an alphanumeric character.

This matches the constraint many platforms place on process names when they are reused as DNS components (for example, in service names or URLs). Turning on strict mode in CI means a Procfile that lints clean locally cannot accidentally violate that constraint downstream.

```bash
procfile-util check --strict
```

A Procfile with a process type like `Web_1` passes the default parser but fails strict mode.
