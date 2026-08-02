---
title: gzip
aliases:
category: archive
difficulty: beginner
keywords:
- compress
- gz
- compression
---

# Introduction

`gzip` فایل را با الگوریتم DEFLATE فشرده می‌کند و معمولاً پسوند `.gz` می‌گذارد. فایل اصلی به‌طور پیش‌فرض جایگزین می‌شود.

# Syntax

```
gzip [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-k` | نگه داشتن اصل |
| `-d` | decompress (مثل gunzip) |
| `-c` | خروجی به stdout |
| `-1` … `-9` | سطح فشرده‌سازی |
| `-t` | تست یکپارچگی |

# Examples

```bash
gzip file.txt          # → file.txt.gz
gzip -k big.log
gzip -c file.txt > file.txt.gz
gzip -d file.txt.gz
```

# Common mistakes

- فراموش کردن `-k` و از دست دادن فایل اصلی.

# Tips

- با tar: `tar -czf archive.tar.gz dir/`

# Related commands

- `gunzip` — باز کردن
- `zcat` — خواندن بدون باز کردن کامل
- `bzip2` / `xz` — فشرده‌سازی قوی‌تر
