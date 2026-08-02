---
title: lsblk
aliases:
category: filesystem
difficulty: beginner
keywords:
- block
- disk
- partition
---

# Introduction

`lsblk` دیوایس‌های بلوکی (دیسک، پارتیشن، LVM، ...) را به صورت درختی نشان می‌دهد.

# Syntax

```
lsblk [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-f` | فایل‌سیستم و UUID |
| `-a` | همه (شامل خالی) |
| `-o COLS` | ستون‌های سفارشی |
| `-e 7` | حذف loop (مثال) |

# Examples

```bash
lsblk
lsblk -f
lsblk -o NAME,SIZE,TYPE,MOUNTPOINT,FSTYPE
```

# Common mistakes

- گیج شدن بین نام دیوایس (`/dev/sda1`) و mountpoint.

# Tips

- قبل از `mount` یا پارتیشن‌بندی حتماً `lsblk` بزنید.

# Related commands

- `blkid` — UUID و برچسب
- `fdisk -l` — جدول پارتیشن
- `df` — فضای mount شده
