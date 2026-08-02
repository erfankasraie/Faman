#!/usr/bin/env bash
# پیکربندی شل بعد از نصب faman — zsh/bash completion + rc snippet
# Usage:
#   PREFIX=/usr/local bash scripts/install-shell.sh
#   PREFIX=$HOME/.local USER_INSTALL=1 PLAIN_DEFAULT=1 bash scripts/install-shell.sh
set -euo pipefail

PREFIX="${PREFIX:-/usr/local}"
USER_INSTALL="${USER_INSTALL:-0}"
PLAIN_DEFAULT="${PLAIN_DEFAULT:-0}"
LOCALE_PREF="${LOCALE_PREF:-en}"
NO_SHELL_RC="${NO_SHELL_RC:-0}"
FAMAN_BIN="${FAMAN_BIN:-$PREFIX/bin/faman}"

log() { printf '\033[1;36m==>\033[0m %s\n' "$*"; }
ok()  { printf '\033[1;32m✓\033[0m %s\n' "$*"; }
warn(){ printf '\033[1;33m!\033[0m %s\n' "$*"; }

if [[ ! -x "$FAMAN_BIN" ]]; then
  if command -v faman >/dev/null 2>&1; then
    FAMAN_BIN="$(command -v faman)"
  else
    warn "faman پیدا نشد — اول باینری را نصب کنید"
    exit 1
  fi
fi

install_completions() {
  log "نصب completion..."

  # zsh
  mkdir -p "$HOME/.zsh/completions"
  if "$FAMAN_BIN" completion zsh > "$HOME/.zsh/completions/_faman" 2>/dev/null; then
    ok "zsh: ~/.zsh/completions/_faman"
  else
    warn "تولید completion zsh ناموفق (نسخه faman را به‌روز کنید)"
  fi

  # bash (user)
  mkdir -p "$HOME/.local/share/bash-completion/completions"
  if "$FAMAN_BIN" completion bash > "$HOME/.local/share/bash-completion/completions/faman" 2>/dev/null; then
    ok "bash: ~/.local/share/bash-completion/completions/faman"
  fi

  # system bash completion if writable
  if [[ -d /etc/bash_completion.d && -w /etc/bash_completion.d ]]; then
    "$FAMAN_BIN" completion bash > /etc/bash_completion.d/faman 2>/dev/null && ok "bash: /etc/bash_completion.d/faman" || true
  fi
}

install_zshrc() {
  [[ "$NO_SHELL_RC" -eq 1 ]] && return 0
  local rc="$HOME/.zshrc"
  local marker='# ── faman'
  if [[ -f "$rc" ]] && grep -qF "$marker" "$rc" 2>/dev/null; then
    ok "zshrc قبلاً تنظیم شده"
    return 0
  fi

  local lang_val="en_US.UTF-8"
  [[ "$LOCALE_PREF" == "fa" ]] && lang_val="fa_IR.UTF-8"

  [[ -f "$rc" ]] || touch "$rc"

  {
    echo ""
    echo "$marker ──────────────────────────────────────────"
    echo "export LANG=\"\${LANG:-$lang_val}\""
    echo "export LC_ALL=\"\${LC_ALL:-$lang_val}\""
    if [[ "$USER_INSTALL" -eq 1 ]] || [[ "$PREFIX" == "$HOME/.local" ]]; then
      echo "export PATH=\"$PREFIX/bin:\$PATH\""
    fi
    if [[ "$PLAIN_DEFAULT" -eq 1 ]]; then
      echo "export FAMAN_PLAIN=1"
    else
      echo "# export FAMAN_PLAIN=1"
    fi
    echo "fpath=(\$HOME/.zsh/completions \$fpath)"
    echo "# اگر Oh My Zsh ندارید، خط بعد را از حالت نظر خارج کنید:"
    echo "# autoload -Uz compinit && compinit"
    echo "alias f='faman'"
    echo "alias fs='faman search'"
    echo "alias fp='FAMAN_PLAIN=1 faman'"
    echo "# ───────────────────────────────────────────────────"
  } >> "$rc"
  ok "به‌روزرسانی ~/.zshrc"
  warn "برای اعمال: exec zsh   یا   source ~/.zshrc"
  warn "اگر Oh My Zsh دارید، بلوک faman را بالای source oh-my-zsh.sh بگذارید"
}

install_bashrc() {
  [[ "$NO_SHELL_RC" -eq 1 ]] && return 0
  local rc="$HOME/.bashrc"
  local marker='# faman — UTF-8'
  if [[ -f "$rc" ]] && grep -qF "$marker" "$rc" 2>/dev/null; then
    ok "bashrc قبلاً تنظیم شده"
    return 0
  fi
  local lang_val="en_US.UTF-8"
  [[ "$LOCALE_PREF" == "fa" ]] && lang_val="fa_IR.UTF-8"
  [[ -f "$rc" ]] || touch "$rc"
  {
    echo ""
    echo "$marker و نمایش فارسی"
    echo "export LANG=\"\${LANG:-$lang_val}\""
    echo "export LC_ALL=\"\${LC_ALL:-$lang_val}\""
    if [[ "$USER_INSTALL" -eq 1 ]] || [[ "$PREFIX" == "$HOME/.local" ]]; then
      echo "export PATH=\"$PREFIX/bin:\$PATH\""
    fi
    if [[ "$PLAIN_DEFAULT" -eq 1 ]]; then
      echo "export FAMAN_PLAIN=1"
    fi
    echo "alias f='faman'"
    echo "alias fs='faman search'"
  } >> "$rc"
  ok "به‌روزرسانی ~/.bashrc"
}

install_completions
install_zshrc
install_bashrc

log "راهنما: https://github.com/erfankasraie/Faman/blob/main/docs/zsh.md"
