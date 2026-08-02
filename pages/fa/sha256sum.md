---
title: sha256sum
aliases:
category: security
difficulty: beginner
keywords:
- hash
- checksum
- integrity
---

# Introduction

`sha256sum` چک‌سام SHA-256 فایل را محاسبه یا تأیید می‌کند؛ برای صحت دانلودها رایج است.

# Syntax

```
sha256sum [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-c` | بررسی از روی فایل checksum |
| `-b` | حالت باینری |

# Examples

```bash
sha256sum image.iso
sha256sum image.iso > SHA256SUMS
sha256sum -c SHA256SUMS
echo -n 'hello' | sha256sum
```

# Common mistakes

- مقایسه دستی با کپی ناقص hash.

# Tips

- الگوریتم‌های دیگر: `sha1sum`, `md5sum` (ضعیف‌تر), `b2sum`.

# Related commands

- `md5sum`
- `gpg --verify` — امضای دیجیتال
