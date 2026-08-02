---
title: pkill
aliases:
category: process
difficulty: intermediate
keywords:
- kill
- signal
- process
---

# Introduction

`pkill` بر اساس نام الگو به فرایندها سیگنال می‌فرستد؛ بدون نیاز به دانستن PID.

# Syntax

```
pkill [OPTIONS] PATTERN
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-SIGNAL` | نوع سیگنال (پیش‌فرض TERM) |
| `-u USER` | فقط یک کاربر |
| `-f` | تطبیق خط فرمان |
| `-x` | تطبیق دقیق نام |
| `-l` | لیست سیگنال‌ها |

# Examples

```bash
pkill nginx
pkill -9 firefox
pkill -u guest
pkill -f 'python app.py'
```

# Common mistakes

- الگوی خیلی کلی که فرایندهای ناخواسته را می‌کشد.
- `kill -9` فوری بدون امتحان TERM.

# Tips

- اول با `pgrep -a pattern` بررسی کنید.

# Related commands

- `pgrep` — فقط پیدا کردن
- `kill` — با PID
- `killall` — بر اساس نام دقیق‌تر در بعضی سیستم‌ها
