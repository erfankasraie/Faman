---
title: timeout
aliases:
category: process
difficulty: intermediate
keywords:
- limit
- kill
- duration
---

# Introduction

`timeout` یک فرمان را حداکثر برای مدت مشخص اجرا می‌کند و در صورت طولانی شدن، سیگنال می‌فرستد.

# Syntax

```
timeout [OPTIONS] DURATION COMMAND [ARG]...
```

DURATION مثل `10s`, `5m`, `1h`.

# Options

| گزینه | توضیح |
|-------|--------|
| `-s SIG` | سیگنال (پیش‌فرض TERM) |
| `-k TIME` | اگر بعد از SIG زنده ماند، بعد از TIME سیگنال KILL |
| `-v` | verbose |
| `--preserve-status` | کد خروج فرمان اصلی |

# Examples

```bash
timeout 10s ping 8.8.8.8
timeout -k 5s 30s long-job.sh
timeout 5m make test
```

# Common mistakes

- واحد زمان را ننوشتن در بعضی نسخه‌های قدیمی.

# Tips

- در اسکریپت شبکه برای جلوگیری از hang عالی است.

# Related commands

- `kill` · `time` · `nice`
