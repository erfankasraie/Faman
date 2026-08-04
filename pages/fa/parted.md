---
title: parted
aliases:
category: filesystem
difficulty: advanced
keywords:
- partition
- gpt
- disk
---

# Introduction

`parted` پارتیشن‌بندی با پشتیبانی خوب از **GPT** و اسکریپت‌پذیری است. مانند fdisk خطرناک است.

# Syntax

```
sudo parted [OPTIONS] [DEVICE [COMMAND...]]
```

# Options

| گزینه / فرمان | توضیح |
|----------------|--------|
| `-l` | لیست |
| `print` | جدول پارتیشن |
| `mkpart` | پارتیشن جدید |
| `rm N` | حذف |
| `resizepart` | تغییر اندازه |
| `unit GiB` | واحد |

# Examples

```bash
sudo parted -l
sudo parted /dev/sdb print
sudo parted /dev/sdb --script mklabel gpt mkpart primary 1MiB 50%
```

# Common mistakes

- `--script` بدون بررسی دوباره.
- واحد MiB/GiB را اشتباه گرفتن.

# Tips

- بعد از تغییر: `partprobe` یا reboot.

# Related commands

- `fdisk` · `gdisk` · `lsblk` · `mkfs`
