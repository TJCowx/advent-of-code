#!/bin/bash

# Usage: ./scripts/newGoDay.sh 2025 1

YEAR=$1
DAY=$2

if [[ -z "$YEAR" || -z "$DAY" ]]; then
  echo "Usage: $0 <year> <day>"
  exit 1
fi

# Pad day to 2 digits
DAY_PAD=$(printf "%02d" "$DAY")
DAY_NO_LEAD=$DAY  # original number without leading zeros

# Root directory (parent of scripts/)
ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"

YEAR_DIR="$ROOT_DIR/$YEAR"
DAY_DIR="$YEAR_DIR/day$DAY_PAD"

GO_TEMPLATE="$ROOT_DIR/templates/go/new_day.txt"
TEST_TEMPLATE="$ROOT_DIR/templates/go/new_day_test.txt"

GO_FILE="$DAY_DIR/day$DAY_PAD.go"
TEST_FILE="$DAY_DIR/day${DAY_PAD}_test.go"
INPUT_FILE="$DAY_DIR/input.txt"
TEST_INPUT_FILE="$DAY_DIR/test-input.txt"

# Create directories if needed
mkdir -p "$DAY_DIR"

# Copy templates
cp "$GO_TEMPLATE" "$GO_FILE"
cp "$TEST_TEMPLATE" "$TEST_FILE"

# Replace placeholders in the copied files
sed -i "s/<DAY>/$DAY_PAD/g" "$GO_FILE" "$TEST_FILE"
sed -i "s/<DAY_NO_LEAD>/$DAY_NO_LEAD/g" "$GO_FILE" "$TEST_FILE"

# Create empty input files
touch "$INPUT_FILE"
touch "$TEST_INPUT_FILE"

echo "Created day$DAY_PAD in $YEAR/"
echo "- $GO_FILE"
echo "- $TEST_FILE"
echo "- $INPUT_FILE"
echo "- $TEST_INPUT_FILE"
