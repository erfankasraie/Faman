# Changelog

## [Unreleased]

### Content
- ۲۸ صفحه جدید برای پوشش بیشتر دستورات رایج man لینوکس:
  - سیستم/سخت‌افزار: `dmesg`, `lscpu`, `lsusb`, `lspci`, `cal`
  - پردازش: `htop`, `killall`, `nice`, `renice`, `strace`, `at`
  - متن: `printf`, `paste`, `join`, `comm`, `column`, `split`, `bc`
  - فایل‌سیستم: `dd`, `mktemp`, `readlink`, `shred`
  - شبکه/امنیت: `openssl`, `ssh-keygen`, `gpg`, `whois`, `nslookup`, `netstat`
- مجموع صفحات از ۱۳۵+ به ۱۶۳+ رسید

## [0.1.4-pre] — 2026-08-04

### Added
- **`faman update` واقعی**
  - `--check` — مقایسه نسخه با GitHub (release/tag)
  - `--pages` — دانلود `pages/fa` از شاخه `main` به مسیر کاربر
  - راهنمای به‌روزرسانی باینری (get.sh / install.ps1 / Releases)
- عمق man برای دستورات سنگین: `find`, `grep`, `sed`, `awk`, `tar`, `rsync`, `curl`
- مدیران بسته: `dpkg`, `snap`, `flatpak`, `zypper`, `apk`, `brew`, `rpm`, `yay`, `nix`, `pip`, `npm`, `cargo`
- محیط مجازی: `venv`, `virtualenv`, `conda`, `pyenv`, `nvm`, `asdf`, `direnv`, `poetry`, `pipenv`, `podman`, cookbook `examples-environments`
- نصب آسان: `scripts/get.sh` (پیش‌فرض `--user`)

### Docs
- README، ROADMAP، pages-index، install، windows هم‌تراز با وضعیت فعلی

### Content
- حدود **۱۳۵+** صفحه فارسی

## [0.1.3-pre] — 2026-08-0x

- تگ/پیش‌انتشار میانی روی GitHub

## [0.1.2-pre] — 2026-08-02

### Added
- پشتیبانی ویندوز روی `main`، `install.ps1`، `docs/windows.md`
- لوگوی مرکزی CLI
- zsh / fzf / bat
- installer با RTL و فلگ‌های پیشرفته
- workflow انتشار کراس‌پلتفرم

### Fixed
- CI golangci-lint
- رندر امن‌تر فارسی

## [0.1.0] — earlier

- CLI اولیه، parser، renderer، search، CI
