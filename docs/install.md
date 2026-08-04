# نصب faman

## انتخاب پلتفرم

| سیستم | راهنما |
|--------|--------|
| همهٔ لینوکس‌ها (جدول distro) | [linux-distros.md](linux-distros.md) |
| macOS (Intel / Apple Silicon) | [macos.md](macos.md) |
| ویندوز / WSL | [windows.md](windows.md) |
| ساخت بسته `.deb` / `.exe` / آرشیو | [packaging.md](packaging.md) |

---

## سریع‌ترین راه

### لینوکس و macOS

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash
export PATH="$HOME/.local/bin:$PATH"
faman version
```

لینوکس + فونت/UTF-8:

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash -s -- --rtl
```

### ویندوز (PowerShell)

```powershell
irm https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.ps1 | iex
```

### از Releases (باینری آماده)

از [Releases](https://github.com/erfankasraie/Faman/releases):

| فایل | سیستم |
|------|--------|
| `*-linux-amd64.tar.gz` | لینوکس 64-bit معمولی |
| `*-linux-arm64.tar.gz` | لینوکس ARM64 |
| `*-darwin-arm64.tar.gz` | مک Apple Silicon |
| `*-darwin-amd64.tar.gz` | مک Intel |
| `*-windows-amd64.zip` | ویندوز (`faman.exe`) |
| `*_amd64.deb` | اوبونتو/دبیان (اگر ضمیمه شده باشد) |

**لینوکس / مک:**

```bash
tar -xzf faman-*-linux-amd64.tar.gz   # یا darwin-arm64 و …
cd faman-*
export FAMAN_PAGES="$PWD/pages/fa"
./faman ls
```

**ویندوز:**

```powershell
Expand-Archive faman-*-windows-amd64.zip
cd faman-*-windows-amd64
$env:FAMAN_PAGES = "$PWD\pages\fa"
.\faman.exe ls
```

---

## به‌روزرسانی

```bash
faman update --check
faman update --pages
# باینری: دوباره get.sh / install.ps1 / دانلود آرشیو
```

جزئیات: [update.md](update.md)

---

## اسکریپت کامل `install.sh`

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash -s -- --user --with-rtl
bash scripts/install.sh --help
```

فلگ‌ها: `--user`, `--with-rtl`, `--prefix`, `--plain-default`, `--uninstall`, `--dry-run`, …

---

## بعد از نصب

فونت و UTF-8: [terminal-persian.md](terminal-persian.md)

```bash
FAMAN_PLAIN=1 faman ls   # اگر حروف خراب است
```

## حذف

```bash
rm -f ~/.local/bin/faman && rm -rf ~/.local/share/faman
sudo rm -f /usr/local/bin/faman && sudo rm -rf /usr/local/share/faman
```
