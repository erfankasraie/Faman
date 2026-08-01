---
title: git
aliases:
category: development
difficulty: intermediate
keywords:
- version-control
- source
- commit
- repository
---

# Introduction

`git` سیستم کنترل نسخه توزیع‌شده است که تقریباً استاندارد صنعت نرم‌افزار شده. این صفحه خلاصه‌ای از پرکاربردترین دستورات است.

# Syntax

```
git <command> [<args>]
```

# دستورات اصلی

## شروع کار

```bash
git init                  # ساخت مخزن جدید
git clone URL             # کلون کردن مخزن
```

## وضعیت و تاریخچه

```bash
git status                # وضعیت فایل‌ها
git log                   # تاریخچه کامیت‌ها
git log --oneline --graph # نمای فشرده
git diff                  # تغییرات stage‌نشده
git diff --staged         # تغییرات stage‌شده
```

## تغییرات

```bash
git add FILE              # اضافه به stage
git add .                 # همه تغییرات
git commit -m "message"   # ثبت کامیت
git commit -am "message"  # add + commit برای فایل‌های tracked
```

## شاخه‌ها

```bash
git branch                # لیست شاخه‌ها
git branch NAME           # ساخت شاخه
git checkout NAME         # جابه‌جایی
git switch NAME           # جابه‌جایی (جدیدتر)
git merge BRANCH          # ادغام
git branch -d NAME        # حذف شاخه
```

## ارتباط با ریموت

```bash
git remote -v             # لیست ریموت‌ها
git fetch                 # دریافت تغییرات بدون ادغام
git pull                  # fetch + merge
git push                  # ارسال به ریموت
git push -u origin BRANCH # اولین push با upstream
```

# Examples

```bash
# گردش کار روزانه
git status
git add .
git commit -m "Fix login bug"
git push

# ساخت و سوییچ به شاخه جدید
git switch -c feature/login

# بازگردانی تغییرات stage‌نشده
git restore file.txt
```

# Common mistakes

- `git add .` بدون بررسی `git status`.
- کامیت روی `main`/`master` به جای شاخه feature.
- force push روی شاخه‌های اشتراکی.
- فراموش کردن `.gitignore`.

# Tips

- پیام کامیت را واضح و به زبان حال بنویسید.
- از `git stash` برای کنار گذاشتن موقت تغییرات استفاده کنید.
- `git config --global` برای تنظیمات دائمی.
- ابزارهایی مثل `lazygit` یا `tig` تجربه بهتری می‌دهند.
- همیشه قبل از push مهم `git log` و `git diff` را چک کنید.

# Related commands

- `gh` — GitHub CLI
- `gitk` / `git gui`
- `delta` — diff زیباتر
- `pre-commit` — هوک‌های کیفیت کد
