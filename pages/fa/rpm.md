---
title: rpm
aliases:
category: package
difficulty: intermediate
keywords:
- redhat
- fedora
- package
- low-level
---

# Introduction

`rpm` قالب و ابزار سطح‌پایین بسته در اکوسیستم Red Hat (Fedora، RHEL، openSUSE و …) است. برای کار روزمره معمولاً `dnf` یا `zypper` بهتر است.

# Syntax

```
rpm [OPTIONS] PACKAGE_FILE
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-i FILE.rpm` | نصب |
| `-U FILE.rpm` | ارتقا/نصب |
| `-e PKG` | حذف |
| `-q PKG` | پرسش (نصب است؟) |
| `-qa` | همهٔ بسته‌ها |
| `-ql PKG` | لیست فایل‌ها |
| `-qf FILE` | مالک فایل |

# Examples

```bash
sudo rpm -ivh package.rpm
rpm -qa | grep nginx
rpm -ql nginx
rpm -qf /usr/bin/curl
sudo rpm -e package-name
```

# Common mistakes

- نصب با `rpm -i` بدون حل وابستگی؛ `dnf install ./file.rpm` وابستگی را بهتر مدیریت می‌کند.

# Tips

- تأیید امضا و مخزن را به `dnf`/`zypper` بسپارید.
- `--nodeps` خطرناک است مگر عیب‌یابی.

# Related commands

- `dnf` / `yum`
- `zypper`
- `dpkg` — معادل دبیان
