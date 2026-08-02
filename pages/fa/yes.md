---
title: yes
aliases:
category: shell
difficulty: beginner
keywords:
- repeat
- pipe
- confirm
---

# Introduction

`yes` یک رشته را تا ابد تکرار می‌کند؛ معمولاً برای پاسخ خودکار `y` به دستورات تعاملی.

# Syntax

```
yes [STRING]
```

# Options

بدون آرگومان `y` چاپ می‌کند.

# Examples

```bash
yes | head -n 5
yes y | command-that-asks
yes '' | head   # خطوط خالی
```

# Common mistakes

- رها کردن `yes` بدون pipe و پر شدن ترمینال.

# Tips

- توقف با `Ctrl+C`.

# Related commands

- `seq`
- `xargs`
