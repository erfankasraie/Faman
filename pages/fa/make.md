---
title: make
aliases:
category: development
difficulty: intermediate
keywords:
- build
- makefile
- automation
---

# Introduction

`make` بر اساس `Makefile` فقط هدف‌های لازم را می‌سازد؛ استاندارد build در پروژه‌های C/C++ و بسیاری ابزارها.

# Syntax

```
make [OPTIONS] [TARGET]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-j N` | موازی با N job |
| `-n` | چاپ بدون اجرا |
| `-f FILE` | Makefile دیگر |
| `-C DIR` | تغییر پوشه |
| `-s` | ساکت |

# Examples

```bash
make
make build
make -j$(nproc)
make clean
make -n install
```

نمونه Makefile ساده:

```makefile
.PHONY: build test clean

build:
	go build -o app .

test:
	go test ./...

clean:
	rm -f app
```

# Common mistakes

- استفاده از فاصله به‌جای Tab در دستورات Makefile.

# Tips

- `.PHONY` برای هدف‌هایی که فایل نیستند.

# Related commands

- `cmake` — تولید Makefile/Ninja
- `ninja` — build سریع
- `go build` / `npm run` — ابزار زبان‌محور
