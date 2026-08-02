#!/usr/bin/env bash
# faman installer
#
# Quick:
#   curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash
#   curl -fsSL ... | bash -s -- --with-rtl
#
# Advanced examples:
#   bash scripts/install.sh --user --with-rtl --plain-default
#   bash scripts/install.sh --prefix=/opt/faman --branch=main
#   bash scripts/install.sh --from-dir=. --skip-deps
#   bash scripts/install.sh --uninstall
#   bash scripts/install.sh --rtl-only --no-mlterm --locale=fa
#
# Config file (optional): ~/.config/faman/install.env
#   PREFIX=/usr/local
#   WITH_RTL=1
#   PLAIN_DEFAULT=1
set -euo pipefail

# ── defaults (overridden by config file, then CLI, then env) ───────────
REPO_URL="${FAMAN_REPO_URL:-https://github.com/erfankasraie/Faman.git}"
BRANCH="${FAMAN_BRANCH:-main}"
PREFIX="${PREFIX:-/usr/local}"
WITH_RTL=0
RTL_ONLY=0
SKIP_DEPS=0
USER_INSTALL=0
NO_MLTERM=0
NO_SHELL_RC=0
PLAIN_DEFAULT=0
LOCALE_PREF="en"   # en | fa
FROM_DIR=""
UNINSTALL=0
DRY_RUN=0
VERBOSE=0

# load optional config
CONFIG_FILE="${FAMAN_INSTALL_CONFIG:-$HOME/.config/faman/install.env}"
if [[ -f "$CONFIG_FILE" ]]; then
  # shellcheck disable=SC1090
  set -a
  # only allow KEY=VALUE lines
  while IFS= read -r line || [[ -n "$line" ]]; do
    [[ "$line" =~ ^[[:space:]]*# ]] && continue
    [[ -z "${line// /}" ]] && continue
    if [[ "$line" =~ ^[A-Za-z_][A-Za-z0-9_]*= ]]; then
      eval "export $line"
    fi
  done < "$CONFIG_FILE"
  set +a
  WITH_RTL="${WITH_RTL:-0}"
  PLAIN_DEFAULT="${PLAIN_DEFAULT:-0}"
  NO_MLTERM="${NO_MLTERM:-0}"
  NO_SHELL_RC="${NO_SHELL_RC:-0}"
  USER_INSTALL="${USER_INSTALL:-0}"
  LOCALE_PREF="${LOCALE_PREF:-en}"
  PREFIX="${PREFIX:-/usr/local}"
  BRANCH="${BRANCH:-main}"
  REPO_URL="${REPO_URL:-$REPO_URL}"
fi

usage() {
  cat <<'EOF'
faman installer

Basic:
  --with-rtl          faman + fonts + locale + RTL helpers
  --rtl-only          only fonts/locale/RTL (no faman binary)
  --skip-deps         do not install system packages (git/go/...)

Paths & source:
  --prefix=DIR        install root (default: /usr/local)
  --user              install to ~/.local (no sudo for binary/pages)
  --branch=NAME       git branch/tag (default: main)
  --repo=URL          git remote (default: erfankasraie/Faman)
  --from-dir=PATH     build from existing checkout instead of clone

RTL / shell:
  --no-mlterm         skip mlterm package
  --no-shell-rc       do not modify ~/.bashrc or ~/.zshrc
  --locale=en|fa      preferred UTF-8 locale (default: en)
  --plain-default     export FAMAN_PLAIN=1 in shell rc

Other:
  --uninstall         remove faman binary + share pages from PREFIX
  --dry-run           print actions only
  --verbose           more logs
  -h, --help          this help

Environment:
  FAMAN_REPO_URL  FAMAN_BRANCH  PREFIX  FAMAN_INSTALL_CONFIG

Config file (optional):
  ~/.config/faman/install.env
  Example:
    PREFIX=$HOME/.local
    USER_INSTALL=1
    WITH_RTL=1
    PLAIN_DEFAULT=1
    LOCALE_PREF=fa
    NO_MLTERM=1
EOF
}

while [[ $# -gt 0 ]]; do
  case "$1" in
    --with-rtl|--rtl) WITH_RTL=1; shift ;;
    --rtl-only) RTL_ONLY=1; WITH_RTL=1; SKIP_DEPS=1; shift ;;
    --skip-deps) SKIP_DEPS=1; shift ;;
    --user) USER_INSTALL=1; PREFIX="${HOME}/.local"; shift ;;
    --no-mlterm) NO_MLTERM=1; shift ;;
    --no-shell-rc) NO_SHELL_RC=1; shift ;;
    --plain-default) PLAIN_DEFAULT=1; shift ;;
    --uninstall) UNINSTALL=1; shift ;;
    --dry-run) DRY_RUN=1; shift ;;
    --verbose|-v) VERBOSE=1; shift ;;
    --locale=*) LOCALE_PREF="${1#*=}"; shift ;;
    --locale) LOCALE_PREF="${2:-en}"; shift 2 ;;
    --prefix=*) PREFIX="${1#*=}"; shift ;;
    --prefix) PREFIX="${2:-/usr/local}"; shift 2 ;;
    --branch=*) BRANCH="${1#*=}"; shift ;;
    --branch) BRANCH="${2:-main}"; shift 2 ;;
    --repo=*) REPO_URL="${1#*=}"; shift ;;
    --repo) REPO_URL="${2:-}"; shift 2 ;;
    --from-dir=*) FROM_DIR="${1#*=}"; shift ;;
    --from-dir) FROM_DIR="${2:-}"; shift 2 ;;
    -h|--help) usage; exit 0 ;;
    *)
      echo "unknown option: $1" >&2
      usage >&2
      exit 2
      ;;
  esac
