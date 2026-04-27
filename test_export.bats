#!/usr/bin/env bats

export SYSTEM_NAME="$(uname -s | tr '[:upper:]' '[:lower:]')"
export PROCFILE_BIN="build/$SYSTEM_NAME/procfile-util-amd64"
export TEST_FIXTURE="fixtures/export.Procfile"

setup_file() {
  make prebuild $PROCFILE_BIN
}

teardown_file() {
  make clean
}

setup() {
  WORK_DIR="$(mktemp -d)"
}

teardown() {
  rm -rf "$WORK_DIR"
}

@test "[export] systemd produces units accepted by systemd-analyze" {
  run $PROCFILE_BIN export --procfile $TEST_FIXTURE --format systemd \
    --location "$WORK_DIR" --app payments --formation web=2,worker=3
  [ "$status" -eq 0 ]

  for f in "$WORK_DIR"/etc/systemd/system/*.service "$WORK_DIR"/etc/systemd/system/*.target; do
    run systemd-analyze verify "$f"
    echo "file: $f"
    echo "output: $output"
    [ "$status" -eq 0 ]
    [[ "$output" != *"Failed to parse"* ]]
    [[ "$output" != *"marked executable"* ]]
  done
}

@test "[export] systemd-user produces units accepted by systemd-analyze" {
  run $PROCFILE_BIN export --procfile $TEST_FIXTURE --format systemd-user \
    --location "$WORK_DIR" --app payments --formation web=2,worker=3
  [ "$status" -eq 0 ]

  for f in "$WORK_DIR"/home/*/.config/systemd/user/*.service; do
    run systemd-analyze verify "$f"
    echo "file: $f"
    echo "output: $output"
    [ "$status" -eq 0 ]
    [[ "$output" != *"Failed to parse"* ]]
    [[ "$output" != *"marked executable"* ]]
  done
}

@test "[export] sysv init scripts pass shellcheck and bash -n" {
  run $PROCFILE_BIN export --procfile $TEST_FIXTURE --format sysv \
    --location "$WORK_DIR" --app payments --formation web=2,worker=3
  [ "$status" -eq 0 ]

  for f in "$WORK_DIR"/etc/init.d/*; do
    run bash -n "$f"
    echo "file: $f"
    echo "output: $output"
    [ "$status" -eq 0 ]
    run shellcheck -S error "$f"
    echo "output: $output"
    [ "$status" -eq 0 ]
  done
}

@test "[export] upstart conf files contain required stanzas" {
  run $PROCFILE_BIN export --procfile $TEST_FIXTURE --format upstart \
    --location "$WORK_DIR" --app payments --formation web=2,worker=3
  [ "$status" -eq 0 ]

  for f in "$WORK_DIR"/etc/init/*.conf; do
    run grep -q '^start on' "$f"
    echo "file: $f"
    [ "$status" -eq 0 ]
  done
}

@test "[export] runit run scripts pass shellcheck and bash -n" {
  run $PROCFILE_BIN export --procfile $TEST_FIXTURE --format runit \
    --location "$WORK_DIR" --app payments --formation web=2,worker=3
  [ "$status" -eq 0 ]

  for f in "$WORK_DIR"/service/*/run "$WORK_DIR"/service/*/log/run; do
    run bash -n "$f"
    echo "file: $f"
    echo "output: $output"
    [ "$status" -eq 0 ]
    run shellcheck -S error "$f"
    echo "output: $output"
    [ "$status" -eq 0 ]
  done
}

@test "[export] launchd plists are well-formed" {
  run $PROCFILE_BIN export --procfile $TEST_FIXTURE --format launchd \
    --location "$WORK_DIR" --app payments --formation web=2,worker=3
  [ "$status" -eq 0 ]

  for f in "$WORK_DIR"/Library/LaunchDaemons/*.plist; do
    if command -v plutil >/dev/null 2>&1; then
      run plutil -lint "$f"
    elif command -v plistutil >/dev/null 2>&1; then
      run plistutil -i "$f" -o /dev/null
    else
      run xmllint --noout "$f"
    fi
    echo "file: $f"
    echo "output: $output"
    [ "$status" -eq 0 ]
  done
}
