# پیکربندی zsh برای faman

## سریع

```bash
PREFIX=$HOME/.local USER_INSTALL=1 bash scripts/install-shell.sh
exec zsh
faman <TAB>
faman-status        # اگر ماژول source شده باشد
```

---

## بلوک `~/.zshrc` (بالای Oh My Zsh)

```zsh
# ── faman ──────────────────────────────────────────
export PATH="$HOME/.local/bin:$PATH"
export LANG="${LANG:-en_US.UTF-8}"
export LC_ALL="${LC_ALL:-en_US.UTF-8}"
# export FAMAN_PLAIN=1

fpath=($HOME/.zsh/completions $fpath)

[[ -f $HOME/.config/faman/faman.zsh ]] && source $HOME/.config/faman/faman.zsh
# [[ -f $HOME/.config/faman/faman-fzf.zsh ]] && source $HOME/.config/faman/faman-fzf.zsh
# ───────────────────────────────────────────────────
```

بدون OMZ بعد از `fpath`:

```zsh
autoload -Uz compinit && compinit
```

### تولید / تازه‌سازی completion

```zsh
mkdir -p ~/.zsh/completions
faman completion zsh > ~/.zsh/completions/_faman
# یا:
faman-recomplete
```

جزئیات خطاها: [completion.md](completion.md)

---

## Aliasها و ابزارهای جدید (ماژول `faman.zsh`)

| میانبر | معادل |
|--------|--------|
| `f` | `faman` |
| `fs` | `faman search` |
| `fl` | `faman list` |
| `fcats` | `faman categories` |
| `frand` | `faman random` |
| `fdoc` | `faman doctor` |
| `fup` / `fupp` / `fupv` | `update` / `--pages` / `--pages --verify` |
| `fp` | `FAMAN_PLAIN=1 faman` |
| `ff` | انتخاب صفحه با fzf |
| `fm name` | show یا search |
| `fls` | لیست نام صفحات |
| `faman-recomplete` | بازتولید Tab completion |
| `faman-status` | تشخیص محیط |

کلیدها: `Alt-f` → fzf · `Alt-m` → man برای آخرین فرمان history

---

## Oh My Zsh / Prezto / Zinit

همان قانون: **`fpath` و source ماژول قبل از initialize completion framework.**

بعد از ارتقای باینری همیشه:

```zsh
faman-recomplete
```

---

## مرتبط

- [completion.md](completion.md) — عیب‌یابی Tab
- [fzf.md](fzf.md) · [bat-preview.md](bat-preview.md) · [zsh-advanced.md](zsh-advanced.md)
