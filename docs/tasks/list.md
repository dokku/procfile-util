# list

Print every process type declared in a Procfile, one name per line. The output is the same set of names that `dokku ps:scale` would accept.

## Synopsis

```bash
procfile-util list [global flags]
```

`list` takes no command-specific flags. See the [global flags](../getting-started.md#global-flags) for `--procfile`, `--delimiter`, `--default-port`, and `--strict`.

## Behavior

- Exits `0` after printing the names.
- Exits `1` if the Procfile fails to parse or contains no entries.

## Examples

Given:

```
web: bundle exec puma -p $PORT
worker: bundle exec sidekiq
clock: bundle exec clockwork clock.rb
```

```bash
$ procfile-util list
web
worker
clock
```

Iterate over every process type in a shell loop:

```bash
for type in $(procfile-util list); do
  echo "scaling $type"
  dokku ps:scale my-app "$type=2"
done
```

## See also

- [`exists`](exists.md) - check whether a single process type is defined without enumerating all of them.
- [Dokku integration](../dokku-integration.md#debugging-process-type-not-found) - using `list` to debug missing process types.
