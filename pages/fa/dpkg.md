---
title: dpkg
aliases:
category: package
difficulty: intermediate
keywords:
- deb
- debian
- package
- low-level
---

# Introduction

`dpkg` مدیر سطح‌پایین بسته‌های `.deb` در Debian/Ubuntu است. معمولاً مستقیم از `apt` استفاده می‌کنید؛ `dpkg` برای نصب فایل محلی و عیب‌یابی است.

# Syntax

```
dpkg [OPTIONS] ACTION
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-i FILE.deb` | نصب بسته محلی |
| `-r PKG` | حذف (فایل‌های تنظیمات می‌ماند) |
| `-P PKG` | purge کامل |
| `-l` | لیست بسته‌های نصب‌شده |
| `-L PKG` | فایل‌های متعلق به بسته |
| `-S FILE` | کدام بسته مالک فایل است |
| `--configure -a` | ادامه پیکربندی ناقص |

# Examples

```bash
# نصب deb محلی
sudo dpkg -i package_1.0_amd64.deb

# اگر وابستگی کم بود:
sudo apt-get install -f

# لیست و جستجو
dpkg -l | grep nginx
dpkg -L curl
dpkg -S /usr/bin/curl
```

# Common mistakes

- نصب با `dpkg -i` بدون رفع وابستگی (`apt -f install`).
- قاطی کردن نام بسته با نام فایل `.deb`.

# Tips

- برای ریپو و وابستگی‌ها همیشه `apt` را ترجیح دهید.
- وضعیت خراب: `sudo dpkg --configure -a`.

# Related commands

- `apt` — لایهٔ سطح‌بالا
- `apt-cache` — جستجوی متادیتا
- `snap` / `flatpak`
