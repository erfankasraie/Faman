---
title: less
aliases:
- pager
category: text
difficulty: beginner
keywords:
- view
- page
- scroll
- file
---

# Introduction

دستور `less` یک pager است که به شما اجازه می‌دهد فایل‌های بلند را صفحه‌به‌صفحه مشاهده کنید. برخلاف `more`، می‌توانید به عقب برگردید.

# Syntax

```
less [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-N` | نمایش شماره خط |
| `-S` | خطوط بلند را نشکند (افقی اسکرول) |
| `-i` | جستجو بدون حساسیت به حروف |
| `-X` | بعد از خروج صفحه را پاک نکند |
| `+F` | حالت follow (مثل tail -f) |

# Examples

```bash
# مشاهده یک فایل
less large.log

# با شماره خط
less -N script.py

# دنبال کردن انتهای فایل (مثل tail -f)
less +F /var/log/syslog
```

# کلیدهای مهم داخل less

| کلید | عمل |
|------|-----|
| `Space` / `f` | صفحه بعد |
| `b` | صفحه قبل |
| `j` / `↓` | خط بعد |
| `k` / `↑` | خط قبل |
| `/pattern` | جستجو به جلو |
| `?pattern` | جستجو به عقب |
| `n` | نتیجه بعدی جستجو |
| `g` | اول فایل |
| `G` | آخر فایل |
| `q` | خروج |

# Common mistakes

- فراموش کردن خروج با `q`.
- استفاده از `cat` برای فایل‌های بلند به جای `less`.

# Tips

- متغیر `LESS` می‌تواند گزینه‌های پیش‌فرض را تنظیم کند.
- `less` می‌تواند خروجی دستورات را بگیرد: `dmesg | less`
- برای فایل‌های فشرده: `less file.gz` (اگر lesspipe نصب باشد)

# Related commands

- `more` — pager ساده‌تر
- `cat` — نمایش کامل
- `head` / `tail`
- `bat` — با highlighting
- `view` — vim در حالت فقط‌خواندنی
