---
title: dirname
aliases:
category: shell
difficulty: beginner
keywords:
- path
- directory
---

# Introduction

`dirname` بخش پوشه یک مسیر را برمی‌گرداند.

# Syntax

```
dirname PATH
```

# Options

گزینه خاصی معمولاً لازم نیست.

# Examples

```bash
dirname /var/log/syslog    # /var/log
dirname ./script.sh        # .
```

# Common mistakes

- برای مسیر بدون `/` نتیجه `.` است.

# Tips

- ترکیب: `cd "$(dirname "$0")"` در اسکریپت‌ها.

# Related commands

- `basename`
- `realpath`
