<p align="center">
  <img src="assets/logo/app-icon-light.png" alt="faman logo" width="120"/>
</p>

<h1 align="center">faman</h1>

<p align="center">
  <strong>صفحات راهنمای فارسی لینوکس — Persian Manual Pages</strong>
</p>

<p align="center">
  <a href="https://github.com/erfankasraie/Faman/actions/workflows/ci.yml"><img src="https://github.com/erfankasraie/Faman/actions/workflows/ci.yml/badge.svg" alt="CI"/></a>
  <a href="https://github.com/erfankasraie/Faman/releases"><img src="https://img.shields.io/github/v/release/erfankasraie/Faman?include_prereleases&label=release" alt="Release"/></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="MIT"/></a>
  <img src="https://img.shields.io/badge/pages-112+-green" alt="pages"/>
</p>

> faman is **not** a translator — it **teaches** Linux commands in Persian.

---

## نصب آسان (یک خط)

### لینوکس / macOS

کپی کن و بزن — نصب در `~/.local` (معمولاً **بدون sudo**):

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash
```

با فونت فارسی + UTF-8 (پیشنهادی روی Ubuntu):

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash -s -- --rtl
```

بعد یک‌بار ترمینال را باز/بسته کن، یا:

```bash
export PATH="$HOME/.local/bin:$PATH"
faman ls
```

### ویندوز (PowerShell)

```powershell
irm https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.ps1 | iex
```

سپس `faman.exe` را از `%LOCALAPPDATA%\faman\bin` اجرا کن (یا PATH را طبق پیام اسکریپت تنظیم کن).

### گزینه‌های بیشتر

| کار | دستور |
|-----|--------|
| نصب کاربر (پیش‌فرض get.sh) | `.../get.sh \| bash` |
| + فونت / RTL | `.../get.sh \| bash -s -- --rtl` |
| نصب سیستم‌واید | `.../get.sh \| bash -s -- --system` |
| حذف | `curl -fsSL .../install.sh \| bash -s -- --uninstall` |
| راهنمای کامل | [docs/install.md](docs/install.md) · [docs/windows.md](docs/windows.md) |

نسخه: **0.1.2-pre** · [Releases](https://github.com/erfankasraie/Faman/releases)

---

## استفاده

```bash
faman ls
faman search docker
faman grep
faman version
```

حروف خراب؟

```bash
FAMAN_PLAIN=1 faman ls
```

راهنمای ترمینال فارسی: [docs/terminal-persian.md](docs/terminal-persian.md)

---

## امکانات

- **۱۱۲+** صفحه فارسی برای دستورات لینوکس
- جستجو، completion، اسکریپت zsh/fzf/bat
- لینوکس · ویندوز · macOS
- لوگو فقط در CLI (ترجمه‌ها تمیز می‌مانند)

---

## مسیر صفحات

`FAMAN_PAGES` → کنار باینری → `/usr/local/share/faman/pages/fa` یا `%LOCALAPPDATA%\faman\pages\fa` → پوشهٔ توسعه

---

## مستندات

[install](docs/install.md) · [windows](docs/windows.md) · [zsh](docs/zsh.md) · [fzf](docs/fzf.md) · [ROADMAP](ROADMAP.md) · [CHANGELOG](CHANGELOG.md)

## ساخت از سورس

```bash
git clone https://github.com/erfankasraie/Faman.git && cd Faman
go build -o faman ./cmd/faman
export FAMAN_PAGES="$PWD/pages/fa"
./faman ls
```

## License

MIT — [LICENSE](LICENSE)
