---
title: host
aliases:
category: network
difficulty: beginner
keywords:
- dns
- lookup
- dig
---

# Introduction

`host` ابزار سادهٔ DNS برای تبدیل نام↔IP و مشاهده رکوردهاست؛ سبک‌تر از `dig` برای کار روزمره.

# Syntax

```
host [OPTIONS] NAME [SERVER]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-t TYPE` | نوع رکورد: A, AAAA, MX, NS, TXT, … |
| `-a` | همه رکوردها (verbose) |
| `-v` | جزئیات |

# Examples

```bash
host example.com
host -t MX example.com
host 8.8.8.8
host example.com 1.1.1.1
```

# Common mistakes

- کش DNS محلی نتیجه را کهنه نشان می‌دهد.

# Tips

- برای دیباگ حرفه‌ای: `dig +trace`.

# Related commands

- `dig` · `nslookup` · `ping`
