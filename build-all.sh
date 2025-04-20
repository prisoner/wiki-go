#!/bin/bash
echo "Building Wiki-Go for multiple platforms..."

BUILD_DIR="build"
mkdir -p "$BUILD_DIR"

find "$BUILD_DIR" -mindepth 1 ! -name ".gitkeep" ! -path "$BUILD_DIR/data" ! -path "$BUILD_DIR/data/*" -delete

# Get version from git tag, fallback to "dev" if not available
VERSION=$(git describe --tags --always 2>/dev/null || echo "dev")
echo "Building version: $VERSION"

# Set ldflags for version
LDFLAGS="-X 'wiki-go/internal/version.Version=$VERSION'"

echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -ldflags "$LDFLAGS" -o $BUILD_DIR/wiki-go-linux-amd64 .

echo "Build complete! Packages are available in the '$BUILD_DIR' directory."
