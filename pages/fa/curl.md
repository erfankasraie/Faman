---
title: curl
aliases:
category: network
difficulty: intermediate
keywords:
- http
- download
- api
- transfer
---

# Introduction

`curl` داده را از یا به سرور با پروتکل‌های مختلف (بیشتر HTTP/HTTPS) منتقل می‌کند. ابزار استاندارد کار با API است.

# Syntax

```
curl [OPTIONS] URL
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-O` | ذخیره با همان نام فایل ریموت |
| `-o FILE` | ذخیره در FILE |
| `-L` | دنبال redirect |
| `-I` | فقط هدرها |
| `-X METHOD` | متد HTTP |
| `-H HEADER` | هدر سفارشی |
| `-d DATA` | بدنه درخواست |
| `-u USER:PASS` | احراز هویت |
| `-sS` | ساکت با نمایش خطا |
| `-f` | خطا روی HTTP error |
| `-k` | نادیده گرفتن TLS (نامناسب برای production) |

# Examples

```bash
# دانلود
curl -LO https://example.com/file.tar.gz

# GET ساده
curl -s https://api.github.com/zen

# POST JSON
curl -sS -X POST -H 'Content-Type: application/json' \
  -d '{"name":"faman"}' https://httpbin.org/post

# هدرها
curl -I https://example.com

# با کد وضعیت
curl -sS -o /dev/null -w '%{http_code}\n' https://example.com
```

# Common mistakes

- فراموش کردن `-L` وقتی سایت redirect می‌شود.
- استفاده از `-k` به‌صورت عادت.
- ننوشتن `-f` در اسکریپت و ادامه بعد از 404.

# Tips

- برای JSON: `curl -s ... | jq .`
- در اسکریپت‌ها `-sS` ترکیب خوبی است.
- فایل تنظیمات: `~/.curlrc`

# Related commands

- `wget` — دانلود فایل‌محور
- `httpie` — سینتکس دوستانه
- `jq` — پردازش JSON
- `ssh` — دسترسی به سرور
