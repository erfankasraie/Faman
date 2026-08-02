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
  <img src="https://img.shields.io/badge/pages-124+-green" alt="pages"/>
</p>

> faman is **not** a translator — it **teaches** Linux commands in Persian.

---

## نصب آسان (یک خط)

### لینوکس / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash
```

با فونت فارسی + UTF-8:

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash -s -- --rtl
```

```bash
export PATH="$HOME/.local/bin:$PATH"
faman ls
```

### ویندوز (PowerShell)

```powershell
irm https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.ps1 | iex
```

جزئیات: [docs/install.md](docs/install.md) · [docs/windows.md](docs/windows.md) · نسخه **0.1.2-pre** · [Releases](https://github.com/erfankasraie/Faman/releases)

---

## استفاده

```bash
faman ls
faman apt
faman search docker
faman version
```

حروف خراب؟ `FAMAN_PLAIN=1 faman ls` — [terminal-persian.md](docs/terminal-persian.md)

---

## مدیران بسته

| دستور | کاربرد |
|--------|--------|
| `apt` `dpkg` | Debian / Ubuntu |
| `dnf` `rpm` | Fedora / RHEL |
| `pacman` `yay` | Arch / AUR |
| `zypper` | openSUSE |
| `apk` | Alpine |
| `snap` `flatpak` | بسته‌های اسنپ / فلت‌پک |
| `brew` `nix` | Homebrew / Nix |
| `pip` `npm` `cargo` | Python / Node / Rust |

فهرست کامل صفحات: [docs/pages-index.md](docs/pages-index.md)

---

## امکانات

- **۱۲۴+** صفحه فارسی
- جستجو، completion، zsh/fzf/bat
- لینوکس · ویندوز · macOS

## مستندات

[install](docs/install.md) · [windows](docs/windows.md) · [zsh](docs/zsh.md) · [ROADMAP](ROADMAP.md) · [CHANGELOG](CHANGELOG.md)

## License

MIT — [LICENSE](LICENSE)
