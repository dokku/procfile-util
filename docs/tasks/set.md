# set

Add a new process type or replace the command of an existing one. The Procfile is rewritten in place by default.

## Synopsis

```bash
procfile-util set --process-type <name> --command "<command>" [flags]
```

## Flags

| Flag | Required | Description |
| ---- | -------- | ----------- |
| `-p, --process-type` | Yes | The process type to add or replace. |
| `-c, --command` | Yes | The command to associate with the process type. Quote the value to keep your shell from interpreting it. |
| `-w, --write-path` | No | Write the result to this path instead of overwriting the source Procfile. |
| `-s, --stdout` | No | Print the result to stdout. Mutually exclusive with `--write-path`. |

See the [global flags](../getting-started.md#global-flags) for the rest.

## Behavior

- Exits `0` after writing the updated Procfile.
- Exits `1` if the input Procfile fails to parse, if both `--write-path` and `--stdout` are set, or if the destination cannot be written.
- The new entry is placed first; existing entries follow in their original order, with any earlier entry that shared the same process type removed.

> **Caveat:** Comments and blank lines from the original Procfile are not preserved in the output. If you rely on Procfile comments, generate the output to a different path and reconcile manually.

## Examples

Given:

```
web: bundle exec puma -p $PORT
```

Replace the `web` command in place:

```bash
procfile-util set --process-type web --command "bundle exec puma -p \$PORT --workers 4"
```

Add a new `worker` process and write to a different file:

```bash
procfile-util set --process-type worker --command "bundle exec sidekiq" --write-path Procfile.next
```

Preview the change without touching the file:

```bash
procfile-util set --process-type worker --command "bundle exec sidekiq" --stdout
```

## See also

- [`delete`](delete.md) - remove a process type entirely.
- [Procfile format](../procfile-format.md) - the syntax `set` writes back out.
