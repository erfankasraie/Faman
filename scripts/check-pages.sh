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

# Exact front-matter keys (start of line)
required_fm=(title: category: difficulty:)

# Section patterns: heading line must match (allow suffixes like "# Options / Commands")
required_patterns=(
  '^# Introduction([[:space:]].*)?$'
  '^# Syntax([[:space:]].*)?$'
  '^# Options([[:space:]/].*)?$'
  '^# Examples([[:space:]].*)?$'
  '^# Common mistakes([[:space:]].*)?$'
  '^# Tips([[:space:]].*)?$'
  '^# Related commands([[:space:]].*)?$'
)
section_names=(
  "Introduction"
  "Syntax"
  "Options"
  "Examples"
  "Common mistakes"
  "Tips"
  "Related commands"
)
MIN_BYTES=280

errors=0
warnings=0
count=0

shopt -s nullglob
for f in "$PAGES"/*.md; do
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

  i=0
  for pat in "${required_patterns[@]}"; do
    if ! grep -qE "$pat" "$f"; then
      echo "FAIL  $base  missing section: ${section_names[$i]}"
      errors=$((errors + 1))
    fi
    i=$((i + 1))
  done

  if ! grep -q '^keywords:' "$f"; then
    echo "WARN  $base  no keywords:"
    warnings=$((warnings + 1))
  fi
done

echo
echo "checked $count pages  errors=$errors  warnings=$warnings"

if [[ "$count" -eq 0 ]]; then
  echo "error: no markdown pages found in $PAGES" >&2
  exit 1
fi
if [[ "$errors" -gt 0 ]]; then
  exit 1
fi
if [[ "$STRICT" -eq 1 && "$warnings" -gt 0 ]]; then
  echo "strict mode: treating warnings as errors"
  exit 1
fi
exit 0
