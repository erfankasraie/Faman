---
title: cut
aliases:
category: text
difficulty: beginner
keywords:
- columns
- fields
- extract
---

# Introduction

`cut` بخش‌هایی از هر خط را بر اساس بایت، کاراکتر یا فیلد جدا می‌کند. برای استخراج ستون ساده سریع است.

# Syntax

```
cut [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-d C` | جداکننده |
| `-f N` | شماره فیلد(ها) |
| `-c N` | بازه کاراکتر |
| `-b N` | بازه بایت |
| `--complement` | مکمل انتخاب |

# Examples

```bash
# فیلد اول با جداکننده :
cut -d: -f1 /etc/passwd

# فیلدهای ۱ و ۷
cut -d: -f1,7 /etc/passwd

# کاراکترهای ۱ تا ۵
cut -c1-5 file.txt
```

# Common mistakes

- برای فیلدهای با فاصلهٔ چندتایی، `awk` معمولاً بهتر از `cut` است.

# Tips

- ترکیب: `grep pattern file | cut -d, -f2`

# Related commands

- `awk` — انعطاف بیشتر
- `paste` — چسباندن ستون‌ها
