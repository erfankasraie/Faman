---
title: tree
aliases:
category: filesystem
difficulty: beginner
keywords:
- directory
- list
- hierarchy
---

# Introduction

`tree` ساختار درختی پوشه‌ها و فایل‌ها را نمایش می‌دهد. در بعضی توزیع‌ها باید جدا نصب شود (`apt install tree`).

# Syntax

```
tree [OPTIONS] [DIRECTORY]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-L N` | حداکثر عمق |
| `-d` | فقط پوشه‌ها |
| `-a` | فایل‌های مخفی |
| `-h` | اندازه خوانا |
| `-C` | رنگ |
| `-I pattern` | نادیده گرفتن |
| `-f` | مسیر کامل |

# Examples

```bash
tree -L 2
tree -d -L 3 /etc
tree -I 'node_modules|.git' -L 3
```

# Common mistakes

- روی درخت خیلی بزرگ بدون `-L` خروجی غیرقابل‌استفاده.

# Tips

- جایگزین ساده: `find . -maxdepth 2`.

# Related commands

- `ls` · `find` · `du`
