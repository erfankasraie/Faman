---
title: sudo
aliases:
category: system
difficulty: beginner
keywords:
- root
- privilege
- admin
---

# Introduction

`sudo` اجازه می‌دهد یک دستور را با دسترسی کاربر دیگر (معمولاً root) اجرا کنید؛ بدون ورود کامل به حساب root.

# Syntax

```
sudo [OPTIONS] COMMAND
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-u USER` | اجرا به‌عنوان کاربر دیگر |
| `-i` | شل login شبیه root |
| `-s` | شل با دسترسی بالا |
| `-k` | باطل کردن اعتبار رمز |
| `-l` | لیست مجوزهای شما |
| `-E` | حفظ محیط |

# Examples

```bash
sudo apt update
sudo -u www-data whoami
sudo -i
sudo tail /var/log/syslog
echo value | sudo tee /etc/config
```

# Common mistakes

- `sudo cd /root` بی‌فایده است چون `cd` builtin شل است.
- `sudo echo > /etc/file` — redirection قبل از sudo انجام می‌شود؛ از `tee` استفاده کنید.

# Tips

- تنظیمات در `/etc/sudoers` — فقط با `visudo` ویرایش کنید.

# Related commands

- `su` — تعویض کاربر
- `whoami` / `id` — هویت فعلی
