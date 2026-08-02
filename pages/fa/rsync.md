---
title: rsync
aliases:
category: filesystem
difficulty: intermediate
keywords:
- sync
- copy
- backup
- remote
---

# Introduction

`rsync` فایل‌ها را همگام‌سازی می‌کند؛ فقط تفاوت‌ها را منتقل می‌کند و برای بکاپ و کپی شبکه عالی است.

# Syntax

```
rsync [OPTIONS] SRC... DEST
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-a` | archive (مجوز، زمان، recursive، symlink) |
| `-v` | verbose |
| `-z` | فشرده‌سازی هنگام انتقال |
| `-h` | اندازه‌های خوانا |
| `-P` | پیشرفت + ادامه دانلود ناتمام |
| `-n` | dry-run (شبیه‌سازی) |
| `--delete` | حذف در مقصد اگر در مبدأ نیست |
| `-e ssh` | انتقال روی SSH |

# Examples

```bash
# کپی محلی
rsync -avh /src/ /backup/src/

# dry-run
rsync -avhn --delete /src/ /backup/src/

# به سرور
rsync -avzP -e ssh ./project/ user@host:/var/www/project/

# از سرور
rsync -avzP user@host:/var/log/app.log ./
```

# Common mistakes

- اسلش انتهای مسیر مهم است: `src` در مقابل `src/`.
- `--delete` بدون dry-run می‌تواند داده پاک کند.

# Tips

- همیشه اول `-n` بزنید.
- برای بکاپ: `-a` معمولاً پایه خوبی است.

# Related commands

- `scp` — کپی ساده روی SSH
- `cp` — کپی محلی
- `tar` — آرشیو
