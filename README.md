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
  <img src="https://img.shields.io/badge/pages-135+-green" alt="pages"/>
</p>

> faman is **not** a translator — it **teaches** Linux commands in Persian.

نسخه: **`0.1.4-pre`**

---

## نصب آسان

### لینوکس / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash
# با فونت و RTL:
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash -s -- --rtl
export PATH="$HOME/.local/bin:$PATH"
```

### ویندوز (PowerShell)

```powershell
irm https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.ps1 | iex
```

باینری‌های آماده: [Releases](https://github.com/erfankasraie/Faman/releases)

جزئیات: [docs/install.md](docs/install.md) · [docs/windows.md](docs/windows.md)

---

## استفاده

```bash
faman ls
faman find
faman search docker
faman update --check
faman update --pages
faman version
```

حروف خراب؟ `FAMAN_PLAIN=1 faman ls` — [docs/terminal-persian.md](docs/terminal-persian.md)

---

## به‌روزرسانی

| هدف | دستور |
|------|--------|
| بررسی نسخه | `faman update` یا `faman update --check` |
| تازه‌سازی صفحات | `faman update --pages` |
| باینری جدید | دوباره `get.sh` / `install.ps1` یا دانلود از Releases |

صفحات معمولاً در `~/.local/share/faman/pages/fa` (لینوکس) یا `%LOCALAPPDATA%\faman\pages\fa` (ویندوز) نصب می‌شوند.

---

## امکانات

- **۱۳۵+** صفحه فارسی (از جمله manهای عمیق مثل find/grep/tar/rsync/curl)
- مدیران بسته و محیط مجازی (apt…cargo، venv…poetry)
- جستجو، completion، zsh/fzf/bat
- لینوکس · ویندوز · macOS

فهرست: [docs/pages-index.md](docs/pages-index.md)

---

## مستندات

| موضوع | لینک |
|--------|------|
| نصب | [docs/install.md](docs/install.md) |
| ویندوز | [docs/windows.md](docs/windows.md) |
| فارسی در ترمینال | [docs/terminal-persian.md](docs/terminal-persian.md) |
| zsh / fzf | [docs/zsh.md](docs/zsh.md) · [docs/fzf.md](docs/fzf.md) |
| رودمپ | [ROADMAP.md](ROADMAP.md) |
| تغییرات | [CHANGELOG.md](CHANGELOG.md) |
| انتشار | [.github/RELEASE_INSTRUCTIONS.md](.github/RELEASE_INSTRUCTIONS.md) |

---

## ساخت از سورس

```bash
git clone https://github.com/erfankasraie/Faman.git && cd Faman
go build -ldflags "-X github.com/faman-project/faman/internal/app.version=0.1.4-pre" -o faman ./cmd/faman
export FAMAN_PAGES="$PWD/pages/fa"
./faman version
```

نیاز: Go **1.22+**

## License

MIT — [LICENSE](LICENSE)
