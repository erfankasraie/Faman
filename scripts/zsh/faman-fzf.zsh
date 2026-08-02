# ═══════════════════════════════════════════════════════════
# faman + fzf
#   source ~/.config/faman/faman-fzf.zsh
# Requires: fzf, faman
# Optional: bat (prettier preview)
# ═══════════════════════════════════════════════════════════

(( $+commands[faman] )) || return 0
(( $+commands[fzf] ))  || return 0

# ── resolve pages directory ────────────────────────────────
_faman_pages_dir() {
  local d
  for d in \
    "${FAMAN_PAGES:-}" \
    "${HOME}/.local/share/faman/pages/fa" \
    /usr/local/share/faman/pages/fa \
    /usr/share/faman/pages/fa \
    /opt/faman/pages/fa; do
    [[ -n "$d" && -d "$d" ]] && { print -r -- "$d"; return 0; }
  done
  return 1
}

_faman_list_pages() {
  local dir
  dir="$(_faman_pages_dir)" || return 1
  # name + category from front matter (best-effort)
  local f title cat
  for f in "$dir"/*.md(N); do
    title="${f:t:r}"
    cat="$(grep -m1 '^category:' "$f" 2>/dev/null | sed 's/^category:[[:space:]]*//')"
    if [[ -n "$cat" ]]; then
      printf '%-16s  %s\n' "$title" "$cat"
    else
      printf '%s\n' "$title"
    fi
  done | sort -u
}

_faman_preview_cmd() {
  local dir
  dir="$(_faman_pages_dir)" || { print -r -- 'echo no pages'; return; }
  if (( $+commands[bat] )); then
    print -r -- "bat --style=plain --color=always --language=markdown ${dir}/{1}.md 2>/dev/null || echo 'missing'"
  else
    print -r -- "head -n 80 ${dir}/{1}.md 2>/dev/null || echo 'missing'"
  fi
}

# default fzf flags (override with FAMAN_FZF_OPTS)
_faman_fzf_opts() {
  local preview
  preview="$(_faman_preview_cmd)"
  # shellcheck disable=SC2086
  print -r -- \
    --ansi --reverse --height=60% --border \
    --prompt='faman ❯ ' \
    --header='Enter: open · Ctrl-P: plain · Ctrl-Y: yank name · Ctrl-S: search mode' \
    --preview="$preview" \
    --preview-window='right:55%:wrap' \
    ${FAMAN_FZF_OPTS:-}
}

# ── interactive page picker ────────────────────────────────
faman-fzf() {
  local dir line name
  dir="$(_faman_pages_dir)" || {
    print -u2 "faman pages not found (set FAMAN_PAGES)"
    return 1
  }

  line="$(_faman_list_pages | FZF_DEFAULT_OPTS='' fzf \
    $(_faman_fzf_opts) \
    --bind="ctrl-p:execute(FAMAN_PLAIN=1 faman {1})+abort" \
    --bind="ctrl-y:execute-silent(printf %s {1} | clipcopy 2>/dev/null || printf %s {1} | xclip -selection clipboard 2>/dev/null || printf %s {1} | wl-copy 2>/dev/null || printf %s {1})+abort" \
  )" || return 1

  name="${${(s: :)line}[1]}"
  [[ -z "$name" ]] && return 1
  faman "$name"
}

# plain variant
faman-fzf-plain() {
  FAMAN_PLAIN=1 faman-fzf
}

# ── search then pick ───────────────────────────────────────
faman-fzf-search() {
  local q="${*:-}"
  local pick name
  # re-query: start with argument as initial query
  pick="$(_faman_list_pages | FZF_DEFAULT_OPTS='' fzf \
    $(_faman_fzf_opts) \
    --query="$q" \
    --prompt='faman search ❯ ' \
  )" || return 1
  name="${${(s: :)pick}[1]}"
  [[ -n "$name" ]] && faman "$name"
}

# live filter using faman search output (text results)
faman-fzf-grep() {
  local q="${1:-.}"
  local pick name
  pick="$(faman search "$q" 2>/dev/null \
    | sed -n 's/^[[:space:]]*\([^[:space:]]*\).*/\1/p' \
    | grep -v '^$' \
    | sort -u \
    | FZF_DEFAULT_OPTS='' fzf \
        --ansi --reverse --height=50% --border \
        --prompt='search hits ❯ ' \
        --preview="$(_faman_preview_cmd)" \
        --preview-window='right:55%:wrap' \
        ${FAMAN_FZF_OPTS:-} \
  )" || return 1
  name="${${(s: :)pick}[1]}"
  [[ -n "$name" ]] && faman "$name"
}

# ── zle widgets ────────────────────────────────────────────
_faman_fzf_widget() {
  faman-fzf
  zle reset-prompt 2>/dev/null || true
}
zle -N _faman_fzf_widget

_faman_fzf_search_widget() {
  # use current buffer as query
  local q="$BUFFER"
  BUFFER=
  zle reset-prompt
  faman-fzf-search "$q"
  zle reset-prompt 2>/dev/null || true
}
zle -N _faman_fzf_search_widget

# Keybindings (customize if they clash):
#   Alt-f  → page picker with preview
#   Alt-s  → search picker (query = current line)
#   Ctrl-G f is not standard; we use Alt
bindkey '^[f' _faman_fzf_widget         # Alt-f
bindkey '^[s' _faman_fzf_search_widget  # Alt-s

# Optional: Ctrl-F if Alt is awkward in your terminal
# bindkey '^F' _faman_fzf_widget

# ── aliases ────────────────────────────────────────────────
alias ff='faman-fzf'
alias ffp='faman-fzf-plain'
alias ffs='faman-fzf-search'
alias ffg='faman-fzf-grep'
