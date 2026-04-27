# delete

Remove a process type from a Procfile. The Procfile is rewritten in place by default.

## Synopsis

```bash
procfile-util delete --process-type <name> [flags]
```

## Flags

| Flag | Required | Description |
| ---- | -------- | ----------- |
| `-p, --process-type` | Yes | The process type to remove. |
| `-w, --write-path` | No | Write the result to this path instead of overwriting the source Procfile. |
| `-s, --stdout` | No | Print the result to stdout. Mutually exclusive with `--write-path`. |

See the [global flags](../getting-started.md#global-flags) for the rest.

## Behavior

- Exits `0` after writing the updated Procfile.
- Exits `1` if the input Procfile fails to parse, if both `--write-path` and `--stdout` are set, or if the destination cannot be written.
- Deleting a process type that is not present is a no-op: the Procfile is rewritten with no entries removed and exit status is still `0`.

> **Caveat:** Comments and blank lines from the original Procfile are not preserved in the output. If you rely on Procfile comments, generate the output to a different path and reconcile manually.

## Examples

Given:

```
web: bundle exec puma -p $PORT
worker: bundle exec sidekiq
```

Remove `worker` in place:

```bash
procfile-util delete --process-type worker
```

Write the result to a different file:

```bash
procfile-util delete --process-type worker --write-path Procfile.next
```

Preview the change on stdout without touching the source file:

```bash
procfile-util delete --process-type worker --stdout
```

## See also

- [`set`](set.md) - add or replace a process type.
- [Procfile format](../procfile-format.md) - the syntax `delete` writes back out.
