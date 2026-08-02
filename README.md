<p align="center">
  <img src="assets/logo/faman-logo-full.svg" alt="faman logo" width="120"/>
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

### فقط faman

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash
```

### faman + فونت فارسی + UTF-8 + کمک RTL (پیشنهادی روی Ubuntu)

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash -s -- --with-rtl
```

### فقط فونت / locale / RTL (اگر faman را دارید)

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/setup-rtl.sh | bash
```

بعد از `--with-rtl` یک‌بار در **GNOME Terminal** فونت پروفایل را روی `Vazirmatn` یا `DejaVu Sans Mono` بگذارید (اسکریپت خودش GUI را عوض نمی‌کند).

راهنمای کامل نصب: [docs/install.md](docs/install.md)  
نمایش فارسی در ترمینال: [docs/terminal-persian.md](docs/terminal-persian.md)

---

## Quick Start

```bash
faman ls
faman search docker
faman grep
FAMAN_PLAIN=1 faman echo   # اگر حروف خراب دیده می‌شود
```

---

## Features

- `faman <command>` — راهنمای فارسی دستور
- `faman search <query>` — جستجو
- `faman version` / `help` / `update`
- بیش از ۱۰۰ صفحه فارسی
- رندر امن‌تر برای متن فارسی (`FAMAN_PLAIN` / `FAMAN_WRAP`)

---

## نصب دستی (بدون اسکریپت)

```bash
git clone https://github.com/erfankasraie/Faman.git
cd Faman
make build
sudo make install
# یا:
# bash scripts/install.sh --with-rtl
```

### وابستگی‌ها

| توزیع | دستور |
|--------|--------|
| Ubuntu / Debian | `sudo apt install git golang-go make` |
| Fedora | `sudo dnf install git golang make` |
| Arch | `sudo pacman -S git go make` |

اگر Go قدیمی بود: `sudo snap install go --classic`

### مسیر صفحات

1. `FAMAN_PAGES`
2. کنار باینری
3. `/usr/local/share/faman/pages/fa`
4. `/usr/share/faman/pages/fa`
5. پوشه جاری (توسعه)

---

## Uninstall

```bash
sudo rm -f /usr/local/bin/faman
sudo rm -rf /usr/local/share/faman
```

---

## مشارکت

[CONTRIBUTING.md](CONTRIBUTING.md) — قالب صفحات و استانداردها.

## Roadmap

جزئیات: [ROADMAP.md](ROADMAP.md)

## License

MIT — [LICENSE](LICENSE)

---

ساخته شده با ♥ برای جامعه لینوکس فارسی
