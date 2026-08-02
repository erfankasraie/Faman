---
title: realpath
aliases:
category: shell
difficulty: beginner
keywords:
- path
- absolute
- resolve
---

# Introduction

`realpath` مسیر مطلق و resolve‌شده (بدون `.` و `..` و معمولاً با دنبال کردن symlink) را چاپ می‌کند.

# Syntax

```
realpath [OPTIONS] PATH...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-s` | بدون resolve symlink |
| `-e` | همه اجزا باید وجود داشته باشند |
| `-m` | اجزای ناموجود مجاز |

# Examples

```bash
realpath .
realpath ../foo/bar
realpath -s some-link
```

# Common mistakes

- روی سیستم‌های خیلی قدیمی ممکن است نصب نباشد (`readlink -f` جایگزین).

# Tips

- در اسکریپت برای مسیر پایدار از realpath استفاده کنید.

# Related commands

- `readlink`
- `dirname` / `basename`
