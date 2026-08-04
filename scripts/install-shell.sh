#!/usr/bin/env bash
# Configure shell integration for faman: completion, zshrc/bashrc, optional modules.
#
#   PREFIX=$HOME/.local USER_INSTALL=1 bash scripts/install-shell.sh
#   PLAIN_DEFAULT=1 bash scripts/install-shell.sh
#   NO_SHELL_RC=1 bash scripts/install-shell.sh   # only completion files
#
set -euo pipefail

PREFIX="${PREFIX:-/usr/local}"
USER_INSTALL="${USER_INSTALL:-0}"
PLAIN_DEFAULT="${PLAIN_DEFAULT:-0}"
LOCALE_PREF="${LOCALE_PREF:-en}"
NO_SHELL_RC="${NO_SHELL_RC:-0}"
FAMAN_BIN="${FAMAN_BIN:-$PREFIX/bin/faman}"
INSTALL_MODULES="${INSTALL_MODULES:-1}"

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
  log "نصب / بازتولید completion..."

  mkdir -p "$HOME/.zsh/completions"
  if "$FAMAN_BIN" completion zsh > "$HOME/.zsh/completions/_faman" 2>/dev/null; then
    ok "zsh: ~/.zsh/completions/_faman"
  else
    warn "تولید completion zsh ناموفق (نسخه faman را به‌روز کنید)"
  fi

  mkdir -p "$HOME/.local/share/bash-completion/completions"
  if "$FAMAN_BIN" completion bash > "$HOME/.local/share/bash-completion/completions/faman" 2>/dev/null; then
    ok "bash: ~/.local/share/bash-completion/completions/faman"
  fi

  if [[ -d /etc/bash_completion.d && -w /etc/bash_completion.d ]]; then
    "$FAMAN_BIN" completion bash > /etc/bash_completion.d/faman 2>/dev/null && ok "bash: /etc/bash_completion.d/faman" || true
  fi

  # Invalidate zsh completion cache so new commands appear
  rm -f "$HOME/.zcompdump" "$HOME"/.zcompdump* 2>/dev/null || true
  ok "کش zcompdump پاک شد — بعد از source/exec، Tab تازه می‌شود"
}

install_zsh_modules() {
  [[ "$INSTALL_MODULES" -eq 1 ]] || return 0
  local src_dir
  src_dir="$(cd "$(dirname "$0")" 2>/dev/null && pwd)/zsh" || true
  if [[ ! -d "$src_dir" ]]; then
    # when piped from curl, scripts may be elsewhere — try repo relative
    if [[ -d "$PWD/scripts/zsh" ]]; then
      src_dir="$PWD/scripts/zsh"
    else
      warn "scripts/zsh یافت نشد — ماژول‌ها را دستی از ریپو کپی کنید"
      return 0
    fi
  fi
  mkdir -p "$HOME/.config/faman"
  for f in faman.zsh faman-fzf.zsh faman-bat.zsh; do
    if [[ -f "$src_dir/$f" ]]; then
      cp "$src_dir/$f" "$HOME/.config/faman/$f"
      ok "module: ~/.config/faman/$f"
    fi
  done
}

install_zshrc() {
  [[ "$NO_SHELL_RC" -eq 1 ]] && return 0
  local rc="$HOME/.zshrc"
  local marker='# ── faman'
  if [[ -f "$rc" ]] && grep -qF "$marker" "$rc" 2>/dev/null; then
    ok "zshrc قبلاً تنظیم شده (برای aliasهای جدید: source scripts/zsh/faman.zsh)"
    return 0
  fi

  local lang_val="en_US.UTF-8"
  [[ "$LOCALE_PREF" == "fa" ]] && lang_val="fa_IR.UTF-8"

  [[ -f "$rc" ]] || touch "$rc"

  {
    echo ""
    echo "$marker ──────────────────────────────────────────"
    echo "# این بلوک را بالای source oh-my-zsh.sh نگه دارید"
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
    echo "# بدون Oh My Zsh:"
    echo "# autoload -Uz compinit && compinit"
    echo "[[ -f \$HOME/.config/faman/faman.zsh ]] && source \$HOME/.config/faman/faman.zsh"
    echo "# [[ -f \$HOME/.config/faman/faman-fzf.zsh ]] && source \$HOME/.config/faman/faman-fzf.zsh"
    echo "# ───────────────────────────────────────────────────"
  } >> "$rc"
  ok "به‌روزرسانی ~/.zshrc"
  warn "برای اعمال: exec zsh   یا   source ~/.zshrc"
  warn "اگر Oh My Zsh دارید، بلوک faman باید بالای source oh-my-zsh.sh باشد"
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
    echo "alias fl='faman list'"
    echo "alias fdoc='faman doctor'"
  } >> "$rc"
  ok "به‌روزرسانی ~/.bashrc"
}

install_completions
install_zsh_modules
install_zshrc
install_bashrc

log "راهنما: docs/zsh.md · docs/completion.md"
log "بعد از آپدیت باینری: faman-recomplete  (اگر ماژول zsh را source کرده باشید)"
