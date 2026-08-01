<p align="center">
  <img src="assets/logo/faman-logo-full.svg" alt="faman logo" width="120"/>
</p>

<h1 align="center">faman</h1>

<p align="center">
  <strong>صفحات راهنمای فارسی لینوکس — Persian Manual Pages</strong>
</p>

<p align="center">
  <img src="assets/logo/faman-icon-24.svg" alt="faman icon" width="24"/>
  &nbsp;
  <a href="https://github.com/faman-project/faman/actions"><img src="https://github.com/faman-project/faman/actions/workflows/ci.yml/badge.svg" alt="CI"/></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="MIT"/></a>
  <a href="https://goreportcard.com/report/github.com/faman-project/faman"><img src="https://goreportcard.com/badge/github.com/faman-project/faman" alt="Go Report Card"/></a>
</p>

> faman is **not** a translator.  
> faman **teaches** Linux commands in Persian.

هر صفحه طوری نوشته شده که هم برای مبتدی قابل فهم باشد و هم برای حرفه‌ای‌ها مفید بماند.

---

## Vision

ساخت معادل فارسی و مدرن صفحات man لینوکس — با تمرکز روی یادگیری واقعی، مثال‌های کاربردی، اشتباهات رایج و نکات حرفه‌ای.

پروژه کاملاً متن‌باز است و انتظار می‌رود با مشارکت جامعه رشد کند.

---

## Features

- `faman <command>` — نمایش راهنمای فارسی دستور
- `faman search <query>` — جستجو در تمام صفحات
- `faman version` — نمایش نسخه
- `faman help` — راهنما
- `faman update` — به‌روزرسانی (placeholder برای آینده)
- خروجی زیبا با رنگ‌ها و پشتیبانی از عرض ترمینال
- معماری تمیز و قابل توسعه
- تست‌ها و CI کامل

---

## Installation

### روش سریع (پیشنهادی برای اکثر کاربران)

```bash
# کلون و نصب
git clone https://github.com/faman-project/faman.git
cd faman
make build
sudo install -m 755 faman /usr/local/bin/faman
sudo mkdir -p /usr/local/share/faman
sudo cp -r pages /usr/local/share/faman/
```

بعد از این دستورات، `faman` از هر جایی قابل اجراست.

---

### نصب روی توزیع‌های مختلف لینوکس

#### Debian / Ubuntu / Linux Mint / Pop!_OS

```bash
# وابستگی‌ها
sudo apt update
sudo apt install -y git golang-go make

# نصب faman
git clone https://github.com/faman-project/faman.git
cd faman
make build
sudo install -m 755 faman /usr/local/bin/faman
sudo mkdir -p /usr/local/share/faman
sudo cp -r pages /usr/local/share/faman/
```

