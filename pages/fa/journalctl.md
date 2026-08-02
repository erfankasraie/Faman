---
title: journalctl
aliases:
category: system
difficulty: intermediate
keywords:
- logs
- systemd
- journal
---

# Introduction

`journalctl` لاگ‌های جمع‌آوری‌شده توسط systemd-journald را نمایش و فیلتر می‌کند.

# Syntax

```
journalctl [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-u UNIT` | فقط یک سرویس |
| `-f` | follow (مثل tail -f) |
| `-n N` | آخرین N خط |
| `-b` | بوت فعلی |
| `-p PRIORITY` | سطح اولویت |
| `--since` / `--until` | بازه زمانی |
| `-xe` | توضیح خطاهای اخیر |

# Examples

```bash
journalctl -u nginx -f
journalctl -b
journalctl -p err -b
journalctl --since "1 hour ago"
journalctl -xe
```

# Common mistakes

- جستجو فقط در `/var/log` در حالی که سرویس لاگ را به journal می‌فرستد.

# Tips

- ترکیب با `systemctl status` برای عیب‌یابی سرویس.

# Related commands

- `systemctl` — کنترل سرویس
- `dmesg` — لاگ کرنل
