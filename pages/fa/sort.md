---
title: sort
aliases:
category: text
difficulty: beginner
keywords:
- order
- sort
- lines
---

# Introduction

`sort` خطوط یک فایل یا ورودی را مرتب می‌کند. پایهٔ بسیاری از pipelineهای متنی است.

# Syntax

```
sort [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-r` | معکوس |
| `-n` | مرتب‌سازی عددی |
| `-h` | اندازهٔ human-readable (مثل 2K, 1G) |
| `-k N` | مرتب بر اساس فیلد N |
| `-t C` | جداکننده فیلد |
| `-u` | خطوط یکتا (مثل uniq) |
| `-f` | بدون حساسیت به حروف |

# Examples

```bash
sort names.txt
sort -n numbers.txt
sort -r -n scores.txt
sort -t: -k3 -n /etc/passwd
du -h | sort -h
```

# Common mistakes

- استفاده از sort معمولی برای اعداد چندرقمی بدون `-n` (مثلاً 10 قبل از 2 می‌آید).

# Tips

- ترکیب کلاسیک: `sort file | uniq`
- برای فایل‌های خیلی بزرگ، sort ممکن است از دیسک موقت استفاده کند.

# Related commands

- `uniq` — حذف تکرارهای پشت‌سرهم
- `awk` — مرتب‌سازی منطقی‌تر
