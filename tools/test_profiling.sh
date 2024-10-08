#!/bin/bash

set -e

export LLVM_PROFILE_FILE="${PWD}/profile/%p-%m.profraw"

rm -rf _test_objects

# run normal tests
RUSTFLAGS="-C instrument-coverage" cargo test --all-targets

# run tests that require to be run in a single thread
RUSTFLAGS="-C instrument-coverage" cargo test --all-targets -- --test-threads=1  --ignored


files=$(RUSTFLAGS="-C instrument-coverage" cargo test --tests --no-run --message-format=json | jq -r "select(.profile.test == true) | .filenames[]" | grep -v dSYM - )

for file in $files;
do 
  printf "%s %s " -object $file >> _test_objects
done
