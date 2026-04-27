# check

Validate that a Procfile is well-formed. This is the same parser Dokku runs against the file you push, so any error you see here is an error Dokku would also reject.

## Synopsis

```bash
procfile-util check [global flags]
```

`check` takes no command-specific flags. See the [global flags](../getting-started.md#global-flags) for `--procfile`, `--delimiter`, `--default-port`, and `--strict`.

## Behavior

- Exits `0` and prints `valid procfile detected <names>` when every line parses.
- Exits `1` and prints the parse error otherwise. Common failures: malformed lines, duplicated process types, an empty file with no entries, and (with `--strict`) names that are not DNS labels.

## Examples

Validate `./Procfile`:

```bash
procfile-util check
```

Validate a file at a different path:

```bash
procfile-util check --procfile config/Procfile.staging
```

Apply DNS-label rules - the constraint Dokku uses internally:

```bash
procfile-util check --strict
```

Read the Procfile from stdin (use `-P -`):

```bash
cat Procfile | procfile-util check --procfile -
```

## See also

- [Procfile format](../procfile-format.md) - the grammar `check` enforces.
- [CI usage](../ci-usage.md#lint-the-procfile-on-every-pull-request) - using `check` as a CI gate.
