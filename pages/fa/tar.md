---
title: tar
aliases:
category: archive
difficulty: intermediate
keywords:
- archive
- compress
- backup
- tarball
---

# Introduction

`tar` (*tape archive*) فایل‌ها را در یک **آرشیو** جمع می‌کند. امروز معمولاً با فشرده‌سازی gzip/xz/zstd همراه است (`tar.gz`, `tar.xz`). یک ابزار برای بسته‌بندی، بکاپ و انتقال درخت پروژه است.

# Syntax

```
tar [OPTIONS] [ARCHIVE] [FILE...]
```

سبک قدیمی حروف بدون `-` هم رایج است: `tar czf a.tgz dir/`.

# Options

## عملیات اصلی (یکی را انتخاب کنید)

| گزینه | توضیح |
|-------|--------|
| `-c` / `--create` | ساخت آرشیو |
| `-x` / `--extract` | استخراج |
| `-t` / `--list` | فهرست محتوا |
| `-r` / `--append` | افزودن به آرشیو بدون فشرده |
| `-u` / `--update` | افزودن فقط جدیدترها |
| `-d` / `--diff` | مقایسه با فایل‌سیستم |

## فایل آرشیو و فشرده‌سازی

| گزینه | توضیح |
|-------|--------|
| `-f FILE` | نام آرشیو (معمولاً لازم) |
| `-z` | gzip |
| `-J` | xz |
| `-j` | bzip2 |
| `--zstd` | zstd (GNU tar جدید) |
| `-a` | حدس فشرده‌سازی از پسوند |
| `-C DIR` | تغییر پوشه قبل از عمل |

## گزینش و نمایش

| گزینه | توضیح |
|-------|--------|
| `-v` | verbose |
| `-p` | حفظ مجوزها (استخراج) |
| `--exclude=PATTERN` | حذف از آرشیو |
| `--exclude-vcs` | نادیده گرفتن .git و … |
| `--strip-components=N` | حذف N جزء اول مسیر هنگام استخراج |
| `--numeric-owner` | uid/gid عددی |
| `-h` | دنبال symlink محتوا |
| `--overwrite` | بازنویسی |
| `--keep-old-files` | روی فایل موجود خطا بده نه overwrite |

# Examples

## ساخت

```bash
# tar.gz
tar -czvf project.tar.gz project/

# xz فشرده‌تر، کندتر
tar -cJvf project.tar.xz project/

# از روی پسوند
tar -cavf project.tar.zst project/

# بدون .git و node_modules
tar -czvf src.tar.gz --exclude='.git' --exclude='node_modules' src/
```

## فهرست و استخراج

```bash
tar -tzvf project.tar.gz
tar -xzvf project.tar.gz
tar -xzvf project.tar.gz -C /opt/

# فقط یک فایل از داخل آرشیو
tar -xzvf project.tar.gz project/README.md

# حذف یک سطح پوشهٔ بالایی
tar -xzvf project.tar.gz --strip-components=1 -C /dest/
```

## بکاپ و stdin/stdout

```bash
# آرشیو به stdout و فشرده‌سازی جدا
tar -cf - dir/ | gzip > dir.tar.gz

# استخراج از stdin
gzip -dc dir.tar.gz | tar -xvf -

# چند مسیر
tar -czvf etc-home.tar.gz /etc /home/alice
```

## remote با ssh

```bash
tar -czf - project/ | ssh user@host 'tar -xzf - -C /var/www/'
```

# Common mistakes

- فراموش کردن `-f` و تفسیر اشتباه آرگومان بعدی به‌عنوان آرشیو.
- استخراج به‌عنوان root روی `/` بدون نگاه به مسیرهای داخل tar (خطر path traversal در tarهای ناشناس).
- `tar -x` روی آرشیو gzip بدون `-z` (بعضی tarها خودکار تشخیص می‌دهند، تکیه نکنید).
- exclude بعد از مسیر فایل‌ها گاهی دیر اعمال می‌شود — exclude را قبل از فایل‌ها بگذارید.

# Tips

- همیشه یک‌بار `-t` قبل از استخراج آرشیو ناشناس.
- برای همگام‌سازی تکراری: `rsync` مناسب‌تر از tar مکرر است.
- مجوز و مالک: روی سیستم‌های مختلف `-p` و `--same-owner` را بشناسید.

# Related commands

- `gzip` / `xz` / `zstd` — فقط فشرده‌سازی
- `zip` / `unzip` — آرشیو رایج ویندوز
- `rsync` — همگام‌سازی
- `cpio` — آرشیو جایگزین
