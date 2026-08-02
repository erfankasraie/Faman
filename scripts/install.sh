#!/usr/bin/env bash
# faman installer — یک‌خطی، با گزینه فونت و RTL
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash
#   curl -fsSL ... | bash -s -- --with-rtl
#   curl -fsSL ... | bash -s -- --rtl-only
#   ./scripts/install.sh --with-rtl --prefix=/usr/local
set -euo pipefail

REPO_URL="${FAMAN_REPO_URL:-https://github.com/erfankasraie/Faman.git}"
PREFIX="${PREFIX:-/usr/local}"
WITH_RTL=0
RTL_ONLY=0
SKIP_DEPS=0

while [[ $# -gt 0 ]]; do
  case "$1" in
    --with-rtl|--rtl) WITH_RTL=1; shift ;;
    --rtl-only) RTL_ONLY=1; WITH_RTL=1; SKIP_DEPS=1; shift ;;
    --skip-deps) SKIP_DEPS=1; shift ;;
    --yes|-y) shift ;;
    --prefix=*) PREFIX="${1#*=}"; shift ;;
    --prefix) PREFIX="${2:-/usr/local}"; shift 2 ;;
    -h|--help)
      cat <<'EOF'
faman installer

  --with-rtl     نصب faman + فونت فارسی + locale UTF-8 + کمک RTL
  --rtl-only     فقط فونت / locale / RTL (بدون نصب faman)
  --skip-deps    نصب وابستگی‌های سیستم را رد کن
  --prefix=DIR   مسیر نصب (پیش‌فرض: /usr/local)
EOF
      exit 0
      ;;
    *) shift ;;
  esac
done

log()  { printf '\033[1;36m==>\033[0m %s\n' "$*"; }
ok()   { printf '\033[1;32m✓\033[0m %s\n' "$*"; }
warn() { printf '\033[1;33m!\033[0m %s\n' "$*"; }
die()  { printf '\033[1;31m✗\033[0m %s\n' "$*" >&2; exit 1; }

need_sudo() {
  if [[ "$(id -u)" -eq 0 ]]; then
    "$@"
  else
    sudo "$@"
  fi
}

detect_pm() {
  if command -v apt-get >/dev/null 2>&1; then echo apt
  elif command -v dnf >/dev/null 2>&1; then echo dnf
  elif command -v pacman >/dev/null 2>&1; then echo pacman
  elif command -v zypper >/dev/null 2>&1; then echo zypper
  elif command -v apk >/dev/null 2>&1; then echo apk
  else echo unknown
  fi
}

install_deps() {
  local pm
  pm="$(detect_pm)"
  log "مدیر بسته: $pm"
  case "$pm" in
    apt)
      need_sudo apt-get update -qq
      need_sudo apt-get install -y git make curl ca-certificates
      if ! command -v go >/dev/null 2>&1; then
        need_sudo apt-get install -y golang-go 2>/dev/null || true
      fi
      if command -v go >/dev/null 2>&1; then
        local ver major minor
        ver="$(go env GOVERSION 2>/dev/null || go version | awk '{print $3}')"
        major="$(echo "$ver" | sed -E 's/.*go([0-9]+).*/\1/')"
        minor="$(echo "$ver" | sed -E 's/.*go[0-9]+\.([0-9]+).*/\1/')"
        if [[ "${major:-0}" -eq 1 && "${minor:-0}" -lt 21 ]]; then
          warn "نسخه Go پایین است؛ تلاش برای snap..."
          if command -v snap >/dev/null 2>&1; then
            need_sudo snap install go --classic || true
            export PATH="/snap/bin:$PATH"
          fi
        fi
      else
        if command -v snap >/dev/null 2>&1; then
          need_sudo snap install go --classic
          export PATH="/snap/bin:$PATH"
        else
          die "go پیدا نشد. sudo apt install golang-go یا https://go.dev/dl"
        fi
      fi
      ;;
    dnf) need_sudo dnf install -y git golang make curl ;;
    pacman) need_sudo pacman -Sy --needed --noconfirm git go make curl ;;
    zypper) need_sudo zypper install -y git go make curl ;;
    apk) need_sudo apk add --no-cache git go make curl ;;
    *) warn "توزیع ناشناخته — فرض بر نصب بودن git/go" ;;
  esac
  command -v git >/dev/null || die "git لازم است"
  command -v go >/dev/null || die "go لازم است"
  ok "وابستگی‌ها آماده (go $(go env GOVERSION 2>/dev/null || go version))"
}

install_faman() {
  local tmp
  tmp="$(mktemp -d)"
  # shellcheck disable=SC2064
  trap "rm -rf '$tmp'" EXIT

  log "کلون مخزن..."
  git clone --depth 1 "$REPO_URL" "$tmp/faman"
  cd "$tmp/faman"

  log "ساخت باینری..."
  go build -ldflags="-s -w" -o faman ./cmd/faman

  log "نصب در $PREFIX ..."
  need_sudo install -d "$PREFIX/bin"
  need_sudo install -m 755 faman "$PREFIX/bin/faman"
  need_sudo install -d "$PREFIX/share/faman"
  need_sudo rm -rf "$PREFIX/share/faman/pages"
  need_sudo cp -a pages "$PREFIX/share/faman/"

  ok "faman نصب شد: $PREFIX/bin/faman"
  "$PREFIX/bin/faman" version || true
}

