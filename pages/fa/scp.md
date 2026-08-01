---
title: scp
aliases:
category: network
difficulty: intermediate
keywords:
- copy
- remote
- secure
- transfer
---

# Introduction

دستور `scp` (Secure Copy) برای کپی امن فایل بین ماشین‌های محلی و راه دور از طریق SSH استفاده می‌شود.

# Syntax

```
scp [OPTIONS] SOURCE DEST
```

SOURCE و DEST می‌توانند به فرم `user@host:path` باشند.

# Options

| گزینه | توضیح |
|-------|--------|
| `-r` | بازگشتی (پوشه‌ها) |
| `-P PORT` | پورت SSH (توجه: P بزرگ) |
| `-i KEY` | کلید خصوصی |
| `-C` | فشرده‌سازی |
| `-p` | حفظ زمان و مجوز |
| `-v` | verbose |
| `-q` | ساکت |

# Examples

```bash
# کپی فایل به سرور
scp file.txt user@host:/home/user/

# کپی از سرور به محلی
scp user@host:/var/log/app.log ./

# کپی پوشه
scp -r project/ user@host:~/backups/

# با پورت خاص
scp -P 2222 file.txt user@host:~/

# با کلید خاص
scp -i ~/.ssh/mykey file.txt user@host:~/
```

# Common mistakes

- استفاده از `-p` به جای `-P` برای پورت.
- فراموش کردن `-r` برای پوشه‌ها.
- مسیرهای دارای فاصله بدون کوتیشن.

# Tips

- برای همگام‌سازی و انتقال‌های بزرگ، `rsync` روی SSH بسیار بهتر است.
- `scp` در نسخه‌های جدید OpenSSH از پروتکل SFTP استفاده می‌کند.
- می‌توانید بین دو سرور راه دور هم کپی کنید (با محدودیت).

# Related commands

- `rsync` — همگام‌سازی پیشرفته
- `sftp` — انتقال تعاملی
- `ssh` — اتصال شل
- `rclone` — برای فضای ابری
