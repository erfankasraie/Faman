---
title: podman
aliases:
category: environment
difficulty: intermediate
keywords:
- container
- docker
- rootless
- isolation
---

# Introduction

`podman` موتور کانتینر سازگار با دستورات Docker است؛ اغلب **rootless** و بدون daemon دائمی. برای ایزوله‌سازی محیط اجرا (نه فقط Python) استفاده می‌شود.

# Syntax

```
podman [OPTIONS] COMMAND
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `run` | اجرای کانتینر |
| `build` | ساخت ایمیج |
| `ps` | کانتینرهای در حال اجرا |
| `images` | ایمیج‌ها |
| `pull` / `push` | دریافت/ارسال |
| `rm` / `rmi` | حذف کانتینر/ایمیج |

# Examples

```bash
podman run --rm -it alpine sh
podman build -t myapp .
podman ps -a
podman images

# بسیاری از دستورات شبیه docker:
alias docker=podman   # در صورت نیاز
```

# Common mistakes

- انتظار socket دقیق Docker بدون `podman-docker` یا تنظیمات سازگاری.
- حجم‌ها و SELinux روی Fedora/RHEL (`:Z` روی volume).

# Tips

- برای توسعه: ترکیب با `podman compose` یا `docker-compose` سازگار.
- محیط مجازی زبان‌ها را با venv/nvm حل کنید؛ podman برای کل runtime.

# Related commands

- `docker`
- `buildah` / `skopeo`
- `nix` — ایزولاسیون از نوع دیگر
