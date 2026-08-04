# تکمیل خودکار (completion) — خطاهای رایج

## نصب درست

```bash
mkdir -p ~/.zsh/completions
faman completion zsh > ~/.zsh/completions/_faman
```

در `~/.zshrc` **قبل از** Oh My Zsh / compinit:

```zsh
fpath=($HOME/.zsh/completions $fpath)
# بدون OMZ:
# autoload -Uz compinit && compinit
```

```zsh
exec zsh
faman <TAB>          # نام صفحات
faman list --cat <TAB>
faman search --cat <TAB>
```

با ماژول پیشرفته:

```zsh
source ~/.config/faman/faman.zsh
faman-recomplete     # بعد از هر آپدیت باینری
faman-status         # تشخیص مشکل
```

---

## خطاهای رایج و راه‌حل

| مشکل | علت | راه‌حل |
|------|-----|--------|
| Tab هیچی نشان نمی‌دهد | `fpath` بعد از `compinit` / بعد از OMZ | بلوک faman را **بالای** `source $ZSH/oh-my-zsh.sh` بگذار |
| فقط فایل‌های دایرکتوری complete می‌شود | completion نصب نشده یا cache کهنه | `faman-recomplete` یا حذف `~/.zcompdump*` و `compinit` |
| دستورهای جدید (`list`, `doctor`) نیست | فایل `_faman` از نسخهٔ قدیمی است | دوباره `faman completion zsh > ~/.zsh/completions/_faman` |
| `command not found: faman` موقع complete | باینری در PATH شل تعاملی نیست | `export PATH="$HOME/.local/bin:$PATH"` در zshrc |
| insecure directories | مالکیت `~/.zsh` برای user دیگر | `compaudit` و `chmod`/`chown`؛ یا `compinit -u` |
| bash: complete کار نمی‌کند | bash-completion فعال نیست | `source /usr/share/bash-completion/bash_completion` |
| Fish / PowerShell | فقط اسکریپت generate شده | `faman completion fish > …` طبق `faman completion -h` |
| در CI / اسکریپت غیرتعاملی | completion فقط برای شل تعاملی | طبیعی است؛ نیازی نیست |

### تشخیص سریع

```zsh
echo $fpath | tr ' ' '\n' | grep completions
ls -la ~/.zsh/completions/_faman
head -5 ~/.zsh/completions/_faman
which faman
faman-status   # اگر ماژول source شده
```

### Oh My Zsh

1. `fpath=($HOME/.zsh/completions $fpath)` قبل از `source $ZSH/oh-my-zsh.sh`
2. نیازی به `compinit` دستی نیست
3. بعد از عوض کردن `_faman`: `rm -f ~/.zcompdump* && exec zsh`

### پیش‌فرض deb

بستهٔ `.deb` completion باینری bash را در `/usr/share/bash-completion/completions/faman` می‌گذارد. برای zsh همچنان `install-shell.sh` یا `faman-recomplete` را بزنید.

---

## چه چیزی complete می‌شود؟

- آرگومان اصلی: نام صفحات (`faman ls`)
- `search` / `list`: پیشنهاد عنوان و `--cat` دسته‌ها
- زیرفرمان‌ها: `list`, `categories`, `random`, `doctor`, `update`, …
