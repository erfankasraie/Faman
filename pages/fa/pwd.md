---
title: pwd
aliases:
- print-working-directory
category: filesystem
difficulty: beginner
keywords:
- path
- directory
- current
---

# Introduction

دستور `pwd` مسیر کامل پوشه جاری (Print Working Directory) را نمایش می‌دهد.

# Syntax

```
pwd [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-L` | مسیر منطقی (پیش‌فرض، symbolic link را دنبال نمی‌کند) |
| `-P` | مسیر فیزیکی (symbolic link را resolve می‌کند) |

# Examples

```bash
# نمایش مسیر جاری
pwd

# مسیر فیزیکی
pwd -P
```

# Common mistakes

- فکر کردن که `pwd` همیشه مسیر واقعی را نشان می‌دهد (در صورت وجود symlink ممکن است منطقی باشد).

# Tips

- در اسکریپت‌ها برای دانستن محل اجرا بسیار مفید است.
- متغیر محیطی `$PWD` معمولاً همان خروجی `pwd` است.
- برای رفتن به پوشه خانه: `cd ~` یا `cd`

# Related commands

- `cd` — تغییر پوشه جاری
- `ls` — لیست محتویات
- `realpath` — مسیر مطلق واقعی
- `readlink` — خواندن symbolic link
