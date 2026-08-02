---
title: nc
aliases:
- netcat
- ncat
category: network
difficulty: intermediate
keywords:
- socket
- debug
- port
---

# Introduction

`nc` (netcat) ابزار همه‌کاره TCP/UDP است؛ تست پورت، انتقال ساده داده و دیباگ شبکه.

# Syntax

```
nc [OPTIONS] HOST PORT
nc -l PORT
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-l` | listen |
| `-u` | UDP |
| `-v` | verbose |
| `-z` | اسکن بدون ارسال داده |
| `-w SEC` | timeout |

# Examples

```bash
nc -zv example.com 80
nc -l 1234
nc host 1234 < file.bin
```

# Common mistakes

- باز گذاشتن listen روی پورت حساس در شبکه عمومی.

# Tips

- نسخه‌های openbsd/gnu/ncat کمی فرق دارند.

# Related commands

- `ss` / `telnet`
- `curl`
- `nmap`
