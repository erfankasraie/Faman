---
title: mount
aliases:
category: filesystem
difficulty: intermediate
keywords:
- disk
- partition
- filesystem
---

# Introduction

`mount` یک فایل‌سیستم را به درخت دایرکتوری متصل (mount) می‌کند.

# Syntax

```
mount [OPTIONS] DEVICE DIR
mount                     # لیست mountها
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-t TYPE` | نوع فایل‌سیستم |
| `-o OPTS` | گزینه‌ها (ro, rw, noexec, ...) |
| `-a` | همه از fstab |

# Examples

```bash
mount
sudo mount /dev/sdb1 /mnt/data
sudo mount -o ro /dev/cdrom /mnt/cd
sudo mount -a
```

# Common mistakes

- mount بدون ثبت در `/etc/fstab` بعد از ریبوت از بین می‌رود.
- mount روی پوشه غیرخالی (محتوا موقتاً پنهان می‌شود).

# Tips

- برای پایدارسازی از fstab یا systemd mount unit استفاده کنید.

# Related commands

- `umount` — جدا کردن
- `lsblk` — دیدن دیوایس‌ها
- `findmnt` — نمای درختی mountها
