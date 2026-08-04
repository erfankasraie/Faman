---
title: journalctl
aliases:
category: system
difficulty: intermediate
keywords:
- systemd
- log
- journal
- debug
---

# Introduction

`journalctl` لاگ‌های جمع‌آوری‌شده توسط **systemd-journald** را می‌خواند: سرویس‌ها، کرنل، بوت.

# Syntax

```
journalctl [OPTIONS] [MATCHES]
```

# Options

## فیلتر منبع

| گزینه | توضیح |
|-------|--------|
| `-u UNIT` | فقط این سرویس |
| `-t SYSLOG_ID` | شناسه syslog |
| `-k` / `--dmesg` | پیام‌های کرنل |
| `_PID=123` | مطابقت فیلد |
| `PRIORITY=3` | سطح (0 emerg … 7 debug) |

## زمان و بوت

| گزینه | توضیح |
|-------|--------|
| `-b` | بوت فعلی |
| `-b -1` | بوت قبلی |
| `-S` / `--since` | از زمان |
| `-U` / `--until` | تا زمان |
| `--list-boots` | فهرست بوت‌ها |

## خروجی

| گزینه | توضیح |
|-------|--------|
| `-f` | follow (مثل tail -f) |
| `-n N` | N خط آخر |
| `-e` | پرش به انتها |
| `-o json` / `json-pretty` / `cat` | قالب |
| `-p err` | اولویت err و بالاتر |
| `--no-pager` | بدون pager |
| `-x` | توضیح catalog |

# Examples

## سرویس

```bash
journalctl -u nginx -e
journalctl -u nginx -f
journalctl -u ssh -S "1 hour ago"
journalctl -u docker -p err --no-pager
```

## سیستم و کرنل

```bash
journalctl -b
journalctl -k -b
journalctl --list-boots
journalctl -b -1 -p warning
```

## بازهٔ زمانی

```bash
journalctl --since "2026-08-04 10:00:00" --until "2026-08-04 12:00:00"
journalctl -S yesterday
journalctl -S "-30m"
```

## JSON برای پردازش

```bash
journalctl -u myapp -o json-pretty -n 20
journalctl -u myapp -o cat
```

# Common mistakes

- بدون دسترسی: گاهی نیاز به عضویت در گروه `systemd-journal` یا `sudo`.
- ژورنال persistent نیست اگر `/var/log/journal` نباشد (بسته به distro).
- `-f` روی سیستم شلوغ بدون فیلتر `-u`.

# Tips

- عادت خوب: `systemctl status` + `journalctl -u ... -e`.
- سقف دیسک journal: `SystemMaxUse` در `journald.conf`.

# Related commands

- `systemctl` · `dmesg` · `tail` · `logger`
