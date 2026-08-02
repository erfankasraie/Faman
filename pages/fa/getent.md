---
title: getent
aliases:
category: system
difficulty: intermediate
keywords:
- nss
- passwd
- hosts
---

# Introduction

`getent` از Name Service Switch داده می‌گیرد؛ کاربران، گروه‌ها، hosts و ... حتی اگر در LDAP باشند.

# Syntax

```
getent DATABASE [KEY...]
```

# Options

دیتابیس‌های رایج: `passwd`, `group`, `hosts`, `services`, `shadow`.

# Examples

```bash
getent passwd alice
getent group sudo
getent hosts example.com
getent services ssh
```

# Common mistakes

- فقط خواندن `/etc/passwd` وقتی کاربر در LDAP است.

# Tips

- برای عیب‌یابی NSS مفیدتر از cat فایل‌های محلی است.

# Related commands

- `id`
- `getent passwd` ≈ لیست کاربران قابل دیدن
