---
title: cp
aliases:
- copy
category: filesystem
difficulty: beginner
keywords:
- copy
- files
- directories
---

# Introduction

دستور `cp` برای کپی کردن فایل‌ها و پوشه‌ها استفاده می‌شود.

برخلاف جابه‌جایی (`mv`)، فایل اصلی سر جایش می‌ماند و یک کپی جدید ساخته می‌شود.

# Syntax

```
cp [OPTIONS] SOURCE DEST
cp [OPTIONS] SOURCE... DIRECTORY
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-r` یا `-R` | کپی بازگشتی پوشه‌ها |
| `-i` | قبل از بازنویسی فایل موجود سؤال کند |
| `-v` | نمایش فایل‌هایی که کپی می‌شوند (verbose) |
| `-u` | فقط اگر فایل مبدأ جدیدتر باشد کپی کن |
| `-p` | حفظ مشخصات فایل (مالک، مجوز، زمان) |
| `-a` | آرشیو (معادل `-dR --preserve=all`) |
| `-n` | هرگز فایل موجود را بازنویسی نکن |

# Examples

```bash
# کپی یک فایل
cp file.txt file_backup.txt

# کپی فایل به پوشه دیگر
cp report.pdf ~/Documents/

# کپی چندین فایل به یک پوشه
cp *.jpg ~/Pictures/

# کپی کامل یک پوشه
cp -r projects/ projects_backup/

# کپی با حفظ مشخصات
cp -a important/ important_backup/
```

# Common mistakes

- فراموش کردن `-r` هنگام کپی پوشه → خطا می‌دهد.
- بازنویسی تصادفی فایل‌های مهم بدون `-i`.
- استفاده از `cp -r source dest` وقتی `dest` از قبل وجود دارد (محتوا داخل آن کپی می‌شود، نه جایگزینی کامل).

# Tips

- همیشه برای پوشه‌ها از `-a` یا `-r` استفاده کنید.
- برای کپی ایمن: `cp -iv`
- اگر می‌خواهید پیشرفت را ببینید، از `rsync -avh --progress` به جای `cp` استفاده کنید.

# Related commands

- `mv` — جابه‌جایی / تغییر نام
- `rsync` — کپی پیشرفته و همگام‌سازی
- `install` — کپی با تنظیم مجوز
- `dd` — کپی سطح پایین
