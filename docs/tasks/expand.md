# expand

Print the entire Procfile with shell variables substituted in every command. The output uses the same `process: command` shape as the input, so it can be diffed against snapshots or piped into another tool.

## Synopsis

```bash
procfile-util expand [flags]
```

## Flags

| Flag | Default | Description |
| ---- | ------- | ----------- |
| `-p, --process-type` | empty | If set, only the matching process type is emitted. Without it, every entry is printed. |
| `-a, --allow-getenv` | `false` | Allow the current shell environment to provide variable values. |
| `-e, --env-file` | empty | Path to a dotenv file. Values in this file take precedence over `--allow-getenv` and the built-in defaults. |

See the [global flags](../getting-started.md#global-flags) for the rest. The output uses whatever character is set as `--delimiter`.

## Behavior

- Exits `0` after printing the expanded Procfile.
- Exits `1` if the Procfile fails to parse, if the env file cannot be read, or if any expansion produces an error.

The expansion rules are identical to [`show`](show.md). See [Variable expansion](../variable-expansion.md) for the source-priority order.

## Examples

Given `.env.ci`:

```
RAILS_ENV=ci
PORT=4000
```

And:

```
web: bundle exec puma -p $PORT
worker: bundle exec sidekiq -e $RAILS_ENV
```

Render every entry against the env file:

```bash
$ procfile-util expand --env-file .env.ci
web: bundle exec puma -p 4000
worker: bundle exec sidekiq -e ci
```

Restrict to a single process while still using the env file:

```bash
$ procfile-util expand --process-type worker --env-file .env.ci
worker: bundle exec sidekiq -e ci
```

Snapshot test against an expected output:

```bash
procfile-util expand --env-file .env.ci > rendered.procfile
diff -u expected.procfile rendered.procfile
```

## See also

- [`show`](show.md) - print one process command rather than the whole file.
- [Variable expansion](../variable-expansion.md) - precedence rules for env files, getenv, and defaults.
- [CI usage](../ci-usage.md#snapshot-the-rendered-procfile) - using `expand` in a snapshot test.
