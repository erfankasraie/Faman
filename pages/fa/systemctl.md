---
title: systemctl
aliases:
- service
category: system
difficulty: intermediate
keywords:
- systemd
- service
- daemon
- unit
---

# Introduction

`systemctl` ابزار اصلی کنترل **systemd** است: شروع، توقف، فعال‌سازی و وضعیت سرویس‌ها (unitها).

# Syntax

```
systemctl [OPTIONS] COMMAND [UNIT...]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `status UNIT` | وضعیت سرویس |
| `start UNIT` | شروع |
| `stop UNIT` | توقف |
| `restart UNIT` | راه‌اندازی مجدد |
| `reload UNIT` | بارگذاری مجدد تنظیمات |
| `enable UNIT` | شروع خودکار در بوت |
| `disable UNIT` | غیرفعال کردن شروع خودکار |
| `list-units` | لیست unitهای فعال |
| `daemon-reload` | بارگذاری مجدد unit فایل‌ها |

# Examples

```bash
# وضعیت nginx
systemctl status nginx

# شروع و فعال‌سازی
sudo systemctl start nginx
sudo systemctl enable nginx

# راه‌اندازی مجدد
sudo systemctl restart ssh

# لیست سرویس‌های در حال اجرا
systemctl list-units --type=service --state=running
```

# Common mistakes

- فراموش کردن `daemon-reload` بعد از ویرایش فایل unit.
- اشتباه گرفتن `enable` (بوت) با `start` (همین الان).

# Tips

- لاگ سرویس: `journalctl -u nginx -f`
- نام unit معمولاً با `.service` تمام می‌شود ولی اغلب می‌توان آن را حذف کرد.

# Related commands

- `journalctl` — لاگ‌های systemd
- `service` — رابط قدیمی‌تر
- `systemctl cat UNIT` — محتوای unit فایل
