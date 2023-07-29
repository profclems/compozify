#!/bin/sh
# Usage: [sudo] [BINDIR=/usr/local/bin] ./install.sh [<BINDIR>]
#
# Example:
#     1. sudo ./install.sh /usr/local/bin
#     2. sudo ./install.sh /usr/bin
#     3. ./install.sh $HOME/usr/bin
#     4. BINDIR=$HOME/usr/bin ./install.sh
#
# Default BINDIR=/usr/bin

set -euf

if [ -n "${DEBUG-}" ]; then
    set -x
fi

: "${BINDIR:=/usr/bin}"

if [ $# -gt 0 ]; then
  BINDIR=$1
fi

_can_install() {
  if [ ! -d "${BINDIR}" ]; then
    mkdir -p "${BINDIR}" 2> /dev/null
  fi
  [ -d "${BINDIR}" ] && [ -w "${BINDIR}" ]
}

if ! _can_install && [ "$(id -u)" != 0 ]; then
  printf "Run script as sudo\n"
  exit 1
fi

if ! _can_install; then
  printf -- "Can't install to %s\n" "${BINDIR}"
  exit 1
fi

machine=$(uname -m)

case ${machine} in
    x86_64)
        machine="amd64"
        ;;
    aarch64)
        machine="arm64"
        ;;
esac

case $(uname -s) in
    Linux)
        os="linux"
        format="tar.gz"
        ;;
    Darwin)
        os="macOS"
        format="zip"
        ;;
    *)
        printf "OS not supported\n"
        exit 1
        ;;
esac

printf "Fetching latest version\n"
latest="$(curl -sL 'https://api.github.com/repos/profclems/compozify/releases/latest' | grep 'tag_name' | grep --only-matching 'v[0-9\.]\+' | cut -c 2-)"
tempFolder="/tmp/compozify_v${latest}"
filename="compozify_${latest}_${os}_${machine}"

printf -- "Found version %s\n" "${latest}"

printf -- "Creating temp folder %s\n" "${tempFolder}"
mkdir -p "${tempFolder}" 2> /dev/null

printf -- "Downloading %s.%s\n" "${filename}" "${format}"
if [ "${format}" = "tar.gz" ]; then
  curl -sL "https://github.com/profclems/compozify/releases/download/v${latest}/${filename}.tar.gz" | tar -C "${tempFolder}" -xvzf -
else
  curl -sL "https://github.com/profclems/compozify/releases/download/v${latest}/${filename}.zip" | bsdtar -xvf - -C "${tempFolder}"
fi

srcDir="${tempFolder}/${filename}"
printf -- "Installing %s...\n" "${srcDir}/bin/compozify"
install -m755 "${srcDir}/bin/compozify" "${BINDIR}/compozify"

printf -- "Installing manpages...\n"
install -d /usr/local/man/man1/
install -m644 "${srcDir}/share/man/man1"/compozify.1 /usr/local/man/man1/
install -m644 "${srcDir}/share/man/man1"/compozify-convert.1 /usr/local/man/man1/


printf "Cleaning up temp files\n"
rm -rf "${tempFolder}"

printf -- "Successfully installed compozify into %s/\n" "${BINDIR}"