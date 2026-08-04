---
title: renice
aliases:
category: process
difficulty: intermediate
keywords:
- process
- priority
- scheduling
---

# Introduction

`renice` اولویت (مقدار nice) یک پردازش **در حال اجرا** را تغییر می‌دهد؛ برخلاف `nice` که فقط هنگام شروع یک برنامه جدید کاربرد دارد.

# Syntax

```
renice [-n] PRIORITY -p PID...
renice [-n] PRIORITY -u USER...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-p PID` | تعیین پردازش با شماره PID |
| `-u USER` | اعمال روی تمام پردازش‌های یک کاربر |
| `-g GROUP` | اعمال روی یک گروه پردازشی |

# Examples

```bash
# کاهش اولویت پردازش با PID مشخص
renice -n 15 -p 4821

# افزایش اولویت (نیاز به sudo)
sudo renice -n -5 -p 4821

# اعمال روی تمام پردازش‌های یک کاربر
sudo renice -n 10 -u ali

# پیداکردن PID یک برنامه قبل از renice
pgrep -f myscript.py
```

# Common mistakes

- فراموش‌کردن اینکه بدون `sudo` فقط می‌توانید اولویت را کاهش دهید (عدد بزرگ‌تر بدهید)، نه افزایش.
- اشتباه بین `nice` (برای شروع برنامه جدید) و `renice` (برای برنامه‌ای که از قبل اجرا شده).

# Tips

- برای پیدا کردن PID پیش از اجرای renice از `pgrep` یا `htop` کمک بگیرید.
- تغییر اولویت یک پردازش در حال اجرا برای نجات‌دادن سیستم از یک برنامه سنگین‌ که فراموش کرده‌اید nice کنید بسیار مفید است.

# Related commands

- `nice` — اجرای برنامه جدید با اولویت مشخص
- `top` / `htop` — مشاهده و تغییر اولویت به‌صورت تعاملی
- `pgrep` — پیداکردن PID بر اساس نام
