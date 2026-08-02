---
title: timedatectl
aliases:
category: system
difficulty: beginner
keywords:
- time
- timezone
- ntp
---

# Introduction

`timedatectl` ساعت، منطقه زمانی و همگام‌سازی NTP را در systemd کنترل می‌کند.

# Syntax

```
timedatectl [COMMAND]
```

# Options

| فرمان | توضیح |
|-------|--------|
| `status` | وضعیت |
| `list-timezones` | لیست مناطق |
| `set-timezone ZONE` | تنظیم منطقه |
| `set-ntp true/false` | NTP |

# Examples

```bash
timedatectl
timedatectl list-timezones | grep Tehran
sudo timedatectl set-timezone Asia/Tehran
sudo timedatectl set-ntp true
```

# Common mistakes

- تنظیم دستی ساعت وقتی NTP روشن است و دوباره برمی‌گردد.

# Tips

- بعد از تغییر timezone سرویس‌ها را در نظر بگیرید.

# Related commands

- `date`
- `hwclock`
