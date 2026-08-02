---
title: ps
aliases:
- processes
category: process
difficulty: beginner
keywords:
- process
- pid
- running
- task
---

# Introduction

دستور `ps` (process status) لیست فرایندهای در حال اجرا را نشان می‌دهد. برای دیدن PID، مصرف CPU/حافظه و دستور اجرا شده استفاده می‌شود.

# Syntax

```
ps [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `aux` | همه فرایندها با جزئیات (سبک BSD) |
| `-ef` | همه فرایندها (سبک System V) |
| `-u USER` | فقط فرایندهای یک کاربر |
| `-p PID` | فقط یک PID خاص |
| `-C name` | فرایندهایی با نام دستور |
| `--forest` | نمایش درختی |

# Examples

```bash
# لیست ساده فرایندهای ترمینال جاری
ps

# همه فرایندها با جزئیات
ps aux

# فقط فرایندهای کاربر فعلی
ps -u $USER

# پیدا کردن یک برنامه
ps aux | grep nginx

# نمایش درختی
ps -ef --forest
```

# Common mistakes

- اشتباه گرفتن `ps aux` با `ps -aux` (خط تیره در ترکیب BSD معمولاً لازم نیست).
- فکر کردن که `ps` به‌صورت زنده به‌روز می‌شود — برای مانیتور زنده از `top` یا `htop` استفاده کنید.

# Tips

- ستون‌های مهم: `PID`, `%CPU`, `%MEM`, `COMMAND`.
- برای کشتن فرایند بعد از پیدا کردن PID از `kill` استفاده کنید.
- `pgrep` و `pkill` برای جستجو/کشتن بر اساس نام راحت‌ترند.

# Related commands

- `top` / `htop` — مانیتور زنده
- `kill` — ارسال سیگنال به فرایند
- `pgrep` — جستجوی PID بر اساس نام
- `pstree` — درخت فرایندها
