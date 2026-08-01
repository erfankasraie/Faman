---
title: find
aliases:
- search-files
category: filesystem
difficulty: intermediate
keywords:
- search
- files
- locate
- recursive
---

# Introduction

دستور `find` قدرتمندترین ابزار جستجوی فایل در لینوکس است. می‌تواند بر اساس نام، نوع، اندازه، زمان، مجوز و بسیاری معیارهای دیگر جستجو کند.

# Syntax

```
find [PATH...] [EXPRESSION]
```

# Options

| گزینه/تست | توضیح |
|-----------|--------|
| `-name PATTERN` | جستجو بر اساس نام (حساس به حروف) |
| `-iname PATTERN` | جستجو بدون حساسیت به حروف |
| `-type f/d/l` | فقط فایل / پوشه / لینک |
| `-size +10M` | بزرگ‌تر از ۱۰ مگابایت |
| `-mtime -7` | تغییر یافته در ۷ روز گذشته |
| `-user USER` | متعلق به کاربر خاص |
| `-perm 644` | با مجوز خاص |
| `-exec CMD {} \;` | اجرای دستور روی نتایج |
| `-delete` | حذف نتایج |
| `-maxdepth N` | حداکثر عمق جستجو |

# Examples

```bash
# پیدا کردن تمام فایل‌های .txt
find . -name "*.txt"

# پیدا کردن فایل‌های بزرگ‌تر از ۱۰۰ مگابایت
find /home -type f -size +100M

# پیدا کردن و حذف فایل‌های موقت
find /tmp -name "*.tmp" -delete

# پیدا کردن فایل‌های تغییر یافته در ۲۴ ساعت گذشته
find . -mtime -1

# اجرای دستور روی نتایج
find . -name "*.log" -exec ls -lh {} \;

# جستجو بدون حساسیت به حروف
find . -iname "readme*"
```

# Common mistakes

- فراموش کردن کوتیشن دور patternهایی که `*` دارند.
- استفاده از `-exec` بدون `\;` یا `+`.
- جستجوی کل `/` بدون `-maxdepth` یا محدودیت → بسیار کند.

# Tips

- برای جستجوی سریع‌تر نام فایل از `locate` یا `plocate` استفاده کنید (نیاز به دیتابیس به‌روز).
- `-exec ... +` کارآمدتر از `\;` است.
- می‌توانید چندین شرط را با `-a` (و) و `-o` (یا) ترکیب کنید.
- برای عملکرد بهتر از `-maxdepth` استفاده کنید.

# Related commands

- `locate` / `plocate` — جستجوی سریع بر اساس نام
- `grep` — جستجو داخل محتوای فایل‌ها
- `fd` — جایگزین مدرن و سریع‌تر find
- `tree` — نمایش درختی
