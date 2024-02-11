#!/usr/bin/env bats

export SYSTEM_NAME="$(uname -s | tr '[:upper:]' '[:lower:]')"
export PROCFILE_BIN="build/$SYSTEM_NAME/procfile-util-amd64"

setup_file() {
  make prebuild $PROCFILE_BIN
}

teardown_file() {
  make clean
}

@test "[lax] comments" {
  run $PROCFILE_BIN check -P fixtures/comments.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "valid procfile detected 2custom, cron, custom, release, web, wor-ker"

  run $PROCFILE_BIN list -P fixtures/comments.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "2custom
cron
custom
release
web
wor-ker"
}

@test "[lax] forwardslash-comments" {
  run $PROCFILE_BIN check -P fixtures/forwardslash-comments.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "valid procfile detected web, worker, worker-2"

  run $PROCFILE_BIN show -P fixtures/forwardslash-comments.Procfile -p worker
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "node worker.js"

  run $PROCFILE_BIN list -P fixtures/forwardslash-comments.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "web
worker
worker-2"
}

@test "[lax] multiple" {
  run $PROCFILE_BIN check -P fixtures/multiple.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "valid procfile detected release, web, webpacker, worker"

  run $PROCFILE_BIN list -P fixtures/multiple.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "release
web
webpacker
worker"
}

@test "[lax] port" {
  run $PROCFILE_BIN check -P fixtures/port.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "valid procfile detected web, worker"

  run $PROCFILE_BIN show -P fixtures/port.Procfile -p web
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "node web.js --port 5000"

  run $PROCFILE_BIN list -P fixtures/port.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "web
worker"
}

@test "[strict] comments" {
  run $PROCFILE_BIN check -S -P fixtures/comments.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "valid procfile detected 2custom, cron, custom, release, web, wor-ker"

  run $PROCFILE_BIN list -S -P fixtures/comments.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "2custom
cron
custom
release
web
wor-ker"
}

@test "[strict] forwardslash-comments" {
  run $PROCFILE_BIN check -S -P fixtures/forwardslash-comments.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "valid procfile detected web, worker, worker-2"

  run $PROCFILE_BIN show -S -P fixtures/forwardslash-comments.Procfile -p worker
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "node worker.js"

  run $PROCFILE_BIN list -S -P fixtures/forwardslash-comments.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "web
worker
worker-2"
}

@test "[strict] multiple" {
  run $PROCFILE_BIN check -S -P fixtures/multiple.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "valid procfile detected release, web, webpacker, worker"

  run $PROCFILE_BIN list -S -P fixtures/multiple.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "release
web
webpacker
worker"
}

@test "[strict] port" {
  run $PROCFILE_BIN check -S -P fixtures/port.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "valid procfile detected web, worker"

  run $PROCFILE_BIN show -S -P fixtures/port.Procfile -p web
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "node web.js --port 5000"

  run $PROCFILE_BIN list -S -P fixtures/port.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  assert_output "web
worker"
}

flunk() {
  {
    if [[ "$#" -eq 0 ]]; then
      cat -
    else
      echo "$*"
    fi
  }
  return 1
}

assert_equal() {
  if [[ "$1" != "$2" ]]; then
    {
      echo "expected: '$1'"
      echo "actual:   '$2'"
    } | flunk
  fi
}

assert_exit_status() {
  exit_status="$1"
  if [[ "$status" -ne "$exit_status" ]]; then
    {
      echo "expected exit status: $exit_status"
      echo "actual exit status:   $status"
    } | flunk
    flunk
  fi
}

assert_failure() {
  if [[ "$status" -eq 0 ]]; then
    flunk "expected failed exit status"
  elif [[ "$#" -gt 0 ]]; then
    assert_output "$1"
  fi
}

assert_success() {
  if [[ "$status" -ne 0 ]]; then
    flunk "command failed with exit status $status"
  elif [[ "$#" -gt 0 ]]; then
    assert_output "$1"
  fi
}

assert_output() {
  local expected
  if [[ $# -eq 0 ]]; then
    expected="$(cat -)"
  else
    expected="$1"
  fi
  assert_equal "$expected" "$output"
}

assert_output_contains() {
  local input="$output"
  local expected="$1"
  local count="${2:-1}"
  local found=0
  until [ "${input/$expected/}" = "$input" ]; do
    input="${input/$expected/}"
    found=$((found + 1))
  done
  assert_equal "$count" "$found"
}

assert_output_not_exists() {
  [[ -z "$output" ]] || flunk "expected no output, found some"
}
