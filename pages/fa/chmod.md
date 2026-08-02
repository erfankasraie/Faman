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

`chmod` (change mode) مجوز دسترسی فایل و پوشه را تغییر می‌دهد؛ هم با عدد (octal) و هم با نماد (`u+x`).

# Syntax

```
chmod [OPTIONS] MODE FILE...
chmod [OPTIONS] OCTAL-MODE FILE...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-R` | بازگشتی |
| `-v` | verbose |
| `-c` | فقط موارد تغییر‌یافته |
| `--reference=RFILE` | کپی مجوز از فایل دیگر |

حالت عددی: `4=r`، `2=w`، `1=x` → مثلاً `755 = rwxr-xr-x`.

حالت سمبلیک: `u/g/o/a` و `+/-/=` مثلاً `chmod u+x script.sh`.

# Examples

```bash
chmod u+x script.sh
chmod 755 script.sh
chmod go-w file.txt
chmod -R 755 project/
find . -type d -exec chmod 755 {} +
find . -type f -exec chmod 644 {} +
```

# Common mistakes

- `777` همه‌جا (ریسک امنیتی).
- فراموش کردن `-R` برای درخت پوشه.
- اشتباه `u` (مالک) با `o` (دیگران).

# Tips

- دیدن مجوز: `ls -l` یا `stat`.
- اسکریپت‌ها معمولاً `755` یا `700`.
- مجوز پیش‌فرض فایل‌های جدید: `umask`.

# Related commands

- `chown` / `chgrp`
- `ls -l` / `stat`
- `umask`
