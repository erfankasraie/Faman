#!/usr/bin/env bash
# faman installer — یک‌خطی، با گزینه فونت و RTL
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash
#   curl -fsSL ... | bash -s -- --with-rtl
#   ./scripts/install.sh --with-rtl --prefix /usr/local
set -euo pipefail

REPO_URL="${FAMAN_REPO_URL:-https://github.com/erfankasraie/Faman.git}"
PREFIX="${PREFIX:-/usr/local}"
WITH_RTL=0
SKIP_DEPS=0
YES=0

for arg in "$@"; do
  case "$arg" in
    --with-rtl|--rtl) WITH_RTL=1 ;;
    --skip-deps) SKIP_DEPS=1 ;;
    --yes|-y) YES=1 ;;
    --prefix=*) PREFIX="${arg#*=}" ;;
    --prefix) shift || true; PREFIX="${1:-/usr/local}" ;;
    -h|--help)
      cat <<'EOF'
faman installer

  --with-rtl    نصب فونت فارسی + locale UTF-8 + ابزارهای کمکی RTL
  --skip-deps   نصب وابستگی‌های سیستم را رد کن
  --prefix DIR  مسیر نصب (پیش‌فرض: /usr/local)
  -y, --yes     بدون تأیید تعاملی
EOF
      exit 0
      ;;
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
        if apt-cache show golang-go >/dev/null 2>&1; then
          need_sudo apt-get install -y golang-go
        fi
      fi
      # Go خیلی قدیمی؟ snap
      if command -v go >/dev/null 2>&1; then
        local major minor
        major="$(go env GOVERSION 2>/dev/null | sed -E 's/^go([0-9]+).*/\1/' || echo 0)"
        minor="$(go env GOVERSION 2>/dev/null | sed -E 's/^go[0-9]+\.([0-9]+).*/\1/' || echo 0)"
        if [[ "${major:-0}" -lt 1 ]] || { [[ "${major:-0}" -eq 1 && "${minor:-0}" -lt 21 ]]; }; then
          warn "نسخه Go پایین است؛ تلاش برای نصب از snap..."
          if command -v snap >/dev/null 2>&1; then
            need_sudo snap install go --classic || true
            export PATH="/snap/bin:$PATH"
          else
            warn "لطفاً Go 1.22+ را از https://go.dev/dl نصب کنید"
          fi
        fi
      else
        if command -v snap >/dev/null 2>&1; then
          need_sudo snap install go --classic
          export PATH="/snap/bin:$PATH"
        else
          die "go پیدا نشد. نصب کنید: sudo apt install golang-go  یا  https://go.dev/dl"
        fi
      fi
      ;;
    dnf)
      need_sudo dnf install -y git golang make curl
      ;;
    pacman)
      need_sudo pacman -Sy --needed --noconfirm git go make curl
      ;;
    zypper)
      need_sudo zypper install -y git go make curl
      ;;
    apk)
      need_sudo apk add --no-cache git go make curl
      ;;
    *)
      warn "توزیع ناشناخته — فرض می‌کنیم git و go از قبل نصب‌اند"
      ;;
  esac
  command -v git >/dev/null || die "git لازم است"
  command -v go >/dev/null || die "go لازم است"
  ok "وابستگی‌ها آماده‌اند (go $(go env GOVERSION 2>/dev/null || go version))"
}

install_faman() {
  local tmp
  tmp="$(mktemp -d)"
  trap 'rm -rf "$tmp"' EXIT

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
      # فونت‌ها
      need_sudo apt-get install -y fonts-noto-core fonts-noto-ui-core fonts-dejavu-core fonts-dejavu-mono 2>/dev/null || true
      # Vazirmatn اگر در ریپو باشد
      need_sudo apt-get install -y fonts-vazirmatn 2>/dev/null || \
        need_sudo apt-get install -y fonts-vazir 2>/dev/null || \
        warn "بسته Vazirmatn در apt نبود — از fallback Noto/DejaVu استفاده می‌شود"

      # locale
      need_sudo apt-get install -y locales 2>/dev/null || true
      if command -v locale-gen >/dev/null 2>&1; then
        need_sudo locale-gen en_US.UTF-8 fa_IR.UTF-8 2>/dev/null || need_sudo locale-gen en_US.UTF-8
      fi

      # ترمینال با پشتیبانی بهتر RTL (اختیاری)
      if [[ "$YES" -eq 1 ]] || [[ "$WITH_RTL" -eq 1 ]]; then
        need_sudo apt-get install -y mlterm 2>/dev/null || warn "mlterm نصب نشد (اختیاری)"
      fi

      # ابزار کمکی فونت
      need_sudo apt-get install -y fontconfig 2>/dev/null || true
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
      warn "نصب خودکار فونت برای این توزیع پشتیبانی نشده — docs/terminal-persian.md را ببینید"
      ;;
  esac

  # snippet شل
  local rc snippet
  snippet='# faman — UTF-8 و نمایش فارسی
export LANG="${LANG:-en_US.UTF-8}"
export LC_ALL="${LC_ALL:-en_US.UTF-8}"
# اگر حروف خراب بود: FAMAN_PLAIN=1 faman ls
'

  for rc in "$HOME/.bashrc" "$HOME/.zshrc"; do
    if [[ -f "$rc" ]] || [[ "$rc" == "$HOME/.bashrc" ]]; then
      if [[ -f "$rc" ]] && grep -q 'faman — UTF-8' "$rc" 2>/dev/null; then
        ok "قبلاً در $rc تنظیم شده"
      else
        if [[ ! -f "$rc" ]]; then
          touch "$rc"
        fi
        printf '\n%s\n' "$snippet" >> "$rc"
        ok "تنظیمات به $rc اضافه شد"
      fi
    fi
  done

  # راهنمای فونت GNOME Terminal
  mkdir -p "$HOME/.config/faman"
  cat > "$HOME/.config/faman/terminal-font-hint.txt" <<'EOF'
فونت پیشنهادی ترمینال برای فارسی:
  Vazirmatn  یا  Noto Sans Mono  یا  DejaVu Sans Mono

GNOME Terminal:
  ☰ → Preferences → Profile → Custom font → انتخاب فونت بالا

VS Code / Cursor (settings.json):
  "terminal.integrated.fontFamily": "Vazirmatn, DejaVu Sans Mono, monospace"

WezTerm (~/.config/wezterm/wezterm.lua):
  bidi_enabled = true
  font = wezterm.font_with_fallback({ "Vazirmatn", "DejaVu Sans Mono" })

mlterm (اگر نصب شد): از منوی اپلیکیشن‌ها باز کنید — برای RTL قوی‌تر است.
EOF
  ok "راهنمای فونت: ~/.config/faman/terminal-font-hint.txt"
}

main() {
  log "نصب faman"
  if [[ "$SKIP_DEPS" -eq 0 ]]; then
    install_deps
  fi
  install_faman
  if [[ "$WITH_RTL" -eq 1 ]]; then
    install_rtl_stack
  else
    warn "برای فونت و RTL:  bash scripts/install.sh --with-rtl"
    warn "یا:  curl -fsSL .../install.sh | bash -s -- --with-rtl"
  fi

  cat <<EOF

┌─────────────────────────────────────────────
│  نصب تمام شد
│
│  امتحان:
│    faman ls
│    faman search docker
│    FAMAN_PLAIN=1 faman echo
│
│  راهنمای فارسی در ترمینال:
│    https://github.com/erfankasraie/Faman/blob/main/docs/terminal-persian.md
└─────────────────────────────────────────────
EOF
}

main
