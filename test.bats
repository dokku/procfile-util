#!/usr/bin/env bats

export SYSTEM_NAME="$(uname -s | tr '[:upper:]' '[:lower:]')"
export PROCFILE_BIN="build/$SYSTEM_NAME/procfile-util"

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
  [[ "$output" == "valid procfile detected 2custom, cron, custom, release, web, wor-ker" ]]
}

@test "[lax] multiple" {
  run $PROCFILE_BIN check -P fixtures/multiple.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  [[ "$output" == "valid procfile detected release, web, webpacker, worker" ]]
}

@test "[lax] port" {
  run $PROCFILE_BIN check -P fixtures/port.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  [[ "$output" == "valid procfile detected web, worker" ]]

  run $PROCFILE_BIN show -P fixtures/port.Procfile -p web
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  [[ "$output" == "node web.js --port 5000" ]]
}

@test "[strict] comments" {
  run $PROCFILE_BIN check -S -P fixtures/comments.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  [[ "$output" == "valid procfile detected 2custom, cron, custom, release, web, wor-ker" ]]
}

@test "[strict] multiple" {
  run $PROCFILE_BIN check -S -P fixtures/multiple.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  [[ "$output" == "valid procfile detected release, web, webpacker, worker" ]]
}

@test "[strict] port" {
  run $PROCFILE_BIN check -S -P fixtures/port.Procfile
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  [[ "$output" == "valid procfile detected web, worker" ]]

  run $PROCFILE_BIN show -S -P fixtures/port.Procfile -p web
  echo "output: $output"
  echo "status: $status"
  [[ "$status" -eq 0 ]]
  [[ "$output" == "node web.js --port 5000" ]]
}
