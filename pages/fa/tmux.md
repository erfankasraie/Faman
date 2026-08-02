---
title: tmux
aliases:
- terminal-multiplexer
category: shell
difficulty: intermediate
keywords:
- session
- multiplex
- terminal
---

# Introduction

`tmux` نشست ترمینال پایدار می‌سازد؛ با بستن SSH کارها قطع نمی‌شوند و چند پنل/پنجره دارید.

# Syntax

```
tmux [command] [args]
```

# Options

| فرمان | توضیح |
|-------|--------|
| `tmux` / `tmux new` | نشست جدید |
| `tmux ls` | لیست نشست‌ها |
| `tmux attach -t NAME` | اتصال مجدد |
| `tmux kill-session -t NAME` | حذف نشست |

پیشوند پیش‌فرض: `Ctrl+b` سپس کلید.

| کلید | کار |
|------|-----|
| `c` | پنجره جدید |
| `"` | تقسیم افقی |
| `%` | تقسیم عمودی |
| `d` | detach |
| `n` / `p` | پنجره بعدی/قبلی |
| `←↑↓→` | جابه‌جایی بین پنل |

# Examples

```bash
tmux new -s work
# ... کار ...
# Ctrl+b d  → detach

tmux ls
tmux attach -t work
```

# Common mistakes

- فراموش کردن detach و بستن ترمینال (نشست می‌ماند ولی باید attach کنید).

# Tips

- جایگزین محبوب: `zellij` یا `screen`.
- تنظیمات در `~/.tmux.conf`.

# Related commands

- `screen` — قدیمی‌تر
- `nohup` — فقط یک فرایند
- `ssh` — ورود از راه دور
