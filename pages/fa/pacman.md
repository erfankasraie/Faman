---
title: pacman
aliases:
category: package
difficulty: intermediate
keywords:
- package
- arch
- install
---

# Introduction

`pacman` مدیر بسته Arch Linux و مشتقات (Manjaro و ...) است.

# Syntax

```
pacman [OPTIONS] [PACKAGES]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-S PKG` | نصب/همگام‌سازی |
| `-Syu` | به‌روزرسانی سیستم |
| `-R PKG` | حذف |
| `-Rs PKG` | حذف با وابستگی‌های بلااستفاده |
| `-Ss KEY` | جستجو در ریپو |
| `-Qs KEY` | جستجو در نصب‌شده‌ها |
| `-Si PKG` | اطلاعات |

# Examples

```bash
sudo pacman -Syu
sudo pacman -S git curl
pacman -Ss nginx
sudo pacman -Rs oldpackage
```

# Common mistakes

- اجرای `-Sy` بدون `-u` (به‌روزرسانی ناقص).
- ویرایش جزئی سیستم بدون مطالعه Arch Wiki.

# Tips

- AUR با helperهایی مثل `yay` یا `paru`.

# Related commands

- `yay` / `paru` — AUR
- `apt` / `dnf` — سایر توزیع‌ها
