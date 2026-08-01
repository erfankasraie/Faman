---
title: ssh
aliases:
category: network
difficulty: intermediate
keywords:
- remote
- shell
- secure
- login
---

# Introduction

دستور `ssh` (Secure Shell) برای اتصال امن به سرورهای راه دور و اجرای دستورات روی آن‌ها استفاده می‌شود.

# Syntax

```
ssh [OPTIONS] [USER@]HOST [COMMAND]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-p PORT` | پورت (پیش‌فرض ۲۲) |
| `-i KEY` | فایل کلید خصوصی |
| `-L` | Local port forwarding |
| `-R` | Remote port forwarding |
| `-D` | Dynamic (SOCKS) |
| `-X` / `-Y` | X11 forwarding |
| `-v` | verbose (اشکال‌زدایی) |
| `-N` | فقط forward، شل باز نکن |
| `-f` | به پس‌زمینه برو |

# Examples

```bash
# اتصال ساده
ssh user@192.168.1.100

# با پورت خاص
ssh -p 2222 user@example.com

# با کلید خاص
ssh -i ~/.ssh/mykey user@example.com

# اجرای یک دستور و خروج
ssh user@host "uptime"

# Port forwarding محلی
ssh -L 8080:localhost:80 user@host

# کپی کلید عمومی (برای ورود بدون پسورد)
ssh-copy-id user@host
```

# Common mistakes

- فراموش کردن تنظیم مجوز کلید خصوصی (`chmod 600 ~/.ssh/id_rsa`).
- استفاده از پسورد به جای کلید در محیط‌های production.
- باز گذاشتن پورت ۲۲ روی اینترنت بدون محدودیت.

# Tips

- از `~/.ssh/config` برای تعریف hostهای پرکاربرد استفاده کنید.
- همیشه کلیدهای ed25519 را ترجیح دهید.
- برای اتصالات مکرر از `ControlMaster` و multiplexing استفاده کنید.
- `ssh-agent` را برای مدیریت کلیدها راه‌اندازی کنید.

# Related commands

- `scp` — کپی فایل روی SSH
- `sftp` — انتقال فایل تعاملی
- `ssh-keygen` — ساخت کلید
- `ssh-copy-id` — نصب کلید عمومی
- `mosh` — جایگزین مقاوم به قطعی شبکه
