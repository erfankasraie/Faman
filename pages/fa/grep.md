---
title: grep
aliases:
- egrep
- fgrep
category: text
difficulty: intermediate
keywords:
- search
- regex
- text
- pattern
---

# Introduction

`grep` خطوطی را چاپ می‌کند که با یک **الگو** (متن ثابت یا regex) مطابقت دارند. نام از *Global Regular Expression Print* می‌آید. نسخه‌های رایج: GNU grep روی لینوکس.

`egrep` ≈ `grep -E`، `fgrep` ≈ `grep -F` (امروز بهتر است همان فلگ‌ها را بنویسید).

# Syntax

```
grep [OPTIONS] PATTERN [FILE...]
grep [OPTIONS] -e PATTERN ... [FILE...]
grep [OPTIONS] -f PATTERN_FILE [FILE...]
```

اگر FILE نباشد، از stdin می‌خواند. کد خروج: `0` match، `1` بدون match، `≥2` خطا.

# Options

## مطابقت الگو

| گزینه | توضیح |
|-------|--------|
| `-F` | الگو را **ثابت** بگیر (بدون regex) — سریع و امن |
| `-E` | regex توسعه‌یافته (`+ ? | ()`) |
| `-G` | regex پایه (پیش‌فرض) |
| `-P` | Perl-compatible regex (اگر پشتیبانی شود) |
| `-e PATTERN` | چند الگو؛ یا وقتی الگو با `-` شروع می‌شود |
| `-f FILE` | الگوها از فایل (هر خط یکی) |
| `-i` | بدون حساسیت حروف |
| `-w` | فقط کلمهٔ کامل |
| `-x` | کل خط باید match شود |
| `-v` | خطوط **غیر**مطابق |
| `-y` | مترادف قدیمی `-i` |

## خروجی

| گزینه | توضیح |
|-------|--------|
| `-n` | شماره خط |
| `-H` / `-h` | با/بدون نام فایل |
| `-l` | فقط نام فایل‌های دارای match |
| `-L` | فایل‌های بدون match |
| `-c` | تعداد خطوط match در هر فایل |
| `-o` | فقط قسمت match‌شده، نه کل خط |
| `-b` | آفست بایت |
| `--color=auto` | رنگ‌آمیزی |
| `-m N` | بعد از N match در هر فایل متوقف شو |
| `-q` / `--silent` | خروجی نه؛ فقط کد خروج |
| `-s` | پیام «فایل نیست» را سرکوب کن |

## context

| گزینه | توضیح |
|-------|--------|
| `-A N` | N خط بعد |
| `-B N` | N خط قبل |
| `-C N` | N خط قبل و بعد |

## بازگشتی و فیلتر فایل

| گزینه | توضیح |
|-------|--------|
| `-r` | بازگشتی روی پوشه‌ها |
| `-R` | مثل `-r` + دنبال symlink |
| `--include=GLOB` | فقط این نام‌ها |
| `--exclude=GLOB` | این نام‌ها نه |
| `--exclude-dir=DIR` | این پوشه‌ها نه |
| `-a` | باینری را مثل متن بخوان |
| `-I` | باینری را نادیده بگیر |
| `-z` | خط با `\0` تمام می‌شود (با `-print0`) |

# Examples

## پایه

```bash
grep 'error' app.log
grep -n 'error' app.log
grep -i 'error' app.log
grep -w 'Error' app.log
grep -v 'DEBUG' app.log
grep -c 'error' app.log
```

## چند فایل و بازگشتی

```bash
grep -n 'TODO' src/*.py
grep -r 'TODO' src/
grep -rl 'password' /etc/
grep -r --include='*.go' --exclude-dir=vendor 'FIXME' .
grep -r --exclude-dir=.git --exclude-dir=node_modules 'API_KEY' .
```

## regex

```bash
# توسعه‌یافته: error یا warning
grep -E 'error|warning|fail' app.log

# IP ساده
grep -Eo '([0-9]{1,3}\.){3}[0-9]{1,3}' access.log

# الگو با خط تیره
grep -e '--debug' *

# چند الگو
grep -E -e 'error' -e 'critical' app.log
grep -f patterns.txt app.log
```

## context و اسکریپت

```bash
grep -C 3 'Exception' app.log
grep -A 5 'panic' app.log

# فقط بگو هست یا نه
grep -q 'ready' /tmp/status && echo ok

# با find برای کنترل دقیق‌تر فایل‌ها
find . -name '*.log' -print0 | xargs -0 grep -Hn 'error'
```

## stdin و pipeline

```bash
ps aux | grep '[n]ginx'    # ترفند: خود grep در خروجی نیاید
journalctl -u nginx | grep -i error
curl -s https://api.example.com/data | grep -o '"id":[0-9]*'
```

# Common mistakes

- الگوی دارای `|` یا `*` بدون کوتیشن یا بدون `-E`.
- `grep -r /` بدون `--exclude-dir` → کند و پرحجم.
- استفاده از `grep` برای JSON ساخت‌یافته به‌جای `jq`.
- انتظار match روی باینری بدون `-a`.
- `grep pattern *` وقتی هیچ فایلی match glob نشود و شل خطا بدهد (`nullglob`).

# Tips

- متن ثابت و دارای کاراکتر خاص regex: `-F`.
- در اسکریپت: `-q` برای if، `-s` برای فایل‌های ممکن‌ناموجود.
- `git grep` برای جستجو داخل ریپو سریع‌تر و آگاه به gitignore است.
- alias رایج: `alias grep='grep --color=auto'`.

# Related commands

- `rg` (ripgrep) — سریع‌تر، پیش‌فرض خوب
- `ag` / `ack` — جستجوی کد
- `find` + `xargs` — انتخاب فایل سپس grep
- `awk` / `sed` — پردازش خطی پیشرفته‌تر
- `jq` — JSON
