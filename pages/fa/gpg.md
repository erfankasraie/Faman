---
title: gpg
aliases:
- gnupg
category: security
difficulty: advanced
keywords:
- encryption
- signature
- privacy
- pgp
---

# Introduction

`gpg` (GNU Privacy Guard) برای رمزنگاری فایل‌ها و ایمیل، امضای دیجیتال، و تأیید صحت فایل‌ها (مثلاً نصب‌کننده‌های نرم‌افزار) استفاده می‌شود؛ پیاده‌سازی آزاد استاندارد OpenPGP.

# Syntax

```
gpg [OPTIONS] [FILE]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `--gen-key` | ساخت جفت‌کلید جدید |
| `--encrypt` / `-e` | رمزنگاری فایل برای یک گیرنده |
| `--decrypt` / `-d` | رمزگشایی فایل |
| `--sign` | امضای دیجیتال فایل |
| `--verify` | تأیید امضای یک فایل |
| `--import` | وارد‌کردن یک کلید عمومی |
| `--list-keys` | لیست کلیدهای موجود در keyring |

# Examples

```bash
# ساخت جفت‌کلید جدید (تعاملی)
gpg --gen-key

# رمزنگاری یک فایل برای یک گیرنده خاص
gpg --encrypt --recipient user@example.com file.txt

# رمزگشایی یک فایل دریافت‌شده
gpg --decrypt file.txt.gpg > file.txt

# امضای دیجیتال یک فایل (تولید فایل .sig جدا)
gpg --detach-sign file.tar.gz

# تأیید صحت یک فایل دانلودشده با امضای آن
gpg --verify file.tar.gz.sig file.tar.gz

# وارد‌کردن کلید عمومی یک پروژه برای تأیید بسته‌ها
gpg --import project-key.asc
```

# Common mistakes

- گم‌کردن یا فراموش‌کردن passphrase کلید خصوصی؛ برخلاف رمز عبور، کلید GPG قابل بازیابی نیست.
- امضای فایل با `--sign` (که فایل را در خروجی باینری می‌پیچد) به‌جای `--detach-sign` وقتی فقط یک فایل امضای جداگانه لازم است.
- عدم بکاپ‌گیری از کلید خصوصی و revocation certificate؛ اگر کلید گم شود، داده‌های رمزشده با آن غیرقابل‌بازیابی می‌شوند.

# Tips

- برای بررسی امنیت یک دانلود (مثلاً یک توزیع لینوکس)، همیشه امضای GPG منتشرشده را با `gpg --verify` چک کنید.
- بلافاصله پس از ساخت کلید، یک revocation certificate بسازید (`gpg --gen-revoke`) و آن را جای امنی نگه دارید.

# Related commands

- `openssl` — رمزنگاری و گواهی‌های TLS/SSL (استاندارد متفاوت از OpenPGP)
- `sha256sum` — بررسی صحت فایل با هش (بدون امضای دیجیتال)
- `ssh-keygen` — ساخت کلید برای SSH
