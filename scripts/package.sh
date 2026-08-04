#!/usr/bin/env bash
# Build release artifacts: tar.gz, zip (.exe), and optional .deb
#
#   ./scripts/package.sh                 # version from internal/app or VERSION=
#   VERSION=0.1.4-pre ./scripts/package.sh
#   ./scripts/package.sh --deb-only
#
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

VERSION="${VERSION:-}"
if [[ -z "$VERSION" ]]; then
  VERSION="$(grep -E '^\s*version = ' internal/app/root.go | head -1 | sed -E 's/.*"([^"]+)".*/\1/')"
fi
VERSION="${VERSION#v}"
LDFLAGS="-s -w -X github.com/faman-project/faman/internal/app.version=${VERSION}"
OUT="${OUT:-dist}"
mkdir -p "$OUT"

DO_ARCHIVES=1
DO_DEB=1
for a in "$@"; do
  case "$a" in
    --deb-only) DO_ARCHIVES=0; DO_DEB=1 ;;
    --no-deb) DO_DEB=0 ;;
    --help|-h)
      echo "Usage: VERSION=x.y.z $0 [--deb-only] [--no-deb]"
      exit 0
      ;;
  esac
done

build() {
  local goos=$1 goarch=$2 out=$3
  echo "==> $out"
  CGO_ENABLED=0 GOOS="$goos" GOARCH="$goarch" go build -trimpath -ldflags "$LDFLAGS" -o "$OUT/$out" ./cmd/faman
}

pack_unix() {
  local name=$1 bin=$2
  local dir="faman-${VERSION}-${name}"
  rm -rf "$OUT/pkg/$dir"
  mkdir -p "$OUT/pkg/$dir"
  cp "$OUT/$bin" "$OUT/pkg/$dir/faman"
  cp -a pages "$OUT/pkg/$dir/"
  cp README.md LICENSE "$OUT/pkg/$dir/" 2>/dev/null || true
  tar -C "$OUT/pkg" -czf "$OUT/${dir}.tar.gz" "$dir"
  echo "    $OUT/${dir}.tar.gz"
}

if [[ "$DO_ARCHIVES" -eq 1 ]]; then
  build linux amd64 faman-linux-amd64
  build linux arm64 faman-linux-arm64
  build windows amd64 faman-windows-amd64.exe
  build darwin amd64 faman-darwin-amd64
  build darwin arm64 faman-darwin-arm64

  pack_unix linux-amd64 faman-linux-amd64
  pack_unix linux-arm64 faman-linux-arm64
  pack_unix darwin-amd64 faman-darwin-amd64
  pack_unix darwin-arm64 faman-darwin-arm64

  dir="faman-${VERSION}-windows-amd64"
  rm -rf "$OUT/pkg/$dir"
  mkdir -p "$OUT/pkg/$dir"
  cp "$OUT/faman-windows-amd64.exe" "$OUT/pkg/$dir/faman.exe"
  cp -a pages "$OUT/pkg/$dir/"
  cp README.md LICENSE "$OUT/pkg/$dir/" 2>/dev/null || true
  (cd "$OUT/pkg" && zip -qr "../${dir}.zip" "$dir")
  echo "    $OUT/${dir}.zip"
fi

if [[ "$DO_DEB" -eq 1 ]]; then
  if ! command -v dpkg-deb >/dev/null 2>&1; then
    echo "! dpkg-deb not found — skip .deb (install dpkg-dev on Debian/Ubuntu)"
  else
    if [[ ! -f "$OUT/faman-linux-amd64" ]]; then
      build linux amd64 faman-linux-amd64
    fi
    debroot="$OUT/deb-root"
    rm -rf "$debroot"
    mkdir -p "$debroot/usr/local/bin" "$debroot/usr/local/share/faman" "$debroot/DEBIAN"
    cp "$OUT/faman-linux-amd64" "$debroot/usr/local/bin/faman"
    chmod 755 "$debroot/usr/local/bin/faman"
    cp -a pages "$debroot/usr/local/share/faman/"
    # Debian version: no leading v; pre → ~pre for ordering
    debver="${VERSION//-pre/~pre}"
    cat > "$debroot/DEBIAN/control" <<EOF
Package: faman
Version: ${debver}
Section: utils
Priority: optional
Architecture: amd64
Maintainer: Erfan Kasraie <erfankasraiee@gmail.com>
Description: Persian manual pages for Linux commands
 faman teaches common Linux commands in Persian from the terminal.
EOF
    deb="$OUT/faman_${VERSION}_amd64.deb"
    dpkg-deb --build "$debroot" "$deb"
    echo "    $deb"
  fi
fi

(cd "$OUT" && sha256sum faman-*.tar.gz faman-*.zip faman_*.deb 2>/dev/null > SHA256SUMS || true)
echo "Done. Artifacts in $OUT/"
ls -lh "$OUT"/faman-* "$OUT"/SHA256SUMS 2>/dev/null || true
