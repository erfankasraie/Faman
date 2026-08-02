# تنظیمات پیشرفته zsh برای faman

پایه: [zsh.md](zsh.md)

## ماژول آماده

فایل `scripts/zsh/faman.zsh` را کپی و source کنید:

```bash
mkdir -p ~/.config/faman
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/zsh/faman.zsh \
  -o ~/.config/faman/faman.zsh
```

در `~/.zshrc` (با Oh My Zsh: **قبل از** `source $ZSH/oh-my-zsh.sh` برای `fpath`، خود source می‌تواند بعد باشد):

```zsh
fpath=($HOME/.zsh/completions $fpath)
# … oh-my-zsh …
source ~/.config/faman/faman.zsh
```

---

## چه چیزهایی داخل ماژول است؟

| قابلیت | کلید / دستور | توضیح |
|---------|----------------|--------|
| Alias | `f` `fs` `fp` `fv` | میانبر |
| `fm cmd` | function | صفحه؛ اگر نبود search |
| `fmp cmd` | function | همیشه plain |
| `fls` | function | لیست نام صفحات |
| `faman-status` | function | تشخیص PATH/completion/locale |
| fzf صفحه | **Alt-f** | انتخاب فازی صفحه (نیاز به `fzf`) |
| `fzf-faman-search q` | function | جستجو + fzf |
| آخرین دستور | **Alt-m** | `faman <اولین‌کلمه-تاریخچه>` |

---

## fzf

```bash
# Ubuntu
sudo apt install fzf

# یا git
git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf && ~/.fzf/install
```

بعد از source کردن `faman.zsh`:

- **Alt-f** → لیست صفحات → Enter → `faman <page>`
- `fzf-faman-search docker`

اگر Alt-f با ترمینال تداخل داشت، در zshrc:

```zsh
bindkey '^F' _faman_fzf_page    # Ctrl-F
```

---

## Completion پیشرفته

```zsh
mkdir -p ~/.zsh/completions
faman completion zsh > ~/.zsh/completions/_faman

# کش کهنه
rm -f ~/.zcompdump*
autoload -Uz compinit && compinit
```

گزینه‌های مفید zsh completion:

```zsh
zstyle ':completion:*' menu select
zstyle ':completion:*' list-colors ''
zstyle ':completion:*:descriptions' format '%F{cyan}-- %d --%f'
# تکمیل حساس به حروف کوچک/بزرگ کمتر سخت‌گیر
zstyle ':completion:*' matcher-list 'm:{a-z}={A-Za-z}'
```

---

## Oh My Zsh — پلاگین محلی

```bash
mkdir -p ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/faman
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/zsh/faman.zsh \
  -o ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/faman/faman.plugin.zsh
```

در `~/.zshrc`:

```zsh
plugins=(git fzf faman)   # faman را اضافه کنید
```

و مطمئن شوید completion موجود است:

```zsh
mkdir -p ~/.zsh/completions
faman completion zsh > ~/.zsh/completions/_faman
fpath=($HOME/.zsh/completions $fpath)  # قبل از oh-my-zsh
```

---

## تاریخچه و راحتی

```zsh
HISTFILE=~/.zsh_history
HISTSIZE=50000
SAVEHIST=50000
setopt SHARE_HISTORY HIST_IGNORE_DUPS HIST_IGNORE_SPACE
```

جستجوی تاریخچه با fzf (اگر از اسکریپت fzf استفاده می‌کنید `Ctrl-R` از قبل هست).

---

## حالت «هر دستور ناشناس → faman» (آگاهانه)

داخل `faman.zsh` بخش `command_not_found_handler` به‌صورت خاموش است. فعال کردن:

```zsh
command_not_found_handler() {
  if faman "$1" >/dev/null 2>&1; then
    faman "$1"
    return 0
  fi
  print -u2 "zsh: command not found: $1"
  return 127
}
```

فقط وقتی sensible است که اشتباه تایپ دستور لینوکس را به صفحهٔ راهنما وصل کنید؛ برای همهٔ commandهای غایب ممکن است مزاحم باشد.

---

## تشخیص مشکل

```zsh
faman-status
```

خروجی نمونه: مسیر باینری، نسخه، `LANG`، `FAMAN_PLAIN`، وضعیت `_faman`.

---

## جمع‌بندی کپی‌پیست

```zsh
# ~/.zshrc — ترتیب مهم است
export PATH="$HOME/.local/bin:$PATH"
fpath=($HOME/.zsh/completions $fpath)

# Oh My Zsh (اختیاری)
# export ZSH="$HOME/.oh-my-zsh"
# plugins=(git fzf faman)
# source $ZSH/oh-my-zsh.sh

autoload -Uz compinit && compinit   # اگر OMZ ندارید

source ~/.config/faman/faman.zsh
```
