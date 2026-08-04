---
title: ip
aliases:
category: network
difficulty: intermediate
keywords:
- address
- route
- link
- netlink
---

# Introduction

`ip` (iproute2) ابزار مدرن تنظیم **شبکهٔ لینوکس** است و جایگزین بسیاری از کارهای `ifconfig`/`route` قدیمی شده: آدرس، لینک، مسیر، همسایه ARP.

# Syntax

```
ip [OPTIONS] OBJECT COMMAND
ip OBJECT help
```

OBJECTهای رایج: `link`, `address` (`addr`), `route`, `neigh`, `rule`.

# Options

## سراسری

| گزینه | توضیح |
|-------|--------|
| `-4` / `-6` | فقط IPv4 / IPv6 |
| `-s` | آمار |
| `-d` | جزئیات بیشتر |
| `-c` | رنگ |
| `-br` | خروجی مختصر |
| `-j` | JSON (نسخه‌های جدید) |

## `ip link`

| فرمان | کار |
|--------|-----|
| `ip link show` | لیست اینترفیس |
| `ip link set DEV up\|down` | بالا/پایین |
| `ip link set DEV mtu N` | MTU |
| `ip link add` / `del` | ساخت/حذف (veth, bridge, …) |

## `ip address`

| فرمان | کار |
|--------|-----|
| `ip addr show` | آدرس‌ها |
| `ip addr add 192.0.2.10/24 dev eth0` | افزودن |
| `ip addr del ... dev eth0` | حذف |
| `ip -br addr` | خلاصه خوانا |

## `ip route`

| فرمان | کار |
|--------|-----|
| `ip route show` | جدول مسیر |
| `ip route get 1.1.1.1` | کدام مسیر استفاده می‌شود |
| `ip route add default via GATEWAY` | default gateway |
| `ip route add 10.0.0.0/8 via GW dev DEV` | مسیر خاص |
| `ip route del ...` | حذف |

## `ip neigh`

| فرمان | کار |
|--------|-----|
| `ip neigh show` | جدول ARP/ND |
| `ip neigh flush dev eth0` | پاکسازی |

# Examples

## وضعیت سریع

```bash
ip -br link
ip -br addr
ip route
ip route get 8.8.8.8
```

## تنظیم موقت آدرس

```bash
sudo ip addr add 192.168.100.5/24 dev eth0
sudo ip link set eth0 up
sudo ip route add default via 192.168.100.1
# بعد از reboot معمولاً از بین می‌رود مگر در NetworkManager/netplan ذخیره شود
```

## عیب‌یابی

```bash
ip -s link show eth0
ip neigh show
ping -c 2 192.168.100.1
```

# Common mistakes

- تغییر با `ip` را دائمی فرض کردن بدون کانفیگ distro (netplan, NM, systemd-networkd).
- اشتباه mask: `/24` در برابر subnet غلط.
- `ifconfig` را هنوز در اسکریپت‌های جدید ترجیح دادن.

# Tips

- برای DNS جدا: `resolvectl` (systemd) یا فایل‌های distro.
- فایروال: `nft` / `ufw` نه خود `ip`.
- `ip -j addr` برای اسکریپت.

# Related commands

- `ss` · `ping` · `nmcli` · `resolvectl` · `ethtool` · `ufw`
