---
title: nslookup
aliases:
category: network
difficulty: beginner
keywords:
- dns
- domain
- lookup
- network
---

# Introduction

`nslookup` یک ابزار ساده و قدیمی برای استعلام رکوردهای DNS است؛ نام دامنه را به آدرس IP (یا برعکس) تبدیل می‌کند.

برای مبتدیان: اگر می‌خواهید بدانید یک دامنه به چه IP‌ای اشاره می‌کند، `nslookup domain.com` سریع‌ترین راه است.

# Syntax

```
nslookup [DOMAIN or IP] [DNS_SERVER]
```

# Options

| گزینه / حالت | توضیح |
|----------------|--------|
| `DOMAIN` | نام دامنه برای resolve |
| `IP` | جستجوی معکوس |
| `DNS_SERVER` | سرور DNS جایگزین (مثلاً `8.8.8.8`) |
| `-type=TYPE` | نوع رکورد: `A`, `AAAA`, `MX`, `NS`, `TXT`, … |
| حالت تعاملی | بدون آرگومان: prompt برای چند query |

# Examples

```bash
# پیداکردن IP یک دامنه
nslookup example.com

# استعلام از یک سرور DNS خاص (مثلاً DNS گوگل)
nslookup example.com 8.8.8.8

# جستجوی معکوس (پیداکردن نام دامنه از روی IP)
nslookup 93.184.216.34

# استعلام نوع خاصی از رکورد (مثل رکورد MX ایمیل)
nslookup -type=MX example.com
```

# Common mistakes

- تصور اینکه `nslookup` همیشه دقیق‌ترین ابزار است؛ ابزار `dig` اطلاعات فنی‌تر و کامل‌تری می‌دهد و در محیط‌های حرفه‌ای ترجیح داده می‌شود.
- عدم توجه به کش DNS محلی که ممکن است نتیجه قدیمی نشان دهد.

# Tips

- برای عیب‌یابی سریع «چرا سایت بالا نمی‌آید»، ابتدا `nslookup domain.com` را برای اطمینان از صحت DNS اجرا کنید.
- برای بررسی‌های دقیق‌تر و پیشرفته DNS (مثل TTL، رکوردهای متعدد)، `dig` جایگزین بهتری است.

# Related commands

- `dig` — ابزار پیشرفته‌تر استعلام DNS
- `whois` — اطلاعات ثبت دامنه
- `ping` — بررسی در‌دسترس‌بودن یک هاست
- `getent` — استعلام از طریق تنظیمات NSS سیستم
