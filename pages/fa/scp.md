---
title: scp
aliases:
category: network
difficulty: intermediate
keywords:
- copy
- ssh
- remote
- transfer
---

# Introduction

`scp` (secure copy) فایل را روی SSH کپی می‌کند. برای همگام‌سازی پیشرفته‌تر `rsync` بهتر است.

# Syntax

```
scp [OPTIONS] SRC DEST
```

SRC/DEST می‌تواند `user@host:path` باشد.

# Options

| گزینه | توضیح |
|-------|--------|
| `-r` | بازگشتی (پوشه) |
| `-P PORT` | پورت SSH |
| `-i KEY` | کلید خصوصی |
| `-C` | فشرده‌سازی |
| `-p` | حفظ زمان/مجوز |

# Examples

```bash
scp file.txt user@server:/tmp/
scp -r project/ user@server:~/app/
scp -P 2222 user@server:/var/log/app.log .
scp -i ~/.ssh/id_ed25519 file.tar.gz user@host:~
```

# Common mistakes

- `-P` بزرگ برای پورت (در ssh کوچک `-p` است).
- فراموش کردن `-r` برای پوشه.

# Tips

- برای resume و sync: `rsync -avz -e ssh`.

# Related commands

- `rsync`
- `sftp`
- `ssh`
