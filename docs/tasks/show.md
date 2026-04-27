# show

Print the command for a single process type, with shell variables expanded. Use it to preview the exact command a deploy target would execute.

## Synopsis

```bash
procfile-util show --process-type <name> [flags]
```

## Flags

| Flag | Default | Description |
| ---- | ------- | ----------- |
| `-p, --process-type` | (required) | The process type whose command should be printed. |
| `-a, --allow-getenv` | `false` | Allow values from the current shell environment to participate in expansion. |
| `-e, --env-file` | empty | Path to a dotenv file. Values in this file take precedence over `--allow-getenv` and the built-in defaults. |

See the [global flags](../getting-started.md#global-flags) for the rest. `--default-port` is the value substituted for `$PORT` when no other source provides one.

## Behavior

- Exits `0` and prints the expanded command.
- Exits `1` if the Procfile fails to parse, if the requested process type is missing, or if the env file cannot be read.

For details on which sources are consulted and in what order, see [Variable expansion](../variable-expansion.md).

## Examples

Given:

```
web: bundle exec puma -p $PORT --environment $RAILS_ENV
```

Default expansion - `PORT` falls back to `5000`, `RAILS_ENV` is unset:

```bash
$ procfile-util show --process-type web
bundle exec puma -p 5000 --environment
```

Override the port:

```bash
$ procfile-util show --process-type web --default-port 8080
bundle exec puma -p 8080 --environment
```

Pull `RAILS_ENV` from the current shell:

```bash
$ RAILS_ENV=production procfile-util show --process-type web --allow-getenv
bundle exec puma -p 5000 --environment production
```

Pull every variable from a dotenv file:

```bash
$ procfile-util show --process-type web --env-file .env.production
bundle exec puma -p 4000 --environment production
```

## See also

- [`expand`](expand.md) - same expansion rules, applied to every entry in the Procfile.
- [Variable expansion](../variable-expansion.md) - precedence rules for env files, getenv, and defaults.
