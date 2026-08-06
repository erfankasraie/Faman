---
title: getent
aliases:
category: system
difficulty: advanced
keywords:
- nss
- database
- user
- dns
---

# Introduction

`getent` اطلاعات را از پایگاه‌های‌داده NSS سیستم (Name Service Switch) می‌خواند — این پایگاه‌ها می‌توانند فایل محلی (مثل `/etc/passwd`)، LDAP، یا DNS باشند، بسته به تنظیمات `/etc/nsswitch.conf`. برخلاف خواندن مستقیم `/etc/passwd`، `getent` دقیقاً همان چیزی را نشان می‌دهد که سیستم واقعاً هنگام resolve کاربران/گروه‌ها/هاست‌ها استفاده می‌کند.

# Syntax

```
getent DATABASE [KEY]
```

# Options

`getent` گزینه‌ی خط‌فرمانی خاصی ندارد؛ آرگومان اول نام پایگاه‌داده (مثل `passwd`, `hosts`) و آرگومان دوم (اختیاری) کلید جستجوست.

پایگاه‌های‌داده رایج: `passwd`, `group`, `hosts`, `services`, `protocols`

# Examples

```bash
# بررسی اطلاعات یک کاربر خاص (شامل LDAP اگر تنظیم شده باشد)
getent passwd ali

# لیست تمام کاربران قابل‌مشاهده سیستم
getent passwd

# بررسی گروه‌های سیستم
getent group developers

# resolve نام دامنه (مشابه nslookup ولی از طریق تنظیمات NSS واقعی سیستم)
getent hosts example.com

# بررسی پورت یک سرویس شناخته‌شده
getent services ssh
```

# Common mistakes

- خواندن مستقیم `/etc/passwd` برای بررسی وجود یک کاربر در محیط‌هایی که از LDAP/NIS هم استفاده می‌کنند؛ در این حالت `/etc/passwd` فقط کاربران محلی را نشان می‌دهد، نه کاربران متمرکز شبکه — `getent passwd` صحیح و کامل است.
- تصور اینکه `getent hosts` همیشه از DNS واقعی می‌خواند؛ اگر `/etc/hosts` مقدار override داشته باشد، همان اولویت دارد (طبق ترتیب در `nsswitch.conf`).

# Tips

- برای عیب‌یابی «چرا این کاربر/گروه در سیستم دیده نمی‌شود» در محیط‌های سازمانی با LDAP/AD، همیشه `getent` را به‌جای خواندن مستقیم فایل‌های محلی چک کنید.
- `getent hosts` روش خوبی برای دیدن دقیقاً همان resolve‌ای است که برنامه‌های سیستم می‌بینند (شامل `/etc/hosts` و DNS با هم).

# Related commands

- `nslookup` / `dig` — استعلام مستقیم DNS (بدون درنظرگرفتن `/etc/hosts` یا NSS)
- `id` — اطلاعات کاربر جاری
- `cat /etc/nsswitch.conf` — بررسی ترتیب منابع NSS