done

log()  { printf '\033[1;36m==>\033[0m %s\n' "$*"; }
ok()   { printf '\033[1;32m✓\033[0m %s\n' "$*"; }
warn() { printf '\033[1;33m!\033[0m %s\n' "$*"; }
die()  { printf '\033[1;31m✗\033[0m %s\n' "$*" >&2; exit 1; }
dbg()  { [[ "$VERBOSE" -eq 1 ]] && printf '\033[2m  · %s\033[0m\n' "$*" || true; }

run() {
  if [[ "$DRY_RUN" -eq 1 ]]; then
    printf '\033[2m[dry-run]\033[0m %s\n' "$*"
    return 0
  fi
  dbg "run: $*"
  "$@"
}

need_sudo() {
  if [[ "$USER_INSTALL" -eq 1 ]]; then
    # user mode: only sudo for system packages, not for PREFIX under $HOME
    "$@"
    return
  fi
  if [[ "$(id -u)" -eq 0 ]]; then
    "$@"
  else
    sudo "$@"
  fi
}

sys_pkg() {
  # always may need root for packages
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

uninstall_faman() {
  log "حذف faman از $PREFIX"
  run need_sudo rm -f "$PREFIX/bin/faman"
  run need_sudo rm -rf "$PREFIX/share/faman"
  ok "حذف شد (shell rc دست نخورده)"
}

install_deps() {
  local pm
  pm="$(detect_pm)"
  log "مدیر بسته: $pm"
  case "$pm" in
    apt)
      run sys_pkg apt-get update -qq
      run sys_pkg apt-get install -y git make curl ca-certificates
      if ! command -v go >/dev/null 2>&1; then
        run sys_pkg apt-get install -y golang-go 2>/dev/null || true
      fi
      if command -v go >/dev/null 2>&1; then
        local ver major minor
        ver="$(go env GOVERSION 2>/dev/null || go version | awk '{print $3}')"
        major="$(echo "$ver" | sed -E 's/.*go([0-9]+).*/\1/')"
        minor="$(echo "$ver" | sed -E 's/.*go[0-9]+\.([0-9]+).*/\1/')"
        if [[ "${major:-0}" -eq 1 && "${minor:-0}" -lt 21 ]]; then
          warn "Go قدیمی ($ver)؛ تلاش snap..."
          if command -v snap >/dev/null 2>&1; then
            run sys_pkg snap install go --classic || true
            export PATH="/snap/bin:$PATH"
          fi
        fi
      else
        if command -v snap >/dev/null 2>&1; then
          run sys_pkg snap install go --classic
          export PATH="/snap/bin:$PATH"
        else
          die "go پیدا نشد"
        fi
      fi
      ;;
    dnf) run sys_pkg dnf install -y git golang make curl ;;
    pacman) run sys_pkg pacman -Sy --needed --noconfirm git go make curl ;;
    zypper) run sys_pkg zypper install -y git go make curl ;;
    apk) run sys_pkg apk add --no-cache git go make curl ;;
    *) warn "توزیع ناشناخته — git/go باید از قبل باشند" ;;
  esac
  if [[ "$DRY_RUN" -eq 0 ]]; then
    command -v git >/dev/null || die "git لازم است"
    command -v go >/dev/null || die "go لازم است"
    ok "deps OK ($(go env GOVERSION 2>/dev/null || go version))"
  fi
}

