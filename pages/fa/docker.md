---
title: docker
aliases:
category: containers
difficulty: intermediate
keywords:
- container
- image
- virtualization
---

# Introduction

`docker` پلتفرم محبوب کانتینرها است که به شما اجازه می‌دهد برنامه‌ها را به صورت ایزوله و قابل‌حمل اجرا کنید.

# Syntax

```
docker <command> [<args>]
```

# دستورات اصلی

## تصاویر (Images)

```bash
docker images                 # لیست تصاویر
docker pull IMAGE             # دانلود تصویر
docker build -t NAME .        # ساخت تصویر از Dockerfile
docker rmi IMAGE              # حذف تصویر
```

## کانتینرها

```bash
docker ps                     # کانتینرهای در حال اجرا
docker ps -a                  # همه کانتینرها
docker run IMAGE              # اجرای کانتینر
docker run -d --name NAME IMAGE  # اجرا در پس‌زمینه
docker run -it IMAGE sh       # تعاملی
docker stop ID/NAME           # توقف
docker start ID/NAME          # شروع
docker rm ID/NAME             # حذف کانتینر
docker logs ID/NAME           # مشاهده لاگ
docker exec -it ID/NAME sh    # ورود به کانتینر در حال اجرا
```

## شبکه‌ها و حجم‌ها

```bash
docker network ls
docker volume ls
docker run -v HOST:CONTAINER IMAGE
docker run -p HOST_PORT:CONTAINER_PORT IMAGE
```

# Examples

```bash
# اجرای nginx
docker run -d -p 8080:80 --name web nginx

# ساخت و اجرای تصویر خودتان
docker build -t myapp .
docker run -d -p 3000:3000 myapp

# ورود به کانتینر
docker exec -it web bash

# پاکسازی
docker system prune
```

# Common mistakes

- اجرای همه چیز با `--privileged` یا `-v /:/host`.
- فراموش کردن `-d` و پر شدن ترمینال.
- نگه داشتن کانتینرهای متوقف‌شده و تصاویر بلااستفاده (پر شدن دیسک).
- استفاده از `latest` tag در production.

# Tips

- همیشه از `.dockerignore` استفاده کنید.
- برای orchestration از Docker Compose یا Kubernetes استفاده کنید.
- `docker compose` (با فاصله) نسخه جدید است.
- منابع را با `--memory` و `--cpus` محدود کنید.
- برای دیباگ: `docker logs -f` و `docker inspect`.

# Related commands

- `docker compose` — مدیریت چندکانتینری
- `podman` — جایگزین بدون daemon
- `kubectl` — Kubernetes
- `nerdctl`
- `buildah`
