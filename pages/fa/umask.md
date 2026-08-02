---
title: umask
aliases:
category: permissions
difficulty: intermediate
keywords:
- permissions
- default
- security
---

# Introduction

`umask` ماسک مجوز پیش‌فرض فایل‌ها و پوشه‌های جدید را تعیین می‌کند.

# Syntax

```
umask [MODE]
```

# Options

بدون آرگومان مقدار فعلی را نشان می‌دهد. مقدار رایج: `022` یا `002`.

# Examples

```bash
umask
umask 022
umask 077    # فقط مالک دسترسی دارد
```

# Common mistakes

- اشتباه گرفتن umask با chmod؛ umask از مجوز پایه کم می‌کند.

# Tips

- پایه فایل معمولاً ۶۶۶ و پوشه ۷۷۷ است؛ umask از آن کم می‌شود.

# Related commands

- `chmod`
- `chmod` روی فایل ساخته‌شده
