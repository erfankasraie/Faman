---
title: blkid
aliases:
category: filesystem
difficulty: intermediate
keywords:
- uuid
- disk
- partition
- label
---

# Introduction

`blkid` ویژگی‌های بلاک‌دیوایس (UUID، TYPE، LABEL) را نشان می‌دهد؛ مفید برای `/etc/fstab`.

# Syntax

```
blkid [OPTIONS] [DEVICE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-o mode` | خروجی: `full`, `value`, `export` |
| `-s TAG` | فقط یک تگ (UUID, TYPE, …) |
| `-U UUID` | پیدا کردن دیوایس با UUID |
| `-L LABEL` | پیدا کردن با label |

# Examples

```bash
sudo blkid
sudo blkid /dev/sda1
blkid -s UUID -o value /dev/sda1
```

# Common mistakes

- بدون sudo بعضی دیوایس‌ها دیده نمی‌شوند.

# Tips

- در fstab ترجیحاً `UUID=` به‌جای `/dev/sdX`.

# Related commands

- `lsblk` · `findmnt` · `mount` · `fdisk`
