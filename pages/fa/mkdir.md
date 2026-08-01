---
title: mkdir
aliases:
- md
category: filesystem
difficulty: beginner
keywords:
- directory
- folder
- create
---

# Introduction

دستور `mkdir` برای ساخت پوشه (directory) جدید استفاده می‌شود.

# Syntax

```
mkdir [OPTIONS] DIRECTORY...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-p` | ساخت پوشه‌های والد در صورت نیاز (parents) |
| `-v` | نمایش پوشه‌های ساخته‌شده |
| `-m` | تنظیم مجوز همزمان با ساخت |

# Examples

```bash
# ساخت یک پوشه
mkdir projects

# ساخت چندین پوشه
mkdir dir1 dir2 dir3

# ساخت مسیر کامل (پوشه‌های والد هم ساخته می‌شوند)
mkdir -p ~/projects/web/frontend/src

# ساخت با مجوز خاص
mkdir -m 755 public_html
```

# Common mistakes

- فراموش کردن `-p` وقتی مسیر والد وجود ندارد → خطا.
- ساخت پوشه با نامی که قبلاً وجود دارد → خطا (مگر با `-p`).

# Tips

- همیشه برای مسیرهای تو در تو از `-p` استفاده کنید.
- می‌توانید در یک دستور چندین سطح بسازید.
- برای ساخت ساختار پروژه از اسکریپت یا `mkdir -p` استفاده کنید.

# Related commands

- `rmdir` — حذف پوشه خالی
- `rm -r` — حذف پوشه غیرخالی
- `touch` — ساخت فایل خالی
- `install -d` — ساخت پوشه با مجوز
