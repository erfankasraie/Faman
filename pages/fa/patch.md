---
title: patch
aliases:
category: text
difficulty: intermediate
keywords:
- diff
- apply
- update
---

# Introduction

`patch` تغییرات یک فایل diff را روی فایل‌های اصلی اعمال می‌کند.

# Syntax

```
patch [OPTIONS] < FILE.patch
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-pN` | حذف N بخش از مسیر |
| `-R` | برعکس کردن patch |
| `-b` | پشتیبان |
| `--dry-run` | شبیه‌سازی |

# Examples

```bash
patch -p1 < fix.patch
patch -p0 < changes.diff
patch -R -p1 < fix.patch
```

# Common mistakes

- مقدار غلط `-p` و پیدا نشدن فایل‌ها.

# Tips

- در پروژه‌های git معمولاً `git apply` راحت‌تر است.

# Related commands

- `diff`
- `git apply` / `git am`
