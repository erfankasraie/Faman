---
title: dig
aliases:
category: network
difficulty: intermediate
keywords:
- dns
- lookup
- resolve
---

# Introduction

`dig` (domain information groper) ابزار قدرتمند پرس‌وجوی DNS است.

# Syntax

```
dig [OPTIONS] [@SERVER] NAME [TYPE]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `+short` | فقط پاسخ کوتاه |
| `+trace` | مسیر از root |
| `-x IP` | reverse lookup |
| `TYPE` | A, AAAA, MX, NS, TXT, CNAME |

# Examples

```bash
dig example.com
dig example.com A +short
dig @8.8.8.8 example.com
dig -x 8.8.8.8
dig example.com MX
```

# Common mistakes

- تکیه فقط به کش محلی بدون پرس‌وجوی سرور مشخص.

# Tips

- برای عیب‌یابی DNS اول `+short` و بعد جزئیات کامل.

# Related commands

- `nslookup` — ساده‌تر/قدیمی‌تر
- `host` — خلاصه
- `resolvectl` — در systemd-resolved
