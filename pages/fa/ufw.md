---
title: ufw
aliases:
category: network
difficulty: intermediate
keywords:
- firewall
- security
- port
---

# Introduction

`ufw` (Uncomplicated Firewall) رابط ساده فایروال روی Ubuntu/Debian است.

# Syntax

```
ufw [OPTIONS] COMMAND
```

# Options

| فرمان | توضیح |
|-------|--------|
| `enable` / `disable` | روشن/خاموش |
| `status` | وضعیت |
| `allow PORT` | اجازه |
| `deny PORT` | مسدود |
| `delete RULE` | حذف قانون |
| `reload` | اعمال مجدد |

# Examples

```bash
sudo ufw status verbose
sudo ufw allow 22/tcp
sudo ufw allow OpenSSH
sudo ufw allow 80,443/tcp
sudo ufw enable
sudo ufw deny 23
```

# Common mistakes

- فعال کردن ufw قبل از allow کردن SSH و قفل شدن بیرون سرور.

# Tips

- اول SSH را allow کنید، بعد enable.

# Related commands

- `iptables` / `nft` — سطح پایین‌تر
- `firewall-cmd` — firewalld (RHEL/Fedora)
- `ss` — پورت‌های باز
