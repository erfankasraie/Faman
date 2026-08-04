---
title: crontab
aliases:
- cron
category: system
difficulty: intermediate
keywords:
- schedule
- job
- timer
- automation
---

# Introduction

`crontab` جدول زمان‌بندی **cron** کاربر را ویرایش/نمایش می‌دهد: اجرای دورهٔ فرمان در پس‌زمینه.

روی سیستم‌های systemd گاهی `systemd.timer` جایگزین مدرن است؛ cron هنوز بسیار رایج است.

# Syntax

```
crontab [-u USER] FILE
crontab [-u USER] [-e | -l | -r]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-e` | ویرایش جدول در ادیتور |
| `-l` | چاپ جدول |
| `-r` | حذف کل جدول (خطرناک) |
| `-u USER` | برای کاربر دیگر (نیاز root) |
| `-i` | تأیید قبل از `-r` |

## فرمت هر خط

```
┌─ دقیقه (0-59)
│ ┌─ ساعت (0-23)
│ │ ┌─ روز ماه (1-31)
│ │ │ ┌─ ماه (1-12)
│ │ │ │ ┌─ روز هفته (0-7، 0 و 7 = یکشنبه)
│ │ │ │ │
* * * * *  command
```

| مثال زمان | معنی |
|-----------|------|
| `0 * * * *` | هر ساعت سر صفر |
| `*/5 * * * *` | هر ۵ دقیقه |
| `0 9 * * 1-5` | ۹ صبح روزهای کاری |
| `0 0 1 * *` | اول هر ماه |
| `@daily` | معادل روزانه (نسخه‌های Vixie/cronie) |
| `@reboot` | بعد از بوت |

متغیرهای رایج در crontab: `PATH=`, `MAILTO=`, `SHELL=`.

# Examples

## ویرایش و لیست

```bash
crontab -l
crontab -e
```

نمونه خطوط:

```cron
PATH=/usr/local/bin:/usr/bin:/bin
MAILTO=""

*/10 * * * * /home/alice/bin/backup.sh >> /home/alice/logs/backup.log 2>&1
0 3 * * * /usr/bin/find /tmp -type f -mtime +7 -delete
@reboot /home/alice/bin/onboot.sh
```

## نصب از فایل

```bash
cat > /tmp/my.cron <<'EOF'
0 6 * * * /usr/local/bin/sync.sh
EOF
crontab /tmp/my.cron
crontab -l
```

## سیستم‌واید

```bash
# معمولاً:
ls /etc/cron.d/ /etc/cron.daily/
# یا /etc/crontab با فیلد user اضافه
```

# Common mistakes

- `%` در فرمان cron معنی خاص دارد؛ باید escape شود `\%`.
- PATH کوتاه cron → همیشه مسیر مطلق باینری.
- `crontab -r` به‌جای `-e` و پاک شدن همه jobها.
- فرض اینکه خروجی را می‌بینید؛ بدون redirect در ایمیل root می‌رود یا دور ریخته می‌شود.

# Tips

- لاگ را صریح به فایل بفرستید: `>> log 2>&1`.
- اول با `*/5` تست کنید بعد بازه را باز کنید.
- جایگزین systemd: `systemctl edit --full my.timer`.

# Related commands

- `systemctl` · `journalctl` · `at` · `systemd-run`
