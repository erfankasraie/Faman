---
title: sed
aliases:
- stream-editor
category: text
difficulty: intermediate
keywords:
- edit
- replace
- stream
- regex
---

# Introduction

`sed` (stream editor) متن را خط‌به‌خط پردازش می‌کند؛ بدون باز کردن فایل در ویرایشگر. بیشتر برای جایگزینی، حذف و تبدیل خودکار استفاده می‌شود.

# Syntax

```
sed [OPTIONS] 'SCRIPT' [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-i` | ویرایش درجا روی فایل |
| `-i.bak` | ویرایش درجا + پشتیبان |
| `-n` | فقط خطوط چاپ‌شده با `p` را نشان بده |
| `-e SCRIPT` | چند اسکریپت |
| `-E` / `-r` | regex توسعه‌یافته |

# Examples

```bash
# جایگزینی اولین رخداد در هر خط
sed 's/foo/bar/' file.txt

# جایگزینی همه رخدادها
sed 's/foo/bar/g' file.txt

# ویرایش درجا
sed -i 's/localhost/127.0.0.1/g' config.ini

# حذف خطوط خالی
sed '/^$/d' file.txt

# چاپ فقط خطوط ۲ تا ۱۰
sed -n '2,10p' file.txt

# حذف خط شامل pattern
sed '/DEBUG/d' app.log
```

# Common mistakes

- فراموش کردن `/g` و تعجب از اینکه فقط اولین match عوض شده.
- استفاده از `-i` بدون پشتیبان روی فایل مهم.
- گیج شدن با delimiter وقتی خود pattern شامل `/` است — می‌توان نوشت: `s|foo/bar|baz|g`.

# Tips

- برای ویرایش‌های پیچیده، گاهی `awk` یا یک اسکریپت کوتاه خواناتر است.
- تست اول بدون `-i`، بعد اعمال.

# Related commands

- `awk` — پردازش ستونی و منطقی
- `grep` — فقط جستجو
- `tr` — تبدیل کاراکتر
