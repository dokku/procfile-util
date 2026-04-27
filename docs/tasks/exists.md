# exists

Test whether a Procfile declares a given process type. The answer is communicated through the exit code, which makes this command easy to use as a guard in shell scripts.

## Synopsis

```bash
procfile-util exists --process-type <name> [global flags]
```

## Flags

| Flag | Required | Description |
| ---- | -------- | ----------- |
| `-p, --process-type` | Yes | The process type to look for. |

See the [global flags](../getting-started.md#global-flags) for the rest.

## Behavior

- Exit `0` when the named process type is present.
- Exit `1` when the Procfile parses successfully but the name is missing, when the Procfile fails to parse, or when no `--process-type` is supplied.

The command does **not** print `true`/`false` text. Read the exit code.

## Examples

Given:

```
web: bundle exec puma -p $PORT
worker: bundle exec sidekiq
```

```bash
$ procfile-util exists --process-type web; echo $?
0

$ procfile-util exists --process-type clock; echo $?
No matching process entry found
1
```

Gate a CI step on a process type being present:

```bash
if procfile-util exists --process-type worker; then
  bundle exec rspec spec/workers
fi
```

## See also

- [`list`](list.md) - enumerate every process type when you want all the names.
- [CI usage](../ci-usage.md#gate-downstream-jobs-on-a-process-type) - longer worked example.
