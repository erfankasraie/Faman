# ═══════════════════════════════════════════════════════════
# faman + fzf (+ bat preview)
#   source ~/.config/faman/faman-bat.zsh   # optional overrides
#   source ~/.config/faman/faman-fzf.zsh
# ═══════════════════════════════════════════════════════════

(( $+commands[faman] )) || return 0
(( $+commands[fzf] ))  || return 0

# bat → batcat on Debian/Ubuntu
if (( !$+commands[bat] )) && (( $+commands[batcat] )); then
  alias bat=batcat
fi

: "${FAMAN_BAT_THEME:=ansi}"
: "${FAMAN_BAT_STYLE:=plain,header,grid}"
: "${FAMAN_BAT_LINES:=120}"
: "${FAMAN_BAT_WRAP:=character}"
: "${FAMAN_BAT_EXTRA:=}"

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
  local dir f title cat
  dir="$(_faman_pages_dir)" || return 1
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
  local dir batbin
  dir="$(_faman_pages_dir)" || { print -r -- 'echo "no pages dir"'; return; }
  if (( $+commands[bat] )) || (( $+commands[batcat] )); then
    batbin=bat
    (( $+commands[bat] )) || batbin=batcat
    # {1} = first field from list (page name)
    print -r -- "$batbin --color=always --paging=never --language=markdown \
--theme=${FAMAN_BAT_THEME} --style=${FAMAN_BAT_STYLE} \
--wrap=${FAMAN_BAT_WRAP} --line-range=:${FAMAN_BAT_LINES} \
${FAMAN_BAT_EXTRA} ${dir}/{1}.md 2>/dev/null || echo missing"
  else
    print -r -- "head -n ${FAMAN_BAT_LINES} ${dir}/{1}.md 2>/dev/null || echo missing"
  fi
}

_faman_fzf_opts() {
  local preview
  preview="$(_faman_preview_cmd)"
  print -r -- \
    --ansi --reverse --height=60% --border \
    --prompt='faman ❯ ' \
    --header='Enter open · Ctrl-P plain · Ctrl-Y yank · Ctrl-B bat-full' \
    --preview="$preview" \
    --preview-window="${FAMAN_FZF_PREVIEW_WINDOW:-right:55%:wrap}" \
    ${FAMAN_FZF_OPTS:-}
}

faman-fzf() {
  local dir line name
  dir="$(_faman_pages_dir)" || {
    print -u2 "faman pages not found (set FAMAN_PAGES)"
    return 1
  }

  line="$(_faman_list_pages | FZF_DEFAULT_OPTS='' fzf \
    $(_faman_fzf_opts) \
    --bind="ctrl-p:execute(FAMAN_PLAIN=1 faman {1})+abort" \
    --bind="ctrl-b:execute(bat --color=always --theme=${FAMAN_BAT_THEME} --style=${FAMAN_BAT_STYLE} --language=markdown ${dir}/{1}.md 2>/dev/null || batcat --color=always --language=markdown ${dir}/{1}.md 2>/dev/null || less ${dir}/{1}.md)" \
    --bind="ctrl-y:execute-silent(printf %s {1} | clipcopy 2>/dev/null || printf %s {1} | xclip -selection clipboard 2>/dev/null || printf %s {1} | wl-copy 2>/dev/null || printf %s {1})+abort" \
  )" || return 1

  name="${${(s: :)line}[1]}"
  [[ -z "$name" ]] && return 1
  faman "$name"
}

faman-fzf-plain() { FAMAN_PLAIN=1 faman-fzf; }

faman-fzf-search() {
  local q="${*:-}" pick name
  pick="$(_faman_list_pages | FZF_DEFAULT_OPTS='' fzf \
    $(_faman_fzf_opts) \
    --query="$q" \
    --prompt='faman search ❯ ' \
  )" || return 1
  name="${${(s: :)pick}[1]}"
  [[ -n "$name" ]] && faman "$name"
}

faman-fzf-grep() {
  local q="${1:-.}" pick name
  pick="$(faman search "$q" 2>/dev/null \
    | sed -n 's/^[[:space:]]*\([^[:space:]]*\).*/\1/p' \
    | grep -v '^$' | sort -u \
    | FZF_DEFAULT_OPTS='' fzf \
        --ansi --reverse --height=50% --border \
        --prompt='search hits ❯ ' \
        --preview="$(_faman_preview_cmd)" \
        --preview-window="${FAMAN_FZF_PREVIEW_WINDOW:-right:55%:wrap}" \
        ${FAMAN_FZF_OPTS:-} \
  )" || return 1
  name="${${(s: :)pick}[1]}"
  [[ -n "$name" ]] && faman "$name"
}

_faman_fzf_widget() {
  faman-fzf
  zle reset-prompt 2>/dev/null || true
}
zle -N _faman_fzf_widget

_faman_fzf_search_widget() {
  local q="$BUFFER"
  BUFFER=
  zle reset-prompt
  faman-fzf-search "$q"
  zle reset-prompt 2>/dev/null || true
}
zle -N _faman_fzf_search_widget

bindkey '^[f' _faman_fzf_widget
bindkey '^[s' _faman_fzf_search_widget

alias ff='faman-fzf'
alias ffp='faman-fzf-plain'
alias ffs='faman-fzf-search'
alias ffg='faman-fzf-grep'
