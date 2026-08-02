---
title: stat
aliases:
category: filesystem
difficulty: intermediate
keywords:
- metadata
- inode
- timestamp
---

# Introduction

`stat` جزئیات کامل metadata فایل را نشان می‌دهد: اندازه، inode، مجوز، زمان‌ها و غیره.

# Syntax

```
stat [OPTIONS] FILE...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-c FORMAT` | قالب سفارشی (GNU) |
| `-f` | اطلاعات فایل‌سیستم |
| `-t` | خروجی فشرده |
| `-L` | دنبال کردن symlink |

# Examples

```bash
stat file.txt
stat -c '%n %s %y' file.txt
stat -c '%a %A %U' file.txt
```

# Common mistakes

- اشتباه گرفتن atime/mtime/ctime.

# Tips

- `%s` اندازه، `%Y` mtime یونیکس، `%a` مجوز عددی.

# Related commands

- `ls -l` — نمای خلاصه
- `touch` — تغییر زمان
