---
title: seq
aliases:
category: shell
difficulty: beginner
keywords:
- numbers
- range
- loop
---

# Introduction

`seq` دنباله اعداد را چاپ می‌کند؛ مفید برای حلقه‌ها و تست.

# Syntax

```
seq [OPTIONS] LAST
seq FIRST LAST
seq FIRST INCREMENT LAST
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-s STR` | جداکننده |
| `-w` | عرض یکسان با صفر |
| `-f FORMAT` | قالب printf |

# Examples

```bash
seq 5
seq 2 10
seq 0 2 10
seq -s, 1 5
seq -w 1 10
```

# Common mistakes

- در bash معمولاً `{1..5}` هم همان کار را می‌کند.

# Tips

- `for i in $(seq 1 10); do ...; done`

# Related commands

- `yes`
- حلقه bash brace expansion
