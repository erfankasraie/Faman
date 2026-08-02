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
---

# Introduction

`crontab` زمان‌بندی اجرای دستورات را برای کاربر مدیریت می‌کند (سرویس cron).

# Syntax

```
crontab [-e|-l|-r]
```

قالب هر خط:

```
minute hour day month weekday command
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-e` | ویرایش crontab |
| `-l` | نمایش |
| `-r` | حذف همه |
| `-u USER` | برای کاربر دیگر (root) |

# Examples

```bash
crontab -e
crontab -l

# هر روز ساعت ۰۳:۳۰
30 3 * * * /usr/local/bin/backup.sh

# هر ۵ دقیقه
*/5 * * * * /usr/bin/php /var/www/cron.php

# دوشنبه‌ها ساعت ۹
0 9 * * 1 /home/user/scripts/report.sh
```

# Common mistakes

- استفاده از دستور نسبی بدون PATH کامل.
- ننوشتن لاگ و نفهمیدن چرا job اجرا نشده.

# Tips

- خروجی را redirect کنید: `>> /var/log/job.log 2>&1`
- در systemd می‌توان از timer unit هم استفاده کرد.

# Related commands

- `systemctl` — timerهای systemd
- `at` — اجرای یک‌باره
- `journalctl` — لاگ
