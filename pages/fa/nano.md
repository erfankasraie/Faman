---
title: nano
aliases:
category: editor
difficulty: beginner
keywords:
- editor
- text
- edit
- file
---

# Introduction

`nano` یک ویرایشگر متن ساده و کاربرپسند در ترمینال است. برای ویرایش سریع فایل‌های تنظیمات و اسکریپت‌ها عالی است.

# Syntax

```
nano [OPTIONS] [FILE]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `+LINE` | باز کردن در خط مشخص |
| `-v` | فقط خواندنی (view) |
| `-B` | پشتیبان‌گیری خودکار |
| `-c` | نمایش مداوم موقعیت نشانگر |
| `-m` | پشتیبانی ماوس |

# Examples

```bash
# باز کردن / ساخت فایل
nano notes.txt

# پرش به خط ۵۰
nano +50 /etc/hosts

# فقط مشاهده
nano -v README.md
```

کلیدهای مهم داخل nano:

- `Ctrl+O` — ذخیره
- `Ctrl+X` — خروج
- `Ctrl+W` — جستجو
- `Ctrl+K` — بریدن خط
- `Ctrl+U` — چسباندن

# Common mistakes

- خروج بدون ذخیره و از دست دادن تغییرات.
- گیج شدن با میانبرها چون با ویرایشگرهای گرافیکی فرق دارند.

# Tips

- در پایین صفحه میانبرها همیشه نمایش داده می‌شوند (`^` یعنی Ctrl).
- برای کار حرفه‌ای‌تر بعداً `vim` یا `nvim` را یاد بگیرید.

# Related commands

- `vim` / `nvim` — ویرایشگر قدرتمند
- `cat` — فقط مشاهده
- `less` — پیمایش فایل
