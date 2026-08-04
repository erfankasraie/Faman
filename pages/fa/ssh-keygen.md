---
title: ssh-keygen
aliases:
category: network
difficulty: intermediate
keywords:
- ssh
- key
- authentication
- security
---

# Introduction

`ssh-keygen` جفت‌کلید رمزنگاری (کلید خصوصی و عمومی) برای احراز هویت SSH بدون رمز عبور می‌سازد؛ روش استاندارد و امن برای اتصال به سرورها.

# Syntax

```
ssh-keygen [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-t TYPE` | نوع کلید (`ed25519` توصیه‌شده، یا `rsa`) |
| `-b BITS` | طول کلید (برای rsa، مثلاً `4096`) |
| `-f FILE` | مسیر ذخیره کلید |
| `-C COMMENT` | برچسب توضیحی (معمولاً ایمیل) |
| `-p` | تغییر passphrase یک کلید موجود |

# Examples

```bash
# ساخت یک کلید مدرن و امن (توصیه‌شده)
ssh-keygen -t ed25519 -C "user@example.com"

# ساخت کلید RSA با طول بالا (سازگاری بیشتر با سیستم‌های قدیمی)
ssh-keygen -t rsa -b 4096 -C "user@example.com"

# کپی کلید عمومی به یک سرور برای ورود بدون رمز عبور
ssh-copy-id user@server.com

# نمایش fingerprint یک کلید (برای تأیید هویت)
ssh-keygen -lf ~/.ssh/id_ed25519.pub

# تغییر passphrase یک کلید موجود
ssh-keygen -p -f ~/.ssh/id_ed25519
```

# Common mistakes

- به‌اشتراک‌گذاشتن یا ارسال فایل کلید **خصوصی** (بدون پسوند `.pub`)؛ فقط فایل `.pub` باید به دیگران یا سرور داده شود.
- ساخت کلید بدون passphrase روی یک لپ‌تاپ شخصی که ریسک امنیتی در صورت سرقت دستگاه دارد.
- فراموش‌کردن `chmod 600` روی فایل کلید خصوصی که باعث خطای «bad permissions» در SSH می‌شود.

# Tips

- الگوریتم `ed25519` نسبت به `rsa` سریع‌تر، امن‌تر و کلید کوتاه‌تری دارد؛ برای کلیدهای جدید توصیه می‌شود.
- برای مدیریت چند کلید برای سرورهای مختلف، فایل `~/.ssh/config` را تنظیم کنید تا نیازی به تایپ `-i` هر بار نباشد.

# Related commands

- `ssh` — اتصال به سرور راه دور
- `scp` / `rsync` — انتقال فایل روی SSH
- `ssh-copy-id` — کپی کلید عمومی به سرور
