---
title: basename
aliases:
category: filesystem
difficulty: beginner
keywords:
- path
- filename
- script
---

# Introduction

`basename` نام فایل را از یک مسیر کامل استخراج می‌کند (بخش پوشه‌ها را حذف می‌کند). خیلی در اسکریپت‌نویسی برای گرفتن فقط نام فایل (بدون مسیر یا حتی بدون پسوند) استفاده می‌شود.

# Syntax

```
basename PATH [SUFFIX]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `SUFFIX` | پسوندی که در صورت وجود از انتهای نام حذف شود |
| `-s SUFFIX` | مشابه بالا |
| `-a` | پردازش چند مسیر همزمان |

# Examples

```bash
# استخراج فقط نام فایل از مسیر کامل
basename /home/user/documents/report.pdf
# → report.pdf

# حذف پسوند مشخص هم
basename /home/user/report.pdf .pdf
# → report

# پردازش چند مسیر همزمان
basename -a /a/file1.txt /b/file2.txt

# استفاده در اسکریپت برای پیداکردن نام اسکریپت در حال اجرا
echo "این اسکریپت $(basename "$0") است"

# استخراج نام پوشه (نه فایل) از یک مسیر
basename /home/user/documents/
# → documents
```

# Common mistakes

- انتظار داشتن حذف خودکار پسوند بدون مشخص‌کردن آن؛ `basename` فقط پسوندی را حذف می‌کند که دقیقاً به آن بدهید (مثلاً `.pdf`)، نه هر پسوندی.
- فراموش‌کردن که `basename` روی یک مسیر با `/` در انتها، همان پوشه آخر را برمی‌گرداند نه رشته خالی.

# Tips

- ترکیب رایج در اسکریپت‌ها: `filename=$(basename "$filepath")` برای گرفتن فقط نام فایل جهت نمایش یا لاگ.
- برای گرفتن نام فایل بدون هیچ پسوندی وقتی پسوند را نمی‌دانید: `basename "$file" | cut -d'.' -f1` یا با `${file%.*}` در bash.

# Related commands

- `dirname` — استخراج مسیر پوشه (برعکس basename)
- `realpath` — گرفتن مسیر کامل و مطلق
- `readlink -f` — حل‌کردن کامل مسیر شامل لینک‌ها
