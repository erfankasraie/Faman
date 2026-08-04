# ═══════════════════════════════════════════════════════════
# faman — advanced zsh module
#   source ~/.config/faman/faman.zsh
# fzf extras: source ~/.config/faman/faman-fzf.zsh  (docs/fzf.md)
# ═══════════════════════════════════════════════════════════

(( $+commands[faman] )) || return 0

: "${LANG:=en_US.UTF-8}"
: "${LC_ALL:=en_US.UTF-8}"
export LANG LC_ALL

# Completion path MUST be on fpath before compinit / Oh My Zsh
[[ -d ${HOME}/.zsh/completions ]] && fpath=(${HOME}/.zsh/completions $fpath)

# ── aliases (core + new CLI) ───────────────────────────────
alias f='faman'
alias fs='faman search'
alias fl='faman list'
alias fcats='faman categories'
alias frand='faman random'
alias fdoc='faman doctor'
alias fup='faman update'
alias fupp='faman update --pages'
alias fupv='faman update --pages --verify'
alias fv='faman version'
alias fh='faman help'
alias fp='FAMAN_PLAIN=1 faman'
alias fn='FAMAN_PLAIN=0 faman'

# Show page or search if missing
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

fmp() { FAMAN_PLAIN=1 faman "$@"; }

# List page titles from known dirs (for fzf / scripts)
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
  # fallback via CLI
  faman list --names 2>/dev/null && return 0
  print -u2 "pages dir not found"
  return 1
}

# Regenerate zsh completion after upgrading faman
faman-recomplete() {
  mkdir -p "${HOME}/.zsh/completions"
  if ! command -v faman >/dev/null 2>&1; then
    print -u2 "faman not in PATH"
    return 1
  fi
  faman completion zsh > "${HOME}/.zsh/completions/_faman"
  rm -f "${HOME}/.zcompdump" "${HOME}/.zcompdump"*(N)
  autoload -Uz compinit && compinit -u
  print -P "%F{green}completion refreshed%f → ~/.zsh/completions/_faman"
  print "  test: faman <TAB>   or   faman list --cat <TAB>"
}

# minimal fzf fallback if faman-fzf.zsh not loaded
if (( $+commands[fzf] )) && ! typeset -f faman-fzf >/dev/null 2>&1; then
  faman-fzf() {
    local page
    page="$(fls 2>/dev/null | sort -u | fzf --prompt='faman> ' --height=40% --reverse)" || return
    [[ -n "$page" ]] && faman "$page"
  }
  alias ff='faman-fzf'
  _faman_fzf_page() { faman-fzf; zle reset-prompt 2>/dev/null || true; }
  zle -N _faman_fzf_page
  bindkey '^[f' _faman_fzf_page 2>/dev/null || true
fi

_faman_last_word() {
  local w
  w="${${(z)history[$((HISTCMD-1))]}[1]}"
  [[ -n "$w" ]] && BUFFER="faman $w" && CURSOR=$#BUFFER
  zle redisplay
}
zle -N _faman_last_word
bindkey '^[m' _faman_last_word 2>/dev/null || true

faman-status() {
  print -P "%F{cyan}faman%f binary: $(command -v faman 2>/dev/null || echo missing)"
  faman version 2>/dev/null || true
  print "LANG=$LANG  LC_ALL=$LC_ALL"
  print "FAMAN_PLAIN=${FAMAN_PLAIN:-unset}  FAMAN_WRAP=${FAMAN_WRAP:-unset}"
  print "FAMAN_PAGES=${FAMAN_PAGES:-unset}"
  (( $+commands[fzf] )) && print -P "fzf: %F{green}$(command -v fzf)%f" || print -P "fzf: %F{red}not installed%f"
  typeset -f faman-fzf >/dev/null 2>&1 && print -P "faman-fzf: %F{green}loaded%f" || print "faman-fzf: not loaded (docs/fzf.md)"
  if [[ -f ${HOME}/.zsh/completions/_faman ]]; then
    print -P "completion file: %F{green}~/.zsh/completions/_faman%f"
  else
    print -P "completion file: %F{red}missing%f — run: faman-recomplete"
  fi
  if (( ${fpath[(I)${HOME}/.zsh/completions]} )); then
    print -P "fpath completions: %F{green}ok%f"
  else
    print -P "fpath completions: %F{red}NOT on fpath%f — add before compinit / oh-my-zsh"
  fi
  print "commands: list categories random doctor update search --cat"
}