install_rtl_stack() {
  log "نصب پشته نمایش فارسی / RTL..."
  local pm
  pm="$(detect_pm)"

  case "$pm" in
    apt)
      need_sudo apt-get update -qq
      need_sudo apt-get install -y fonts-noto-core fonts-noto-ui-core fonts-dejavu-core fonts-dejavu-mono fontconfig locales 2>/dev/null || true
      need_sudo apt-get install -y fonts-vazirmatn 2>/dev/null || \
        need_sudo apt-get install -y fonts-vazir 2>/dev/null || \
        warn "Vazirmatn در apt نبود — Noto/DejaVu کافی است"
      if command -v locale-gen >/dev/null 2>&1; then
        need_sudo locale-gen en_US.UTF-8 fa_IR.UTF-8 2>/dev/null || need_sudo locale-gen en_US.UTF-8
      fi
      need_sudo apt-get install -y mlterm 2>/dev/null || warn "mlterm نصب نشد (اختیاری)"
      fc-cache -f >/dev/null 2>&1 || true
      ;;
    dnf)
      need_sudo dnf install -y google-noto-sans-fonts google-noto-sans-arabic-fonts dejavu-sans-mono-fonts fontconfig 2>/dev/null || true
      need_sudo dnf install -y mlterm 2>/dev/null || true
      fc-cache -f >/dev/null 2>&1 || true
      ;;
    pacman)
      need_sudo pacman -Sy --needed --noconfirm noto-fonts noto-fonts-extra ttf-dejavu fontconfig 2>/dev/null || true
      need_sudo pacman -Sy --needed --noconfirm mlterm 2>/dev/null || true
      fc-cache -f >/dev/null 2>&1 || true
      ;;
    *)
      warn "نصب خودکار فونت برای این توزیع نیست — docs/terminal-persian.md"
      ;;
  esac

  local rc snippet
  snippet='# faman — UTF-8 و نمایش فارسی
export LANG="${LANG:-en_US.UTF-8}"
export LC_ALL="${LC_ALL:-en_US.UTF-8}"
# اگر حروف خراب بود: FAMAN_PLAIN=1 faman ls
'

  for rc in "$HOME/.bashrc" "$HOME/.zshrc"; do
    if [[ -f "$rc" ]] && grep -q 'faman — UTF-8' "$rc" 2>/dev/null; then
      ok "قبلاً در $rc تنظیم شده"
      continue
    fi
    if [[ "$rc" == "$HOME/.bashrc" ]] || [[ -f "$rc" ]]; then
      [[ -f "$rc" ]] || touch "$rc"
      printf '\n%s\n' "$snippet" >> "$rc"
      ok "تنظیمات به $rc اضافه شد"
    fi
  done

  mkdir -p "$HOME/.config/faman"
  cat > "$HOME/.config/faman/terminal-font-hint.txt" <<'EOF'
فونت پیشنهادی ترمینال برای فارسی:
  Vazirmatn  یا  Noto Sans Mono  یا  DejaVu Sans Mono

GNOME Terminal (Ubuntu):
  ☰ → Preferences → Profile → Custom font → فونت بالا

VS Code / Cursor (settings.json):
  "terminal.integrated.fontFamily": "Vazirmatn, DejaVu Sans Mono, monospace"

WezTerm (~/.config/wezterm/wezterm.lua):
  return {
    bidi_enabled = true,
    font = require("wezterm").font_with_fallback({ "Vazirmatn", "DejaVu Sans Mono" }),
  }

mlterm (اگر نصب شد): از منوی اپلیکیشن‌ها — RTL قوی‌تر.
EOF
  ok "راهنمای فونت: ~/.config/faman/terminal-font-hint.txt"
}

main() {
  if [[ "$RTL_ONLY" -eq 1 ]]; then
    log "فقط پشته RTL / فونت"
    install_rtl_stack
  else
    log "نصب faman"
    if [[ "$SKIP_DEPS" -eq 0 ]]; then
      install_deps
    fi
    install_faman
    if [[ "$WITH_RTL" -eq 1 ]]; then
      install_rtl_stack
    else
      warn "فونت و RTL: curl ... | bash -s -- --with-rtl"
    fi
  fi

  cat <<EOF

┌─────────────────────────────────────────────
│  تمام شد
│
│  امتحان:
│    faman ls
│    faman search docker
│    FAMAN_PLAIN=1 faman echo
│
│  یک‌بار فونت ترمینال را دستی تنظیم کنید:
│    cat ~/.config/faman/terminal-font-hint.txt
│
│  docs: https://github.com/erfankasraie/Faman/blob/main/docs/install.md
└─────────────────────────────────────────────
EOF
}

main
