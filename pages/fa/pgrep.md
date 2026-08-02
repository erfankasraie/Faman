---
title: pgrep
aliases:
category: process
difficulty: beginner
keywords:
- pid
- process
- search
---

# Introduction

`pgrep` PID فرایندها را بر اساس نام یا ویژگی‌ها پیدا می‌کند؛ ساده‌تر از `ps | grep`.

# Syntax

```
pgrep [OPTIONS] PATTERN
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-l` | نام فرایند هم نشان بده |
| `-a` | خط فرمان کامل |
| `-u USER` | فقط یک کاربر |
| `-f` | تطبیق روی خط فرمان کامل |
| `-c` | فقط تعداد |

# Examples

```bash
pgrep nginx
pgrep -l ssh
pgrep -u www-data
pgrep -af python
```

# Common mistakes

- بدون `-f` فقط نام executable تطبیق می‌خورد نه آرگومان‌ها.

# Tips

- ترکیب با kill: `kill $(pgrep nginx)` — با احتیاط.

# Related commands

- `pkill` — سیگنال بر اساس نام
- `pidof` — فقط نام دقیق
- `ps` — لیست کامل
