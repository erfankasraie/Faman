---
title: zip
aliases:
category: archive
difficulty: beginner
keywords:
- compress
- archive
- zip
---

# Introduction

دستور `zip` برای ساخت فایل‌های فشرده با فرمت ZIP استفاده می‌شود. این فرمت در ویندوز بسیار رایج است.

# Syntax

```
zip [OPTIONS] ARCHIVE FILE...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-r` | بازگشتی (برای پوشه‌ها) |
| `-e` | رمزگذاری با پسورد |
| `-x PATTERN` | نادیده گرفتن |
| `-u` | به‌روزرسانی فایل‌های موجود در آرشیو |
| `-9` | حداکثر فشرده‌سازی |
| `-0` | بدون فشرده‌سازی (فقط آرشیو) |

# Examples

```bash
# فشرده‌سازی چند فایل
zip archive.zip file1.txt file2.txt

# فشرده‌سازی پوشه
zip -r project.zip project/

# با رمز عبور
zip -e secret.zip confidential.txt

# نادیده گرفتن برخی فایل‌ها
zip -r backup.zip project/ -x "*.git*" -x "*node_modules*"
```

# Common mistakes

- فراموش کردن `-r` برای پوشه‌ها.
- استفاده از `zip` وقتی `tar.gz` مناسب‌تر است (برای لینوکس).

# Tips

- برای سازگاری با ویندوز، ZIP انتخاب خوبی است.
- برای بکاپ لینوکسی معمولاً `tar.gz` یا `tar.xz` بهتر است.
- می‌توانید فایل‌ها را به آرشیو موجود اضافه کنید.

# Related commands

- `unzip` — استخراج
- `tar` — آرشیو لینوکسی
- `7z` — فشرده‌سازی قوی‌تر
- `gzip`
