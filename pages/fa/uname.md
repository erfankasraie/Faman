---
title: uname
aliases:
category: system
difficulty: beginner
keywords:
- kernel
- system
- architecture
---

# Introduction

`uname` اطلاعات پایه سیستم‌عامل و کرنل را چاپ می‌کند.

# Syntax

```
uname [OPTIONS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-a` | همه اطلاعات |
| `-s` | نام کرنل |
| `-r` | نسخه کرنل |
| `-m` | معماری |
| `-n` | hostname |

# Examples

```bash
uname -a
uname -r
uname -m
```

# Common mistakes

- اشتباه گرفتن خروجی uname با نسخه توزیع — برای آن از `/etc/os-release` استفاده کنید.

# Tips

- نسخه توزیع: `cat /etc/os-release`

# Related commands

- `hostnamectl`
- `lsb_release -a` (در بعضی توزیع‌ها)
