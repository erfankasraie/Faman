<p align="center">
  <img src="assets/logo/app-icon-light.png" alt="faman logo" width="120"/>
</p>

<h1 align="center">faman</h1>

<p align="center">
  <strong>صفحات راهنمای فارسی لینوکس — Persian Manual Pages</strong>
</p>

<p align="center">
  <a href="https://github.com/erfankasraie/Faman/actions/workflows/ci.yml"><img src="https://github.com/erfankasraie/Faman/actions/workflows/ci.yml/badge.svg" alt="CI"/></a>
  <a href="https://github.com/erfankasraie/Faman/releases"><img src="https://img.shields.io/github/v/release/erfankasraie/Faman?label=release" alt="Release"/></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="MIT"/></a>
  <img src="https://img.shields.io/badge/pages-176+-green" alt="pages"/>
  <img src="https://img.shields.io/badge/version-0.2.2-blue" alt="version"/>
</p>

> faman is **not** a translator — it **teaches** Linux commands in Persian.


---

## نصب آسان

### لینوکس / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash
export PATH="$HOME/.local/bin:$PATH"
```

### Debian / Ubuntu (بعد از ریلیز v0.2.2)

```bash
sudo dpkg -i faman_0.2.2_amd64.deb
```

### ویندوز

```powershell
irm https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.ps1 | iex
```

آرتیفکت‌ها: [Releases](https://github.com/erfankasraie/Faman/releases)

| سیستم | سند |
|--------|-----|
| لینوکس | [docs/linux-distros.md](docs/linux-distros.md) |
| macOS | [docs/macos.md](docs/macos.md) |
| ویندوز | [docs/windows.md](docs/windows.md) |
| zsh / Tab | [docs/zsh.md](docs/zsh.md) · [docs/completion.md](docs/completion.md) |
| بسته‌بندی | [docs/packaging.md](docs/packaging.md) |

---

## استفاده

```bash
faman ls
faman list --cat network
faman search port --cat network
faman categories
faman random
faman doctor
faman update --pages --verify
faman version
```

---

## امکانات

- ۱۷۶+ صفحه فارسی
- `list` · `categories` · `random` · `doctor` · `update` (+ SHA256 verify)
- completion نام صفحات · `search --cat`
- tar.gz · zip · **`.deb`** · لینوکس / macOS / ویندوز

[CHANGELOG](CHANGELOG.md) · [ROADMAP](ROADMAP.md) · [ریلیز 0.2.2](docs/release-0.2.2.md)

## License

MIT — [LICENSE](LICENSE)
