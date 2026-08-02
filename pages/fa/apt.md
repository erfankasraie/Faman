---
title: apt
aliases:
- apt-get
category: package
difficulty: beginner
keywords:
- package
- install
- debian
- ubuntu
---

# Introduction

`apt` ابزار مدیریت بسته در توزیع‌های مبتنی بر Debian/Ubuntu است. نصب، به‌روزرسانی و حذف نرم‌افزار را ساده می‌کند.

# Syntax

```
apt [OPTIONS] COMMAND
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `update` | به‌روزرسانی فهرست بسته‌ها |
| `upgrade` | ارتقای بسته‌های نصب‌شده |
| `install PKG` | نصب بسته |
| `remove PKG` | حذف بسته |
| `purge PKG` | حذف همراه فایل‌های تنظیمات |
| `search KEYWORD` | جستجو |
| `show PKG` | اطلاعات بسته |
| `autoremove` | حذف وابستگی‌های بلااستفاده |

# Examples

```bash
# به‌روزرسانی فهرست و سیستم
sudo apt update
sudo apt upgrade

# نصب
sudo apt install curl git

# جستجو
apt search nginx

# حذف
sudo apt remove package-name
sudo apt autoremove
```

# Common mistakes

- اجرای `upgrade` بدون `update` اول.
- استفاده از `apt` بدون `sudo` برای عملیات سیستمی.

# Tips

- `apt` رابط کاربرپسندتر از `apt-get` است؛ برای اسکریپت‌ها گاهی `apt-get` پایدارتر در نظر گرفته می‌شود.
- قبل از نصب بزرگ: `apt show package`.

# Related commands

- `dpkg` — مدیریت سطح پایین‌تر deb
- `snap` / `flatpak` — بسته‌های جایگزین
- `aptitude` — رابط پیشرفته‌تر
