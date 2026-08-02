---
title: umount
aliases:
category: filesystem
difficulty: intermediate
keywords:
- unmount
- disk
- detach
---

# Introduction

`umount` فایل‌سیستم mount شده را جدا می‌کند. توجه: املای دستور `umount` است نه unmount.

# Syntax

```
umount [OPTIONS] DIR|DEVICE
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-l` | lazy unmount |
| `-f` | اجباری (با احتیاط) |
| `-a` | همه |

# Examples

```bash
sudo umount /mnt/data
sudo umount /dev/sdb1
sudo umount -l /mnt/stuck
```

# Common mistakes

- «target is busy» وقتی فایل یا شل هنوز داخل mount باز است.

# Tips

- پیدا کردن فرایند استفاده‌کننده: `lsof +f -- /mnt/data` یا `fuser -m /mnt/data`

# Related commands

- `mount`
- `lsof` / `fuser`
- `findmnt`
