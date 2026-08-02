# تنظیمات پیش‌نمایش bat برای faman + fzf

## نصب bat

```bash
sudo apt install bat
```

روی Ubuntu گاهی دستور `batcat` است:

```bash
# اختیاری
mkdir -p ~/.local/bin
ln -sf /usr/bin/batcat ~/.local/bin/bat
```

ماژول‌های faman هر دو نام را می‌شناسند.

---

## بارگذاری

```bash
mkdir -p ~/.config/faman
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/zsh/faman-bat.zsh \
  -o ~/.config/faman/faman-bat.zsh
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/zsh/faman-fzf.zsh \
  -o ~/.config/faman/faman-fzf.zsh
```

`~/.zshrc`:

```zsh
# اختیاری — قبل از fzf
export FAMAN_BAT_THEME=TwoDark
export FAMAN_BAT_STYLE=plain,header,grid
export FAMAN_BAT_LINES=120

source ~/.config/faman/faman-bat.zsh   # اختیاری
source ~/.config/faman/faman-fzf.zsh
```

---

## متغیرهای محیطی

| متغیر | پیش‌فرض | معنی |
|--------|---------|------|
| `FAMAN_BAT_THEME` | `ansi` | تم رنگ bat |
| `FAMAN_BAT_STYLE` | `plain,header,grid` | اجزای UI |
| `FAMAN_BAT_LINES` | `120` | حداکثر خط در preview |
| `FAMAN_BAT_WRAP` | `character` | نحوهٔ شکستن خط |
| `FAMAN_BAT_EXTRA` | خالی | فلگ اضافه برای bat |
| `FAMAN_FZF_PREVIEW_WINDOW` | `right:55%:wrap` | پنل fzf |
| `FAMAN_FZF_OPTS` | خالی | فلگ اضافه fzf |

### تم‌های رایج

```zsh
faman-bat-themes          # لیست کامل (اگر faman-bat.zsh لود شده)

export FAMAN_BAT_THEME=ansi
export FAMAN_BAT_THEME=TwoDark
export FAMAN_BAT_THEME=Dracula
export FAMAN_BAT_THEME=GitHub
export FAMAN_BAT_THEME=Monokai Extended
export FAMAN_BAT_THEME=Solarized (dark)
```

تعویض در همان نشست:

```zsh
faman-bat-theme Dracula
```

### استایل

مقادیر قابل ترکیب با ویرگول:

- `plain` — بدون خط شماره اجباری شلوغ
- `header` — نام فایل
- `grid` — قاب
- `numbers` — شماره خط
- `changes` — برای diff

مثال مینیمال:

```zsh
export FAMAN_BAT_STYLE=plain
```

مثال کامل‌تر:

```zsh
export FAMAN_BAT_STYLE=numbers,header,grid
```

### فقط چند خط اول

```zsh
export FAMAN_BAT_LINES=40
```

### پنل بزرگ‌تر در fzf

```zsh
export FAMAN_FZF_PREVIEW_WINDOW='right:65%:wrap'
# یا پایین:
export FAMAN_FZF_PREVIEW_WINDOW='down:40%:wrap'
```

---

## کلیدها داخل fzf

| کلید | کار |
|------|-----|
| Enter | `faman <page>` |
| Ctrl-P | plain |
| Ctrl-B | باز کردن کامل با bat در ترمینال (pager) |
| Ctrl-Y | کپی نام |

---

## پیش‌نمایش مستقیم بدون fzf

```zsh
faman-bat ls          # یا: fbat grep
```

---

## نمونه بلوک پیشنهادی (تیره)

```zsh
export FAMAN_BAT_THEME=TwoDark
export FAMAN_BAT_STYLE=plain,header
export FAMAN_BAT_LINES=100
export FAMAN_FZF_PREVIEW_WINDOW='right:58%:wrap'
export FAMAN_FZF_OPTS='--info=inline --pointer=▶'
```

## نمونه (روشن / GitHub)

```zsh
export FAMAN_BAT_THEME=GitHub
export FAMAN_BAT_STYLE=numbers,grid
```

---

## عیب‌یابی

| مشکل | کار |
|------|-----|
| پیش‌نمایش بدون رنگ | نصب `bat` / `batcat` |
| `bat: command not found` | `alias bat=batcat` یا لینک در `~/.local/bin` |
| تم نامعتبر | `bat --list-themes` |
| فارسی در preview به‌هم‌ریخته | محدودیت ترمینال؛ `FAMAN_BAT_STYLE=plain` و فونت مناسب |

مرتبط: [fzf.md](fzf.md)
