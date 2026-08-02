# ═══════════════════════════════════════════════════════════
# faman — advanced zsh module
# Source from ~/.zshrc (before or after Oh My Zsh as noted):
#   source /path/to/faman.zsh
# یا:
#   source ~/.config/faman/faman.zsh
# ═══════════════════════════════════════════════════════════

# ── 0) guard ───────────────────────────────────────────────
(( $+commands[faman] )) || return 0

# ── 1) environment ─────────────────────────────────────────
: "${LANG:=en_US.UTF-8}"
: "${LC_ALL:=en_US.UTF-8}"
export LANG LC_ALL

# Uncomment for safer Persian rendering by default:
# export FAMAN_PLAIN=1
# export FAMAN_WRAP=0

# ── 2) completion path ─────────────────────────────────────
[[ -d ${HOME}/.zsh/completions ]] && fpath=(${HOME}/.zsh/completions $fpath)

# ── 3) aliases ─────────────────────────────────────────────
alias f='faman'
alias fs='faman search'
alias fv='faman version'
alias fh='faman help'
alias fp='FAMAN_PLAIN=1 faman'
alias fn='FAMAN_PLAIN=0 faman'   # force rich render

# ── 4) helper functions ────────────────────────────────────

# fm <cmd>  — open page; if missing, search
fm() {
  if [[ $# -eq 0 ]]; then
    faman help
    return
  fi
  if ! faman "$1" 2>/dev/null; then
    print -P "%F{yellow}صفحه نبود — جستجو…%f"
    faman search "$*"
  fi
}

# fmp <cmd> — always plain (good for broken RTL terminals)
fmp() {
  FAMAN_PLAIN=1 faman "$@"
}

# fls — list all known page basenames (from pages dir if readable)
fls() {
  local d
  for d in \
    "${FAMAN_PAGES:-}" \
    "${HOME}/.local/share/faman/pages/fa" \
    /usr/local/share/faman/pages/fa \
    /usr/share/faman/pages/fa; do
    [[ -n "$d" && -d "$d" ]] || continue
    print -l -- "$d"/*.md(.N:t:r) 2>/dev/null
    return 0
  done
  faman search "" 2>/dev/null || print -u2 "pages dir not found"
}

# ── 5) fzf integration (optional) ──────────────────────────
if (( $+commands[fzf] )); then
  # Ctrl-F F — fuzzy pick a page then open it
  _faman_fzf_page() {
    local page
    page="$(fls 2>/dev/null | sort -u | fzf --prompt='faman> ' --height=40% --reverse)" || return
    [[ -n "$page" ]] && faman "$page"
  }
  zle -N _faman_fzf_page
  # Ctrl-Alt-F (works in many terminals); change if conflict
  bindkey '^[f' _faman_fzf_page 2>/dev/null || true

  # fzf-search: type query, pick from search results titles
  fzf-faman-search() {
    local q="${*:-}"
    local pick
    pick="$(faman search "$q" 2>/dev/null | sed -n 's/^[[:space:]]*\([^ ]*\).*/\1/p' | sort -u | fzf --prompt='search> ' --height=40% --reverse)" || return
    [[ -n "$pick" ]] && faman "$pick"
  }
fi

# ── 6) widget: insert last command as faman arg ────────────
# Usage: run something, then Esc-m to open «faman <last-word>»
_faman_last_word() {
  local w
  w="${${(z)history[$((HISTCMD-1))]}[1]}"
  [[ -n "$w" ]] && BUFFER="faman $w" && CURSOR=$#BUFFER
  zle redisplay
}
zle -N _faman_last_word
bindkey '^[m' _faman_last_word 2>/dev/null || true   # Alt-m

# ── 7) smart default: faman when command is unknown? (off)
# Uncomment only if you want aggressive behavior:
# command_not_found_handler() {
#   if faman "$1" >/dev/null 2>&1; then
#     faman "$1"
#     return 0
#   fi
#   print -u2 "zsh: command not found: $1"
#   return 127
# }

# ── 8) status helper ───────────────────────────────────────
faman-status() {
  print -P "%F{cyan}faman%f binary: $(command -v faman)"
  faman version 2>/dev/null || true
  print "LANG=$LANG  LC_ALL=$LC_ALL"
  print "FAMAN_PLAIN=${FAMAN_PLAIN:-unset}  FAMAN_WRAP=${FAMAN_WRAP:-unset}"
  print "FAMAN_PAGES=${FAMAN_PAGES:-unset}"
  [[ -f ${HOME}/.zsh/completions/_faman ]] \
    && print -P "completion: %F{green}~/.zsh/completions/_faman%f" \
    || print -P "completion: %F{red}missing%f — run: faman completion zsh > ~/.zsh/completions/_faman"
}
