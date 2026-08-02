---
title: screen
aliases:
category: shell
difficulty: intermediate
keywords:
- session
- multiplex
- terminal
---

# Introduction

`screen` مالتی‌پلکسر قدیمی ترمینال است؛ شبیه tmux ولی با کلیدهای متفاوت.

# Syntax

```
screen [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-S NAME` | نام نشست |
| `-ls` | لیست |
| `-r NAME` | اتصال مجدد |
| `-d NAME` | detach از راه دور |

پیشوند: `Ctrl+a` سپس کلید (`d` برای detach، `c` پنجره جدید).

# Examples

```bash
screen -S build
screen -ls
screen -r build
```

# Common mistakes

- قاطی کردن پیشوند با tmux (`Ctrl+b`).

# Tips

- برای کار جدید معمولاً tmux پیشنهاد می‌شود.

# Related commands

- `tmux`
- `nohup`
