---
title: yay
aliases:
category: package
difficulty: intermediate
keywords:
- arch
- aur
- pacman
- helper
---

# Introduction

`yay` یک AUR helper برای Arch Linux است: علاوه بر مخازن رسمی، بسته‌های **AUR** را هم می‌سازد و نصب می‌کند (روی `pacman`).

# Syntax

```
yay [OPTIONS] [PACKAGES]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `yay PKG` | نصب از repo یا AUR |
| `-Syu` | به‌روزرسانی سیستم + AUR |
| `-S PKG` | نصب |
| `-Rns PKG` | حذف تمیز |
| `-Ss KEY` | جستجو |
| `-Yc` | پاکسازی |

# Examples

```bash
yay -Syu
yay package-name
yay -Ss keyword
yay -Rns package-name
```

# Common mistakes

- نصب کور از AUR بدون خواندن PKGBUILD.
- جایگزین کامل دانستن به‌جای `pacman` برای کارهای حساس سیستمی.

# Tips

- اول `pacman -Syu`؛ yay برای نرم‌افزار خارج از repo رسمی.
- همیشه به امتیاز و نظر AUR نگاه کنید.

# Related commands

- `pacman`
- `makepkg`
- `paru` — helper جایگزین
