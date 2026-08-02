---
title: ping
aliases:
category: network
difficulty: beginner
keywords:
- network
- icmp
- latency
- connectivity
---

# Introduction

`ping` با بسته‌های ICMP بررسی می‌کند که یک میزبان در شبکه پاسخ می‌دهد یا نه و تأخیر (latency) را اندازه می‌گیرد.

# Syntax

```
ping [OPTIONS] HOST
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-c N` | فقط N بسته بفرست |
| `-i SEC` | فاصله بین بسته‌ها |
| `-W SEC` | مهلت انتظار پاسخ |
| `-s SIZE` | اندازه بسته |
| `-4` / `-6` | اجبار IPv4 یا IPv6 |

# Examples

```bash
# پینگ معمولی
ping google.com

# فقط ۴ بسته
ping -c 4 8.8.8.8

# پینگ محلی
ping -c 3 192.168.1.1
```

# Common mistakes

- نتیجه نگرفتن را فقط «قطع بودن اینترنت» دانستن — فایروال ممکن است ICMP را بلاک کند.
- رها کردن `ping` بدون `-c` در اسکریپت‌ها.

# Tips

- برای تست DNS جدا از شبکه: `ping` به IP و به نام دامنه را مقایسه کنید.
- `ping -c 1 -W 1 host && echo up` برای چک سریع.

# Related commands

- `traceroute` / `tracepath` — مسیر بسته‌ها
- `curl` — تست HTTP
- `ss` / `netstat` — سوکت‌های باز
