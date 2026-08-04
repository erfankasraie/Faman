---
title: zstd
aliases:
category: archive
difficulty: intermediate
keywords:
- compress
- zstandard
- tar
---

# Introduction

`zstd` (Zstandard) فشرده‌سازی سریع با نسبت خوب است؛ انتخاب رایج برای بکاپ و توزیع مدرن.

# Syntax

```
zstd [OPTIONS] [FILE...]
unzstd FILE.zst
zstdcat FILE.zst
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-d` | decompress |
| `-k` | نگه داشتن ورودی |
| `-c` | stdout |
| `-1` … `-19` | سطح (۳ پیش‌فرض؛ ۱۹ خیلی کند/قوی) |
| `-T0` | همه هسته‌ها |
| `-f` | overwrite |
| `-l` | لیست فریم |

# Examples

```bash
zstd file.bin
zstd -T0 -10 dataset.csv
zstd -d file.bin.zst
tar --zstd -cf a.tar.zst dir/
tar --zstd -xf a.tar.zst
```

# Common mistakes

- سطح خیلی بالا روی CPU ضعیف برای فایل‌های عظیم.
- قاطی کردن پسوند `.zst` با `.gz`.

# Tips

- `zstdmt` گاهی به‌عنوان alias چندنخی نصب می‌شود.
- برای pipe: `zstd -c file | ssh host 'zstd -d > file'`.

# Related commands

- `xz` · `gzip` · `tar` · `pigz`
