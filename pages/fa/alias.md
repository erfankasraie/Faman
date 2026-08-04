---
title: alias
aliases:
category: shell
difficulty: beginner
keywords:
- shortcut
- shell
- customization
- bashrc
---

# Introduction

`alias` یک میان‌بر (نام مستعار) برای یک دستور طولانی یا پرتکرار می‌سازد؛ مثلاً به‌جای تایپ یک دستور طولانی هر بار، یک نام کوتاه تعریف می‌کنید.

# Syntax

```
alias NAME='COMMAND'
alias                    # نمایش تمام aliasهای فعلی
unalias NAME             # حذف یک alias
```

# Options

`alias` گزینه خط‌فرمانی خاصی ندارد؛ سینتکس آن خودِ تعریف `NAME=COMMAND` است. برای حذف از `unalias NAME` یا `unalias -a` (حذف همه) استفاده کنید.

# Examples

```bash
# ساخت یک alias ساده
alias ll='ls -la'

# alias با چند دستور ترکیبی
alias update='sudo apt update && sudo apt upgrade -y'

# alias برای جلوگیری از اشتباهات رایج (تأیید قبل از حذف)
alias rm='rm -i'

# نمایش تمام aliasهای فعال
alias

# حذف یک alias
unalias ll

# ذخیره دائمی: افزودن به فایل تنظیمات پوسته
echo "alias ll='ls -la'" >> ~/.bashrc
source ~/.bashrc
```

# Common mistakes

- تعریف alias مستقیم در ترمینال بدون افزودن به `~/.bashrc` (یا `~/.zshrc`)؛ در این حالت با بستن ترمینال، alias از بین می‌رود.
- فراموش‌کردن `source ~/.bashrc` بعد از ویرایش فایل تنظیمات؛ بدون آن، alias جدید تا باز‌کردن ترمینال جدید اعمال نمی‌شود.
- ساخت alias با نامی که یک دستور واقعی سیستم را بازنویسی می‌کند (مثل `alias ls='ls --color'`) بدون آگاهی، که می‌تواند رفتار غیرمنتظره در اسکریپت‌ها ایجاد کند (البته aliasها معمولاً در اسکریپت‌های غیرتعاملی فعال نیستند).

# Tips

- برای دیدن این‌که یک دستور alias است یا فایل واقعی: `type COMMAND`
- الگوی متداول: aliasهای شخصی را در یک فایل جدا (مثل `~/.bash_aliases`) نگه دارید و در `~/.bashrc` آن را source کنید تا مدیریت راحت‌تر باشد.
- alias فقط در پوسته‌های تعاملی کار می‌کند، نه در اسکریپت‌های اجراشده با `bash script.sh` (مگر تنظیم خاص انجام شود).

# Related commands

- `type` — بررسی این‌که یک نام، alias/function/فایل است
- `unalias` — حذف یک alias
- `function` — تعریف یک تابع کامل‌تر از alias (با پشتیبانی آرگومان)
