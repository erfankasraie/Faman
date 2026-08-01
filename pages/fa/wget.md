---
title: wget
aliases:
category: network
difficulty: beginner
keywords:
- download
- http
- recursive
---

# Introduction

دستور `wget` ابزاری غیرتعاملی برای دانلود فایل از وب است. برای دانلود‌های ساده و بازگشتی بسیار محبوب است.

# Syntax

```
wget [OPTIONS] URL...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-O FILE` | ذخیره با نام دلخواه |
| `-c` | ادامه دانلود ناقص |
| `-b` | اجرا در پس‌زمینه |
| `-q` | ساکت |
| `-r` | دانلود بازگشتی |
| `-np` | به پوشه والد نرو |
| `-k` | تبدیل لینک‌ها برای مشاهده آفلاین |
| `-P DIR` | ذخیره در پوشه خاص |
| `--limit-rate=RATE` | محدودیت سرعت |
| `-U AGENT` | تنظیم User-Agent |
| `--no-check-certificate` | نادیده گرفتن SSL |

# Examples

```bash
# دانلود ساده
wget https://example.com/file.zip

# با نام دلخواه
wget -O myfile.zip https://example.com/file.zip

# ادامه دانلود
wget -c https://example.com/large.iso

# دانلود کل یک سایت (با احتیاط)
wget -r -np -k https://example.com/docs/

# محدودیت سرعت
wget --limit-rate=500k https://example.com/file.iso
```

# Common mistakes

- استفاده از `-r` بدون `-np` و دانلود کل سایت.
- فراموش کردن `-c` برای فایل‌های بزرگ.

# Tips

- برای دانلود‌های موازی و قدرتمندتر از `aria2c` استفاده کنید.
- `wget` به طور پیش‌فرض نام فایل را از URL می‌گیرد.
- می‌توانید لیستی از URLها را از فایل بخوانید: `wget -i urls.txt`

# Related commands

- `curl` — انعطاف‌پذیرتر برای API
- `aria2c` — دانلود موازی
- `httpie`
- `scp` / `rsync`
