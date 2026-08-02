---
title: snap
aliases:
category: package
difficulty: beginner
keywords:
- ubuntu
- package
- sandbox
- install
---

# Introduction

`snap` سامانهٔ بسته‌بندی Canonical است: برنامه‌ها به‌صورت اسنپ، ایزوله و با به‌روزرسانی خودکار نصب می‌شوند. روی Ubuntu پیش‌فرض رایج است.

# Syntax

```
snap <command> [args]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `find KEY` | جستجو |
| `install NAME` | نصب |
| `remove NAME` | حذف |
| `list` | اسنپ‌های نصب‌شده |
| `info NAME` | جزئیات |
| `refresh` | به‌روزرسانی |
| `services` | سرویس‌های اسنپ |

# Examples

```bash
snap find code
sudo snap install code --classic
snap list
sudo snap remove code
sudo snap refresh
```

# Common mistakes

- فراموش کردن `--classic` برای ابزارهایی که به دسترسی کامل نیاز دارند (مثل IDE).
- انتظار مسیرهای FHS سنتی؛ اسنپ‌ها زیر `/snap` هستند.

# Tips

- کانال‌ها: `--channel=candidate` یا `edge`.
- اگر `snap` کند بود، برای برخی نرم‌افزارها `flatpak` یا deb جایگزین است.

# Related commands

- `flatpak`
- `apt`
- `dpkg`
