---
title: curl
aliases:
category: network
difficulty: intermediate
keywords:
- download
- http
- api
- transfer
---

# Introduction

دستور `curl` ابزاری قدرتمند برای انتقال داده از یا به سرور است. از پروتکل‌های متعددی مثل HTTP، HTTPS، FTP، SFTP و غیره پشتیبانی می‌کند.

# Syntax

```
curl [OPTIONS] URL
```

# Options پرکاربرد

| گزینه | توضیح |
|-------|--------|
| `-O` | ذخیره با همان نام فایل ریموت |
| `-o FILE` | ذخیره با نام دلخواه |
| `-L` | دنبال کردن redirect |
| `-C -` | ادامه دانلود ناقص |
| `-I` | فقط هدرها (HEAD) |
| `-X METHOD` | متد HTTP (GET, POST, ...) |
| `-H "Header"` | اضافه کردن هدر |
| `-d DATA` | ارسال داده (POST) |
| `-u USER:PASS` | احراز هویت |
| `-k` | نادیده گرفتن گواهی SSL |
| `-s` | حالت ساکت |
| `-v` | verbose |
| `-w` | فرمت خروجی سفارشی |
| `--json` | ارسال JSON (نسخه‌های جدید) |

# Examples

```bash
# دانلود فایل
curl -O https://example.com/file.zip

# دانلود با نام دلخواه
curl -o myfile.zip https://example.com/file.zip

# دنبال کردن redirect
curl -LO https://github.com/.../releases/latest

# درخواست GET ساده
curl https://api.example.com/users

# درخواست POST با JSON
curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"Ali"}' https://api.example.com/users

# ادامه دانلود
curl -C - -O https://example.com/large.iso

# فقط هدرها
curl -I https://example.com
```

# Common mistakes

- فراموش کردن `-L` وقتی سایت redirect می‌کند.
- استفاده از `-k` در محیط production (غیرامن).
- فراموش کردن کوتیشن دور هدرها و داده‌ها.

# Tips

- برای دانلود‌های بزرگ و قابل‌ازسرگیری `wget` یا `aria2c` گاهی راحت‌تر است.
- `curl` برای کار با APIها عالی است.
- می‌توانید خروجی را مستقیماً به فایل یا پایپ بدهید.
- از `jq` برای پردازش JSON خروجی استفاده کنید.

# Related commands

- `wget` — دانلود ساده‌تر
- `httpie` — رابط کاربرپسندتر برای HTTP
- `aria2c` — دانلود موازی و قدرتمند
- `ssh` / `scp`
