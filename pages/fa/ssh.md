---
title: ssh
aliases:
- secure-shell
category: network
difficulty: intermediate
keywords:
- remote
- shell
- key
- tunnel
- openssh
---

# Introduction

`ssh` (OpenSSH) یک شل امن روی شبکه است: ورود به سرور، اجرای فرمان از راه دور، تونل پورت و پایهٔ `scp`/`sftp`/`rsync -e ssh`.

# Syntax

```
ssh [OPTIONS] [user@]host [command]
ssh -f -N -L ...
```

# Options

## اتصال و هویت

| گزینه | توضیح |
|-------|--------|
| `-p PORT` | پورت (پیش‌فرض ۲۲) |
| `-l USER` | نام کاربر (معادل `user@host`) |
| `-i IDENTITY` | فایل کلید خصوصی |
| `-F CONFIG` | فایل کانفیگ جایگزین |
| `-o OPTION=VALUE` | تنظیم OpenSSH (مثلاً `StrictHostKeyChecking`) |
| `-v` / `-vv` | دیباگ |
| `-A` | agent forwarding (با احتیاط) |
| `-X` / `-Y` | X11 forwarding |

## تونل و پس‌زمینه

| گزینه | توضیح |
|-------|--------|
| `-L [bind:]port:host:hostport` | Local port forward |
| `-R [bind:]port:host:hostport` | Remote port forward |
| `-D [bind:]port` | SOCKS dynamic |
| `-N` | فرمان remote اجرا نکن (فقط تونل) |
| `-f` | بعد از احراز هویت پس‌زمینه |
| `-W host:port` | stdio forward (Proxy) |
| `-J jump` | ProxyJump یک‌خطی |
| `-t` / `-T` | اجبار / منع TTY |
| `-q` | ساکت |

## فایل `~/.ssh/config` (خلاصه)

```
Host myserver
  HostName 203.0.113.10
  User deploy
  Port 2222
  IdentityFile ~/.ssh/id_ed25519
  ProxyJump bastion
```

سپس: `ssh myserver`

# Examples

## ورود و فرمان

```bash
ssh alice@server.example.com
ssh -p 2222 alice@server.example.com
ssh alice@server 'uptime; df -h'
```

## کلید

```bash
ssh-keygen -t ed25519 -C "me@laptop"
ssh-copy-id -i ~/.ssh/id_ed25519.pub alice@server
ssh -i ~/.ssh/id_ed25519 alice@server
```

## تونل محلی (دیتابیس روی سرور)

```bash
# localhost:5433 → server به 127.0.0.1:5432
ssh -N -L 5433:127.0.0.1:5432 alice@server
# در ترمینال دیگر:
psql -h 127.0.0.1 -p 5433 -U dbuser
```

## Jump host

```bash
ssh -J bastion.example.com alice@internal.lan
# یا در config: ProxyJump bastion.example.com
```

## کپی با ابزارهای وابسته

```bash
scp -r ./app alice@server:/var/www/
rsync -avz -e ssh ./app/ alice@server:/var/www/app/
```

# Common mistakes

- گذاشتن کلید **خصوصی** روی سرور یا در git.
- `StrictHostKeyChecking=no` دائمی در production.
- agent forwarding (`-A`) روی سرورهای غیرقابل‌اعتماد.
- فراموش کردن `-p` وقتی sshd روی پورت غیر۲۲ است.

# Tips

- برای چند سرور همیشه `~/.ssh/config` بنویسید.
- `ssh-add` برای کلیدهای passphrase‌دار.
- لاگ سرور: `journalctl -u ssh` یا `auth.log`.

# Related commands

- `scp` · `sftp` · `ssh-keygen` · `ssh-copy-id` · `rsync` · `sshd`
