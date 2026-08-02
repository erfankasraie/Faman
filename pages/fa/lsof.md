---
title: lsof
aliases:
category: system
difficulty: intermediate
keywords:
- open-files
- port
- process
---

# Introduction

`lsof` (list open files) نشان می‌دهد کدام فرایند چه فایل‌ها، سوکت‌ها یا پورت‌هایی را باز کرده است.

# Syntax

```
lsof [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-i` | فایل‌های شبکه |
| `-i :PORT` | پورت خاص |
| `-p PID` | فقط یک فرایند |
| `-u USER` | یک کاربر |
| `+D DIR` | فایل‌های باز زیر یک پوشه |

# Examples

```bash
lsof -i :80
lsof -p 1234
sudo lsof -iTCP -sTCP:LISTEN
lsof +f -- /mnt/data
```

# Common mistakes

- بدون sudo بعضی سوکت‌های سیستم دیده نمی‌شوند.

# Tips

- عیب‌یابی «device busy» هنگام umount.

# Related commands

- `ss` / `netstat` — سوکت‌ها
- `fuser` — فرایند استفاده‌کننده از فایل
- `ps`
