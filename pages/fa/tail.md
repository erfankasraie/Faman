---
title: tail
aliases:
category: text
difficulty: beginner
keywords:
- end
- follow
- log
---

# Introduction

`tail` انتهای فایل را نشان می‌دهد؛ با `-f` برای دنبال کردن لاگ زنده استفاده می‌شود.

# Syntax

```
tail [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-n N` | N خط آخر |
| `-f` | follow |
| `-F` | follow + تلاش مجدد اگر فایل عوض شد |
| `-c N` | N بایت آخر |

# Examples

```bash
tail file.txt
tail -n 50 /var/log/syslog
tail -f /var/log/nginx/access.log
tail -n +10 file.txt    # از خط ۱۰ تا آخر
```

# Common mistakes

- `tail -f` روی فایلی که rotate می‌شود بدون `-F`.

# Tips

- چند فایل: `tail -f a.log b.log`

# Related commands

- `head`
- `less` + `F`
- `journalctl -f`
