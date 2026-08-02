---
title: time
aliases:
category: system
difficulty: beginner
keywords:
- benchmark
- duration
- performance
---

# Introduction

`time` مدت اجرای یک دستور را اندازه می‌گیرد (real/user/sys).

# Syntax

```
time COMMAND
```

# Options

بسته به builtin شل یا `/usr/bin/time` فرق می‌کند. نسخه باینری:

| گزینه | توضیح |
|-------|--------|
| `-v` | جزئیات زیاد |
| `-f FORMAT` | قالب سفارشی |

# Examples

```bash
time ls -R /
/usr/bin/time -v ./myapp
```

# Common mistakes

- اشتباه گرفتن builtin `time` با `/usr/bin/time`.

# Tips

- برای بنچمارک جدی‌تر از `hyperfine` استفاده کنید.

# Related commands

- `date`
- `perf` — پروفایل کرنل
