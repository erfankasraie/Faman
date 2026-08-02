# ═══════════════════════════════════════════════════════════
# faman — bat preview helpers for fzf
#   source ~/.config/faman/faman-bat.zsh
#   (optional; faman-fzf.zsh loads compatible defaults alone)
# ═══════════════════════════════════════════════════════════

# Ubuntu package name is often "batcat"
if (( !$+commands[bat] )) && (( $+commands[batcat] )); then
  alias bat=batcat
fi

# ── user-tunable env (set in ~/.zshrc before sourcing) ─────
# FAMAN_BAT_THEME   default: ansi  (or TwoDark, Dracula, GitHub, …)
# FAMAN_BAT_STYLE   default: plain,header,grid
# FAMAN_BAT_LINES   default: 120   (max lines in preview)
# FAMAN_BAT_WRAP    default: character
# FAMAN_BAT_EXTRA   default: empty (extra CLI flags)

: "${FAMAN_BAT_THEME:=ansi}"
: "${FAMAN_BAT_STYLE:=plain,header,grid}"
: "${FAMAN_BAT_LINES:=120}"
: "${FAMAN_BAT_WRAP:=character}"
: "${FAMAN_BAT_EXTRA:=}"

export FAMAN_BAT_THEME FAMAN_BAT_STYLE FAMAN_BAT_LINES FAMAN_BAT_WRAP FAMAN_BAT_EXTRA

# Global bat defaults (affects all bat usage in this shell)
export BAT_THEME="${BAT_THEME:-$FAMAN_BAT_THEME}"
export BAT_STYLE="${BAT_STYLE:-$FAMAN_BAT_STYLE}"

# Build a safe preview command string for fzf --preview
# Usage:  fzf --preview "$(faman_bat_preview_cmd /path/to/pages/{1}.md)"
# Or with placeholder already in path:
faman_bat_preview_cmd() {
  local target="${1:-"{}"}"
  if (( $+commands[bat] )) || (( $+commands[batcat] )); then
    local batbin=bat
    (( $+commands[bat] )) || batbin=batcat
    # {1} is fzf field — keep as literal for fzf
    print -r -- "$batbin --color=always --paging=never --language=markdown \
--theme='${FAMAN_BAT_THEME}' --style='${FAMAN_BAT_STYLE}' \
--wrap='${FAMAN_BAT_WRAP}' --line-range=':${FAMAN_BAT_LINES}' \
${FAMAN_BAT_EXTRA} ${target}"
  else
    print -r -- "head -n ${FAMAN_BAT_LINES} ${target}"
  fi
}

# List themes
faman-bat-themes() {
  if (( $+commands[bat] )); then
    bat --list-themes
  elif (( $+commands[batcat] )); then
    batcat --list-themes
  else
    print -u2 "bat not installed"
    return 1
  fi
}

# Quick theme switch for current session
faman-bat-theme() {
  if [[ -z "$1" ]]; then
    print "current: $FAMAN_BAT_THEME"
    print "usage: faman-bat-theme <name>   (see faman-bat-themes)"
    return 0
  fi
  export FAMAN_BAT_THEME="$1"
  export BAT_THEME="$1"
  print -P "%F{green}BAT theme → $1%f"
}

# Preview a page by name in the terminal (not fzf)
faman-bat() {
  local name="${1:?usage: faman-bat <page>}"
  local dir f
  for dir in \
    "${FAMAN_PAGES:-}" \
    "${HOME}/.local/share/faman/pages/fa" \
    /usr/local/share/faman/pages/fa \
    /usr/share/faman/pages/fa; do
    [[ -n "$dir" && -f "$dir/${name}.md" ]] || continue
    f="$dir/${name}.md"
    if (( $+commands[bat] )); then
      bat --language=markdown --theme="$FAMAN_BAT_THEME" --style="$FAMAN_BAT_STYLE" \
        --line-range=":$FAMAN_BAT_LINES" ${=FAMAN_BAT_EXTRA} "$f"
    elif (( $+commands[batcat] )); then
      batcat --language=markdown --theme="$FAMAN_BAT_THEME" --style="$FAMAN_BAT_STYLE" \
        --line-range=":$FAMAN_BAT_LINES" ${=FAMAN_BAT_EXTRA} "$f"
    else
      head -n "$FAMAN_BAT_LINES" "$f"
    fi
    return 0
  done
  print -u2 "page not found: $name"
  return 1
}

alias fbat='faman-bat'