install_faman() {
  local src tmp
  tmp=""
  if [[ -n "$FROM_DIR" ]]; then
    src="$(cd "$FROM_DIR" && pwd)"
    [[ -d "$src/cmd/faman" ]] || die "--from-dir نامعتبر: $src"
    log "ساخت از پوشه محلی: $src"
    cd "$src"
  else
    tmp="$(mktemp -d)"
    # shellcheck disable=SC2064
    trap "rm -rf '$tmp'" EXIT
    log "کلون $REPO_URL (branch=$BRANCH)..."
    if [[ "$DRY_RUN" -eq 1 ]]; then
      run git clone --depth 1 --branch "$BRANCH" "$REPO_URL" "$tmp/faman"
      return 0
    fi
    git clone --depth 1 --branch "$BRANCH" "$REPO_URL" "$tmp/faman"
    cd "$tmp/faman"
  fi

  log "go build..."
  if [[ "$DRY_RUN" -eq 1 ]]; then
    run go build -ldflags="-s -w" -o faman ./cmd/faman
  else
    go build -ldflags="-s -w" -o faman ./cmd/faman
  fi

  log "نصب → $PREFIX"
  run need_sudo install -d "$PREFIX/bin"
  run need_sudo install -m 755 faman "$PREFIX/bin/faman"
  run need_sudo install -d "$PREFIX/share/faman"
  if [[ "$DRY_RUN" -eq 0 ]]; then
    need_sudo rm -rf "$PREFIX/share/faman/pages"
    need_sudo cp -a pages "$PREFIX/share/faman/"
  else
    run need_sudo cp -a pages "$PREFIX/share/faman/"
  fi

  if [[ "$USER_INSTALL" -eq 1 ]]; then
    case ":$PATH:" in
      *:"$PREFIX/bin":*) ;;
      *)
        warn "$PREFIX/bin در PATH نیست — به shell rc اضافه می‌شود"
        ;;
    esac
  fi

  ok "faman → $PREFIX/bin/faman"
  if [[ "$DRY_RUN" -eq 0 ]]; then
    "$PREFIX/bin/faman" version || true
  fi
}

install_rtl_stack() {
  log "RTL / fonts / locale..."
  local pm
  pm="$(detect_pm)"

  case "$pm" in
    apt)
      run sys_pkg apt-get update -qq
      run sys_pkg apt-get install -y fonts-noto-core fonts-noto-ui-core fonts-dejavu-core fonts-dejavu-mono fontconfig locales 2>/dev/null || true
      run sys_pkg apt-get install -y fonts-vazirmatn 2>/dev/null || \
        run sys_pkg apt-get install -y fonts-vazir 2>/dev/null || \
        warn "Vazirmatn در apt نبود"
      if command -v locale-gen >/dev/null 2>&1 || [[ "$DRY_RUN" -eq 1 ]]; then
        if [[ "$LOCALE_PREF" == "fa" ]]; then
          run sys_pkg locale-gen fa_IR.UTF-8 en_US.UTF-8 2>/dev/null || run sys_pkg locale-gen en_US.UTF-8
        else
          run sys_pkg locale-gen en_US.UTF-8 fa_IR.UTF-8 2>/dev/null || run sys_pkg locale-gen en_US.UTF-8
        fi
      fi
      if [[ "$NO_MLTERM" -eq 0 ]]; then
        run sys_pkg apt-get install -y mlterm 2>/dev/null || warn "mlterm نصب نشد"
      else
        dbg "skip mlterm"
      fi
      run fc-cache -f >/dev/null 2>&1 || true
      ;;
    dnf)
      run sys_pkg dnf install -y google-noto-sans-fonts google-noto-sans-arabic-fonts dejavu-sans-mono-fonts fontconfig 2>/dev/null || true
      [[ "$NO_MLTERM" -eq 0 ]] && run sys_pkg dnf install -y mlterm 2>/dev/null || true
      run fc-cache -f >/dev/null 2>&1 || true
      ;;
    pacman)
      run sys_pkg pacman -Sy --needed --noconfirm noto-fonts noto-fonts-extra ttf-dejavu fontconfig 2>/dev/null || true
      [[ "$NO_MLTERM" -eq 0 ]] && run sys_pkg pacman -Sy --needed --noconfirm mlterm 2>/dev/null || true
      run fc-cache -f >/dev/null 2>&1 || true
      ;;
    *) warn "فونت خودکار برای این PM پشتیبانی نشده" ;;
  esac

  if [[ "$NO_SHELL_RC" -eq 1 ]]; then
    dbg "skip shell rc"
  else
    local lang_val snippet marker rc
    if [[ "$LOCALE_PREF" == "fa" ]]; then
      lang_val="fa_IR.UTF-8"
    else
      lang_val="en_US.UTF-8"
    fi
    marker='# faman — UTF-8 و نمایش فارسی'
    snippet="$marker
