---
title: apk
aliases:
- apk-tools
category: package
difficulty: beginner
keywords:
- alpine
- package
- docker
- install
---

# Introduction

`apk` مدیر بسته **Alpine Linux** است؛ سبک و رایج داخل کانتینرهای Docker.

# Syntax

```
apk [OPTIONS] COMMAND
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `update` | به‌روزرسانی ایندکس |
| `add PKG` | نصب |
| `del PKG` | حذف |
| `search KEY` | جستجو |
| `info` | اطلاعات |
| `upgrade` | ارتقا |

# Examples

```bash
apk update
apk add curl git
apk search python
apk del curl
apk info -v | head

# در Dockerfile اغلب:
# RUN apk add --no-cache curl
```

# Common mistakes

- فراموش کردن `update` قبل از `add` در ایمیج‌های قدیمی.
- قاطی کردن با Android `apk` (فایل اپ اندروید).

# Tips

- `--no-cache` در Docker تا ایندکس در لایه نماند.
- بستههای `-dev` برای کامپایل.

# Related commands

- `apt` / `dnf` در توزیع‌های دیگر
- `docker`
