---
title: head
aliases:
category: text
difficulty: beginner
keywords:
- beginning
- file
- lines
---

# Introduction

دستور `head` ابتدای یک فایل را نمایش می‌دهد (به صورت پیش‌فرض ۱۰ خط اول).

# Syntax

```
head [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-n N` | نمایش N خط اول |
| `-c N` | نمایش N بایت اول |
| `-q` | عدم نمایش هدر نام فایل |
| `-v` | همیشه هدر نام فایل را نشان بده |

# Examples

```bash
# ۱۰ خط اول
head file.txt

# ۲۰ خط اول
head -n 20 file.txt

# ۵۰ بایت اول
head -c 50 file.txt

# چند فایل همزمان
head -n 5 *.log
```

# Common mistakes

- استفاده از `-n` قدیمی به صورت `head -20` (هنوز کار می‌کند اما فرم استاندارد `-n 20` است).

# Tips

- ترکیب با `tail` برای برش وسط فایل مفید است.
- در پایپ‌لاین‌ها برای محدود کردن خروجی بسیار کاربردی است.

# Related commands

- `tail` — انتهای فایل
- `less` — مشاهده کامل
- `sed` — برش پیشرفته‌تر
- `cat`
