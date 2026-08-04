---
title: arp
aliases:
category: network
difficulty: intermediate
keywords:
- network
- mac
- address
- lan
---

# Introduction

`arp` جدول ARP (Address Resolution Protocol) سیستم را نمایش یا مدیریت می‌کند؛ این جدول نگاشت بین آدرس‌های IP و آدرس‌های فیزیکی (MAC) دستگاه‌های شبکه محلی را نگه می‌دارد.

# Syntax

```
arp [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-a` | نمایش تمام ورودی‌های جدول ARP |
| `-n` | نمایش آدرس IP عددی به‌جای نام هاست |
| `-d IP` | حذف یک ورودی از جدول (نیاز به sudo) |
| `-s IP MAC` | افزودن دستی یک ورودی ثابت |

# Examples

```bash
# نمایش جدول ARP فعلی
arp -a

# نمایش سریع‌تر بدون resolve نام (عددی)
arp -n

# حذف یک ورودی خاص (مثلاً بعد از تغییر کارت شبکه دستگاهی در شبکه)
sudo arp -d 192.168.1.50

# افزودن دستی یک ورودی ثابت
sudo arp -s 192.168.1.50 aa:bb:cc:dd:ee:ff
```

# Common mistakes

- در توزیع‌های جدید `arp` (از پکیج `net-tools`) ممکن است پیش‌فرض نصب نباشد؛ معادل مدرن‌تر `ip neigh` است.
- اشتباه‌گرفتن مشکلات ARP (لایه ۲، شبکه محلی) با مشکلات DNS (که مربوط به نام دامنه است).

# Tips

- معادل مدرن با ابزار `ip`: `ip neigh show`
- برای عیب‌یابی مشکلات اتصال در شبکه محلی (LAN) که IP درست است ولی اتصال برقرار نمی‌شود، بررسی جدول ARP می‌تواند نشان دهد آیا MAC آدرس درست resolve شده یا نه.

# Related commands

- `ip` — ابزار مدرن شبکه (`ip neigh` جایگزین arp)
- `ping` — تست دسترسی به یک هاست
- `netstat` / `ss` — اتصالات و پورت‌های شبکه
