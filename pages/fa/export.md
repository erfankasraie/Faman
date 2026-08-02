---
title: export
aliases:
category: shell
difficulty: beginner
keywords:
- environment
- variable
- shell
---

# Introduction

`export` متغیر را به محیط فرایندهای فرزند می‌فرستد.

# Syntax

```
export NAME=value
export NAME
export -n NAME
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-n` | لغو export |
| `-p` | لیست exportها |

# Examples

```bash
export PATH="$HOME/bin:$PATH"
export EDITOR=vim
export -p | grep PATH
```

# Common mistakes

- فاصله دور `=` : `export A = 1` غلط است.

# Tips

- برای یک اجرا: `ENV=1 command` بدون export.

# Related commands

- `env`
- `printenv`
- `set`
