#!/usr/bin/env bash
# Validate Persian man pages under pages/fa
# Usage: ./scripts/check-pages.sh [--strict]
# Exit 1 if any page fails required checks.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
PAGES="${FAMAN_PAGES_DIR:-$ROOT/pages/fa}"
STRICT=0
[[ "${1:-}" == "--strict" ]] && STRICT=1

if [[ ! -d "$PAGES" ]]; then
  echo "error: pages dir not found: $PAGES" >&2
  exit 1
fi

required_fm=(title: category: difficulty:)
required_h=(
  "# Introduction"
  "# Syntax"
  "# Options"
  "# Examples"
  "# Common mistakes"
  "# Tips"
  "# Related commands"
)
MIN_BYTES=280

errors=0
warnings=0
count=0

while IFS= read -r -d '' f; do
  count=$((count + 1))
  base="$(basename "$f")"
  size="$(wc -c < "$f" | tr -d ' ')"

  if [[ "$size" -lt "$MIN_BYTES" ]]; then
    echo "FAIL  $base  too short (${size} < ${MIN_BYTES} bytes)"
    errors=$((errors + 1))
  fi

  for key in "${required_fm[@]}"; do
    if ! grep -q "^${key}" "$f"; then
      echo "FAIL  $base  missing front matter: $key"
      errors=$((errors + 1))
    fi
  done

  if ! grep -q '^---$' "$f"; then
    echo "FAIL  $base  missing YAML front matter delimiter ---"
    errors=$((errors + 1))
  fi

  for h in "${required_h[@]}"; do
    if ! grep -qFx "$h" "$f"; then
      echo "FAIL  $base  missing section: $h"
      errors=$((errors + 1))
    fi
  done

  # soft checks
  if ! grep -q '^keywords:' "$f"; then
    echo "WARN  $base  no keywords:"
    warnings=$((warnings + 1))
  fi
done < <(find "$PAGES" -maxdepth 1 -type f -name '*.md' -print0 | sort -z)

echo
echo "checked $count pages  errors=$errors  warnings=$warnings"

if [[ "$errors" -gt 0 ]]; then
  exit 1
fi
if [[ "$STRICT" -eq 1 && "$warnings" -gt 0 ]]; then
  echo "strict mode: treating warnings as errors"
  exit 1
fi
exit 0
