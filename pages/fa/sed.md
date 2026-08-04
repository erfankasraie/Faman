---
title: sed
aliases:
category: text
difficulty: intermediate
keywords:
- stream
- edit
- replace
- regex
---

# Introduction

`sed` (*stream editor*) جریان متن را خط‌به‌خط می‌خواند، دستورات ویرایش را اعمال می‌کند و نتیجه را می‌نویسد — بدون باز کردن ادیتور تعاملی. رایج‌ترین کار: **جایگزینی** با `s///`.

# Syntax

```
sed [OPTIONS] 'script' [FILE...]
sed [OPTIONS] -e 'script' -e 'script' [FILE...]
sed [OPTIONS] -f script.sed [FILE...]
```

بدون `-i` فایل اصلی عوض نمی‌شود؛ خروجی روی stdout است.

# Options

| گزینه | توضیح |
|-------|--------|
| `-n` | چاپ خودکار هر خط را خاموش کن (فقط با `p` چاپ کن) |
| `-i` / `-i.bak` | ویرایش درجا (GNU)؛ اختیاری بکاپ) |
| `-e SCRIPT` | افزودن دستور |
| `-f FILE` | اسکریپت از فایل |
| `-r` / `-E` | regex توسعه‌یافته (GNU/BSD) |
| `-s` | هر فایل را جداگانه آدرس‌دهی کن |
| `-u` | بافر کمتر (خروجی فوری‌تر) |

## دستورات اصلی اسکریپت

| دستور | معنی |
|--------|------|
| `s/PATTERN/REPL/flags` | جایگزینی |
| `d` | حذف خط |
| `p` | چاپ خط |
| `a\text` | الحاق بعد از خط |
| `i\text` | درج قبل از خط |
| `c\text` | تعویض کل خط |
| `q` | خروج |
| `y/src/dst/` | ترجمه کاراکتر به‌کاراکتر |
| `=` | چاپ شماره خط |
| `w FILE` | نوشتن خطوط به فایل |

## آدرس‌ها

| آدرس | معنی |
|--------|------|
| `5` | خط ۵ |
| `5,10` | خطوط ۵ تا ۱۰ |
| `5,$` | از ۵ تا آخر |
| `/PAT/` | خطوط match |
| `/A/,/B/` | از match A تا B |
| `5!` | همه جز خط ۵ |

## فلگ‌های `s///`

| فلگ | معنی |
|------|------|
| `g` | همهٔ رخدادها در خط |
| `N` | فقط N-امین رخداد |
| `p` | بعد از جایگزینی موفق چاپ کن (با `-n`) |
| `i` / `I` | بدون حساسیت حروف (GNU) |
| `w FILE` | نوشتن خطوط جایگزین‌شده |

جداکننده می‌تواند عوض شود: `s#/usr/old#/usr/new#g` وقتی `/` زیاد است.

# Examples

## جایگزینی

```bash
# اول فقط stdout
sed 's/foo/bar/' file.txt
sed 's/foo/bar/g' file.txt
sed 's/[0-9]\+/N/g' file.txt
sed -E 's/[0-9]+/N/g' file.txt

# درجا با بکاپ
sed -i.bak 's/localhost/127.0.0.1/g' config.env
```

## حذف و چاپ انتخابی

```bash
# حذف خطوط خالی
sed '/^$/d' file.txt

# حذف خطوط comment
sed '/^#/d' config

# فقط خطوط دارای error
sed -n '/error/p' app.log

# خطوط ۵ تا ۱۰
sed -n '5,10p' file.txt

# از خط matching تا آخر
sed -n '/^## Install/,/^## /p' README.md
```

## چند دستور و اسکریپت

```bash
sed -e 's/\r$//' -e '/^$/d' windows.txt

cat > clean.sed <<'EOF'
s/[[:space:]]\+$//
/^#/d
EOF
sed -f clean.sed input.txt
```

## گروه‌گیری و ارجاع

```bash
# جابه‌جایی دو فیلد جدا با فاصله
sed -E 's/^(\S+)\s+(\S+)/\2 \1/' names.txt

# افزودن پسوند
sed 's/$/.bak/' list.txt
```

## pipeline

```bash
curl -sL https://example.com | sed -n 's/.*<title>\(.*\)<\/title>.*/\1/p'
ps aux | sed -n '1p;/nginx/p'
```

# Common mistakes

- `-i` بدون بکاپ روی فایل مهم.
- فراموش کردن `g` → فقط اولین رخداد هر خط.
- regex پیچیده در basic mode بدون escape کردن `()`.
- انتظار که sed کل فایل را مثل ادیتور چندخطی پیشرفته بفهمد (برای چندخطی سنگین گاهی `awk`/`perl` بهتر است).

# Tips

- قبل از `-i` یک‌بار بدون آن تست کنید.
- برای جایگزینی ثابت بدون regex: هنوز `s///` است ولی متاها را escape کنید یا از ابزار دیگر استفاده کنید.
- `sed` برای streamهای بزرگ حافظهٔ کمی مصرف می‌کند.

# Related commands

- `awk` — فیلد و منطق قوی‌تر
- `grep` — فقط فیلتر/جستجو
- `tr` — ترجمه/حذف کاراکتر ساده
- `perl -pe` — regex قدرتمندتر
- `cut` — برش فیلد ساده
