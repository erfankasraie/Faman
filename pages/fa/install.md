---
title: install
aliases:
category: filesystem
difficulty: intermediate
keywords:
- copy
- permissions
- deploy
- build
---

# Introduction

`install` فایل‌ها را کپی می‌کند و همزمان مالکیت، مجوز (permission) و پوشه مقصد را هم تنظیم می‌کند؛ معمولاً در اسکریپت‌های نصب (`make install`) برای قراردادن باینری‌ها و فایل‌های تنظیمات در جای درست استفاده می‌شود.

# Syntax

```
install [OPTIONS] SOURCE DEST
install [OPTIONS] SOURCE... DIRECTORY
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-m MODE` | تعیین مجوز فایل مقصد (مثل `755`) |
| `-o OWNER` | تعیین مالک فایل |
| `-g GROUP` | تعیین گروه فایل |
| `-d` | ساخت پوشه‌ها به‌جای کپی فایل |
| `-D` | ساخت پوشه‌های والد در صورت نیاز |

# Examples

```bash
# کپی یک باینری با مجوز اجرایی به مسیر سیستم
sudo install -m 755 myprogram /usr/local/bin/

# ساخت یک پوشه با مجوز مشخص
install -d -m 755 /opt/myapp

# کپی چند فایل به یک پوشه با مالک و گروه مشخص
sudo install -o root -g root -m 644 config.conf /etc/myapp/

# کپی با ساخت خودکار پوشه‌های والد در مسیر مقصد
install -D myfile.txt /some/deep/path/myfile.txt
```

# Common mistakes

- استفاده از `cp` به‌جای `install` در اسکریپت‌های نصب که باعث می‌شود مجوزها به‌صورت دستی و جداگانه تنظیم شوند؛ `install` این کار را در یک دستور انجام می‌دهد.
- فراموش‌کردن `sudo` هنگام نصب در مسیرهای سیستمی مثل `/usr/local/bin`.

# Tips

- در فایل‌های `Makefile`، دستور `install` استاندارد صنعتی برای مرحله `make install` است.
- `-D` بسیار مفید است وقتی مسیر مقصد هنوز وجود ندارد و نیاز به ساخت خودکار پوشه‌ها دارید.

# Related commands

- `cp` — کپی ساده بدون تنظیم مجوز/مالک
- `chmod` / `chown` — تنظیم جداگانه مجوز و مالکیت
- `make` — اجرای Makefile که معمولاً از install استفاده می‌کند
