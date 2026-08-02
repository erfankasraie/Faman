---
title: last
aliases:
category: user
difficulty: beginner
keywords:
- login
- history
- audit
---

# Introduction

`last` تاریخچه ورود کاربران را از `wtmp` نشان می‌دهد.

# Syntax

```
last [OPTIONS] [USER]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-n N` | N مورد آخر |
| `-a` | hostname در انتها |
| `-x` | شامل shutdown/reboot |

# Examples

```bash
last
last -n 20
last reboot
last alice
```

# Common mistakes

- خالی بودن خروجی اگر logrotate فایل wtmp را پاک کرده باشد.

# Tips

- برای تلاش‌های ناموفق: `lastb` (نیاز به دسترسی).

# Related commands

- `who`
- `journalctl`
- `lastlog`
