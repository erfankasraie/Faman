---
title: passwd
aliases:
category: user
difficulty: beginner
keywords:
- password
- security
- account
---

# Introduction

`passwd` رمز عبور حساب کاربری را تغییر می‌دهد.

# Syntax

```
passwd [OPTIONS] [USER]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-S` | وضعیت رمز |
| `-l` | قفل حساب (root) |
| `-u` | باز کردن قفل |
| `-d` | حذف رمز (خطرناک) |
| `-e` | منقضی کردن رمز |

# Examples

```bash
passwd
sudo passwd alice
passwd -S
```

# Common mistakes

- انتخاب رمز ضعیف.
- تغییر رمز کاربر دیگر بدون `sudo`.

# Tips

- سیاست رمز در `/etc/login.defs` و PAM تعریف می‌شود.

# Related commands

- `chpasswd` — تغییر دسته‌ای
- `usermod` — تنظیمات حساب
