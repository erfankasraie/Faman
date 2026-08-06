---
title: chgrp
aliases:
category: permissions
difficulty: intermediate
keywords:
- group
- ownership
- permissions
---

# Introduction

`chgrp` گروه مالک (owner group) یک فایل یا پوشه را تغییر می‌دهد؛ برای مدیریت دسترسی گروهی روی سیستم‌های چندکاربره یا سرورهای اشتراکی کاربرد دارد.

# Syntax

```
chgrp [OPTIONS] GROUP FILE...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-R` | اعمال بازگشتی روی پوشه و محتویاتش |
| `-v` | نمایش جزئیات هر تغییر |
| `--reference=FILE` | کپی‌کردن گروه از یک فایل دیگر |

# Examples

```bash
# تغییر گروه یک فایل
sudo chgrp developers project.txt

# تغییر گروه یک پوشه به‌صورت بازگشتی
sudo chgrp -R developers /var/www/myapp

# کپی‌کردن گروه از یک فایل دیگر
sudo chgrp --reference=template.txt newfile.txt

# نمایش جزئیات هنگام تغییر
sudo chgrp -v developers *.log
```

# Common mistakes

- فراموش‌کردن این‌که کاربر باید عضو گروه مقصد باشد (یا root باشد) تا اجازه تغییر گروه فایل به آن را داشته باشد.
- خلط‌کردن `chgrp` (فقط گروه) با `chown` (که هم مالک و هم گروه را تغییر می‌دهد و می‌تواند جایگزین کامل‌تری باشد: `chown user:group file`).

# Tips

- به‌جای اجرای جداگانه `chown` و `chgrp`، می‌توانید هر دو را یک‌جا با `chown user:group file` انجام دهید.
- برای پوشه‌های اشتراکی تیمی، ترکیب `chgrp -R` با `chmod g+s` (setgid) باعث می‌شود فایل‌های جدید هم به‌صورت خودکار همان گروه را بگیرند.

# Related commands

- `chown` — تغییر مالک (و در صورت نیاز گروه هم)
- `chmod` — تغییر مجوز دسترسی
- `groups` — نمایش گروه‌های یک کاربر
