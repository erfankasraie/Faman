---
title: brew
aliases:
- homebrew
category: package
difficulty: beginner
keywords:
- macos
- linux
- package
- install
---

# Introduction

`brew` (Homebrew) مدیر بسته محبوب macOS و نیز **Linuxbrew** روی لینوکس است؛ برای نصب ابزارهای CLI بدون دست زدن به پکیج‌منیجر سیستم.

# Syntax

```
brew COMMAND [FORMULA]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `install FORMULA` | نصب |
| `uninstall FORMULA` | حذف |
| `search KEY` | جستجو |
| `update` | به‌روزرسانی brew |
| `upgrade` | ارتقای فرمول‌ها |
| `list` | نصب‌شده‌ها |
| `info FORMULA` | جزئیات |

# Examples

```bash
brew install git jq
brew search wget
brew update && brew upgrade
brew list
brew uninstall jq
```

# Common mistakes

- قاطی کردن `cask` (اپ‌های GUI مک) با formula معمولی: `brew install --cask firefox`.
- دو brew موازی روی مسیرهای مختلف بدون تنظیم PATH.

# Tips

- روی لینوکس پیش‌فرض اغلب `~/.linuxbrew` یا `/home/linuxbrew/...` است.
- `brew doctor` برای عیب‌یابی.

# Related commands

- `apt` / `dnf` / `pacman` — پکیج سیستم
- `nix` — رویکرد اعلانی
