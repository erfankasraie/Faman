# پیکربندی zsh برای faman

## سریع

```bash
# نصب + تنظیم شل
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh \
  | bash -s -- --user --with-rtl

# فقط completion و zshrc (اگر faman نصب است)
PREFIX=$HOME/.local USER_INSTALL=1 bash <(curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install-shell.sh)
```

یا از کلون:

```bash
PREFIX=$HOME/.local USER_INSTALL=1 bash scripts/install-shell.sh
```

بعد:

```zsh
exec zsh
faman <TAB>
```

---

## بلوک آماده برای `~/.zshrc`

این بلوک را **بالای** `source $ZSH/oh-my-zsh.sh` بگذارید (اگر Oh My Zsh دارید):

```zsh
# ── faman ──────────────────────────────────────────
export PATH="$HOME/.local/bin:$PATH"
export LANG="${LANG:-en_US.UTF-8}"
export LC_ALL="${LC_ALL:-en_US.UTF-8}"
# export FAMAN_PLAIN=1        # اگر حروف خراب بود فعال کنید
# export FAMAN_WRAP=0

fpath=($HOME/.zsh/completions $fpath)

alias f='faman'
alias fs='faman search'
alias fp='FAMAN_PLAIN=1 faman'
# ───────────────────────────────────────────────────
```

بدون Oh My Zsh، بعد از `fpath`:

```zsh
autoload -Uz compinit && compinit
```

### تولید completion

```zsh
mkdir -p ~/.zsh/completions
faman completion zsh > ~/.zsh/completions/_faman
rm -f ~/.zcompdump*
compinit
```

---

## Oh My Zsh

1. `fpath` و بلوک faman **قبل از** `source $ZSH/oh-my-zsh.sh`
2. نیازی به `compinit` دستی نیست
3. `exec zsh` بعد از تغییر

## Prezto / Antigen / Zinit

همان منطق: مسیر `~/.zsh/completions` باید در `fpath` باشد **قبل از** initialize شدن سیستم completion.

مثال Zinit:

```zsh
fpath=($HOME/.zsh/completions $fpath)
# ... zinit load ...
```

---

## متغیرهای محیطی مفید

| متغیر | معنی |
|--------|------|
| `FAMAN_PLAIN=1` | خروجی ساده، بدون wrap تهاجمی |
| `FAMAN_WRAP=0` | بدون شکستن خط |
| `FAMAN_PAGES` | مسیر سفارشی صفحات |
| `LANG` / `LC_ALL` | باید UTF-8 باشد |

---

## عیب‌یابی

| مشکل | کار |
|------|-----|
| `faman: command not found` | `export PATH="$HOME/.local/bin:$PATH"` |
| TAB چیزی نشان نمی‌دهد | `_faman` + `fpath` + `compinit` |
| completion قدیمی | دوباره `faman completion zsh > ~/.zsh/completions/_faman` و `rm ~/.zcompdump*` |
| فارسی خراب | فونت ترمینال + UTF-8 + `FAMAN_PLAIN=1` |

مستندات نصب: [install.md](install.md)  
ترمینال فارسی: [terminal-persian.md](terminal-persian.md)
