---
title: grep
aliases:
- search-text
category: text
difficulty: intermediate
keywords:
- search
- text
- pattern
- regex
---

# Introduction

دستور `grep` برای جستجوی متن داخل فایل‌ها استفاده می‌شود. نام آن از عبارت Global Regular Expression Print گرفته شده است.

# Syntax

```
grep [OPTIONS] PATTERN [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-i` | بدون حساسیت به حروف بزرگ/کوچک |
| `-r` یا `-R` | جستجوی بازگشتی در پوشه‌ها |
| `-n` | نمایش شماره خط |
| `-v` | خطوطی که مطابقت ندارند |
| `-l` | فقط نام فایل‌های دارای مطابقت |
| `-c` | تعداد مطابقت‌ها |
| `-A N` | N خط بعد از مطابقت |
| `-B N` | N خط قبل از مطابقت |
| `-C N` | N خط قبل و بعد |
| `-E` | استفاده از regex گسترش‌یافته |
| `-w` | فقط کلمات کامل |
| `--color=auto` | رنگ‌آمیزی نتایج |

# Examples

```bash
# جستجوی ساده
grep "error" logfile.txt

# بدون حساسیت به حروف + شماره خط
grep -in "error" logfile.txt

# جستجوی بازگشتی
grep -r "TODO" src/

# فقط نام فایل‌ها
grep -rl "password" /etc/

# با context
grep -C 3 "exception" app.log

# استفاده از regex
grep -E "error|warning|fail" app.log
```

# Common mistakes

- فراموش کردن کوتیشن دور patternهای پیچیده.
- استفاده از `grep -r` روی پوشه‌های خیلی بزرگ بدون محدودیت.
- انتظار داشتن regex پیشرفته بدون `-E` یا `-P`.

# Tips

- برای جستجوی پیشرفته‌تر از `rg` (ripgrep) استفاده کنید — بسیار سریع‌تر است.
- می‌توانید خروجی را به `less` بدهید: `grep ... | less`
- در ترکیب با `find`: `find . -name "*.py" | xargs grep -n "def "`
- از `--exclude-dir` برای نادیده گرفتن پوشه‌هایی مثل `.git` و `node_modules` استفاده کنید.

# Related commands

- `rg` (ripgrep) — جایگزین سریع مدرن
- `ag` (the silver searcher)
- `ack`
- `find` — جستجوی فایل
- `sed` / `awk` — پردازش متن
