---
title: base64
aliases:
category: text
difficulty: beginner
keywords:
- encode
- decode
- binary
---

# Introduction

`base64` داده باینری را به متن ASCII و برعکس تبدیل می‌کند.

# Syntax

```
base64 [OPTIONS] [FILE]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-d` / `--decode` | رمزگشایی |
| `-w 0` | بدون شکست خط |

# Examples

```bash
echo -n 'hello' | base64
echo 'aGVsbG8=' | base64 -d
base64 -w 0 file.bin > file.b64
```

# Common mistakes

- فراموش کردن `-n` در echo و اضافه شدن newline به داده.

# Tips

- برای فایل‌های بزرگ پایپ به فایل خروجی بدهید.

# Related commands

- `xxd` / `od` — نمایش هگز
- `openssl enc` — رمزنگاری
