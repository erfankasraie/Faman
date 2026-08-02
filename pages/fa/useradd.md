---
title: useradd
aliases:
category: user
difficulty: intermediate
keywords:
- account
- create
- user
---

# Introduction

`useradd` حساب کاربری جدید در سیستم می‌سازد. در بعضی توزیع‌ها `adduser` رابط دوستانه‌تری است.

# Syntax

```
useradd [OPTIONS] USERNAME
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-m` | ساخت home |
| `-s SHELL` | شل پیش‌فرض |
| `-G GROUPS` | گروه‌های فرعی |
| `-c COMMENT` | توضیح (GECOS) |
| `-d DIR` | مسیر home |
| `-r` | حساب سیستمی |

# Examples

```bash
sudo useradd -m -s /bin/bash alice
sudo passwd alice
sudo useradd -m -G sudo,docker bob
```

# Common mistakes

- فراموش کردن `-m` و نبودن home.
- ندادن رمز با `passwd` بعد از ساخت.

# Tips

- حذف: `userdel -r username`

# Related commands

- `usermod` — ویرایش حساب
- `userdel` — حذف
- `adduser` — رابط سطح بالاتر (Debian)
