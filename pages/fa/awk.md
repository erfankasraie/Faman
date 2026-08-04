---
title: awk
aliases:
- gawk
category: text
difficulty: intermediate
keywords:
- field
- report
- text
- programming
---

# Introduction

`awk` زبان کوچک پردازش متن **خط‌به‌خط و فیلدبه‌فیلد** است. برای گزارش، جمع ستون‌ها، فیلتر شرطی و تبدیل جدول‌های متنی عالی است. روی لینوکس معمولاً GNU Awk (`gawk`) است.

# Syntax

```
awk [OPTIONS] 'program' [FILE...]
awk [OPTIONS] -f program.awk [FILE...]
```

ساختار برنامه:

```
pattern { actions }
```

الگوهای خاص: `BEGIN` (قبل از ورودی)، `END` (بعد از همهٔ خطوط).

# Options

| گزینه | توضیح |
|-------|--------|
| `-F FS` | جداکنندهٔ فیلد ورودی (`-F:` یا `-F'\t'`) |
| `-v VAR=VAL` | متغیر از بیرون |
| `-f FILE` | برنامه از فایل |
| `-o` / `--optimize` | بهینه‌سازی (gawk) |

## متغیرهای داخلی مهم

| متغیر | معنی |
|--------|------|
| `$0` | کل خط |
| `$1` … `$NF` | فیلد ۱ … آخر |
| `NF` | تعداد فیلدها |
| `NR` | شماره رکورد از ابتدای اجرا |
| `FNR` | شماره رکورد در فایل فعلی |
| `FS` | جداکنندهٔ ورودی (پیش‌فرض فاصله) |
| `OFS` | جداکنندهٔ خروجی |
| `RS` / `ORS` | جداکنندهٔ رکورد ورودی/خروجی |
| `FILENAME` | نام فایل فعلی |

## الگوها و کنترل

| الگو / دستور | معنی |
|---------------|------|
| `/regex/` | خط match |
| `NR==1` | شرط روی شماره خط |
| `$3 > 100` | شرط روی فیلد |
| `next` | برو خط بعد |
| `exit` | خروج |
| `print` / `printf` | خروجی |
| `if` / `for` / `while` | کنترل جریان |
| `split(s,a,fs)` | شکستن رشته |
| `substr` / `index` / `length` | رشته |
| `gsub` / `sub` | جایگزینی |

# Examples

## فیلد و جدول

```bash
# ستون ۱ و ۳
awk '{print $1, $3}' data.txt

# CSV ساده
awk -F, '{print $2}' data.csv

# با هدر خروجی
awk 'BEGIN{OFS="\t"} {print $1,$3}' data.txt
```

## فیلتر و گزارش

```bash
# خطوط با بیش از ۵ فیلد
awk 'NF>5' data.txt

# جمع ستون ۲
awk '{s+=$2} END{print s}' nums.txt

# میانگین
awk '{s+=$1; n++} END{if(n) print s/n}' nums.txt

# فقط وقتی ستون ۳ از ۱۰۰ بزرگ‌تر است
awk '$3 > 100 {print $1, $3}' data.txt
```

## لاگ و سیستم

```bash
# IPهای unique از access.log (فرض: فیلد ۱ = IP)
awk '{print $1}' access.log | sort -u

# شمارش کدهای HTTP (فرض فیلد مناسب)
awk '{codes[$9]++} END{for (c in codes) print c, codes[c]}' access.log

# کاربران از /etc/passwd
awk -F: '{print $1, $7}' /etc/passwd
```

## BEGIN/END و اسکریپت

```bash
awk -F: 'BEGIN{print "user\tshell"} {print $1"\t"$7} END{print "done"}' /etc/passwd

cat > sum.awk <<'EOF'
{ s += $1 }
END { print "sum=", s }
EOF
awk -f sum.awk nums.txt
```

## از بیرون متغیر بده

```bash
awk -v min=100 '$3 >= min {print}' data.txt
```

# Common mistakes

- `print $1 $2` بدون ویرگول → چسبیدن فیلدها؛ برای OFS از ویرگول استفاده کنید: `print $1, $2`.
- CSV با نقل‌قول و کاما داخل فیلد → awk ساده کافی نیست؛ `csvkit` / `python`.
- تک‌نقل‌قول شل دور برنامه را فراموش کردن و شل `$1` را بلعیدن.

# Tips

- برای یک جایگزینی ساده گاهی `sed` خواناتر است؛ برای حساب روی ستون‌ها awk بهتر است.
- `column -t` برای قشنگ‌کردن خروجی جدولی.
- GNU awk مستندات بسیار کاملی دارد (`info awk`).

# Related commands

- `sed` — ویرایش جریانی
- `cut` — برش فیلد ساده
- `sort` / `uniq` — مرتب‌سازی و یکتا
- `grep` — فیلتر الگو
- `jq` — JSON
