---
title: shutdown
aliases:
category: system
difficulty: beginner
keywords:
- poweroff
- halt
- reboot
---

# Introduction

`shutdown` سیستم را خاموش یا بعد از تأخیر راه‌اندازی مجدد می‌کند و به کاربران پیام می‌دهد.

# Syntax

```
shutdown [OPTIONS] [TIME] [MESSAGE]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-h` | halt/poweroff |
| `-r` | reboot |
| `-c` | لغو shutdown زمان‌بندی‌شده |
| `now` | فوری |
| `+N` | بعد از N دقیقه |

# Examples

```bash
sudo shutdown -h now
sudo shutdown -r +10 "Reboot in 10 minutes"
sudo shutdown -c
```

# Common mistakes

- فراموش کردن `-c` برای لغو.

# Tips

- روی دسکتاپ گاهی از منوی گرافیکی امن‌تر است.

# Related commands

- `reboot`
- `poweroff`
- `systemctl poweroff`
