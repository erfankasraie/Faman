---
title: which
aliases:
category: shell
difficulty: beginner
keywords:
- path
- locate
- executable
- command
---

# Introduction

`which` مسیر کامل فایل اجرایی یک دستور را که در `PATH` قرار دارد نشان می‌دهد؛ برای فهمیدن این‌که «وقتی این دستور را می‌زنم، دقیقاً کدام فایل اجرا می‌شود» استفاده می‌شود — خصوصاً وقتی چند نسخه از یک برنامه (مثلاً چند نسخه python) روی سیستم نصب است.

# Syntax

```
which [OPTIONS] COMMAND...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-a` | نمایش تمام مسیرهای منطبق در PATH (نه فقط اولی) |

# Examples

```bash
# پیداکردن مسیر یک دستور
which python3

# پیداکردن تمام نسخه‌های نصب‌شده یک دستور در PATH
which -a python3

# بررسی این‌که آیا یک ابزار اصلاً نصب است
which docker || echo "docker نصب نیست"

# بررسی چند دستور همزمان
which git curl wget
```

# Common mistakes

- استفاده از `which` برای بررسی alias یا function شل؛ `which` فقط برنامه‌های واقعی موجود در `PATH` را می‌بیند، نه alias/function تعریف‌شده در `.bashrc`. برای آن‌ها از `type` استفاده کنید.
- تصور اینکه اگر `which` مسیری برگرداند، آن دستور حتماً همان نسخه‌ای است که انتظار دارید؛ اگر چند نسخه نصب باشد، فقط اولین مورد در ترتیب `PATH` نشان داده می‌شود (استفاده از `-a` برای دیدن همه).

# Tips

- `type COMMAND` جامع‌تر از `which` است چون هم alias، هم function، و هم فایل اجرایی را نشان می‌دهد.
- برای دیباگ مشکل «چرا نسخه اشتباه پایتون/node اجرا می‌شود»، `which -a` ترتیب دقیق نسخه‌های موجود در PATH را نشان می‌دهد.

# Related commands

- `type` — نمایش کامل‌تر (شامل alias و function)
- `whereis` — پیداکردن باینری، منبع، و صفحه man یک دستور
- `command -v` — معادل استاندارد POSIX برای which
