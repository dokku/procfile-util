#!/usr/bin/env sh
# Install procfile-util as a standalone binary on the local machine.
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/dokku/procfile-util/master/install.sh | sh
#   VERSION=v0.20.4 curl -fsSL https://raw.githubusercontent.com/dokku/procfile-util/master/install.sh | sh
#   INSTALL_DIR=$HOME/.local/bin curl -fsSL https://raw.githubusercontent.com/dokku/procfile-util/master/install.sh | sh
#
# Environment variables:
#   VERSION      Release tag to install (defaults to the latest release).
#   INSTALL_DIR  Override destination directory (defaults to /usr/local/bin).

set -eu

NAME="procfile-util"
REPO="dokku/procfile-util"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

os="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$os" in
  linux | darwin) ;;
  *)
    echo "error: unsupported OS: $os" >&2
    exit 1
    ;;
esac

arch="$(uname -m)"
case "$arch" in
  x86_64 | amd64) arch="amd64" ;;
  aarch64 | arm64) arch="arm64" ;;
  *)
    echo "error: unsupported architecture: $arch" >&2
    exit 1
    ;;
esac

if [ -z "${VERSION:-}" ]; then
  VERSION="$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" \
    | sed -n 's/.*"tag_name": "\(.*\)".*/\1/p' | head -n1)"
fi

if [ -z "$VERSION" ]; then
  echo "error: could not determine latest version; set VERSION explicitly" >&2
  exit 1
fi

# Release assets are raw binaries named like procfile-util-<os>-<arch>.
asset="${NAME}-${os}-${arch}"
url="https://github.com/${REPO}/releases/download/${VERSION}/${asset}"

tmpdir="$(mktemp -d)"
trap 'rm -rf "$tmpdir"' EXIT INT TERM

echo "downloading ${url}"
curl -fsSL "$url" -o "${tmpdir}/${asset}"

# Use sudo only when the install directory is not writable by the current user.
sudo=""
if [ ! -w "$INSTALL_DIR" ] && [ -d "$INSTALL_DIR" ]; then
  sudo="sudo"
elif [ ! -d "$INSTALL_DIR" ] && [ ! -w "$(dirname "$INSTALL_DIR")" ]; then
  sudo="sudo"
fi

$sudo mkdir -p "$INSTALL_DIR"
$sudo install -m 0755 "${tmpdir}/${asset}" "${INSTALL_DIR}/${NAME}"

echo "installed ${NAME} ${VERSION} to ${INSTALL_DIR}/${NAME}"
echo "try: ${NAME} version"
