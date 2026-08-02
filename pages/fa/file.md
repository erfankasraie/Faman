---
title: file
aliases:
category: filesystem
difficulty: beginner
keywords:
- type
- mime
- detect
---

# Introduction

`file` نوع فایل را از روی محتوا (magic) تشخیص می‌دهد؛ نه فقط از پسوند نام.

# Syntax

```
file [OPTIONS] FILE...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-b` | بدون نام فایل در خروجی |
| `-i` | MIME type |
| `-s` | فایل‌های خاص (مثل device) |
| `-z` | داخل فشرده |

# Examples

```bash
file photo.jpg
file -i document
file /bin/ls
file *
```

# Common mistakes

- اعتماد کامل به پسوند `.txt` در حالی که محتوا چیز دیگری است.

# Tips

- در اسکریپت‌ها: `file -b --mime-type file`

# Related commands

- `stat` — metadata
- `ls -l` — مجوز و اندازه
