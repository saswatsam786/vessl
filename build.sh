#!/bin/bash

# Build script for Vessl Docker CLI Tool

echo "🚀 Building Vessl Docker CLI Tool..."

# Set version
VERSION="1.0.0"
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')

# Create build directory
mkdir -p build

# Build for different platforms
echo "📦 Building for multiple platforms..."

# Linux AMD64
echo "Building for Linux AMD64..."
GOOS=linux GOARCH=amd64 go build -ldflags="-X main.version=$VERSION -X main.buildTime=$BUILD_TIME" -o build/vessl-linux-amd64

# macOS AMD64
echo "Building for macOS AMD64..."
GOOS=darwin GOARCH=amd64 go build -ldflags="-X main.version=$VERSION -X main.buildTime=$BUILD_TIME" -o build/vessl-darwin-amd64

# macOS ARM64
echo "Building for macOS ARM64..."
GOOS=darwin GOARCH=arm64 go build -ldflags="-X main.version=$VERSION -X main.buildTime=$BUILD_TIME" -o build/vessl-darwin-arm64

# Windows AMD64
echo "Building for Windows AMD64..."
GOOS=windows GOARCH=amd64 go build -ldflags="-X main.version=$VERSION -X main.buildTime=$BUILD_TIME" -o build/vessl-windows-amd64.exe

# Make Linux and macOS binaries executable
chmod +x build/vessl-linux-amd64
chmod +x build/vessl-darwin-amd64
chmod +x build/vessl-darwin-arm64

echo "✅ Build complete!"
echo "📁 Binaries created in build/ directory:"
ls -la build/

echo ""
echo "🎯 To install locally:"
echo "sudo cp build/vessl-$(go env GOOS)-$(go env GOARCH) /usr/local/bin/vessl"
echo ""
echo "📦 To create a release, zip the build/ directory and upload to GitHub" 