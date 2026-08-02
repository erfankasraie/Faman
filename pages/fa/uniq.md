---
title: uniq
aliases:
category: text
difficulty: beginner
keywords:
- unique
- duplicate
- count
---

# Introduction

`uniq` خطوط تکراری **پشت‌سرهم** را حذف یا شمارش می‌کند. معمولاً بعد از `sort` استفاده می‌شود.

# Syntax

```
uniq [OPTIONS] [INPUT [OUTPUT]]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-c` | شمارش تکرار |
| `-d` | فقط خطوط تکراری |
| `-u` | فقط خطوط یکتا |
| `-i` | بدون حساسیت به حروف |

# Examples

```bash
sort names.txt | uniq
sort names.txt | uniq -c
sort names.txt | uniq -d
```

# Common mistakes

- انتظار حذف همهٔ تکراری‌ها بدون `sort` اول — `uniq` فقط همسایه‌ها را می‌بیند.

# Tips

- پرتکرارترین‌ها: `sort file | uniq -c | sort -nr`

# Related commands

- `sort` — پیش‌نیاز معمول
- `awk '!seen[$0]++'` — یکتاسازی بدون sort (حافظه بیشتر)
