---
title: du
aliases:
category: filesystem
difficulty: beginner
keywords:
- disk
- usage
- size
- directory
---

# Introduction

`du` (disk usage) فضای اشغال‌شده توسط فایل‌ها و پوشه‌ها را نشان می‌دهد. برخلاف `df` که کل پارتیشن را می‌گوید، `du` روی مسیر مشخص کار می‌کند.

# Syntax

```
du [OPTIONS] [PATH...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-h` | خوانا (K/M/G) |
| `-s` | فقط جمع کل |
| `-a` | همه فایل‌ها |
| `-d N` | عمق حداکثر N |
| `--max-depth=N` | همان |
| `-x` | فقط یک فایل‌سیستم |

# Examples

```bash
du -h
du -sh *
du -h --max-depth=1 /var
du -sh /home/* | sort -h
```

# Common mistakes

- مقایسهٔ مستقیم با `df` بدون در نظر گرفتن رزرو root و فایل‌های حذف‌شده باز.

# Tips

- پیدا کردن پوشه‌های بزرگ: `du -h --max-depth=1 | sort -h`

# Related commands

- `df` — فضای آزاد پارتیشن
- `ncdu` — رابط تعاملی (اگر نصب باشد)
