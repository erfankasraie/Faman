---
title: date
aliases:
category: system
difficulty: beginner
keywords:
- time
- clock
- format
---

# Introduction

`date` تاریخ و ساعت سیستم را نمایش می‌دهد و می‌تواند قالب‌بندی یا (با دسترسی root) تنظیم کند.

# Syntax

```
date [OPTIONS] [+FORMAT]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-u` | UTC |
| `-d STRING` | تفسیر تاریخ دلخواه |
| `-s STRING` | تنظیم ساعت (نیاز به دسترسی) |

# Examples

```bash
date
date -u
date +%Y-%m-%d
date +'%Y-%m-%d %H:%M:%S'
date -d 'next Monday'
```

# Common mistakes

- فراموش کردن `+` قبل از format string.

# Tips

- برای نام فایل لاگ: `date +%Y%m%d_%H%M%S`

# Related commands

- `timedatectl` — مدیریت timezone در systemd
- `touch -t` — زمان فایل
