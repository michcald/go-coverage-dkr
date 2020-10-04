#!/bin/sh

# This script fails if the code coverage is less than the limit and generates a badge in the readme

set -e

if [ ! -f "$COVERAGE_FILE" ]; then
    echo "coverage file -$COVERAGE_FILE- does not exist"
    exit 1
fi

if [ ! -f "$README_FILE" ]; then
    echo "coverage file -$README_FILE- does not exist"
    exit 1
fi

go tool cover -func=${COVERAGE_FILE} -o /devtools-coverage.out

go run /coverage.go
