---
title: jobs
aliases:
category: process
difficulty: beginner
keywords:
- background
- shell
- job-control
---

# Introduction

`jobs` وضعیت jobهای پس‌زمینه/معلق شل فعلی را نشان می‌دهد. یک builtin است.

# Syntax

```
jobs [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-l` | با PID |
| `-p` | فقط PID |
| `-r` | فقط running |
| `-s` | فقط stopped |

# Examples

```bash
sleep 100 &
jobs
jobs -l
fg %1
bg %1
```

کلیدهای مرتبط:

- `Ctrl+Z` — معلق کردن
- `bg` — ادامه در پس‌زمینه
- `fg` — آوردن به پیش‌زمینه

# Common mistakes

- انتظار دیدن فرایندهای سیستم — فقط jobهای همین شل.

# Tips

- `%1` یعنی job شماره ۱.

# Related commands

- `bg` / `fg`
- `nohup` / `disown`
- `ps`