export LANG=\"\${LANG:-$lang_val}\"
export LC_ALL=\"\${LC_ALL:-$lang_val}\"
"
    if [[ "$USER_INSTALL" -eq 1 ]]; then
      snippet+="export PATH=\"$PREFIX/bin:\$PATH\"\n"
    fi
    if [[ "$PLAIN_DEFAULT" -eq 1 ]]; then
      snippet+=$'export FAMAN_PLAIN=1\n'
    else
      snippet+=$'# اگر حروف خراب بود: FAMAN_PLAIN=1 faman ls\n'
    fi

    for rc in "$HOME/.bashrc" "$HOME/.zshrc"; do
      if [[ -f "$rc" ]] && grep -qF "$marker" "$rc" 2>/dev/null; then
        ok "قبلاً در $rc"
        continue
      fi
      if [[ "$rc" == "$HOME/.bashrc" ]] || [[ -f "$rc" ]]; then
        if [[ "$DRY_RUN" -eq 1 ]]; then
          run echo "append snippet → $rc"
        else
          [[ -f "$rc" ]] || touch "$rc"
          printf '\n%s\n' "$snippet" >> "$rc"
          ok "به‌روزرسانی $rc"
        fi
      fi
    done
  fi

  if [[ "$DRY_RUN" -eq 0 ]]; then
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
    # sample advanced config
    if [[ ! -f "$HOME/.config/faman/install.env.example" ]]; then
      cat > "$HOME/.config/faman/install.env.example" <<'EOF'
# کپی به install.env و مقادیر را عوض کنید:
#   cp ~/.config/faman/install.env.example ~/.config/faman/install.env

# PREFIX=/usr/local
# USER_INSTALL=0
# WITH_RTL=1
# PLAIN_DEFAULT=0
# LOCALE_PREF=en
# NO_MLTERM=0
# NO_SHELL_RC=0
# BRANCH=main
# REPO_URL=https://github.com/erfankasraie/Faman.git
EOF
    fi
    ok "راهنما: ~/.config/faman/terminal-font-hint.txt"
  fi
}

main() {
  dbg "PREFIX=$PREFIX USER=$USER_INSTALL RTL=$WITH_RTL BRANCH=$BRANCH"

  if [[ "$UNINSTALL" -eq 1 ]]; then
    uninstall_faman
    exit 0
  fi

  if [[ "$RTL_ONLY" -eq 1 ]]; then
    log "حالت: فقط RTL"
    install_rtl_stack
  else
    log "حالت: نصب faman"
    if [[ "$SKIP_DEPS" -eq 0 ]]; then
      install_deps
    fi
    install_faman
    if [[ "$WITH_RTL" -eq 1 ]]; then
      install_rtl_stack
    else
      warn "برای فونت/RTL: --with-rtl"
    fi
  fi

  cat <<EOF

┌─────────────────────────────────────────────
│  تمام شد  (prefix: $PREFIX)
│
│  faman ls
│  faman search docker
│  FAMAN_PLAIN=1 faman echo
│
│  تنظیمات پیشرفته:
│    scripts/install.sh --help
│    docs/install.md
│    ~/.config/faman/install.env.example
└─────────────────────────────────────────────
EOF
}

main
