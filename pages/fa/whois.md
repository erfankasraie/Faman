---
title: whois
aliases:
category: network
difficulty: beginner
keywords:
- domain
- dns
- registration
- lookup
---

# Introduction

`whois` اطلاعات ثبتی یک دامنه یا آدرس IP را از پایگاه‌داده‌های WHOIS نمایش می‌دهد؛ شامل صاحب دامنه (در صورت عدم privacy)، تاریخ ثبت و انقضا، و سرورهای نام (nameservers).

# Syntax

```
whois [OPTIONS] DOMAIN_OR_IP
```

# Options

| گزینه | توضیح |
|-------|--------|
| `DOMAIN_OR_IP` | دامنه یا آدرس برای استعلام |
| `-h HOST` | سرور whois مشخص (در برخی پیاده‌سازی‌ها) |
| `-p PORT` | پورت سرور whois |

خروجی بین registryها فرق می‌کند؛ فیلدها استاندارد یکسان ندارند.

# Examples

```bash
# دیدن اطلاعات ثبت یک دامنه
whois example.com

# دیدن اطلاعات یک آدرس IP (صاحب رنج IP)
whois 8.8.8.8

# فیلترکردن فقط تاریخ انقضا
whois example.com | grep -i "expir"

# دیدن nameserverهای یک دامنه
whois example.com | grep -i "name server"
```

# Common mistakes

- انتظار داشتن اطلاعات کامل مالک دامنه؛ بسیاری از دامنه‌ها با سرویس privacy protection ثبت شده‌اند و اطلاعات واقعی صاحب نمایش داده نمی‌شود.
- در برخی توزیع‌ها ابزار `whois` به‌صورت پیش‌فرض نصب نیست و باید نصب شود (`apt install whois`).

# Tips

- برای بررسی سریع اینکه یک دامنه چه زمانی منقضی می‌شود (مفید برای یادآوری تمدید)، `whois domain.com | grep -i expir` سریع‌ترین راه است.
- برای اطلاعات فنی‌تر DNS (نه ثبت دامنه)، از `dig` استفاده کنید.

# Related commands

- `dig` — استعلام رکوردهای DNS
- `nslookup` — استعلام ساده DNS
- `curl` — بررسی وضعیت واقعی یک وب‌سایت
