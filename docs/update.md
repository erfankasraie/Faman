# به‌روزرسانی faman و SHA256

## دستورها

| دستور | کار |
|--------|-----|
| `faman update` | وضعیت نسخه + راهنما |
| `faman update --check` | فقط مقایسه با GitHub |
| `faman update --pages` | صفحات از **main** + چاپ SHA256 آرشیو |
| `faman update --pages --verify` | آرشیو **آخرین Release** + تأیید با `SHA256SUMS` |
| `faman update --force` | اجازه بازنویسی مسیر سیستمی |

باینری با این دستور عوض نمی‌شود.

---

## SHA256 چگونه کار می‌کند؟

### ۱) سمت Release (ساخت)

`scripts/package.sh` و workflow ریلیز برای هر آرتیفکت فایل می‌نویسند:

```text
dist/SHA256SUMS
```

قالب GNU:

```text
<a hex>  faman-0.2.0-linux-amd64.tar.gz
<a hex>  faman-0.2.0-windows-amd64.zip
...
```

این فایل همراه ریلیز روی GitHub آپلود می‌شود.

### ۲) تأیید دستی بعد از دانلود از Releases

```bash
# لینوکس
cd ~/Downloads
sha256sum -c SHA256SUMS --ignore-missing
# یا فقط یک فایل:
sha256sum -c SHA256SUMS 2>/dev/null | grep OK
```

```powershell
# ویندوز (PowerShell 7+)
Get-FileHash .\faman-0.2.0-windows-amd64.zip -Algorithm SHA256
# با خط متناظر در SHA256SUMS مقایسه کنید
```

### ۳) تأیید خودکار صفحات (`--verify`)

```bash
faman update --pages --verify
```

مراحل داخلی:

1. تگ آخرین Release را از API می‌گیرد  
2. `SHA256SUMS` و آرشیو مناسب OS را دانلود می‌کند  
3. هش فایل را با `SHA256SUMS` مقایسه می‌کند  
4. فقط در صورت match، `pages/fa` را استخراج می‌کند  

بدون `--verify` منبع **main** است (هش هر commit عوض می‌شود؛ فقط چاپ می‌شود).

---

## مسیر نصب صفحات

- لینوکس/macOS: `~/.local/share/faman/pages/fa`
- ویندوز: `%LOCALAPPDATA%\faman\pages\fa`
- یا `FAMAN_PAGES`

```bash
export FAMAN_PAGES="$HOME/.local/share/faman/pages/fa"
```

---

## باینری

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash
# یا از Releases + sha256sum -c SHA256SUMS
```
