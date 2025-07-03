#!/bin/bash

# gtest: smarter go test wrapper
set -e

# Default test options
TEST_DIRS="./..."
RUN_PATTERN="."
BENCH_PATTERN=""
VERBOSE=0
MEMSTATS=0

usage() {
  echo "Usage: $0 [-d dir] [-r regex] [-b benchmark_regex] [-v] [-m]"
  echo "  -d dir         Directory to test (default: ./...)"
  echo "  -r regex       Regex for test names (default: .)"
  echo "  -b regex       Regex for benchmark names (default: none)"
  echo "  -v             Verbose test output"
  echo "  -m             Include benchmark memory stats"
  exit 1
}

while getopts "d:r:b:vm" opt; do
  case "$opt" in
    d) TEST_DIRS="$OPTARG" ;;
    r) RUN_PATTERN="$OPTARG" ;;
    b) BENCH_PATTERN="$OPTARG" ;;
    v) VERBOSE=1 ;;
    m) MEMSTATS=1 ;;
    *) usage ;;
  esac
done

CMD="go test $TEST_DIRS -run=$RUN_PATTERN"

if [ -n "$BENCH_PATTERN" ]; then
  CMD="$CMD -bench=$BENCH_PATTERN"
  if [ $MEMSTATS -eq 1 ]; then
    CMD="$CMD -benchmem"
  fi
fi

if [ $VERBOSE -eq 1 ]; then
  CMD="$CMD -v"
fi

echo "Running: $CMD"
eval "$CMD"
