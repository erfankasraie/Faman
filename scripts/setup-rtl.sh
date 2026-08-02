#!/usr/bin/env bash
# فقط فونت فارسی + locale + ابزار RTL (بدون نصب مجدد faman)
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/setup-rtl.sh | bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." 2>/dev/null && pwd || true)"
if [[ -n "${ROOT:-}" && -f "$ROOT/scripts/install.sh" ]]; then
  exec bash "$ROOT/scripts/install.sh" --rtl-only "$@"
fi

TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh -o "$TMP/install.sh"
bash "$TMP/install.sh" --rtl-only "$@"
