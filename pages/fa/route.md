---
title: route
aliases:
category: network
difficulty: advanced
keywords:
- network
- routing
- gateway
---

# Introduction

`route` جدول مسیریابی (routing table) کرنل را نمایش یا ویرایش می‌کند؛ مشخص می‌کند بسته‌های شبکه برای رسیدن به مقصدهای مختلف از کدام مسیر/gateway عبور کنند. در سیستم‌های مدرن جای خود را به `ip route` داده است.

# Syntax

```
route [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-n` | نمایش آدرس‌های عددی (سریع‌تر، بدون resolve نام) |
| `add` | افزودن یک مسیر جدید |
| `del` | حذف یک مسیر |

# Examples

```bash
# نمایش جدول مسیریابی فعلی
route -n

# افزودن یک مسیر پیش‌فرض جدید (نیاز به sudo)
sudo route add default gw 192.168.1.1

# افزودن مسیر به یک شبکه خاص از طریق یک gateway مشخص
sudo route add -net 10.0.0.0/24 gw 192.168.1.254

# حذف یک مسیر
sudo route del -net 10.0.0.0/24
```

# Common mistakes

- استفاده از دستور قدیمی `route` در توزیع‌های مدرن که بسته `net-tools` را پیش‌فرض ندارند؛ باید یا نصب شود یا از `ip route` استفاده شود.
- فراموش‌کردن این‌که تغییرات `route` پس از ریبوت سیستم از بین می‌رود، مگر این‌که در تنظیمات دائمی شبکه (مثل netplan یا NetworkManager) ذخیره شود.

# Tips

- معادل مدرن و پیشنهادی: `ip route show` برای نمایش، و `ip route add ...` برای افزودن.
- برای عیب‌یابی «چرا اینترنت وصل نیست ولی شبکه محلی کار می‌کند»، بررسی وجود مسیر پیش‌فرض (`default`) اولین قدم است.

# Related commands

- `ip` — ابزار مدرن شبکه (`ip route` جایگزین route)
- `ping` — تست اتصال به یک مقصد
- `traceroute` — نمایش مسیر عبور بسته‌ها تا مقصد