> **نکته:** در Ubuntu نسخه‌های قدیمی ممکن است نسخه Go پایین باشد. در این صورت از [نصب رسمی Go](https://go.dev/dl/) استفاده کنید یا:
> ```bash
> sudo snap install go --classic
> ```

#### Fedora / RHEL / CentOS Stream / Rocky / AlmaLinux

```bash
sudo dnf install -y git golang make

git clone https://github.com/faman-project/faman.git
cd faman
make build
sudo install -m 755 faman /usr/local/bin/faman
sudo mkdir -p /usr/local/share/faman
sudo cp -r pages /usr/local/share/faman/
```

#### Arch Linux / Manjaro / EndeavourOS

```bash
sudo pacman -S --needed git go make

git clone https://github.com/faman-project/faman.git
cd faman
make build
sudo install -m 755 faman /usr/local/bin/faman
sudo mkdir -p /usr/local/share/faman
sudo cp -r pages /usr/local/share/faman/
```

> در آینده بسته AUR هم اضافه خواهد شد (`yay -S faman`).

#### openSUSE

```bash
sudo zypper install -y git go make

git clone https://github.com/faman-project/faman.git
cd faman
make build
sudo install -m 755 faman /usr/local/bin/faman
sudo mkdir -p /usr/local/share/faman
sudo cp -r pages /usr/local/share/faman/
```

#### Alpine Linux

```bash
sudo apk add git go make

git clone https://github.com/faman-project/faman.git
cd faman
make build
sudo install -m 755 faman /usr/local/bin/faman
sudo mkdir -p /usr/local/share/faman
sudo cp -r pages /usr/local/share/faman/
```

#### Gentoo

```bash
sudo emerge --ask dev-vcs/git dev-lang/go

git clone https://github.com/faman-project/faman.git
cd faman
make build
sudo install -m 755 faman /usr/local/bin/faman
sudo mkdir -p /usr/local/share/faman
sudo cp -r pages /usr/local/share/faman/
```

---

### نصب فقط با Go (بدون make)

```bash
git clone https://github.com/faman-project/faman.git
cd faman
go build -o faman ./cmd/faman
sudo install -m 755 faman /usr/local/bin/faman
sudo mkdir -p /usr/local/share/faman
sudo cp -r pages /usr/local/share/faman/
```

### با `go install`

```bash
go install github.com/faman-project/faman/cmd/faman@latest
```

**مهم:** با این روش فقط باینری نصب می‌شود. صفحات را جداگانه کپی کنید:

```bash
git clone --depth 1 https://github.com/faman-project/faman.git /tmp/faman
sudo mkdir -p /usr/local/share/faman
sudo cp -r /tmp/faman/pages /usr/local/share/faman/
rm -rf /tmp/faman
```

---

### مسیرهای جستجوی صفحات

faman صفحات را به ترتیب زیر پیدا می‌کند:

1. متغیر محیطی `FAMAN_PAGES`
2. کنار فایل اجرایی
3. `/usr/local/share/faman/pages/fa`
4. `/usr/share/faman/pages/fa`
5. `/opt/faman/pages/fa`
6. پوشه جاری (حالت توسعه)

مثال تنظیم دستی:

```bash
export FAMAN_PAGES=/path/to/pages/fa
faman ls
```

---

## Quick Start

```bash
faman ls
faman search docker
faman grep
faman version
faman help
```

---

## Examples

```bash
# راهنمای دستور ls
faman ls

# جستجو
faman search "فایل"
faman search permission

# نسخه
faman version
```

---

## Uninstall

```bash
sudo rm -f /usr/local/bin/faman
sudo rm -rf /usr/local/share/faman
```

---

## Project Structure

```
faman/
├── cmd/faman/          # نقطه ورود
├── internal/
│   ├── app/            # دستورات CLI (cobra)
│   ├── parser/         # پارسر markdown + front matter
│   ├── renderer/       # رندر زیبا با lipgloss/glamour
│   ├── search/         # موتور جستجو
│   └── update/         # به‌روزرسانی (placeholder)
├── pages/fa/           # صفحات راهنما (Markdown)
├── assets/
│   └── logo/           # لوگو و آیکون‌ها
├── scripts/
├── docs/
├── .github/workflows/
└── ...
```

---

## How to contribute

ما از مشارکت استقبال می‌کنیم!

1. یک Issue باز کنید یا یک صفحه جدید پیشنهاد دهید.
2. Fork کنید و شاخه بسازید.
3. تغییرات را اعمال و تست کنید: `go test ./...`
4. Pull Request ارسال کنید.

راهنمای کامل: [CONTRIBUTING.md](CONTRIBUTING.md)

---

## Roadmap

| نسخه | هدف |
|------|------|
| **v0.1** | صفحات استاتیک + CLI پایه |
| **v0.2** | رابط تعاملی TUI |
| **v0.3** | سیستم پلاگین |
| **v0.4** | به‌روزرسانی آنلاین |
| **v0.5** | دستیار هوش مصنوعی |
| **v1.0** | انتشار پایدار |

جزئیات بیشتر در [ROADMAP.md](ROADMAP.md).

---

## Logo & Branding

<p align="center">
  <img src="assets/logo/faman-logo-full.svg" alt="faman logo" width="140"/>
</p>

لوگوی **faman** الهام‌گرفته از بز کوهی ایرانی (Capra aegagrus) روی جام طلایی مارلیک و سفال‌های شهر سوخته است — خلاصه‌شده به ایده‌ی بصری اصلی: دو شاخ بزرگ، باریک‌شونده و رو به بالا، به صورت یک سیلوئت تخت مشکی.

منحنی شاخ‌ها مارپیچ لگاریتمی واقعی هستند (همان منحنی ریاضی رشد شاخ قوچ و کل)، نه قوس دستی — به همین دلیل حتی در اندازه‌های خیلی کوچک هم «شاخ» خوانده می‌شوند نه یک swoosh معمولی. تقارن دقیق است. بدون صورت، بدون چشم، بدون بافت، بدون گرادیان، بدون بج — فقط سیلوئت؛ در روح لوگوهای GitHub / Rust / Docker / Bun.

### سه نسخه رسمی

| نسخه | فایل | کاربرد |
|------|------|--------|
| **Full logo** | [`faman-logo-full.svg`](assets/logo/faman-logo-full.svg) | README، هدر مستندات، سربرگ |
| **Small icon (24×24)** | [`faman-icon-24.svg`](assets/logo/faman-icon-24.svg) | favicon، آواتار GitHub، آیکون اپ، تولبار |
| **Terminal glyph** | [`terminal-glyph.txt`](assets/logo/terminal-glyph.txt) | نمایش در ترمینال |

### فایل‌های موجود

| فایل | توضیح |
|------|--------|
| `assets/logo/faman-logo-full.svg` | لوگوی کامل، برش فشرده (غیرمربعی) |
| `assets/logo/faman-icon-24.svg` | آیکون مربعی با پدینگ، مشکی |
| `assets/logo/faman-icon-inverse.svg` | همان آیکون با fill سفید (پس‌زمینه تیره) |
| `assets/logo/faman-icon-512.png` | PNG شفاف ۵۱۲px |
| `assets/logo/app-icon-light.png` | پیش‌نمایش آیکون اپ (پس‌زمینه روشن) |
| `assets/logo/app-icon-dark.png` | پیش‌نمایش آیکون اپ (پس‌زمینه تیره) |
| `assets/logo/terminal-glyph.txt` | نسخه‌های Unicode و ASCII برای ترمینال |

### استفاده

- **پس‌زمینه روشن:** از `faman-logo-full.svg` یا `faman-icon-24.svg` استفاده کنید.
- **پس‌زمینه تیره:** از `faman-icon-inverse.svg` استفاده کنید، یا در CSS خودتان `fill="currentColor"` بگذارید.
- **حداقل اندازه:** تا ۱۶×۱۶ خوانا است؛ زیر آن ترجیحاً از terminal glyph استفاده کنید.
- **نکنید:** گرادیان، سایه، outline یا قرار دادن داخل بج دایره‌ای اضافه نکنید — سیلوئت خودش کل هویت است.

### Terminal glyph

نزدیک‌ترین تقریب به سیلوئت اصلی (دو شاخ + چانه مرکزی):

```
    ╭─╮ ╭─╮
   ╱   ╲╱   ╲
       ▾
```

نسخه ۲ خطی فشرده:

```
   ╭╮  ╭╮
    ╲▾╱
```

نسخه ASCII-only:

```
   /‾\ /‾\
  /   X   \
      V
```

جایگزین تک‌کاراکتری (برای prompt): `∧` یا `Λ`

```bash
# مثال در prompt
export PS1='∧ faman $ '
```

جزئیات کامل طراحی و راهنمای استفاده: [assets/logo/README.md](assets/logo/README.md)

---

## Screenshots

> بخش اسکرین‌شات — به‌زودی با تصاویر واقعی ترمینال تکمیل می‌شود.

```
$ faman ls
╭──────────────────────────────────────╮
│ ls                                   │
╰──────────────────────────────────────╯
ls  •  filesystem  •  beginner
...
```

---

## License

MIT License — جزئیات در فایل [LICENSE](LICENSE).

---

ساخته شده با ♥ برای جامعه لینوکس فارسی
