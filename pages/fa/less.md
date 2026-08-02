---
title: less
aliases:
category: text
difficulty: beginner
keywords:
- pager
- view
- scroll
---

# Introduction

`less` صفحه‌بند (pager) برای مشاهده فایل‌های بلند است؛ برخلاف `more` می‌توانید عقب و جلو بروید.

# Syntax

```
less [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-N` | شماره خط |
| `-S` | بدون wrap |
| `-i` | جستجوی بدون حساسیت حروف |
| `+F` | شروع در حالت follow |
| `-R` | اجازه رنگ ANSI |

# Examples

```bash
less /var/log/syslog
less -N script.sh
man gcc | less
command | less
```

کلیدها: `q` خروج، `/` جستجو، `g`/`G` اول/آخر، `Space` صفحه بعد، `b` صفحه قبل.

# Common mistakes

- تلاش برای ویرایش داخل less — برای ویرایش از ویرایشگر استفاده کنید.

# Tips

- `less +F file` شبیه `tail -f` با امکان اسکرول.

# Related commands

- `more`
- `bat`
- `tail` / `head`
