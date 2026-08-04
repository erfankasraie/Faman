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

> faman یک مترجم نیست؛ بلکه دستورات لینوکس را به زبان فارسی آموزش می‌دهد.

نسخه: **`0.1.4-pre`**

---

## نصب آسان

### لینوکس / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash
export PATH="$HOME/.local/bin:$PATH"
```

### ویندوز

```powershell
irm https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.ps1 | iex
```

### راهنمای هر سیستم

| سیستم | سند |
|--------|-----|
| Ubuntu، Fedora، Arch، Alpine، WSL، … | [docs/linux-distros.md](docs/linux-distros.md) |
| macOS (Intel و Apple Silicon) | [docs/macos.md](docs/macos.md) |
| ویندوز | [docs/windows.md](docs/windows.md) |
| ساخت `.deb` / `.exe` / آرشیو | [docs/packaging.md](docs/packaging.md) |

باینری آماده: [Releases](https://github.com/erfankasraie/Faman/releases)  
(`linux-amd64/arm64` · `darwin-amd64/arm64` · `windows-amd64.zip`)

---

## استفاده

```bash
faman ls
faman find
faman update --pages
faman version
```

حروف خراب؟ `FAMAN_PLAIN=1 faman ls` — [docs/terminal-persian.md](docs/terminal-persian.md)

---

## امکانات

- ۱۳۵+ صفحه فارسی
- `faman update --check` / `--pages`
- لینوکس · macOS · ویندوز

## مستندات بیشتر

[install](docs/install.md) · [update](docs/update.md) · [ROADMAP](ROADMAP.md) · [CHANGELOG](CHANGELOG.md)

## License

MIT — [LICENSE](LICENSE)
