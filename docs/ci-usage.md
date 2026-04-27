# CI usage

`procfile-util` exits non-zero on every error condition, which makes it a good fit for CI pipelines. Three patterns cover almost everything you might want.

## Lint the Procfile on every pull request

Run [`check`](tasks/check.md) as a build step. It rejects malformed lines, duplicate process types, and (with `--strict`) names that would not survive as DNS labels:

```yaml
# .github/workflows/lint-procfile.yml
name: Lint Procfile
on:
  pull_request:
jobs:
  procfile:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Install procfile-util
        run: |
          curl -fsSL https://raw.githubusercontent.com/dokku/procfile-util/master/install.sh | sudo sh
      - name: Validate Procfile
        run: procfile-util check --strict
```

Pin a specific version with `VERSION=v0.20.4` if you want reproducible builds; otherwise the installer takes the latest GitHub release.

## Gate downstream jobs on a process type

`procfile-util exists --process-type <name>` returns exit code `0` when the process type is present and `1` otherwise. Use it to skip jobs that do not apply to the current Procfile:

```yaml
- name: Run worker tests
  run: |
    if procfile-util exists --process-type worker; then
      bundle exec rspec spec/workers
    else
      echo "no worker process declared, skipping worker tests"
    fi
```

This is more honest than always running the worker tests and silently passing when there is no worker.

## Snapshot the rendered Procfile

`procfile-util expand --env-file .env.ci` produces the Procfile with every variable substituted, in the same form `show` would produce for each entry. Useful when you want to assert that a config change does not silently alter what production runs:

```bash
procfile-util expand --env-file .env.ci > rendered.procfile
diff -u expected.procfile rendered.procfile
```

Commit `expected.procfile` alongside your source. The `diff` step fails the build whenever the rendered output drifts from the snapshot.

## Reproducibility tips

- Pin `VERSION` in your install step so `master` does not silently update.
- Run with `--strict` in CI even if you do not run with it locally; it surfaces problems earlier.
- Cache the binary between runs: a single `procfile-util` is small enough that a basic actions cache keyed on `VERSION` is enough.

## See also

- [`check`](tasks/check.md), [`exists`](tasks/exists.md), [`expand`](tasks/expand.md) - the three commands referenced above.
- [Variable expansion](variable-expansion.md) - what `expand --env-file` actually does.
