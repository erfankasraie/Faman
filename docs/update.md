# به‌روزرسانی faman

## خلاصه

| دستور | کار |
|--------|-----|
| `faman update` | وضعیت نسخه + راهنما |
| `faman update --check` | فقط مقایسه با GitHub |
| `faman update --pages` | دانلود `pages/fa` از `main` |

باینری توسط این دستور **جایگزین نمی‌شود**.

## صفحات

```bash
faman update --pages
```

مقصد پیش‌فرض:

- لینوکس/macOS: `~/.local/share/faman/pages/fa`
- ویندوز: `%LOCALAPPDATA%\faman\pages\fa`
- اگر `FAMAN_PAGES` روی مسیر قابل‌نوشتن باشد، همان‌جا

نصب سیستم‌واید (`/usr/local/share/...`) بدون `--force` بازنویسی نمی‌شود؛ به‌جای آن مسیر کاربر پر می‌شود. بعداً:

```bash
export FAMAN_PAGES="$HOME/.local/share/faman/pages/fa"
```

## باینری

```bash
# لینوکس/macOS
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash

# ویندوز
irm https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.ps1 | iex
```

یا آرتیفکت از [Releases](https://github.com/erfankasraie/Faman/releases).

## نسخه

`faman version` نسخهٔ لینک‌شده در باینری را نشان می‌دهد.  
`faman update --check` آخرین **tag/release** روی GitHub را می‌خواند.

## محدودیت‌های فعلی

- بدون امضای GPG / بدون تأیید SHA اجباری روی آرشیو صفحات
- بدون self-update خودکار باینری
- `--pages` همیشه از شاخه **main** می‌گیرد (نه الزاماً همان تگ release)

این موارد در رودمپ فاز ۳ هستند.
