---
title: tail
aliases:
category: text
difficulty: beginner
keywords:
- end
- file
- lines
- follow
---

# Introduction

دستور `tail` انتهای یک فایل را نمایش می‌دهد (پیش‌فرض ۱۰ خط آخر). حالت follow آن برای دنبال کردن لاگ‌ها بسیار محبوب است.

# Syntax

```
tail [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-n N` | N خط آخر |
| `-c N` | N بایت آخر |
| `-f` | follow — منتظر داده‌های جدید بمان |
| `-F` | مثل `-f` اما فایل را بعد از rotate دوباره باز می‌کند |
| `--pid=PID` | همراه با `-f`، وقتی پروسس تمام شد خارج شو |

# Examples

```bash
# ۱۰ خط آخر
tail file.txt

# ۵۰ خط آخر
tail -n 50 app.log

# دنبال کردن لاگ به صورت زنده
tail -f /var/log/syslog

# دنبال کردن با reopen بعد از rotate
tail -F /var/log/nginx/access.log
```

# Common mistakes

- استفاده از `-f` روی فایلی که rotate می‌شود بدون `-F`.
- فراموش کردن خروج از حالت follow با Ctrl+C.

# Tips

- `tail -f` یکی از پرکاربردترین دستورات برای مانیتورینگ است.
- می‌توانید چند فایل را همزمان follow کنید.
- ترکیب با `grep`: `tail -f app.log | grep ERROR`

# Related commands

- `head` — ابتدای فایل
- `less +F` — follow با قابلیت اسکرول
- `journalctl -f` — برای systemd logs
- `multitail` — follow چند فایل با UI بهتر
