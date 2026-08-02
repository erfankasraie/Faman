---
title: diff
aliases:
category: text
difficulty: beginner
keywords:
- compare
- patch
- changes
---

# Introduction

`diff` تفاوت دو فایل یا پوشه را نشان می‌دهد.

# Syntax

```
diff [OPTIONS] FILE1 FILE2
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-u` | unified (رایج برای patch) |
| `-c` | context |
| `-r` | بازگشتی روی پوشه |
| `-N` | فایل ناموجود را خالی فرض کن |
| `-q` | فقط بگو فرق دارند یا نه |
| `-i` | بدون حساسیت حروف |

# Examples

```bash
diff -u old.txt new.txt
diff -ruN dir1/ dir2/
diff -u a b > changes.patch
```

# Common mistakes

- خواندن خروجی پیش‌فرض بدون `-u` سخت‌تر است.

# Tips

- ابزارهای رنگی: `colordiff` یا `delta`.

# Related commands

- `patch` — اعمال diff
- `git diff` — در مخزن git
- `comm` — مقایسه فایل‌های مرتب
