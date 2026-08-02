---
title: tee
aliases:
category: text
difficulty: beginner
keywords:
- pipe
- copy
- log
---

# Introduction

`tee` ورودی را هم به stdout و هم به یک یا چند فایل می‌نویسد. برای لاگ گرفتن وسط pipeline مفید است.

# Syntax

```
tee [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-a` | append به‌جای overwrite |
| `-i` | نادیده گرفتن SIGINT |

# Examples

```bash
echo hello | tee out.txt
command | tee build.log
command | tee -a history.log
command | tee file1 file2
```

# Common mistakes

- فراموش کردن `-a` و پاک شدن لاگ قبلی.

# Tips

- با sudo: `echo value | sudo tee /etc/some.conf` (چون `sudo echo > file` درست کار نمی‌کند).

# Related commands

- `cat` — فقط خواندن/نوشتن ساده
- `script` — ضبط کل نشست ترمینال
