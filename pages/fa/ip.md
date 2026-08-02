---
title: ip
aliases:
category: network
difficulty: intermediate
keywords:
- network
- interface
- route
- address
---

# Introduction

`ip` ابزار مدرن مدیریت شبکه در لینوکس است (جایگزین بسیاری از کاربردهای `ifconfig` و `route`).

# Syntax

```
ip [OPTIONS] OBJECT COMMAND
```

اشیاء رایج: `addr`, `link`, `route`, `neigh`.

# Options

| دستور | توضیح |
|-------|--------|
| `ip a` / `ip addr` | آدرس‌ها |
| `ip link` | اینترفیس‌ها |
| `ip route` | جدول مسیریابی |
| `ip neigh` | ARP/neighbor |

# Examples

```bash
ip a
ip -br a
ip link set eth0 up
ip route show
ip addr add 192.168.1.10/24 dev eth0
```

# Common mistakes

- تغییرات بدون پایدارسازی در تنظیمات distro بعد از ریبوت از بین می‌روند.

# Tips

- خروجی خلاصه: `ip -br addr`

# Related commands

- `ss` — سوکت‌ها
- `ping` — تست اتصال
- `nmcli` — NetworkManager
