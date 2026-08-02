# پیکربندی fzf برای faman

## پیش‌نیاز

```bash
# Ubuntu / Debian
sudo apt install fzf

# اختیاری — پیش‌نمایش رنگی markdown
sudo apt install bat
# روی بعضی نسخه‌ها باینری `batcat` است:
# alias bat=batcat
```

یا:

```bash
git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf
~/.fzf/install
```

---

## نصب یک‌خطی ماژول

```bash
mkdir -p ~/.config/faman
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/zsh/faman-fzf.zsh \
  -o ~/.config/faman/faman-fzf.zsh
```

در `~/.zshrc`:

```zsh
source ~/.config/faman/faman-fzf.zsh
# اگر faman.zsh را هم دارید، هر دو را source کنید
```

```zsh
exec zsh
```

---

## کلیدها و دستورات

| کلید / دستور | کار |
|--------------|-----|
| **Alt-f** یا `ff` / `faman-fzf` | انتخاب فازی صفحه + **پیش‌نمایش** markdown |
| **Alt-s** یا `ffs query` | همان picker با query اولیه |
| `ffp` | picker سپس باز کردن با `FAMAN_PLAIN=1` |
| `ffg docker` | نتایج `faman search` → fzf → باز کردن |
| **Ctrl-P** داخل fzf | باز کردن همان صفحه به‌صورت plain و خروج |
| **Ctrl-Y** داخل fzf | کپی نام صفحه (xclip / wl-copy / clipcopy) |
| **Enter** | `faman <page>` |

---

## ظاهر fzf

متغیر اختیاری:

```zsh
export FAMAN_FZF_OPTS='--height=80% --layout=reverse --info=inline'
```

تم کلی fzf (همهٔ fzfها، نه فقط faman):

```zsh
export FZF_DEFAULT_OPTS="--height 50% --layout=reverse --border --ansi"
```

---

## پیش‌نمایش

- اگر `bat` نصب باشد: syntax highlight برای markdown
- وگرنه: `head -n 80` از فایل `.md`

مسیر صفحات مثل faman حل می‌شود:

1. `FAMAN_PAGES`
2. `~/.local/share/faman/pages/fa`
3. `/usr/local/share/faman/pages/fa`
4. …

---

## تداخل کلید

اگر **Alt-f** / **Alt-s** کار نکرد (ترمینال یا tmux):

```zsh
# بعد از source ماژول:
bindkey '^F' _faman_fzf_widget          # Ctrl-F
bindkey '^[^F' _faman_fzf_search_widget # Ctrl-Alt-F — وابسته به ترمینال
```

در tmux گاهی نیاز است `Alt` به‌درستی پاس شود (`tmux.conf`: `set -g xterm-keys on`).

---

## Bash (ساده)

```bash
faman-fzf() {
  local dir="${FAMAN_PAGES:-$HOME/.local/share/faman/pages/fa}"
  [[ -d "$dir" ]] || dir=/usr/local/share/faman/pages/fa
  local page
  page="$(find "$dir" -name '*.md' -printf '%f\n' 2>/dev/null \
    | sed 's/\.md$//' \
    | fzf --prompt='faman> ' --height=60% --reverse \
        --preview "head -n 60 $dir/{}.md")" || return
  [[ -n "$page" ]] && faman "$page"
}
```

---

## عیب‌یابی

| مشکل | کار |
|------|-----|
| `ff: command not found` | `source` ماژول در zshrc |
| پیش‌نمایش خالی | مسیر pages / `FAMAN_PAGES` |
| بدون رنگ | `sudo apt install bat` |
| Alt-f بی‌اثر | `bindkey '^F' _faman_fzf_widget` |

مرتبط: [zsh.md](zsh.md) · [zsh-advanced.md](zsh-advanced.md)
