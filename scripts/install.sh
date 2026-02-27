#!/usr/bin/env bash

set -e

REPO="vidya381/devcheck"
BIN="devcheck"
INSTALL_DIR="/usr/local/bin"

# detect OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$OS" in
  linux)   OS="linux" ;;
  darwin)  OS="darwin" ;;
  *)       echo "Unsupported OS: $OS"; exit 1 ;;
esac

# detect arch
ARCH=$(uname -m)
case "$ARCH" in
  x86_64)  ARCH="amd64" ;;
  aarch64 | arm64) ARCH="arm64" ;;
  *)       echo "Unsupported arch: $ARCH"; exit 1 ;;
esac

# get latest release tag
echo "Fetching latest release..."
TAG=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name"' | cut -d'"' -f4)

if [ -z "$TAG" ]; then
  echo "Could not find a release. Check https://github.com/$REPO/releases"
  exit 1
fi

FILENAME="${BIN}-${OS}-${ARCH}"
URL="https://github.com/$REPO/releases/download/$TAG/$FILENAME"

# download
TMP=$(mktemp)
echo "Downloading $BIN $TAG ($OS/$ARCH)..."
curl -fsSL "$URL" -o "$TMP"
chmod +x "$TMP"

# install
if [ -w "$INSTALL_DIR" ]; then
  mv "$TMP" "$INSTALL_DIR/$BIN"
else
  sudo mv "$TMP" "$INSTALL_DIR/$BIN"
fi

echo "Installed $BIN $TAG to $INSTALL_DIR/$BIN"
echo "Run: devcheck --help"
