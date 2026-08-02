---
title: traceroute
aliases:
- tracepath
category: network
difficulty: intermediate
keywords:
- route
- hops
- latency
---

# Introduction

`traceroute` مسیر بسته‌ها تا مقصد را hopبه‌hop نشان می‌دهد و به عیب‌یابی مسیریابی کمک می‌کند.

# Syntax

```
traceroute [OPTIONS] HOST
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-n` | بدون resolve نام |
| `-m N` | حداکثر hop |
| `-q N` | تعداد probe |
| `-I` | ICMP به‌جای UDP |

# Examples

```bash
traceroute example.com
traceroute -n 8.8.8.8
tracepath example.com   # اغلب بدون root
```

# Common mistakes

- ستاره (`*`) همیشه یعنی قطع نیست؛ ممکن است ICMP time-exceeded فیلتر شده باشد.

# Tips

- در بعضی توزیع‌ها به‌جای آن `tracepath` یا `mtr` نصب است.

# Related commands

- `ping`
- `mtr` — ترکیب ping+traceroute
- `ip route`
