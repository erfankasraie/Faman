---
title: flatpak
aliases:
category: package
difficulty: beginner
keywords:
- package
- flathub
- sandbox
- desktop
---

# Introduction

`flatpak` بسته‌بندی دسکتاپ لینوکس با sandbox است. معمولاً با مخزن **Flathub** استفاده می‌شود و روی بسیاری از توزیع‌ها کار می‌کند.

# Syntax

```
flatpak [OPTIONS] COMMAND
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `remote-add` | افزودن مخزن |
| `search KEY` | جستجو |
| `install REF` | نصب |
| `uninstall REF` | حذف |
| `list` | لیست |
| `update` | به‌روزرسانی |
| `run REF` | اجرا |

# Examples

```bash
# افزودن Flathub (یک‌بار)
sudo flatpak remote-add --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo

flatpak search vlc
flatpak install flathub org.videolan.VLC
flatpak run org.videolan.VLC
flatpak update
flatpak uninstall org.videolan.VLC
```

# Common mistakes

- نصب بدون اضافه کردن remote.
- اشتباه گرفتن Application ID طولانی با نام کوتاه.

# Tips

- `--user` برای نصب فقط کاربر بدون sudo.
- ترکیب با `apt`/`dnf` برای ابزارهای سیستمی؛ flatpak بیشتر برای اپ‌های GUI.

# Related commands

- `snap`
- `apt` / `dnf` / `pacman`
