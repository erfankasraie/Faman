---
title: source
aliases:
- dot
category: shell
difficulty: beginner
keywords:
- shell
- script
- environment
---

# Introduction

`source` (یا `.`) فایل را در شل **فعلی** اجرا می‌کند؛ متغیرها و توابع در نشست می‌مانند.

# Syntax

```
source FILE [ARGS]
. FILE [ARGS]
```

# Options

آرگومان‌های بعد از FILE به اسکریپت می‌رسند.

# Examples

```bash
source ~/.bashrc
. ./venv/bin/activate
source scripts/env.sh
```

# Common mistakes

- `bash file.sh` به‌جای `source` — آن‌وقت exportها در شل پدر نمی‌مانند.

# Tips

- برای activate محیط مجازی همیشه source لازم است.

# Related commands

- `export`
- `bash` / `zsh`
