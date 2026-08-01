---
title: unzip
aliases:
category: archive
difficulty: beginner
keywords:
- extract
- decompress
- zip
---

# Introduction

دستور `unzip` برای استخراج محتویات فایل‌های ZIP استفاده می‌شود.

# Syntax

```
unzip [OPTIONS] ARCHIVE [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-l` | فقط لیست محتویات |
| `-d DIR` | استخراج به پوشه خاص |
| `-o` | بازنویسی بدون سؤال |
| `-q` | حالت ساکت |
| `-P PASSWORD` | پسورد |

# Examples

```bash
# استخراج در پوشه جاری
unzip archive.zip

# استخراج به پوشه خاص
unzip archive.zip -d /tmp/extracted/

# فقط لیست محتویات
unzip -l archive.zip

# استخراج فایل خاص
unzip archive.zip file.txt
```

# Common mistakes

- استخراج بدون `-d` و شلوغ شدن پوشه جاری.
- مشکل encoding نام فایل‌های فارسی (گاهی نیاز به `-O` یا ابزارهای دیگر).

# Tips

- قبل از استخراج همیشه با `-l` محتویات را بررسی کنید.
- برای فایل‌های ZIP آسیب‌دیده ابزارهای تخصصی‌تری وجود دارد.
- `bsdtar` یا `7z` هم می‌توانند ZIP را باز کنند.

# Related commands

- `zip` — ساخت آرشیو
- `tar` — آرشیو لینوکسی
- `7z` — پشتیبانی از فرمت‌های بیشتر
- `jar` — برای فایل‌های jar (که در واقع zip هستند)
