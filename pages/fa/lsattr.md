---
title: lsattr
aliases:
category: permissions
difficulty: intermediate
keywords:
- attributes
- filesystem
- immutable
---

# Introduction

`lsattr` ویژگی‌های خاص فایل‌سیستم (مثل immutable یا append-only) که با `chattr` تنظیم شده‌اند را برای فایل‌ها نمایش می‌دهد.

# Syntax

```
lsattr [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-a` | نمایش فایل‌های مخفی هم |
| `-d` | نمایش خود پوشه به‌جای محتویاتش |
| `-R` | نمایش بازگشتی برای زیرپوشه‌ها |
| `-l` | نمایش نام کامل ویژگی‌ها به‌جای حروف مخفف |

# Examples

```bash
# نمایش ویژگی‌های فایل‌های یک پوشه
lsattr /etc/

# بررسی این‌که آیا یک فایل خاص immutable است یا نه
lsattr /etc/important-config.conf

# نمایش بازگشتی برای کل یک پوشه
lsattr -R /etc/critical-configs/

# نمایش با نام کامل ویژگی‌ها (خواناتر)
lsattr -l /etc/passwd
```

# Common mistakes

- گیج‌شدن خروجی حروف مخفف (مثل `i----------------`) بدون دانستن معنی هر حرف؛ استفاده از `-l` خروجی خواناتری می‌دهد.

# Tips

- قبل از تلاش برای حذف یا ویرایش فایلی که به‌طور مرموزی «permission denied» می‌دهد با وجود مجوزهای درست، `lsattr` را چک کنید — ممکن است `+i` (immutable) روی آن تنظیم شده باشد.

# Related commands

- `chattr` — تنظیم ویژگی‌های فایل‌سیستم
- `ls -l` — نمایش مجوزهای معمولی (متفاوت از ویژگی‌های خاص)
