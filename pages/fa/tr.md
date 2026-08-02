---
title: tr
aliases:
category: text
difficulty: beginner
keywords:
- translate
- delete
- characters
---

# Introduction

`tr` کاراکترها را ترجمه، حذف یا فشرده می‌کند. روی stdin کار می‌کند (نه مستقیم روی نام فایل).

# Syntax

```
tr [OPTIONS] SET1 [SET2]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-d` | حذف کاراکترهای SET1 |
| `-s` | فشرده‌سازی تکرارها |
| `-c` | مکمل مجموعه |

# Examples

```bash
# حروف کوچک به بزرگ
echo hello | tr 'a-z' 'A-Z'

# حذف ارقام
echo a1b2 | tr -d '0-9'

# تبدیل فاصله به خط جدید
echo 'a b c' | tr ' ' '\n'

# فشردن فاصله‌های پشت‌سرهم
echo 'a    b' | tr -s ' '
```

# Common mistakes

- دادن نام فایل به‌جای pipe: `tr` فایل نمی‌خواند مگر با `< file`.

# Tips

- برای Unicode پیچیده، ابزارهای دیگر ممکن است قابل‌اعتمادتر باشند.

# Related commands

- `sed` — جایگزینی الگو
- `awk` — منطق پیچیده‌تر
