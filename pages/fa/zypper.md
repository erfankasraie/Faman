---
title: zypper
aliases:
category: package
difficulty: beginner
keywords:
- opensuse
- package
- install
- suse
---

# Introduction

`zypper` مدیر بسته openSUSE و SLE است (روی libzypp).

# Syntax

```
zypper [OPTIONS] COMMAND
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `refresh` / `ref` | به‌روزرسانی مخازن |
| `install` / `in` | نصب |
| `remove` / `rm` | حذف |
| `update` / `up` | به‌روزرسانی بسته‌ها |
| `dist-upgrade` / `dup` | ارتقای توزیع |
| `search` / `se` | جستجو |
| `info` | اطلاعات بسته |

# Examples

```bash
sudo zypper refresh
sudo zypper install git curl
zypper search nginx
sudo zypper update
sudo zypper remove package-name
```

# Common mistakes

- روی openSUSE Leap برای ارتقای نسخه از `dup` به‌جای `up` استفاده کنید وقتی راهنما می‌گوید.
- افزودن مخزن شخص ثالث بدون بررسی GPG.

# Tips

- خلاصه‌ها: `zypper in`، `zypper se`.
- `--non-interactive` برای اسکریپت.

# Related commands

- `rpm` — لایهٔ سطح‌پایین
- `dnf` — فدورا/RHEL
- `apt` — دبیان
