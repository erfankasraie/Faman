# ── faman (managed by scripts/install.sh) ─────────────────
# UTF-8
export LANG="${LANG:-en_US.UTF-8}"
export LC_ALL="${LC_ALL:-en_US.UTF-8}"

# اگر با --user نصب شده:
# export PATH="$HOME/.local/bin:$PATH"

# خروجی ساده فارسی (اختیاری — با --plain-default فعال می‌شود):
# export FAMAN_PLAIN=1

# completion
if [[ -d "$HOME/.zsh/completions" ]]; then
  fpath=($HOME/.zsh/completions $fpath)
fi

alias f='faman' 2>/dev/null
alias fs='faman search' 2>/dev/null
alias fp='FAMAN_PLAIN=1 faman' 2>/dev/null
# ──────────────────────────────────────────────────────────
