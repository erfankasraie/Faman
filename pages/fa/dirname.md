---
title: dirname
aliases:
category: filesystem
difficulty: beginner
keywords:
- path
- directory
- script
---

# Introduction

`dirname` بخش پوشه (مسیر والد) یک مسیر فایل را استخراج می‌کند؛ دقیقاً برعکس `basename` که نام فایل را می‌گیرد.

# Syntax

```
dirname PATH
```

# Options

`dirname` گزینه خاصی ندارد؛ فقط یک مسیر می‌گیرد و بخش پوشه آن را برمی‌گرداند.

# Examples

```bash
# استخراج مسیر پوشه از یک مسیر کامل
dirname /home/user/documents/report.pdf
# → /home/user/documents

# استخراج پوشه والد یک پوشه
dirname /home/user/documents
# → /home/user

# رایج‌ترین کاربرد: پیداکردن مسیر واقعی پوشه‌ی یک اسکریپت در حال اجرا
SCRIPT_DIR=$(dirname "$(readlink -f "$0")")
echo "اسکریپت در $SCRIPT_DIR قرار دارد"

# استفاده برای رفتن به پوشه یک فایل خاص
cd "$(dirname /path/to/some/file.txt)"
```

# Common mistakes

- استفاده تنها از `dirname "$0"` بدون `readlink -f` برای پیداکردن مسیر واقعی اسکریپت؛ اگر اسکریپت از طریق یک لینک نمادین یا با مسیر نسبی اجرا شود، نتیجه ممکن است نادرست یا نسبی باشد. ترکیب با `readlink -f` این مشکل را حل می‌کند.
- گیج‌شدن dirname (پوشه والد) با basename (نام فایل)؛ این دو مکمل هم هستند.

# Tips

- الگوی استاندارد `SCRIPT_DIR=$(dirname "$(readlink -f "$0")")` در ابتدای اسکریپت‌های bash برای پیداکردن مسیر مطلق پوشه‌ی خود اسکریپت (مثلاً برای بارگذاری فایل‌های کمکی کنارش) بسیار پرکاربرد است.

# Related commands

- `basename` — استخراج نام فایل (برعکس dirname)
- `readlink -f` — حل‌کردن کامل مسیر شامل لینک‌ها
- `realpath` — گرفتن مسیر مطلق فایل یا پوشه
