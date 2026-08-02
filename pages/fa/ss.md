---
title: ss
aliases:
category: network
difficulty: intermediate
keywords:
- socket
- port
- listen
- netstat
---

# Introduction

`ss` (socket statistics) وضعیت سوکت‌های شبکه را نشان می‌دهد. جایگزین سریع‌تر و مدرن `netstat` است.

# Syntax

```
ss [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-t` | TCP |
| `-u` | UDP |
| `-l` | فقط listening |
| `-a` | همه |
| `-n` | بدون resolve نام |
| `-p` | نمایش فرایند |
| `-r` | resolve نام |

# Examples

```bash
ss -tulpn
ss -tn
ss -ltn
ss -tp | grep nginx
```

# Common mistakes

- فراموش کردن `-p` وقتی می‌خواهید PID را ببینید (گاهی نیاز به root).

# Tips

- ترکیب محبوب عیب‌یابی پورت: `ss -tulpn`

# Related commands

- `netstat` — قدیمی‌تر
- `lsof -i` — فایل‌ها و پورت‌ها
- `ip` — تنظیمات شبکه
