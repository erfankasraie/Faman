---
title: cat
aliases:
- concatenate
category: text
difficulty: beginner
keywords:
- view
- file
- content
- concatenate
---

# Introduction

دستور `cat` (concatenate) محتوای فایل‌ها را نمایش می‌دهد یا چند فایل را به هم می‌چسباند.

برای فایل‌های کوتاه مناسب است. برای فایل‌های بلند از `less` یا `more` استفاده کنید.

# Syntax

```
cat [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-n` | شماره‌گذاری خطوط |
| `-b` | شماره‌گذاری فقط خطوط غیرخالی |
| `-s` | فشرده کردن خطوط خالی متوالی |
| `-A` | نمایش کاراکترهای خاص |
| `-E` | نمایش `$` در انتهای خطوط |

# Examples

```bash
# نمایش محتوای یک فایل
cat file.txt

# چسباندن چند فایل
cat part1.txt part2.txt > full.txt

# شماره‌گذاری خطوط
cat -n script.sh

# ساخت فایل جدید (با Ctrl+D تمام کنید)
cat > newfile.txt
```

# Common mistakes

- استفاده از `cat` برای فایل‌های خیلی بزرگ → ترمینال پر می‌شود.
- `cat file | less` به جای `less file` (غیرضروری است).

# Tips

- برای فایل‌های کد از `bat` (جایگزین مدرن با syntax highlighting) استفاده کنید.
- `cat` می‌تواند از stdin بخواند.
- برای اضافه کردن به انتهای فایل: `cat >> file.txt`

# Related commands

- `less` / `more` — مشاهده صفحه‌به‌صفحه
- `head` / `tail` — ابتدا و انتهای فایل
- `bat` — جایگزین مدرن
- `tac` — نمایش معکوس
