#!/usr/bin/env bash
# نصب آسان faman — یک خط، بدون sudo (پیش‌فرض ~/.local)
#
#   curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash
#
# با فونت و RTL:
#   curl -fsSL .../get.sh | bash -s -- --rtl
#   یا:  WITH_RTL=1 curl -fsSL .../get.sh | bash
#
# سیستم‌واید (sudo):
#   curl -fsSL .../get.sh | bash -s -- --system
#
set -euo pipefail

RAW="https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh"
ARGS=(--user)

# env shortcuts
if [[ "${WITH_RTL:-0}" == "1" || "${WITH_RTL:-}" == "true" ]]; then
  ARGS+=(--with-rtl)
fi
if [[ "${PLAIN:-0}" == "1" || "${FAMAN_PLAIN_DEFAULT:-0}" == "1" ]]; then
  ARGS+=(--plain-default)
fi

while [[ $# -gt 0 ]]; do
  case "$1" in
    --rtl|--with-rtl|-r) ARGS+=(--with-rtl); shift ;;
    --system|--global) ARGS=(--with-rtl); shift ;; # drop --user; system prefix
    --system-only) ARGS=(); shift ;;
    --plain) ARGS+=(--plain-default); shift ;;
    --help|-h)
      cat <<'EOF'
faman get.sh — نصب آسان

  curl -fsSL .../scripts/get.sh | bash
  curl -fsSL .../scripts/get.sh | bash -s -- --rtl
  curl -fsSL .../scripts/get.sh | bash -s -- --system

پیش‌فرض: --user → ~/.local/bin (بدون sudo برای باینری)
EOF
      exit 0
      ;;
    *)
      # pass through to install.sh
      ARGS+=("$1")
      shift
      ;;
  esac
done

# if user asked --system, ARGS was reset without --user
# re-apply rtl if they used WITH_RTL with --system
if [[ "${ARGS[*]}" != *--user* ]] && [[ "${ARGS[*]}" != *--with-rtl* ]]; then
  if [[ "${WITH_RTL:-0}" == "1" ]]; then
    ARGS+=(--with-rtl)
  fi
fi

echo "==> faman easy install: install.sh ${ARGS[*]}"
curl -fsSL "$RAW" | bash -s -- "${ARGS[@]}"
