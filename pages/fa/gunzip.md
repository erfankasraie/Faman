---
title: gunzip
aliases:
category: archive
difficulty: beginner
keywords:
- decompress
- gz
- extract
---

# Introduction

`gunzip` فایل‌های `.gz` را باز می‌کند. معادل `gzip -d` است.

# Syntax

```
gunzip [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-k` | نگه داشتن .gz |
| `-c` | خروجی به stdout |
| `-t` | تست |

# Examples

```bash
gunzip file.txt.gz
gunzip -c file.txt.gz | less
zcat file.txt.gz
```

# Common mistakes

- تلاش برای gunzip روی `.tar.gz` بدون tar — برای آرشیو از `tar -xzf` استفاده کنید.

# Tips

- مشاهده بدون استخراج: `zless` / `zcat`

# Related commands

- `gzip`
- `tar`
- `zcat` / `zless`
