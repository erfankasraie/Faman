---
title: systemctl
aliases:
category: system
difficulty: intermediate
keywords:
- systemd
- service
- unit
- daemon
---

# Introduction

`systemctl` رابط اصلی **systemd** است: شروع/توقف سرویس، enable هنگام بوت، وضعیت unitها، و کنترل ماشین (reboot).

# Syntax

```
systemctl [OPTIONS] COMMAND [UNIT...]
```

واحدها معمولاً پسوند دارند: `.service`, `.socket`, `.timer`, `.target`.

# Options

## فرمان‌های سرویس

| فرمان | کار |
|--------|-----|
| `start UNIT` | شروع |
| `stop UNIT` | توقف |
| `restart UNIT` | راه‌اندازی مجدد |
| `reload UNIT` | بارگذاری دوبارهٔ کانفیگ (اگر پشتیبانی شود) |
| `status UNIT` | وضعیت و آخرین لاگ |
| `enable UNIT` | استارت در بوت |
| `disable UNIT` | خلاف enable |
| `is-active UNIT` | کد خروج فعال بودن |
| `is-enabled UNIT` | آیا enable است |
| `mask UNIT` | جلوگیری کامل از start |
| `unmask UNIT` | لغو mask |
| `daemon-reload` | بعد از تغییر unit فایل |
| `list-units --type=service` | سرویس‌های بارشده |
| `list-unit-files` | همهٔ unit فایل‌ها |
| `cat UNIT` | مسیر و محتوای unit |
| `edit UNIT` | override |
| `reboot` / `poweroff` | خاموش/ری‌استارت سیستم |

## گزینه‌های سراسری

| گزینه | توضیح |
|-------|--------|
| `--user` | systemd کاربر (نه سیستم) |
| `-l` | خروجی کامل (نه truncate) |
| `--no-pager` | بدون less |
| `--failed` | فقط failed |

# Examples

## روزمره

```bash
sudo systemctl status nginx
sudo systemctl restart nginx
sudo systemctl enable --now nginx   # enable + start
sudo systemctl disable --now nginx
```

## عیب‌یابی

```bash
systemctl --failed
systemctl status ssh
journalctl -u ssh -e
systemctl cat nginx.service
```

## بعد از نوشتن unit جدید

```bash
sudo cp myapp.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now myapp
```

## سطح کاربر

```bash
systemctl --user status podman.socket
systemctl --user enable --now syncthing
```

# Common mistakes

- `start` بدون `enable` → بعد از reboot سرویس بالا نمی‌آید.
- ویرایش unit بدون `daemon-reload`.
- `mask` کردن و فراموش کردن `unmask`.
- اشتباه گرفتن نام: `nginx` در برابر `nginx.service` (معمولاً هر دو ok).

# Tips

- برای لاگ همیشه با `journalctl -u` جفت کنید.
- `systemctl list-timers` برای کارهایی که قبلاً cron بودند.

# Related commands

- `journalctl` · `hostnamectl` · `timedatectl` · `loginctl` · `service` (قدیمی)
