---
title: find
aliases:
category: filesystem
difficulty: intermediate
keywords:
- search
- files
- recursive
- locate
---

# Introduction

`find` فایل‌ها و پوشه‌ها را بر اساس نام، نوع، زمان، اندازه و ... در درخت دایرکتوری جستجو می‌کند.

# Syntax

```
find [PATH...] [EXPRESSION]
```

# Options

| عبارت | توضیح |
|-------|--------|
| `-name PATTERN` | نام (حساس به حروف) |
| `-iname PATTERN` | نام بدون حساسیت |
| `-type f/d/l` | فایل / پوشه / symlink |
| `-size +N` | بزرگ‌تر از N (c/k/M/G) |
| `-mtime -N` | تغییر در N روز اخیر |
| `-maxdepth N` | محدودیت عمق |
| `-exec CMD {} \;` | اجرا روی هر نتیجه |
| `-exec CMD {} +` | اجرا دسته‌ای |
| `-delete` | حذف نتایج |
| `-print0` | جداکننده null برای xargs |

# Examples

```bash
# همه فایل‌های .log
find /var/log -type f -name '*.log'

# فایل‌های بزرگ‌تر از ۱۰۰ مگ
find /home -type f -size +100M

# تغییر یافته در ۲۴ ساعت
find . -mtime -1

# حذف فایل‌های tmp امن
find /tmp -type f -name '*.tmp' -print0 | xargs -0 rm -f

# جستجو تا عمق ۲
find . -maxdepth 2 -type d
```

# Common mistakes

- `-name *.log` بدون کوتیشن (شل زودتر expand می‌کند).
- `-delete` یا `-exec rm` بدون dry-run ذهنی.
- جستجوی `/` بدون `-xdev` یا محدودیت.

# Tips

- با فاصله در نام: همیشه `-print0` + `xargs -0`.
- جایگزین سریع‌تر برای نام: `fd` یا `locate` (دیتابیس).
- `-exec ... {} +` از `;` کارآمدتر است.

# Related commands

- `locate` / `updatedb` — جستجوی ایندکسی
- `fd` — سینتکس مدرن
- `grep -r` — جستجوی داخل محتوا
- `xargs` — آرگومان‌سازی
