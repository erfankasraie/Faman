---
title: docker
aliases:
category: development
difficulty: intermediate
keywords:
- container
- image
- compose
- isolation
---

# Introduction

`docker` کانتینر را از روی **image** اجرا می‌کند: محیط ایزوله با فایل‌سیستم و شبکهٔ جدا. برای اجرای اپ، تست و استقرار سبک رایج است.

# Syntax

```
docker [OPTIONS] COMMAND
docker run [OPTIONS] IMAGE [COMMAND]
```

# Options

## چرخهٔ کانتینر و ایمیج

| دستور | کار |
|--------|-----|
| `docker pull IMAGE` | دریافت ایمیج |
| `docker images` / `docker image ls` | لیست ایمیج |
| `docker build -t NAME .` | ساخت از Dockerfile |
| `docker run ...` | اجرای کانتینر |
| `docker ps` | در حال اجرا |
| `docker ps -a` | همه |
| `docker stop ID` | توقف |
| `docker rm ID` | حذف کانتینر |
| `docker rmi IMAGE` | حذف ایمیج |
| `docker logs -f ID` | لاگ |
| `docker exec -it ID SH` | شل داخل کانتینر |

## گزینه‌های مهم `run`

| گزینه | توضیح |
|-------|--------|
| `--name NAME` | نام |
| `-d` | detached |
| `-it` | تعاملی + TTY |
| `--rm` | حذف بعد از خروج |
| `-p host:container` | پورت |
| `-v host:container` | volume |
| `-e VAR=val` | متغیر محیطی |
| `--network` | شبکه |
| `-w DIR` | workdir |

# Examples

## اجرا و پورت

```bash
docker run --rm -it alpine sh
docker run -d --name web -p 8080:80 nginx
curl -sI http://127.0.0.1:8080
docker logs web
docker stop web && docker rm web
```

## build پروژه

```bash
# در پوشه‌ای که Dockerfile دارد
docker build -t myapp:dev .
docker run --rm -p 3000:3000 myapp:dev
```

## volume و env

```bash
docker run --rm -v "$PWD":/data -w /data alpine ls
docker run --rm -e NODE_ENV=production myapp:dev
```

## پاکسازی

```bash
docker ps -a
docker container prune
docker image prune
# خطرناک‌تر:
# docker system prune -a
```

## Compose (اشاره)

```bash
docker compose up -d
docker compose logs -f
docker compose down
```

# Common mistakes

- اجرای container به‌عنوان root و mount کردن `/` یا docker.sock بدون نیاز.
- tag نکردن ایمیج و گیج شدن بین `latest`ها.
- فراموش کردن `-p` و تعجب که از میزبان سرویس دیده نمی‌شود.
- حجم پر از ایمیج/لایهٔ یتیم — گاه‌گاه prune.

# Tips

- `.dockerignore` مثل `.gitignore`.
- روی سرورهای جدید گاهی **podman** جایگزین است (`faman podman`).
- برای dev دیتابیس: volume نام‌دار نه فقط bind-mount اگر lifecycle مهم است.

# Related commands

- `podman` · `buildah` · `docker compose` · `kubectl` · `faman nix`
