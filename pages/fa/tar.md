---
title: tar
aliases:
- archive
category: archive
difficulty: intermediate
keywords:
- archive
- compress
- backup
---

# Introduction

دستور `tar` (tape archive) برای ساخت و استخراج آرشیو استفاده می‌شود. معمولاً با فشرده‌سازی gzip یا bzip2 ترکیب می‌شود.

# Syntax

```
tar [OPTIONS] [ARCHIVE] [FILE...]
```

# Options اصلی

| گزینه | توضیح |
|-------|--------|
| `-c` | ساخت آرشیو جدید (create) |
| `-x` | استخراج (extract) |
| `-t` | لیست محتویات (list) |
| `-f` | نام فایل آرشیو |
| `-v` | verbose |
| `-z` | استفاده از gzip |
| `-j` | استفاده از bzip2 |
| `-J` | استفاده از xz |
| `-C DIR` | تغییر به پوشه قبل از عملیات |
| `--exclude=PATTERN` | نادیده گرفتن |

# Examples

```bash
# ساخت آرشیو فشرده gzip
tar -czvf archive.tar.gz folder/

# استخراج
tar -xzvf archive.tar.gz

# استخراج به پوشه خاص
tar -xzvf archive.tar.gz -C /tmp/dest/

# لیست محتویات
tar -tzvf archive.tar.gz

# ساخت بدون فشرده‌سازی
tar -cvf archive.tar folder/

# نادیده گرفتن برخی فایل‌ها
tar -czvf backup.tar.gz --exclude='*.log' --exclude='node_modules' project/
```

# Common mistakes

- فراموش کردن `-f` قبل از نام آرشیو.
- ترتیب گزینه‌ها (برخی نسخه‌های قدیمی حساس هستند).
- استخراج بدون `-C` و پر شدن پوشه جاری.

# Tips

- پسوندهای رایج: `.tar.gz` / `.tgz`، `.tar.bz2`، `.tar.xz`
- برای آرشیوهای خیلی بزرگ از `pigz` به جای gzip استفاده کنید.
- `tar` می‌تواند از stdin/stdout کار کند.

# Related commands

- `gzip` / `gunzip`
- `zip` / `unzip`
- `7z`
- `rsync` — برای بکاپ
