---
title: mv
aliases:
- move
- rename
category: filesystem
difficulty: beginner
keywords:
- move
- rename
- files
---

# Introduction

دستور `mv` برای جابه‌جا کردن یا تغییر نام فایل‌ها و پوشه‌ها استفاده می‌شود.

اگر مبدأ و مقصد در یک فایل‌سیستم باشند، عملیات بسیار سریع است (فقط تغییر نام در جدول فایل‌سیستم).

# Syntax

```
mv [OPTIONS] SOURCE DEST
mv [OPTIONS] SOURCE... DIRECTORY
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-i` | قبل از بازنویسی سؤال کند |
| `-v` | نمایش عملیات (verbose) |
| `-n` | بازنویسی نکند |
| `-u` | فقط اگر مبدأ جدیدتر باشد |
| `-b` | قبل از بازنویسی، از فایل مقصد پشتیبان بگیرد |

# Examples

```bash
# تغییر نام فایل
mv oldname.txt newname.txt

# جابه‌جایی فایل به پوشه دیگر
mv report.pdf ~/Documents/

# جابه‌جایی چندین فایل
mv *.log /var/log/archive/

# تغییر نام پوشه
mv old_project/ new_project/
```

# Common mistakes

- استفاده از `mv` به جای `cp` و از دست دادن فایل اصلی.
- جابه‌جایی فایل‌های باز توسط برنامه‌ها (ممکن است مشکل ایجاد کند).
- فراموش کردن اینکه `mv` روی پارتیشن‌های مختلف در واقع کپی + حذف است.

# Tips

- برای تغییر نام دسته‌ای می‌توانید از `rename` یا حلقه bash استفاده کنید.
- همیشه با `-i` کار کنید تا اشتباهات جبران‌ناپذیر کمتر شود.
- `mv` روی همان فایل‌سیستم atomic است.

# Related commands

- `cp` — کپی
- `rm` — حذف
- `rename` — تغییر نام پیشرفته
- `rsync` — همگام‌سازی
