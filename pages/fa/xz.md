---
title: xz
aliases:
category: archive
difficulty: intermediate
keywords:
- compress
- lzma
- tar
---

# Introduction

`xz` فشرده‌سازی با الگوریتم **LZMA2** است؛ نسبت فشرده‌سازی معمولاً بهتر از gzip و نزدیک bzip2/zstd در تنظیمات بالا، با سرعت کمتر از zstd در حالت پیش‌فرض.

# Syntax

```
xz [OPTIONS] [FILE...]
unxz FILE.xz
xzcat FILE.xz
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-d` / `--decompress` | استخراج |
| `-k` | نگه داشتن فایل اصلی |
| `-c` | خروجی stdout |
| `-0` … `-9` | سطح فشرده‌سازی (پیش‌فرض ۶) |
| `-T N` | تعداد نخ |
| `-v` | verbose |
| `-l` | اطلاعات آرشیو |

# Examples

```bash
xz file.txt                 # → file.txt.xz
xz -k -9 big.log
xz -d file.txt.xz
xzcat data.xz | less
tar -cJf archive.tar.xz dir/
tar -xJf archive.tar.xz
```

# Common mistakes

- فراموش کردن `-k` و پاک شدن فایل ورودی.
- انتظار سرعت gzip؛ xz سنگین‌تر است.

# Tips

- برای سرعت بهتر اغلب `zstd` ترجیح داده می‌شود.
- `tar -J` معادل xz است.

# Related commands

- `zstd` · `gzip` · `bzip2` · `tar`
