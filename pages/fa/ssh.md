---
title: ssh
aliases:
category: network
difficulty: intermediate
keywords:
- remote
- shell
- secure
- login
---

# Introduction

`ssh` اتصال امن به شل راه‌دور برقرار می‌کند؛ پایه مدیریت سرور لینوکس.

# Syntax

```
ssh [OPTIONS] [user@]host [command]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-p PORT` | پورت |
| `-i KEY` | فایل کلید |
| `-L` / `-R` | port forward |
| `-N` | بدون اجرای فرمان (فقط تانل) |
| `-v` | verbose |
| `-A` | agent forwarding |

# Examples

```bash
ssh user@192.168.1.10
ssh -p 2222 alice@example.com
ssh -i ~/.ssh/id_ed25519 deploy@prod 'uptime'
ssh -L 8080:localhost:80 user@host
```

پیکربندی در `~/.ssh/config`:

```
Host myserver
  HostName 203.0.113.5
  User deploy
  Port 22
  IdentityFile ~/.ssh/id_ed25519
```

# Common mistakes

- مجوز باز برای کلید خصوصی (`chmod 600`).
- اشتباه `-p` (ssh) با `-P` (scp).

# Tips

- کلید: `ssh-keygen -t ed25519`
- کپی کلید: `ssh-copy-id user@host`

# Related commands

- `scp` / `sftp` / `rsync`
- `ssh-keygen`
- `tmux` روی سرور
