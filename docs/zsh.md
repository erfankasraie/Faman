# پیکربندی zsh برای faman

## نصب خودکار با اسکریپت

```bash
# نصب کاربر + RTL + تنظیم zsh
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh \
  | bash -s -- --user --with-rtl --plain-default
```

اسکریپت در صورت وجود zsh:

- snippet UTF-8 / PATH را به `~/.zshrc` اضافه می‌کند
- completion را در `~/.zsh/completions/_faman` می‌نویسد
- در صورت نیاز `fpath` و `compinit` را یادآوری می‌کند

---

## تنظیم دستی کامل

### ۱) PATH (اگر با `--user` نصب کردید)

در `~/.zshrc`:

```zsh
export PATH="$HOME/.local/bin:$PATH"
```

### ۲) UTF-8 و نمایش فارسی

```zsh
# faman — UTF-8 و نمایش فارسی
export LANG="${LANG:-en_US.UTF-8}"
export LC_ALL="${LC_ALL:-en_US.UTF-8}"

# اختیاری: خروجی ساده بدون wrap تهاجمی
# export FAMAN_PLAIN=1
# export FAMAN_WRAP=0
```

برای locale فارسی:

```zsh
export LANG=fa_IR.UTF-8
export LC_ALL=fa_IR.UTF-8
```

### ۳) تکمیل خودکار (completion)

```zsh
mkdir -p ~/.zsh/completions
faman completion zsh > ~/.zsh/completions/_faman
```

در `~/.zshrc` (قبل از `compinit`):

```zsh
fpath=($HOME/.zsh/completions $fpath)
autoload -Uz compinit
compinit
```

اگر از **Oh My Zsh** استفاده می‌کنید، معمولاً `compinit` از قبل هست؛ فقط `fpath` را **قبل از** `source $ZSH/oh-my-zsh.sh` بگذارید:

```zsh
fpath=($HOME/.zsh/completions $fpath)

export ZSH="$HOME/.oh-my-zsh"
# ZSH_THEME=...
source $ZSH/oh-my-zsh.sh
```

سپس:

```zsh
exec zsh   # یا: source ~/.zshrc
faman <TAB>
faman search <TAB>
```

### ۴) Aliasهای مفید

```zsh
alias f='faman'
alias fs='faman search'
alias fp='FAMAN_PLAIN=1 faman'   # وقتی حروف خراب دیده می‌شود
```

### ۵) پیام خوش‌آمد اختیاری

```zsh
# فقط در نشست تعاملی
if [[ -o interactive ]] && command -v faman >/dev/null; then
  # quiet — چیزی چاپ نکن مگر بخواهی:
  # echo "faman $(faman version 2>/dev/null | head -1)"
fi
```

---

## نمونه بلوک آماده برای `~/.zshrc`

```zsh
# ── faman ──────────────────────────────────────────
export PATH="$HOME/.local/bin:$PATH"
export LANG="${LANG:-en_US.UTF-8}"
export LC_ALL="${LC_ALL:-en_US.UTF-8}"
# export FAMAN_PLAIN=1

fpath=($HOME/.zsh/completions $fpath)

alias f='faman'
alias fs='faman search'
alias fp='FAMAN_PLAIN=1 faman'
# ───────────────────────────────────────────────────
```

اگر Oh My Zsh دارید، این بلوک را **بالای** `source $ZSH/oh-my-zsh.sh` بگذارید و `compinit` جدا لازم نیست.

بدون Oh My Zsh، بعد از `fpath` این را هم بگذارید:

```zsh
autoload -Uz compinit && compinit
```

---

## عیب‌یابی

| مشکل | راه‌حل |
|------|--------|
| `command not found: faman` | PATH و `~/.local/bin` |
| TAB کار نمی‌کند | `fpath` + `compinit` + فایل `_faman` |
| حروف فارسی خراب | `LANG/LC_ALL` + فونت ترمینال + `FAMAN_PLAIN=1` |
| completion کهنه | دوباره `faman completion zsh > ~/.zsh/completions/_faman` |

تولید مجدد completion:

```zsh
faman completion zsh | tee ~/.zsh/completions/_faman >/dev/null
rm -f ~/.zcompdump*
compinit
```
