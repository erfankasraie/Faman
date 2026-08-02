---
title: df
aliases:
- diskfree
category: filesystem
difficulty: beginner
keywords:
- disk
- space
- partition
- mount
---

# Introduction

`df` (disk free) فضای آزاد و استفاده‌شده پارتیشن‌ها و سیستم‌فایل‌های mount شده را نشان می‌دهد.

# Syntax

```
df [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-h` | اندازه خوانا (G, M, K) |
| `-T` | نوع سیستم‌فایل |
| `-i` | نمایش inode به‌جای بلوک |
| `-a` | شامل سیستم‌فایل‌های مجازی |
| `--total` | جمع کل |

# Examples

```bash
# خلاصه خوانا
df -h

# با نوع فایل‌سیستم
df -hT

# فضای مربوط به یک مسیر
df -h /home

# بررسی inode (وقتی «disk full» ولی فضا هست)
df -i
```

# Common mistakes

- نگاه نکردن به inodeها وقتی تعداد فایل‌های کوچک زیاد است.
- گیج شدن با سیستم‌فایل‌های tmpfs و overlay در داکر.

# Tips

- ترکیب محبوب: `df -hT`
- برای فضای یک پوشه خاص از `du` استفاده کنید، نه `df`.

# Related commands

- `du` — فضای اشغال‌شده توسط پوشه‌ها
- `lsblk` — لیست بلاک‌دیوایس‌ها
- `mount` — سیستم‌فایل‌های mount شده
