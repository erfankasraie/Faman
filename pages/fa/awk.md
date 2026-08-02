---
title: awk
aliases:
category: text
difficulty: intermediate
keywords:
- columns
- field
- report
- text
---

# Introduction

`awk` زبان کوچک پردازش متن است؛ برای کار با ستون‌ها، فیلتر کردن و ساخت گزارش از فایل‌های متنی عالی است.

# Syntax

```
awk [OPTIONS] 'PROGRAM' [FILE...]
```

برنامه معمولاً به شکل `pattern { action }` نوشته می‌شود.

# Options

| گزینه | توضیح |
|-------|--------|
| `-F FS` | جداکننده فیلد (مثلاً `:` یا `,`) |
| `-v VAR=val` | تعریف متغیر |
| `-f file` | خواندن برنامه از فایل |

متغیرهای مهم: `$1`, `$2`, … فیلدها؛ `$0` کل خط؛ `NF` تعداد فیلد؛ `NR` شماره خط.

# Examples

```bash
# ستون اول
awk '{print $1}' file.txt

# کاربران و شل از /etc/passwd
awk -F: '{print $1, $7}' /etc/passwd

# خطوط با بیش از ۳ فیلد
awk 'NF > 3' data.txt

# جمع ستون ۳
awk '{sum += $3} END {print sum}' nums.txt

# CSV ساده
awk -F, '{print $2}' data.csv
```

# Common mistakes

- فراموش کردن `-F` وقتی جداکننده فاصله نیست.
- نوشتن منطق خیلی پیچیده در یک خط — بهتر است در فایل اسکریپت باشد.

# Tips

- `BEGIN { }` قبل از خواندن ورودی، `END { }` بعد از آن اجرا می‌شود.
- برای JSON از ابزارهای اختصاصی (`jq`) استفاده کنید.

# Related commands

- `sed` — جایگزینی و ویرایش جریانی
- `cut` — بریدن ستون‌های ساده
- `sort` / `uniq` — مرتب‌سازی و یکتاسازی
