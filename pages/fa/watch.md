---
title: watch
aliases:
category: system
difficulty: beginner
keywords:
- repeat
- monitor
- refresh
---

# Introduction

`watch` یک دستور را در بازه ثابت تکرار و خروجی را روی صفحه نشان می‌دهد.

# Syntax

```
watch [OPTIONS] COMMAND
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-n SEC` | فاصله ثانیه (پیش‌فرض ۲) |
| `-d` | هایلایت تفاوت‌ها |
| `-t` | بدون هدر |

# Examples

```bash
watch -n 1 'df -h'
watch -d 'ps aux | head'
watch -n 5 systemctl status nginx
```

# Common mistakes

- فراموش کردن کوتیشن دور pipeline.

# Tips

- خروج با `Ctrl+C`.

# Related commands

- `top` — مانیتور فرایند
- `journalctl -f` — دنبال لاگ
