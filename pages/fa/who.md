---
title: who
aliases:
category: user
difficulty: beginner
keywords:
- login
- users
- session
---

# Introduction

`who` کاربران لاگین‌شده به سیستم را نشان می‌دهد.

# Syntax

```
who [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-b` | زمان boot |
| `-u` | idle |
| `-a` | همه اطلاعات |

# Examples

```bash
who
who -ab
w          # نسخه پرجزئیات‌تر
```

# Common mistakes

- اشتباه گرفتن با `whoami` (کاربر فعلی).

# Tips

- `w` load و دستور جاری را هم نشان می‌دهد.

# Related commands

- `w`
- `whoami`
- `last`
