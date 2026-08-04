---
title: whereis
aliases:
category: system
difficulty: beginner
keywords:
- locate
- binary
- man
- path
---

# Introduction

`whereis` محل باینری، منبع و man page یک فرمان را در مسیرهای استاندارد جستجو می‌کند (سریع‌تر و محدودتر از `find`).

# Syntax

```
whereis [OPTIONS] name...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-b` | فقط باینری |
| `-m` | فقط man |
| `-s` | فقط source |
| `-B dirs` | محدود کردن جستجوی باینری |

# Examples

```bash
whereis ls
whereis -b python
whereis gcc make
```

# Common mistakes

- انتظار پیدا کردن اسکریپت‌های خارج از مسیرهای استاندارد → `which` / `type`.

# Tips

- `type -a cmd` و `command -v cmd` در شل دقیق‌ترند برای PATH فعلی.

# Related commands

- `which` · `type` · `locate` · `find`
