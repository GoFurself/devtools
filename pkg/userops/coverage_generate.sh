#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e

# Navigate to the directory where your Go code is located (if necessary)
# cd /path/to/your/go/project

# Run tests and generate coverage profile
echo "Running tests and generating coverage profile..."
go test ./... -coverprofile=coverage.out

# Generate HTML coverage report
echo "Generating HTML coverage report..."
go tool cover -html=coverage.out -o coverage.html

echo "Coverage report generated: coverage.html"

# Open the coverage report in the default web browser
echo "Opening the coverage report in the default web browser..."
xdg-open coverage.html