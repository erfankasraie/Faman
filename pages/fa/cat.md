---
title: cat
aliases:
- concatenate
category: text
difficulty: beginner
keywords:
- view
- file
- content
- concatenate
---

# Introduction

`cat` محتوای فایل را چاپ می‌کند یا چند فایل را به هم می‌چسباند. برای فایل‌های کوتاه مناسب است؛ برای فایل بلند از `less` استفاده کنید.

# Syntax

```
cat [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-n` | شماره‌گذاری خطوط |
| `-b` | شماره‌گذاری خطوط غیرخالی |
| `-s` | فشردن خطوط خالی متوالی |
| `-A` | نمایش کاراکترهای خاص |
| `-E` | نمایش `$` انتهای خط |

# Examples

```bash
cat file.txt
cat part1.txt part2.txt > full.txt
cat -n script.sh
cat > newfile.txt    # پایان با Ctrl+D
cat >> file.txt      # append
```

# Common mistakes

- `cat` روی فایل خیلی بزرگ و پر شدن ترمینال.
- `cat file | less` به‌جای `less file`.

# Tips

- برای کد رنگی: `bat`.
- `tac` محتوا را از آخر به اول نشان می‌دهد.

# Related commands

- `less` / `more`
- `head` / `tail`
- `bat`
- `tee`
