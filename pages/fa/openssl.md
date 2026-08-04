---
title: openssl
aliases:
category: security
difficulty: advanced
keywords:
- ssl
- tls
- certificate
- encryption
- security
---

# Introduction

`openssl` مجموعه‌ابزاری قدرتمند برای کار با رمزنگاری است: ساخت گواهی‌های SSL/TLS، رمزنگاری و رمزگشایی فایل‌ها، محاسبه هش، و بررسی اتصال امن سرورها.

# Syntax

```
openssl COMMAND [OPTIONS]
```

# Options

زیردستورهای پرکاربرد:

| زیردستور | توضیح |
|-------|--------|
| `genrsa` | ساخت کلید خصوصی RSA |
| `req` | ساخت CSR یا گواهی خودامضا (self-signed) |
| `x509` | مشاهده و مدیریت گواهی‌ها |
| `s_client` | تست اتصال TLS به یک سرور |
| `enc` | رمزنگاری/رمزگشایی متقارن فایل |
| `dgst` | محاسبه هش (مثل sha256) |

# Examples

```bash
# ساخت یک کلید خصوصی RSA ۲۰۴۸ بیتی
openssl genrsa -out key.pem 2048

# ساخت یک گواهی خودامضا برای تست (معتبر ۳۶۵ روز)
openssl req -x509 -new -key key.pem -out cert.pem -days 365

# دیدن اطلاعات یک گواهی SSL (تاریخ انقضا و صادرکننده)
openssl x509 -in cert.pem -noout -text

# بررسی گواهی SSL یک سایت از راه دور
openssl s_client -connect example.com:443 -servername example.com

# محاسبه هش SHA-256 یک فایل
openssl dgst -sha256 file.txt

# رمزنگاری یک فایل با AES-256
openssl enc -aes-256-cbc -salt -in secret.txt -out secret.enc
```

# Common mistakes

- فراموش‌کردن `-servername` هنگام تست SSL روی سرورهایی که چند دامنه با یک IP دارند (SNI)؛ بدون آن ممکن است گواهی اشتباه برگردد.
- استفاده از الگوریتم‌های قدیمی و ناامن (مثل `-des` یا کلیدهای کوچک) که دیگر توصیه نمی‌شوند.
- گیج‌شدن بین فرمت‌های PEM، DER و PKCS12 هنگام کار با گواهی‌ها.

# Tips

- برای تست سریع اعتبار و تاریخ انقضای گواهی یک سایت: `echo | openssl s_client -connect example.com:443 2>/dev/null | openssl x509 -noout -dates`
- برای اکثر کارهای روزمره (مثل ساخت گواهی Let's Encrypt) از ابزارهایی مثل `certbot` استفاده کنید که خودشان روی openssl ساخته شده‌اند.

# Related commands

- `curl` — می‌تواند برای تست HTTPS هم استفاده شود
- `ssh-keygen` — ساخت کلید برای SSH (نه TLS)
- `gpg` — رمزنگاری و امضای دیجیتال با GPG
