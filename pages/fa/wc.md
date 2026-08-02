---
title: wc
aliases:
category: text
difficulty: beginner
keywords:
- count
- lines
- words
- bytes
---

# Introduction

`wc` (word count) تعداد خط، کلمه و بایت/کاراکتر فایل یا ورودی را می‌شمارد.

# Syntax

```
wc [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-l` | فقط خطوط |
| `-w` | فقط کلمات |
| `-c` | بایت |
| `-m` | کاراکتر |
| `-L` | طول بلندترین خط |

# Examples

```bash
wc file.txt
wc -l file.txt
ls | wc -l
grep -c pattern file   # جایگزین رایج برای شمارش match
```

# Common mistakes

- اشتباه گرفتن `-c` (بایت) با `-m` (کاراکتر) در متن UTF-8.

# Tips

- شمارش فایل‌های یک پوشه: `ls | wc -l`

# Related commands

- `du` — حجم دیسک
- `grep -c` — شمارش تطبیق‌ها
