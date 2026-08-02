---
title: which
aliases:
category: shell
difficulty: beginner
keywords:
- path
- binary
- location
---

# Introduction

`which` مسیر اجرایی دستوری را که شل پیدا می‌کند نشان می‌دهد (جستجو در `PATH`).

# Syntax

```
which [OPTIONS] COMMAND...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-a` | همه مسیرهای مطابق |

# Examples

```bash
which python3
which -a python
which ls
```

# Common mistakes

- `which` ممکن است builtinها و aliasها را مثل `type` نشان ندهد.

# Tips

- برای اطلاعات کامل‌تر در bash: `type -a command`

# Related commands

- `type` — builtin شل
- `whereis` — باینری، سورس، man
- `command -v` — قابل حمل‌تر در اسکریپت
