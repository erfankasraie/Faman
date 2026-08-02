---
title: docker
aliases:
category: development
difficulty: intermediate
keywords:
- container
- image
- devops
---

# Introduction

`docker` پلتفرم اجرای کانتینر است؛ اپلیکیشن را با وابستگی‌هایش ایزوله اجرا می‌کند.

# Syntax

```
docker [OPTIONS] COMMAND
```

# Options

| فرمان | توضیح |
|-------|--------|
| `ps` | کانتینرهای در حال اجرا |
| `ps -a` | همه کانتینرها |
| `images` | لیست ایمیج‌ها |
| `pull IMAGE` | دانلود ایمیج |
| `run` | اجرای کانتینر |
| `exec` | دستور داخل کانتینر |
| `logs` | لاگ |
| `stop` / `start` / `rm` | مدیریت چرخه عمر |
| `build` | ساخت ایمیج از Dockerfile |
| `compose` | چندسرویسی |

# Examples

```bash
docker ps
docker run --rm -it ubuntu:22.04 bash
docker run -d -p 8080:80 --name web nginx
docker logs -f web
docker exec -it web bash
docker stop web && docker rm web
docker build -t myapp:dev .
docker compose up -d
```

# Common mistakes

- اجرای همه‌چیز با root داخل کانتینر بدون نیاز.
- volume را mount نکردن و از دست رفتن داده.
- پر شدن دیسک از ایمیج/کانتینر بلااستفاده.

# Tips

- پاکسازی: `docker system prune`
- وضعیت منابع: `docker stats`
- ترجیح `docker compose` برای چند سرویس.

# Related commands

- `podman` — جایگزین سازگار
- `kubectl` — کوبرنتیز
- `systemctl` — سرویس داکر روی میزبان
