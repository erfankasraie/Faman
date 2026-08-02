---
title: head
aliases:
category: text
difficulty: beginner
keywords:
- beginning
- lines
- preview
---

# Introduction

`head` ابتدای فایل یا ورودی را نشان می‌دهد؛ پیش‌فرض ۱۰ خط اول.

# Syntax

```
head [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-n N` | N خط اول |
| `-c N` | N بایت اول |
| `-q` | بدون هدر نام فایل |
| `-v` | همیشه هدر |

# Examples

```bash
head file.txt
head -n 20 file.txt
head -n 5 *.log
ps aux | head
```

# Common mistakes

- `head -20` در بعضی سیستم‌ها کهنه است؛ ترجیح `-n 20`.

# Tips

- ترکیب کلاسیک: `head` + `tail` برای برش میانی.

# Related commands

- `tail`
- `less`
- `sed -n '1,20p'`
