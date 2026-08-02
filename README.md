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
  <img src="https://img.shields.io/badge/platform-Linux%20%7C%20Windows%20%7C%20macOS-lightgrey" alt="platforms"/>
</p>

> faman is **not** a translator.  
> faman **teaches** Linux commands in Persian.

نسخه فعلی: **`0.1.2-pre`** (پیش‌انتشار)

---

## نصب

### Linux / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash
```

با فونت فارسی، locale UTF-8 و کمک RTL (پیشنهادی روی Ubuntu):

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash -s -- --with-rtl
```

### Windows (PowerShell)

```powershell
irm https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.ps1 | iex
```

### باینری آماده (Pre-release)

از صفحهٔ [Releases](https://github.com/erfankasraie/Faman/releases) یکی از آرشیوها را بگیرید (هر کدام شامل باینری + `pages/fa` است):

| فایل | پلتفرم |
|------|--------|
| `faman-*-linux-amd64.tar.gz` | لینوکس x86_64 |
| `faman-*-linux-arm64.tar.gz` | لینوکس ARM64 |
| `faman-*-windows-amd64.zip` | ویندوز x64 |
| `faman-*-darwin-*.tar.gz` | macOS |

**لینوکس:**

```bash
tar -xzf faman-*-linux-amd64.tar.gz
cd faman-*-linux-amd64
export FAMAN_PAGES="$PWD/pages/fa"
./faman version
./faman ls
```

**ویندوز:**

```powershell
Expand-Archive faman-*-windows-amd64.zip
cd faman-*-windows-amd64
$env:FAMAN_PAGES = "$PWD\pages\fa"
.\faman.exe version
.\faman.exe ls
```

انتشار خودکار: Actions → workflow **Release** (تگ `v*` یا Run workflow).

---

## Quick Start

```bash
faman ls
faman search docker
faman grep
faman version
faman help
```

اگر حروف فارسی خراب دیده شد:

```bash
FAMAN_PLAIN=1 faman ls
# یا
export FAMAN_WRAP=0
```

---

## امکانات

- راهنمای فارسی دستورات لینوکس (`faman <cmd>`)
- جستجو (`faman search`)
- **۱۱۲+** صفحه Markdown فارسی با قالب یکسان
- لوگو و برندینگ فقط در CLI (نیازی به لوگو داخل ترجمه‌ها نیست)
- لینوکس، **ویندوز** و macOS روی یک شاخهٔ `main`
- تکمیل شل (Cobra completion) + اسکریپت‌های zsh / fzf / bat
- نصب‌کنندهٔ پیشرفته با فلگ‌های `--user`، `--with-rtl`، `--plain-default` و …

---

## متغیرهای محیطی

| متغیر | معنی |
|--------|------|
| `FAMAN_PAGES` | مسیر پوشهٔ `pages/fa` |
| `FAMAN_PLAIN=1` | خروجی ساده بدون رنگ/باکس |
| `FAMAN_WRAP=0` | بدون wrap (امن‌تر برای فارسی) |
| `FAMAN_LOGO=0` | مخفی کردن بنر چندخطی |
| `NO_COLOR` | خاموش کردن رنگ |

---

## مسیر جستجوی صفحات

1. `FAMAN_PAGES`
2. `pages/fa` کنار باینری
3. لینوکس: `/usr/local/share/faman/pages/fa` · ویندوز: `%LOCALAPPDATA%\faman\pages\fa`
4. `~/.local/share/faman/pages/fa` / config کاربر
5. پوشهٔ جاری (توسعه)

---

## مستندات

| موضوع | لینک |
|--------|------|
| نصب | [docs/install.md](docs/install.md) |
| ویندوز | [docs/windows.md](docs/windows.md) |
| فارسی در ترمینال | [docs/terminal-persian.md](docs/terminal-persian.md) |
| zsh | [docs/zsh.md](docs/zsh.md) · [zsh-advanced](docs/zsh-advanced.md) |
| fzf / bat | [docs/fzf.md](docs/fzf.md) · [bat-preview](docs/bat-preview.md) |
| رودمپ | [ROADMAP.md](ROADMAP.md) |
| تغییرات | [CHANGELOG.md](CHANGELOG.md) |
| انتشار | [.github/RELEASE_INSTRUCTIONS.md](.github/RELEASE_INSTRUCTIONS.md) |

---

## ساخت از سورس

```bash
git clone https://github.com/erfankasraie/Faman.git
cd Faman
go build -o faman ./cmd/faman
export FAMAN_PAGES="$PWD/pages/fa"
./faman ls
```

نیاز: Go **1.22+**

---

## مشارکت

صفحات جدید با قالب ۷بخشی در `pages/fa/` — جزئیات در [CONTRIBUTING.md](CONTRIBUTING.md).

---

## License

MIT — [LICENSE](LICENSE)

---

ساخته شده با ♥ برای جامعه لینوکس فارسی
