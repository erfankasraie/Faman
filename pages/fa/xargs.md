---
title: xargs
aliases:
category: shell
difficulty: intermediate
keywords:
- arguments
- pipe
- batch
---

# Introduction

`xargs` ورودی استاندارد را به آرگومان‌های یک دستور تبدیل می‌کند. پلی بین `find`/`grep` و دستوراتی که لیست فایل می‌گیرند.

# Syntax

```
... | xargs [OPTIONS] COMMAND
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-n N` | حداکثر N آرگومان در هر اجرا |
| `-p` | تأیید قبل از اجرا |
| `-i` / `-I {}` | جایگزینی placeholder |
| `-0` | جداکننده null (امن با find -print0) |
| `-P N` | اجرای موازی |

# Examples

```bash
# حذف فایل‌های پیدا شده
find . -name '*.tmp' -print0 | xargs -0 rm -f

# یک آرگومان در هر بار
echo a b c | xargs -n1 echo

# با placeholder
find . -name '*.md' | xargs -I{} cp {} /backup/
```

# Common mistakes

- نام فایل با فاصله بدون `-0` و `-print0` می‌شکند.
- دستورات تعاملی داخل xargs رفتار عجیب دارند.

# Tips

- جایگزین مدرن در GNU find: `-exec ... {} +`

# Related commands

- `find` — تولید لیست فایل
- `parallel` — موازی‌سازی پیشرفته‌تر
