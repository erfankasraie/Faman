---
title: nohup
aliases:
category: process
difficulty: intermediate
keywords:
- background
- hangup
- long-running
---

# Introduction

`nohup` دستور را در برابر سیگنال SIGHUP مقاوم می‌کند تا با بستن ترمینال متوقف نشود.

# Syntax

```
nohup COMMAND [ARGS...] &
```

# Options

معمولاً بدون گزینه؛ خروجی پیش‌فرض به `nohup.out` می‌رود.

# Examples

```bash
nohup ./long-job.sh &
nohup python server.py > server.log 2>&1 &
```

# Common mistakes

- فراموش کردن `&` در انتها برای پس‌زمینه.
- نداشتن مسیر لاگ و پر شدن `nohup.out`.

# Tips

- جایگزین مدرن: `systemd-run` یا `tmux`/`screen`.

# Related commands

- `jobs` / `bg` / `fg` — کنترل job شل
- `disown` — جدا کردن job از شل
- `tmux` — نشست پایدار
