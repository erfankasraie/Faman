---
title: tmux
aliases:
category: shell
difficulty: intermediate
keywords:
- terminal
- multiplexer
- session
- pane
---

# Introduction

`tmux` چند **نشست ترمینال** را نگه می‌دارد که با قطع SSH هم زنده می‌مانند: پنجره (window)، قاب (pane)، و session قابل‌اتصال مجدد.

# Syntax

```
tmux [command] [options]
tmux new -s NAME
tmux attach -t NAME
```

کلید پیش‌فرض prefix: **`Ctrl-b`** سپس کلید فرمان.

# Options

## خط فرمان

| فرمان | کار |
|--------|-----|
| `tmux` / `tmux new` | session جدید |
| `tmux new -s NAME` | با نام |
| `tmux ls` | لیست session |
| `tmux attach -t NAME` | اتصال |
| `tmux attach -d -t NAME` | اتصال و detach دیگران |
| `tmux kill-session -t NAME` | کشتن session |
| `tmux kill-server` | همه |

## بعد از prefix (`Ctrl-b`)

| کلید | کار |
|------|-----|
| `d` | detach |
| `c` | window جدید |
| `n` / `p` | window بعد/قبل |
| `,` | نام‌گذاری window |
| `&` | بستن window |
| `%` | شکاف عمودی pane |
| `"` | شکاف افقی |
| `←↑↓→` | جابه‌جایی بین pane |
| `x` | بستن pane |
| `z` | zoom pane |
| `[` | mode کپی/اسکرول |
| `?` | لیست کلیدها |

# Examples

## کار روی سرور

```bash
ssh server
tmux new -s work
# ... کار طولانی ...
# Ctrl-b d   → detach
exit   # از ssh

# بعداً:
ssh server
tmux attach -t work
```

## چند session

```bash
tmux ls
tmux new -s logs
tmux new -s build
tmux attach -t build
```

## اسکریپت

```bash
tmux new-session -d -s job 'npm run build'
tmux attach -t job
```

# Common mistakes

- بستن ترمینال بدون detach → اگر session نباشد پروسه ممکن است بمیرد (با tmux معمولاً session می‌ماند فقط اگر داخل tmux بوده باشید).
- قاطی کردن prefix با `screen` (`Ctrl-a`).
- scroll با چرخ ماوس بدون mode یا تنظیم terminal-overrides.

# Tips

- کانفیگ: `~/.tmux.conf` — مثلاً `set -g mouse on`.
- جایگزین سبک: `zellij`، `screen`.
- برای لاگ بلند: pane جدا + `journalctl -f`.

# Related commands

- `screen` · `ssh` · `nohup` · `jobs`
