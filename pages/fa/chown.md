---
title: chown
aliases:
category: permissions
difficulty: intermediate
keywords:
- owner
- ownership
- user
- group
---

# Introduction

دستور `chown` (change owner) مالک و/یا گروه یک فایل یا پوشه را تغییر می‌دهد.

معمولاً فقط کاربر root می‌تواند مالک را تغییر دهد.

# Syntax

```
chown [OPTIONS] OWNER[:GROUP] FILE...
chown [OPTIONS] :GROUP FILE...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-R` | بازگشتی |
| `-v` | verbose |
| `-h` | روی symbolic link تأثیر بگذارد نه هدف آن |
| `--reference=RFILE` | کپی مالک از فایل دیگر |

# Examples

```bash
# تغییر مالک
sudo chown alice file.txt

# تغییر مالک و گروه
sudo chown alice:developers project/

# فقط تغییر گروه
sudo chown :developers file.txt

# بازگشتی
sudo chown -R www-data:www-data /var/www/html
```

# Common mistakes

- فراموش کردن `sudo` وقتی مجوز کافی ندارید.
- تغییر مالک فایل‌های سیستمی بدون دقت.

# Tips

- بعد از `chown` معمولاً `chmod` هم لازم است.
- برای وب‌سرورها اغلب مالک را به کاربر وب‌سرور (`www-data` یا `nginx`) تغییر می‌دهند.
- از `--from` برای تغییر فقط اگر مالک فعلی خاص باشد استفاده کنید.

# Related commands

- `chmod` — تغییر مجوز
- `chgrp` — فقط تغییر گروه
- `ls -l` — مشاهده مالک
- `id` — دیدن uid/gid کاربر
