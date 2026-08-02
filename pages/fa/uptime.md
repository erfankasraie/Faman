---
title: uptime
aliases:
category: system
difficulty: beginner
keywords:
- load
- runtime
- users
---

# Introduction

`uptime` مدت روشن بودن سیستم، تعداد کاربران و میانگین بار (load average) را نشان می‌دهد.

# Syntax

```
uptime [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-p` | فقط مدت uptime خوانا |
| `-s` | زمان boot |

# Examples

```bash
uptime
uptime -p
uptime -s
```

# Common mistakes

- تفسیر load average بدون در نظر گرفتن تعداد هسته‌ها.

# Tips

- load 1.0 روی سیستم ۴ هسته‌ای یعنی نسبتاً آرام است.

# Related commands

- `top` / `htop`
- `w` — کاربران و load
