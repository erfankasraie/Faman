---
title: history
aliases:
category: shell
difficulty: beginner
keywords:
- shell
- commands
- recall
---

# Introduction

`history` لیست دستورات قبلی شل را نشان می‌دهد. یک builtin در bash/zsh است نه باینری جدا.

# Syntax

```
history [OPTIONS] [N]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `N` | فقط N مورد آخر |
| `-c` | پاک کردن تاریخچه نشست |
| `-d N` | حذف آیتم شماره N |

# Examples

```bash
history
history 20
!123          # اجرای مجدد شماره 123
!!            # تکرار آخرین دستور
!curl         # آخرین دستوری که با curl شروع شده
history | grep apt
```

# Common mistakes

- ذخیره شدن دستورات حساس (رمز) در history — از فاصله اول یا `HISTCONTROL` استفاده کنید.

# Tips

- جستجوی تعاملی: `Ctrl+R`
- فایل تاریخچه bash معمولاً `~/.bash_history` است.

# Related commands

- `alias` — میانبر دائمی‌تر
- `fc` — ویرایش و اجرای مجدد
