<p align="center">
  <img src="assets/logo/app-icon-light.png" alt="faman logo" width="120"/>
</p>

<h1 align="center">faman</h1>

<p align="center">
  <strong>صفحات راهنمای فارسی لینوکس — Persian Manual Pages</strong>
</p>

<p align="center">
  <a href="https://github.com/erfankasraie/Faman/actions"><img src="https://github.com/erfankasraie/Faman/actions/workflows/ci.yml/badge.svg" alt="CI"/></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="MIT"/></a>
</p>

> faman is **not** a translator.  
> faman **teaches** Linux commands in Persian.

---

## نصب با یک کپی‌پیست

### Linux / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash
# با فونت و RTL:
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash -s -- --with-rtl
```

### Windows (PowerShell) — همان `main`

```powershell
irm https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.ps1 | iex
```

راهنما: [docs/windows.md](docs/windows.md) · فارسی در ترمینال: [docs/terminal-persian.md](docs/terminal-persian.md) · نصب کامل: [docs/install.md](docs/install.md)

---

## Quick Start

```bash
faman ls
faman search docker
faman grep
FAMAN_PLAIN=1 faman echo
```

---

## Features

- `faman <command>` — راهنمای فارسی دستور
- `faman search` / `version` / `help` / `completion`
- ۱۰۰+ صفحه فارسی
- لینوکس، macOS و **ویندوز** (مسیر LocalAppData + Windows Terminal)
- `FAMAN_PLAIN` / `FAMAN_WRAP` برای ترمینال‌های ضعیف

---

## مسیر صفحات

1. `FAMAN_PAGES`
2. کنار باینری `pages/fa`
3. لینوکس: `/usr/local/share/faman/pages/fa` — ویندوز: `%LOCALAPPDATA%\faman\pages\fa`
4. پوشه جاری (توسعه)

---

## مشارکت

[CONTRIBUTING.md](CONTRIBUTING.md) · [ROADMAP.md](ROADMAP.md)

## License

MIT — [LICENSE](LICENSE)
