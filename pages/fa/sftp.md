---
title: sftp
aliases:
category: network
difficulty: intermediate
keywords:
- ssh
- transfer
- ftp
---

# Introduction

`sftp` انتقال فایل تعاملی روی **SSH** است (نه FTP کلاسیک). امن‌تر از FTP خام.

# Syntax

```
sftp [OPTIONS] [user@]host
sftp -b batchfile [user@]host
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-P PORT` | پورت SSH |
| `-i KEY` | کلید |
| `-b FILE` | حالت batch |
| `-R` | recursively در بعضی نسخه‌ها برای get/put |

دستورات داخلی: `ls` `cd` `get` `put` `mkdir` `rm` `bye`.

# Examples

```bash
sftp alice@server
# get report.pdf
# put -r ./web public_html
sftp -P 2222 -i ~/.ssh/id_ed25519 alice@server
```

# Common mistakes

- قاطی با `ftp` بدون رمزنگاری.
- `-P` بزرگ برای پورت (مثل scp).

# Tips

- برای همگام‌سازی: `rsync -e ssh`.
- `scp` برای کپی یک‌خطی.

# Related commands

- `scp` · `ssh` · `rsync`
