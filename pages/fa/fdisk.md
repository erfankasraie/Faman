---
title: fdisk
aliases:
category: filesystem
difficulty: advanced
keywords:
- partition
- disk
- gpt
- mbr
---

# Introduction

`fdisk` ابزار کلاسیک پارتیشن‌بندی دیسک (MBR/GPT بسته به نسخه) است. **مخرب** است؛ اشتباه = از دست رفتن داده.

# Syntax

```
sudo fdisk [OPTIONS] DEVICE
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-l` | لیست پارتیشن‌ها |
| `-l DEVICE` | فقط یک دیسک |

در حالت تعاملی: `m` راهنما · `p` چاپ · `n` جدید · `d` حذف · `w` نوشتن · `q` خروج بدون ذخیره.

# Examples

```bash
sudo fdisk -l
sudo fdisk /dev/sdb
```

# Common mistakes

- `w` روی دیسک اشتباه.
- قاطی `sdX` با پارتیشن `sdX1`.

# Tips

- برای GPT گاهی `parted` یا `gdisk` راحت‌تر است.
- قبل از کار: بکاپ و `lsblk`.

# Related commands

- `parted` · `lsblk` · `blkid` · `mkfs`
