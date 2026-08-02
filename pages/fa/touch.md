---
title: touch
aliases:
category: filesystem
difficulty: beginner
keywords:
- create
- timestamp
- empty
---

# Introduction

`touch` زمان دسترسی/تغییر فایل را به‌روز می‌کند؛ اگر فایل نباشد، معمولاً یک فایل خالی می‌سازد.

# Syntax

```
touch [OPTIONS] FILE...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-a` | فقط access time |
| `-m` | فقط modification time |
| `-c` | اگر نبود نساز |
| `-t STAMP` | زمان مشخص |
| `-r FILE` | کپی زمان از فایل دیگر |

# Examples

```bash
touch newfile.txt
touch a.txt b.txt c.txt
touch -c existing.txt
touch -r ref.txt target.txt
```

# Common mistakes

- فکر کردن که touch محتوای فایل را پاک می‌کند — فقط metadata را عوض می‌کند.

# Tips

- ساخت سریع چند فایل خالی برای تست.

# Related commands

- `stat` — جزئیات زمان فایل
- `date` — زمان سیستم
