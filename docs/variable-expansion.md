# Variable expansion

The [`show`](tasks/show.md) and [`expand`](tasks/expand.md) commands resolve shell-style variables (`$PORT`, `${DATABASE_URL}`, and so on) inside a process command before printing it. This page explains where the values come from and the order they are tried in.

## Why expansion matters

A line like `web: bundle exec puma -p $PORT` is what you commit to source control. What actually runs in production is `bundle exec puma -p 5000` (or whatever port the platform assigned). When you debug a deploy, what you usually want to see is the *expanded* form, not the template. That is what `show` and `expand` produce.

## What is expanded

Expansion uses Go's `os.Expand` semantics, so anything matching `$NAME` or `${NAME}` is a candidate. Two names are always populated by `procfile-util`, regardless of any flag:

| Name | Source |
| ---- | ------ |
| `PORT` | The value of `--default-port`, default `5000`. |
| `PS` | The name of the process type currently being expanded (for example, `web`). |

Anything else expands to an empty string unless you opt in with `--env-file`, `--allow-getenv`, or both.

## Sources, in priority order

When a variable is referenced, `procfile-util` checks sources in this order and stops at the first hit:

1. **`--env-file <path>`**, if provided. The file is parsed with [`godotenv`](https://github.com/joho/godotenv), so `KEY=value` and `export KEY=value` lines both work.
2. **The current process environment**, but only when `--allow-getenv` is set. Without that flag, the env is treated as if it were empty.
3. **The built-in defaults** above (`PORT` from `--default-port`, `PS` from the process name).

If no source supplies a value, the variable expands to an empty string.

## Examples

Given this Procfile:

```
web: bundle exec puma -p $PORT --environment $RAILS_ENV
worker: bundle exec sidekiq -e $RAILS_ENV
```

### No flags

```bash
procfile-util show --process-type web
# bundle exec puma -p 5000 --environment
```

`PORT` falls back to the default port. `RAILS_ENV` is unset and expands to an empty string.

### `--default-port`

```bash
procfile-util show --process-type web --default-port 8080
# bundle exec puma -p 8080 --environment
```

### `--allow-getenv`

```bash
RAILS_ENV=production procfile-util show --process-type web --allow-getenv
# bundle exec puma -p 5000 --environment production
```

`RAILS_ENV` is now read from your shell. `--allow-getenv` is opt-in because it lets the surrounding shell influence command rendering, which is convenient locally but a hazard in CI where unrelated environment variables may bleed in.

### `--env-file`

Given a `.env` file containing:

```
RAILS_ENV=staging
PORT=4000
```

```bash
procfile-util show --process-type web --env-file .env
# bundle exec puma -p 4000 --environment staging
```

Values in `.env` win even over the `--default-port` fallback because the env file is the highest-priority source.

### Combining flags

```bash
RAILS_ENV=production procfile-util show \
  --process-type web \
  --env-file .env.shared \
  --allow-getenv \
  --default-port 8080
```

Resolution order for each variable: `.env.shared` first; then the shell environment (because `--allow-getenv` is set); then the built-in defaults (`PORT` falls back to `8080` if it is in neither file nor shell).

## Expanding the whole Procfile at once

`expand` applies the same rules to every entry and prints the full Procfile with substitutions applied:

```bash
procfile-util expand --env-file .env.ci
# web: bundle exec puma -p 4000 --environment staging
# worker: bundle exec sidekiq -e staging
```

This is the form to use in CI when you want to assert that the rendered Procfile matches an expected snapshot.
