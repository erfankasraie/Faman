---
title: locate
aliases:
category: filesystem
difficulty: beginner
keywords:
- search
- index
- find
---

# Introduction

`locate` نام فایل را از دیتابیس از پیش ساخته‌شده سریع پیدا می‌کند (نه پیمایش زنده مثل find).

# Syntax

```
locate [OPTIONS] PATTERN
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-i` | بدون حساسیت حروف |
| `-n N` | محدودیت تعداد |
| `-e` | فقط موجودها |

# Examples

```bash
locate passwd
locate -i readme.md
sudo updatedb    # به‌روزرسانی ایندکس
```

# Common mistakes

- انتظار فایل تازه‌ساخته بدون `updatedb`.

# Tips

- برای جستجوی لحظه‌ای: `find` یا `fd`.

# Related commands

- `find`
- `updatedb`
- `fd`
