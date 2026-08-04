---
title: mktemp
aliases:
category: filesystem
difficulty: intermediate
keywords:
- temporary
- file
- script
---

# Introduction

`mktemp` یک فایل یا پوشه موقت با نام تصادفی و تضمین‌شده منحصربه‌فرد می‌سازد؛ گزینه‌ای امن‌تر از انتخاب دستی نام فایل موقت در اسکریپت‌ها (که ممکن است با فایل دیگری تداخل کند).

# Syntax

```
mktemp [OPTIONS] [TEMPLATE]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-d` | ساخت پوشه موقت به‌جای فایل |
| `-p DIR` | مسیر پایه به‌جای `/tmp` پیش‌فرض |
| `--suffix=STR` | افزودن پسوند به انتهای نام |
| `-u` | فقط نام تولید کند بدون ساخت واقعی فایل (dry-run) |

# Examples

```bash
# ساخت یک فایل موقت و نمایش مسیر آن
tmpfile=$(mktemp)
echo "فایل موقت: $tmpfile"

# ساخت یک پوشه موقت
tmpdir=$(mktemp -d)

# ساخت فایل موقت با پسوند مشخص
mktemp --suffix=.log

# استفاده در یک اسکریپت با پاک‌سازی خودکار در پایان
tmpfile=$(mktemp)
trap 'rm -f "$tmpfile"' EXIT
echo "داده" > "$tmpfile"
```

# Common mistakes

- ساختن نام فایل موقت به‌صورت دستی (مثل `/tmp/myfile$$`) که ریسک امنیتی و تداخل نام دارد؛ همیشه از `mktemp` استفاده کنید.
- فراموش‌کردن پاک‌کردن فایل موقت در پایان اسکریپت (استفاده از `trap ... EXIT` بهترین راه است).

# Tips

- ترکیب `mktemp` با `trap 'rm -f "$tmpfile"' EXIT` تضمین می‌کند فایل موقت حتی در صورت خطای اسکریپت هم پاک شود.
- برای ساخت چند فایل موقت با الگو: `mktemp /tmp/myapp.XXXXXX`

# Related commands

- `trap` — اجرای دستور هنگام خروج یا سیگنال (برای پاک‌سازی)
- `rm` — حذف فایل
- `tmpfs` / `/tmp` — محل معمول فایل‌های موقت در لینوکس
