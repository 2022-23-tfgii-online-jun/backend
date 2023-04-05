#!/bin/bash

# Define colors
RED="\033[31m"
GREEN="\033[32m"
RESET="\033[0m"

echo "Running tests"

# Save the test output in a variable
output=$(go test -v ./internal/pkg/service/... 2>&1)

# Process the output and color lines according to the result
while IFS= read -r line; do
    if [[ $line == *"--- PASS"* ]]; then
        echo -e "${GREEN}[OK]${RESET} $line"
    elif [[ $line == *"--- FAIL"* ]]; then
        echo -e "${RED}[ERROR]${RESET} $line"
    fi
done <<< "$output"

# Check test status and set exit code
if echo "$output" | grep -q "FAIL"; then
    result=1
    echo "There were errors in the tests."
else
    result=0
    echo "All tests were successful."
fi

exit $result
