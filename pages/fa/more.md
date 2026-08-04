---
title: more
aliases:
category: text
difficulty: beginner
keywords:
- pager
- less
- view
---

# Introduction

`more` صفحه‌بند قدیمی ترمینال است. برای جلو رفتن در متن؛ امکانات `less` بیشتر است و معمولاً ترجیح داده می‌شود.

# Syntax

```
more [OPTIONS] [FILE...]
command | more
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-d` | راهنمای کلیدها |
| `-s` | فشردن خطوط خالی پشت‌سرهم |
| `+N` | شروع از خط N |

کلیدها: Space صفحه بعد · Enter خط بعد · `q` خروج · `/` جستجو (بسته به نسخه).

# Examples

```bash
more README.md
dmesg | more
more +50 large.log
```

# Common mistakes

- انتظار scroll عقب مثل less (در many more محدود است).

# Tips

- پیش‌فرض خوب: `export PAGER=less`.

# Related commands

- `less` · `most` · `bat`
