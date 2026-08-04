---
title: scp
aliases:
category: network
difficulty: beginner
keywords:
- copy
- remote
- ssh
- transfer
---

# Introduction

`scp` (*secure copy*) فایل/پوشه را روی کانال **SSH** کپی می‌کند. سینتکس شبیه `cp` است با `user@host:path` برای سمت remote.

برای همگام‌سازی تکراری و resume، معمولاً `rsync -e ssh` بهتر است؛ scp برای کپی یک‌باره ساده کافی است.

# Syntax

```
scp [OPTIONS] SOURCE... DESTINATION
```

SOURCE یا DEST می‌تواند محلی یا `user@host:path` باشد.

# Options

| گزینه | توضیح |
|-------|--------|
| `-r` | کپی بازگشتی پوشه |
| `-p` | حفظ زمان و mode تا حد ممکن |
| `-P PORT` | پورت SSH (**حرف بزرگ P**؛ با ssh فرق دارد) |
| `-i IDENTITY` | کلید خصوصی |
| `-C` | فشرده‌سازی |
| `-q` | ساکت |
| `-v` | verbose |
| `-3` | کپی بین دو remote از طریق ماشین محلی |
| `-o OPTION` | گزینه‌های ssh (مثل `ProxyJump`) |
| `-F CONFIG` | فایل ssh config |

# Examples

## به سرور / از سرور

```bash
scp report.pdf alice@server:/home/alice/
scp alice@server:/var/log/app.log ./
scp -r ./website alice@server:/var/www/site/
```

## پورت و کلید

```bash
scp -P 2222 -i ~/.ssh/id_ed25519 app.tgz alice@server:~/uploads/
scp -o ProxyJump=bastion ./file alice@internal:/tmp/
```

## بین دو مسیر remote (از طریق لپ‌تاپ)

```bash
scp -3 alice@hostA:/data/out.bin bob@hostB:/data/in.bin
```

## چند فایل

```bash
scp a.txt b.txt alice@server:~/inbox/
```

# Common mistakes

- `-p` کوچک برای پورت → اشتباه؛ پورت scp **`-P`** است.
- فراموش کردن `-r` برای پوشه.
- فاصله در مسیر remote بدون کوتیشن.
- انتظار resume بعد از قطع شبکه (scp از اول شروع می‌کند).

# Tips

- همان `~/.ssh/config` برای Host کوتاه کار می‌کند: `scp file myserver:~/`.
- برای درخت بزرگ: `rsync -avzP -e ssh`.
- SFTP تعاملی: دستور `sftp`.

# Related commands

- `ssh` · `rsync` · `sftp` · `tar` · `curl`
