---
title: env
aliases:
category: shell
difficulty: beginner
keywords:
- environment
- variables
- path
---

# Introduction

`env` متغیرهای محیطی را نمایش می‌دهد یا یک دستور را با محیط تغییر‌یافته اجرا می‌کند.

# Syntax

```
env [OPTIONS] [NAME=VALUE]... [COMMAND]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-i` | محیط خالی |
| `-u NAME` | حذف یک متغیر |

# Examples

```bash
env
env | grep PATH
env LANG=C ls
env -i HOME=$HOME bash --noprofile --norc
```

# Common mistakes

- اشتباه گرفتن `export` (پایدار در شل فعلی) با `env` برای یک اجرا.

# Tips

- شِبانگ قابل حمل: `#!/usr/bin/env python3`

# Related commands

- `export` — تنظیم در شل
- `printenv` — چاپ یک متغیر
