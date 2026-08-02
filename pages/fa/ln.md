---
title: ln
aliases:
category: filesystem
difficulty: intermediate
keywords:
- link
- symlink
- hardlink
---

# Introduction

`ln` لینک سخت یا نمادین (symlink) می‌سازد. symlink رایج‌ترین شکل برای ارجاع به مسیر دیگر است.

# Syntax

```
ln [OPTIONS] TARGET LINK_NAME
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-s` | symlink |
| `-f` | بازنویسی لینک موجود |
| `-v` | حالت verbose |
| `-r` | مسیر نسبی برای symlink |

# Examples

```bash
# symlink
ln -s /usr/local/bin/app app

# hard link
ln original.txt hardlink.txt

# بازنویسی
ln -sf /new/target current
```

# Common mistakes

- جابه‌جا نوشتن TARGET و LINK_NAME.
- ساخت symlink با مسیر نسبی اشتباه وقتی بعداً از پوشه دیگر استفاده می‌شود.

# Tips

- دیدن مقصد: `readlink -f path` یا `ls -l`

# Related commands

- `readlink` — مقصد symlink
- `cp -s` — کپی به‌صورت symlink (در برخی سیستم‌ها)
