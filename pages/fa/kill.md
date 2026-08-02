---
title: kill
aliases:
- terminate
category: process
difficulty: intermediate
keywords:
- signal
- pid
- stop
- terminate
---

# Introduction

`kill` سیگنالی به یک فرایند می‌فرستد. برخلاف نامش، لزوماً «نمی‌کشد»؛ سیگنال پیش‌فرض `TERM` است که از فرایند می‌خواهد مؤدبانه تمام شود.

# Syntax

```
kill [OPTIONS] PID...
kill -SIGNAL PID...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-l` | لیست سیگنال‌ها |
| `-s SIGNAL` / `-SIGNAL` | نوع سیگنال |
| `-9` | SIGKILL (اجباری، غیرقابل گرفتن) |
| `-15` | SIGTERM (پیش‌فرض) |
| `-HUP` | بارگذاری مجدد تنظیمات (رایج برای سرویس‌ها) |

# Examples

```bash
# پایان مؤدبانه
kill 1234

# اجباری
kill -9 1234

# چند فرایند
kill -TERM 100 101 102

# پیدا کردن و کشتن بر اساس نام
pkill nginx
killall firefox
```

# Common mistakes

- استفاده فوری از `kill -9` بدون امتحان `TERM`.
- کشتن PID اشتباه (همیشه با `ps` یا `pgrep` تأیید کنید).

# Tips

- ترتیب پیشنهادی: `TERM` → صبر → `KILL`.
- `kill -0 PID` فقط وجود فرایند را چک می‌کند بدون ارسال سیگنال واقعی.

# Related commands

- `pkill` / `killall` — بر اساس نام
- `ps` — پیدا کردن PID
- `top` — مانیتور و کشتن تعاملی
