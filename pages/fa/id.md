---
title: id
aliases:
category: user
difficulty: beginner
keywords:
- uid
- gid
- groups
---

# Introduction

`id` شناسه کاربر (UID)، گروه اصلی (GID) و گروه‌های فرعی را نشان می‌دهد.

# Syntax

```
id [OPTIONS] [USER]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-u` | فقط UID |
| `-g` | فقط GID |
| `-gn` | نام گروه |
| `-n` | نام به‌جای عدد |
| `-G` | همه گروه‌ها |

# Examples

```bash
id
id -un
id www-data
id -G -n
```

# Common mistakes

- اشتباه گرفتن نام کاربری با UID عددی در مجوزها.

# Tips

- در اسکریپت: `if [ "$(id -u)" -eq 0 ]; then ...`

# Related commands

- `whoami` — فقط نام کاربر
- `groups` — لیست گروه‌ها
- `getent passwd`
