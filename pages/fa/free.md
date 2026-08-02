---
title: free
aliases:
- memory
category: system
difficulty: beginner
keywords:
- ram
- memory
- swap
- cache
---

# Introduction

`free` وضعیت حافظه RAM و swap را نشان می‌دهد: کل، استفاده‌شده، آزاد و بافر/کش.

# Syntax

```
free [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-h` | واحد خوانا |
| `-m` | مگابایت |
| `-g` | گیگابایت |
| `-s N` | تکرار هر N ثانیه |
| `-t` | نمایش جمع کل |

# Examples

```bash
# خلاصه خوانا
free -h

# به‌روزرسانی هر ۲ ثانیه
free -h -s 2

# فقط مگابایت
free -m
```

# Common mistakes

- فکر کردن که ستون `used` یعنی حافظه واقعاً تمام شده — لینوکس کش را در `available` حساب می‌کند.
- نادیده گرفتن swap وقتی RAM پر است.

# Tips

- ستون مهم در نسخه‌های جدید: `available`.
- برای جزئیات بیشتر: `cat /proc/meminfo`

# Related commands

- `top` — مصرف حافظه به‌ازای فرایند
- `vmstat` — آمار حافظه و IO
- `swapon -s` — وضعیت swap
