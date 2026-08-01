---
title: chmod
aliases:
category: permissions
difficulty: intermediate
keywords:
- permissions
- mode
- access
---

# Introduction

دستور `chmod` (change mode) مجوزهای دسترسی فایل‌ها و پوشه‌ها را تغییر می‌دهد.

# Syntax

```
chmod [OPTIONS] MODE FILE...
chmod [OPTIONS] OCTAL-MODE FILE...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-R` | بازگشتی روی پوشه‌ها |
| `-v` | نمایش تغییرات |
| `-c` | فقط تغییرات را نشان بده |
| `--reference=RFILE` | کپی مجوز از فایل دیگر |

# حالت‌های مجوز

## عددی (Octal)

| عدد | معنی |
|-----|------|
| 4 | خواندن (r) |
| 2 | نوشتن (w) |
| 1 | اجرا (x) |

ترکیب‌ها: `7` = rwx، `6` = rw-، `5` = r-x، `4` = r--

مثال: `755` = rwxr-xr-x

## سمبلیک

```
u = user (مالک)
g = group
o = others
a = all

+ اضافه کردن
- حذف کردن
= تنظیم دقیق
```

# Examples

```bash
# دادن مجوز اجرا به مالک
chmod u+x script.sh

# تنظیم دقیق 755
chmod 755 script.sh

# حذف نوشتن از group و others
chmod go-w file.txt

# بازگشتی
chmod -R 755 project/

# فقط به پوشه‌ها اجرا بده
find . -type d -exec chmod 755 {} \;
```

# Common mistakes

- دادن `777` به همه چیز (امنیت بسیار ضعیف).
- فراموش کردن `-R` برای پوشه‌ها.
- اشتباه گرفتن `u` و `o`.

# Tips

- از `ls -l` برای دیدن مجوزهای فعلی استفاده کنید.
- برای اسکریپت‌ها معمولاً `755` یا `700` مناسب است.
- از `umask` برای تنظیم مجوز پیش‌فرض فایل‌های جدید استفاده کنید.

# Related commands

- `chown` — تغییر مالک
- `chgrp` — تغییر گروه
- `ls -l` — مشاهده مجوزها
- `umask` — مجوز پیش‌فرض
- `stat` — اطلاعات کامل
