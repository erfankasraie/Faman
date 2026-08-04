---
title: go
aliases:
- golang
category: development
difficulty: intermediate
keywords:
- golang
- build
- compiler
- development
---

# Introduction

`go` ابزار خط‌فرمانی زبان برنامه‌نویسی Go است؛ برای build، اجرا، تست، مدیریت ماژول‌ها (وابستگی‌ها) و فرمت‌کردن کد Go استفاده می‌شود.

# Syntax

```
go COMMAND [ARGUMENTS]
```

# Options

زیردستورهای پرکاربرد:

| زیردستور | توضیح |
|-------|--------|
| `run` | کامپایل و اجرای مستقیم یک فایل/پکیج |
| `build` | ساخت فایل اجرایی |
| `test` | اجرای تست‌ها |
| `mod` | مدیریت ماژول‌ها (`go.mod`, `go.sum`) |
| `get` | افزودن یا به‌روزرسانی یک وابستگی |
| `fmt` | فرمت‌کردن خودکار کد |
| `vet` | بررسی مشکلات احتمالی در کد |
| `install` | ساخت و نصب باینری در `GOPATH/bin` |

# Examples

```bash
# اجرای مستقیم یک برنامه بدون ساخت فایل باینری جداگانه
go run main.go

# ساخت یک فایل اجرایی
go build -o myapp ./cmd/myapp

# اجرای تمام تست‌ها با نمایش جزئیات
go test -v ./...

# افزودن یک وابستگی جدید
go get github.com/spf13/cobra@latest

# پاک‌سازی و به‌روزرسانی go.mod/go.sum بر اساس کد واقعی
go mod tidy

# فرمت‌کردن خودکار تمام فایل‌های پروژه
go fmt ./...

# ساخت برای سیستم‌عامل و معماری دیگر (کراس-کامپایل)
GOOS=linux GOARCH=arm64 go build -o myapp-arm64 ./cmd/myapp
```

# Common mistakes

- فراموش‌کردن `go mod tidy` بعد از افزودن import جدید که باعث خطای «missing go.sum entry» می‌شود.
- گیج‌شدن بین `go run` (برای توسعه سریع، بدون فایل خروجی دائمی) و `go build` (برای ساخت باینری نهایی قابل توزیع).
- تنظیم‌نکردن `GOPATH`/`GOBIN` هنگام استفاده از `go install` که باعث می‌شود باینری در جایی غیرمنتظره قرار بگیرد.

# Tips

- برای پروژه‌های جدید همیشه از Go Modules (`go mod init`) استفاده کنید؛ سیستم قدیمی GOPATH-only دیگر توصیه نمی‌شود.
- `go vet ./...` را قبل از commit اجرا کنید تا خطاهای رایج (مثل فرمت اشتباه Printf) زودتر پیدا شوند.

# Related commands

- `gofmt` — فرمت‌کننده پایه (که `go fmt` از آن استفاده می‌کند)
- `git` — معمولاً همراه با go برای مدیریت نسخه پروژه
- `make` — گاهی برای اسکریپت‌های build پیچیده‌تر روی go استفاده می‌شود
