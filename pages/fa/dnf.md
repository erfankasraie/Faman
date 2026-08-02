---
title: dnf
aliases:
- yum
category: package
difficulty: beginner
keywords:
- package
- fedora
- rhel
- install
---

# Introduction

`dnf` مدیر بسته توزیع‌های Fedora، RHEL، Rocky و AlmaLinux است (جانشین yum).

# Syntax

```
dnf [OPTIONS] COMMAND
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `install PKG` | نصب |
| `remove PKG` | حذف |
| `update` / `upgrade` | به‌روزرسانی |
| `search KEY` | جستجو |
| `info PKG` | اطلاعات |
| `clean all` | پاک کردن کش |

# Examples

```bash
sudo dnf install git
sudo dnf upgrade
dnf search nginx
dnf info curl
sudo dnf remove package
```

# Common mistakes

- اشتباه گرفتن نام بسته با نام باینری.

# Tips

- تاریخچه: `dnf history` و امکان rollback.

# Related commands

- `rpm` — سطح پایین‌تر
- `apt` — در Debian/Ubuntu
- `pacman` — در Arch
