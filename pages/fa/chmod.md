---
title: chmod
aliases:
category: permissions
difficulty: beginner
keywords:
- permission
- mode
- security
- acl
---

# Introduction

`chmod` مجوز دسترسی فایل/پوشه را عوض می‌کند: خواندن (r)، نوشتن (w)، اجرا (x) برای owner، group، others.

# Syntax

```
chmod [OPTIONS] MODE[,MODE]... FILE...
chmod [OPTIONS] OCTAL-MODE FILE...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-R` | بازگشتی روی درخت |
| `-v` | verbose |
| `-c` | فقط وقتی تغییر واقعی شد گزارش بده |
| `--reference=RFILE` | کپی mode از فایل مرجع |
| `-h` | روی symlink (در برخی سیستم‌ها محدود) |

## حالت نمادین

| جزء | معنی |
|------|------|
| `u` `g` `o` `a` | user / group / others / all |
| `+` `-` `=` | افزودن / برداشتن / تنظیم دقیق |
| `r` `w` `x` | خواندن / نوشتن / اجرا |
| `X` | اجرا فقط اگر روی دایرکتوری یا قبلاً x داشته |
| `s` | setuid/setgid |
| `t` | sticky (مثلاً `/tmp`) |

## حالت عددی (octal)

| رقم | معنی (برای یک کلاس) |
|------|---------------------|
| 4 | r |
| 2 | w |
| 1 | x |
| جمع | مثلاً 7 = rwx، 5 = r-x، 6 = rw- |

ترتیب ارقام: **u g o** → `755` = u=rwx, g=rx, o=rx  
پیشوند خاص: `4xxx` setuid، `2xxx` setgid، `1xxx` sticky → مثلاً `4755`.

# Examples

## روزمره

```bash
chmod 644 file.txt          # rw-r--r--
chmod 755 script.sh         # rwxr-xr-x
chmod u+x run.sh            # فقط برای owner اجرا
chmod go-w secret.env       # گروه و others ننویسند
chmod -R u+rwX project/     # پوشه‌ها قابل عبور
```

## اسکریپت و پوشهٔ وب

```bash
find project -type f -exec chmod 644 {} +
find project -type d -exec chmod 755 {} +
chmod 600 ~/.ssh/id_ed25519
chmod 700 ~/.ssh
```

## sticky و اشتراک

```bash
chmod 1777 /tmp/shared-drop   # sticky bit
# یا
chmod a+rwxt /tmp/shared-drop
```

## دیدن مجوز

```bash
ls -l file.txt
stat -c '%a %n' file.txt    # عددی
```

# Common mistakes

- `chmod -R 777` روی درخت پروژه یا `/` — خطر امنیتی.
- `chmod +x` روی فایل دادهٔ تصادفی.
- انتظار که chmod مالک را عوض کند → آن کار `chown` است.
- روی فایل‌سیستم‌های خاص (بعضی mountهای ویندوز/FAT) مجوز POSIX کامل نیست.

# Tips

- ACL ریزدانه‌تر: `setfacl` / `getfacl` (وقتی گروه کافی نیست).
- umask پیش‌فرض مجوز فایل‌های جدید را محدود می‌کند (`umask`).

# Related commands

- `chown` · `chgrp` · `ls -l` · `umask` · `setfacl` · `stat`
