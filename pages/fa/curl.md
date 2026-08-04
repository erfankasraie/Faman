---
title: curl
aliases:
category: network
difficulty: intermediate
keywords:
- http
- https
- api
- download
- transfer
---

# Introduction

`curl` کلاینت انتقال داده از خط فرمان است: HTTP/HTTPS، FTP، SFTP و … . استاندارد کار با API، دانلود و دیباگ درخواست‌های شبکه است.

# Syntax

```
curl [OPTIONS] [URL...]
```

# Options

## URL، متد، redirect

| گزینه | توضیح |
|-------|--------|
| `-X METHOD` / `--request` | GET/POST/PUT/DELETE/… |
| `-L` / `--location` | دنبال کردن redirect |
| `--max-redirs N` | سقف redirect |
| `-I` / `--head` | فقط هدرها (HEAD) |
| `-G` | داده‌های `-d` را به query string تبدیل کن |
| `--url URL` | URL صریح |

## بدنه و هدر

| گزینه | توضیح |
|-------|--------|
| `-H 'Name: value'` | هدر |
| `-d DATA` / `--data` | بدنه POST (application/x-www-form-urlencoded) |
| `--data-raw` | بدون @file interpretation |
| `--data-binary` | باینری |
| `-F name=value` | multipart form |
| `--json DATA` | JSON + هدر مناسب (نسخه‌های جدید) |
| `-T FILE` | آپلود (PUT) |

## خروجی و فایل

| گزینه | توضیح |
|-------|--------|
| `-o FILE` | ذخیره در FILE |
| `-O` | نام از URL |
| `-J` | نام از هدر Content-Disposition |
| `-s` / `--silent` | بدون progress meter |
| `-S` | با `-s` خطاها را نشان بده |
| `-v` | verbose (دیباگ) |
| `-i` | هدر پاسخ + بدنه |
| `-D FILE` | هدرها در FILE |
| `-w FORMAT` | بعد از تمام شدن، فرمت سفارشی |
| `-f` | روی HTTP 4xx/5xx کد خروج خطا |
| `--fail-with-body` | مثل fail ولی بدنه را نگه دار |

## امنیت و اتصال

| گزینه | توضیح |
|-------|--------|
| `-u USER:PASS` | Basic auth |
| `-A STRING` | User-Agent |
| `-e URL` | Referer |
| `-k` / `--insecure` | تأیید TLS را رد کن (خطرناک) |
| `--cacert FILE` | CA سفارشی |
| `--cert` / `--key` | کلاینت‌سرتیفیکیت |
| `-x PROXY` | پروکسی |
| `--connect-timeout N` | تایم‌اوت اتصال |
| `-m N` / `--max-time` | سقف کل زمان |
| `-C -` | ادامه دانلود ناتمام |
| `-Z` | دانلود موازی چند URL (نسخه‌های جدید) |

# Examples

## دانلود و GET

```bash
curl -LO https://example.com/file.tar.gz
curl -sL https://example.com
curl -I https://example.com
```

## API JSON

```bash
curl -sS https://api.github.com/zen

curl -sS -X POST https://httpbin.org/post \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"name":"faman","ok":true}'

# پاسخ را با jq قشنگ کن
curl -sS https://api.github.com/repos/erfankasraie/Faman | jq '.full_name, .stargazers_count'
```

## فرم، آپلود، کد وضعیت

```bash
curl -sS -F 'file=@./report.pdf' https://httpbin.org/post

curl -sS -o /dev/null -w '%{http_code} time=%{time_total}\n' https://example.com

curl -sS -f -o out.json https://api.example.com/data || echo "failed: $?"
```

## دیباگ TLS و پروکسی

```bash
curl -vI https://example.com
curl -x http://proxy:8080 https://example.com
```

## از فایل تنظیمات

```bash
# ~/.curlrc
# silent
# show-error
```

# Common mistakes

- فراموش کردن `-L` پشت CDN/www redirect.
- عادت به `-k` در production.
- در اسکریپت بدون `-f` یا چک `%{http_code}` ادامه دادن بعد از 404.
- JSON با `-d` بدون هدر `Content-Type: application/json`.

# Tips

- اسکریپت‌ها: `-sS` یا `-sS -f`.
- زمان و سرعت: `-w` با متغیرهایی مثل `time_total`, `size_download`.
- کوکی: `-c cookies.txt` ذخیره، `-b cookies.txt` ارسال.

# Related commands

- `wget` — دانلود بازگشتی/سایت
- `http` (HTTPie) — سینتکس دوستانه‌تر
- `jq` — JSON
- `ssh` — دسترسی به سرور
