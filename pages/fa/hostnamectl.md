---
title: hostnamectl
aliases:
category: system
difficulty: beginner
keywords:
- hostname
- systemd
- identity
---

# Introduction

`hostnamectl` نام میزبان و اطلاعات مرتبط را در سیستم‌های systemd مدیریت می‌کند.

# Syntax

```
hostnamectl [OPTIONS] [COMMAND]
```

# Options

| فرمان | توضیح |
|-------|--------|
| `status` | نمایش (پیش‌فرض) |
| `set-hostname NAME` | تنظیم نام |
| `set-chassis TYPE` | نوع دستگاه |

# Examples

```bash
hostnamectl
sudo hostnamectl set-hostname web-01
```

# Common mistakes

- تغییر فقط با `hostname` موقت بدون hostnamectl.

# Tips

- نام در `/etc/hostname` و گاهی hosts هم باید سازگار باشد.

# Related commands

- `hostname`
- `timedatectl`
- `localectl`
