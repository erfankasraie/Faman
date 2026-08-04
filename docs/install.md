# نصب faman

## سریع‌ترین راه (پیشنهادی)

نصب در `~/.local` — معمولاً **بدون sudo**:

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash
export PATH="$HOME/.local/bin:$PATH"
faman version
```

با فونت فارسی + UTF-8 (Ubuntu و مشابه):

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash -s -- --rtl
```

### ویندوز

```powershell
irm https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.ps1 | iex
```

راهنما: [windows.md](windows.md)

### از Releases

از [Releases](https://github.com/erfankasraie/Faman/releases) آرشیو لینوکس/ویندوز/macOS را بگیرید (باینری + `pages/fa`).

```bash
tar -xzf faman-*-linux-amd64.tar.gz
cd faman-*-linux-amd64
export FAMAN_PAGES="$PWD/pages/fa"
./faman ls
```

---

## به‌روزرسانی بعد از نصب

```bash
faman update --check    # نسخه محلی در برابر GitHub
faman update --pages    # فقط صفحات را از main تازه کن
# باینری:
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash
```

---

## اسکریپت کامل `install.sh`

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash
curl -fsSL .../install.sh | bash -s -- --with-rtl
bash scripts/install.sh --help
```

### فلگ‌ها

| فلگ | معنی |
|------|------|
| `--with-rtl` | faman + فونت + locale + کمک RTL |
| `--rtl-only` | فقط فونت/locale/RTL |
| `--skip-deps` | نصب git/go سیستمی را رد کن |
| `--user` | نصب در `~/.local` |
| `--prefix=DIR` | ریشه نصب |
| `--branch=NAME` | شاخه یا تگ git |
| `--plain-default` | `FAMAN_PLAIN=1` در shell rc |
| `--uninstall` | حذف |
| `--dry-run` | فقط چاپ |

### مثال‌ها

```bash
bash scripts/install.sh --user --with-rtl --plain-default
bash scripts/install.sh --from-dir=. --skip-deps
bash scripts/install.sh --user --uninstall
```

### کانفیگ دائمی

`~/.config/faman/install.env` — نمونه بعد از `--with-rtl` در `install.env.example`.

---

## بعد از نصب

### فونت ترمینال

GNOME: Preferences → Custom font → Vazirmatn / DejaVu Sans Mono  
جزئیات: [terminal-persian.md](terminal-persian.md)

### اگر حروف خراب است

```bash
export LANG=en_US.UTF-8
FAMAN_PLAIN=1 faman ls
```

## حذف

```bash
rm -f ~/.local/bin/faman && rm -rf ~/.local/share/faman
# یا سیستم‌واید:
sudo rm -f /usr/local/bin/faman && sudo rm -rf /usr/local/share/faman
```
