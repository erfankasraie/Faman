---
title: ncdu
aliases:
category: filesystem
difficulty: intermediate
keywords:
- disk
- usage
- du
- interactive
---

# Introduction

`ncdu` (NCurses Disk Usage) نسخهٔ تعاملی تحلیل فضای دیسک است؛ راحت‌تر از `du` خام برای پیدا کردن پوشه‌های حجیم.

# Syntax

```
ncdu [OPTIONS] [PATH]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-x` | یک filesystem |
| `-r` | فقط خواندن (بدون delete) |
| `-o FILE` | ذخیره اسکن |
| `-f FILE` | بارگذاری اسکن قبلی |
| `-e` | extended info |

کلیدها: ↑↓ حرکت · Enter ورود · d حذف (اگر مجاز) · q خروج.

# Examples

```bash
ncdu /
ncdu -x /home
ncdu -o scan ~/data && ncdu -f scan
```

# Common mistakes

- اسکن `/` بدون دسترسی → خطاهای permission زیاد (طبیعی).
- حذف تصادفی از داخل UI.

# Tips

- روی سرور: اول `ncdu -x -r`.

# Related commands

- `du` · `df` · `ls`
