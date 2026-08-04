---
title: rsync
aliases:
category: archive
difficulty: intermediate
keywords:
- sync
- backup
- copy
- remote
---

# Introduction

`rsync` فایل‌ها را **همگام** می‌کند: فقط تفاوت‌ها را منتقل می‌کند، مجوز/زمان را می‌تواند حفظ کند، روی SSH کار می‌کند و برای بکاپ و استقرار عالی است.

# Syntax

```
rsync [OPTIONS] SRC... DEST
```

SRC و DEST می‌توانند محلی یا `user@host:path` باشند.

**اسلش انتها مهم است:**
- `rsync -a src/ dest/` → *محتوای* src داخل dest
- `rsync -a src dest/` → پوشهٔ src به‌عنوان زیرپوشه داخل dest

# Options

## پایه و آرشیو

| گزینه | توضیح |
|-------|--------|
| `-a` | archive = `-rlptgoD` (recursive، symlink، مجوز، زمان، گروه، مالک، device) |
| `-r` | بازگشتی |
| `-l` | symlink به‌صورت لینک |
| `-p` | مجوزها |
| `-t` | زمان تغییر |
| `-g` / `-o` | گروه / مالک (معمولاً نیاز به root) |
| `-D` | device و special |
| `-L` | symlink را resolve و کپی محتوا |
| `-H` | hardlink حفظ شود |
| `-A` | ACL |
| `-X` | xattr |

## نمایش و ایمنی

| گزینه | توضیح |
|-------|--------|
| `-v` | verbose |
| `-h` | اندازهٔ خوانا |
| `-P` | `--partial --progress` |
| `-n` / `--dry-run` | شبیه‌سازی |
| `-i` | آیتم‌سازی تغییرات |
| `--delete` | حذف در DEST آنچه در SRC نیست |
| `--delete-excluded` | حذف excludeشده‌ها در dest هم |
| `--remove-source-files` | بعد از کپی موفق از مبدأ پاک کن |
| `-c` | مقایسه با checksum نه فقط size+mtime |
| `--checksum` | همان |

## شبکه و فیلتر

| گزینه | توضیح |
|-------|--------|
| `-z` | فشرده‌سازی حین انتقال |
| `-e ssh` | شل راه دور |
| `--rsh=ssh` | همان |
| `-S` | فایل‌های sparse |
| `--exclude=PAT` | حذف الگو |
| `--exclude-from=FILE` | الگوها از فایل |
| `--include=PAT` | استثنا روی exclude |
| `--max-size=N` / `--min-size=N` | محدودیت اندازه |
| `--bwlimit=RATE` | محدودیت پهنای باند |
| `--timeout=SEC` | تایم‌اوت IO |

# Examples

## محلی

```bash
# بکاپ با حفظ ویژگی‌ها
rsync -avh /home/alice/ /backup/alice/

# dry-run قبل از delete
rsync -avhn --delete /home/alice/ /backup/alice/
rsync -avh --delete /home/alice/ /backup/alice/
```

## از طریق SSH

```bash
rsync -avzP -e ssh ./project/ user@server:/var/www/project/
rsync -avzP user@server:/var/log/nginx/ ./logs/

# پورت غیر۲۲
rsync -avzP -e 'ssh -p 2222' ./ user@host:/data/
```

## فیلتر

```bash
rsync -av --exclude '.git' --exclude 'node_modules' --exclude '*.tmp' ./src/ ./deploy/

# فقط فایل‌های .c
rsync -av --include '*/' --include '*.c' --exclude '*' ./src/ ./c-only/
```

## محدودیت و ازسرگیری

```bash
rsync -avP --bwlimit=5000 big/ user@host:big/
# -P کمک می‌کند انتقال ناقص از سر گرفته شود
```

# Common mistakes

- اشتباه `src` در برابر `src/` و کپی شدن یک لایه پوشهٔ اضافه.
- `--delete` بدون `-n` روی مسیر اشتباه → پاک شدن داده.
- فراموش کردن اینکه `-a` مالک را هم می‌خواهد کپی کند و روی سرور بدون دسترسی fail می‌شود (`--no-owner --no-group`).

# Tips

- عادت طلایی: همیشه یک‌بار `-n`.
- برای mirror: `-a --delete` با excludeهای صریح.
- لاگ: `--log-file=rsync.log`.

# Related commands

- `scp` — کپی ساده یک‌باره
- `cp -a` — کپی محلی
- `tar` — آرشیو یک‌جا
- `rclone` — ابر و S3
