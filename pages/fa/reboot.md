---
title: reboot
aliases:
category: system
difficulty: beginner
keywords:
- restart
- shutdown
- power
---

# Introduction

`reboot` سیستم را مجدداً راه‌اندازی می‌کند.

# Syntax

```
reboot [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-f` | اجباری (بدون shutdown ملایم) |
| `--help` | راهنما |

روی systemd معمولاً معادل `systemctl reboot` است.

# Examples

```bash
sudo reboot
sudo systemctl reboot
```

# Common mistakes

- reboot روی سرور production بدون اطلاع دیگران.

# Tips

- قبل از reboot سرویس‌های حساس را چک کنید.

# Related commands

- `shutdown`
- `poweroff`
- `systemctl`
