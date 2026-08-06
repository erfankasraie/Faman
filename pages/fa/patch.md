---
title: patch
aliases:
category: development
difficulty: advanced
keywords:
- diff
- version
- code
- apply
---

# Introduction

`patch` تغییرات توصیف‌شده در یک فایل diff (که معمولاً با `diff -u` یا `git diff` ساخته شده) را روی فایل‌های اصلی اعمال می‌کند؛ روشی سنتی برای توزیع و اعمال تغییرات کد بدون نیاز به کل فایل‌های تغییریافته.

# Syntax

```
patch [OPTIONS] < patchfile
patch TARGET_FILE < patchfile
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-p N` | حذف N سطح از مسیر فایل در هدر patch (رایج: `-p1`) |
| `-R` | معکوس‌کردن (revert) یک patch که قبلاً اعمال شده |
| `--dry-run` | فقط بررسی این‌که آیا patch بدون خطا اعمال می‌شود، بدون تغییر واقعی |
| `-b` | ساخت نسخه پشتیبان از فایل اصلی قبل از تغییر |

# Examples

```bash
# اعمال یک patch ساده
patch < changes.patch

# اعمال patch با حذف یک سطح از مسیر (رایج‌ترین حالت برای patch های git)
patch -p1 < changes.patch

# بررسی این‌که آیا patch بدون خطا اعمال می‌شود (بدون تغییر واقعی)
patch --dry-run -p1 < changes.patch

# معکوس‌کردن (revert) یک patch که قبلاً اعمال شده
patch -R -p1 < changes.patch

# اعمال با ساخت نسخه پشتیبان از فایل اصلی
patch -b < changes.patch
```

# Common mistakes

- انتخاب اشتباه عدد `-p`؛ اگر مسیر فایل در diff با `a/` و `b/` شروع شود (فرمت رایج git)، معمولاً `-p1` لازم است تا این پیشوندها حذف شوند.
- اعمال patch بدون `--dry-run` روی فایل‌های حساس؛ همیشه ابتدا با `--dry-run` بررسی کنید که patch بدون تداخل (conflict) اعمال می‌شود.
- استفاده از `patch` در پروژه‌های مدرن مبتنی بر git به‌جای `git apply`؛ اگر پروژه از قبل با git مدیریت می‌شود، `git apply` یا `git am` معمولاً مناسب‌تر و ایمن‌تر هستند.

# Tips

- برای پروژه‌های git، `git apply patchfile` یا `git am patchfile` (که تاریخچه commit را هم حفظ می‌کند) اغلب جایگزین بهتری از `patch` خام است.
- همیشه قبل از اعمال یک patch ناشناس/دانلودی، با `--dry-run` و بررسی محتوای آن، از امن‌بودنش مطمئن شوید.

# Related commands

- `diff` — ساخت فایل patch از تفاوت دو فایل
- `git apply` — معادل مدرن‌تر برای پروژه‌های git
- `git diff` — ساخت diff در فرمت سازگار با git apply
