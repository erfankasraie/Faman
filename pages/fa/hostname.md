---
title: hostname
aliases:
category: system
difficulty: beginner
keywords:
- host
- name
- network
---

# Introduction

`hostname` نام میزبان سیستم را نمایش یا (با دسترسی کافی) تنظیم می‌کند.

# Syntax

```
hostname [OPTIONS] [NAME]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-f` | FQDN |
| `-i` | آدرس‌های IP |
| `-s` | نام کوتاه |

# Examples

```bash
hostname
hostname -f
hostnamectl  # در سیستم‌های systemd ترجیح داده می‌شود
```

# Common mistakes

- تغییر موقت با `hostname newname` بدون پایدارسازی از طریق `hostnamectl` یا فایل تنظیمات.

# Tips

- در systemd: `hostnamectl set-hostname mybox`

# Related commands

- `hostnamectl` — مدیریت مدرن‌تر
- `uname` — اطلاعات کرنل
