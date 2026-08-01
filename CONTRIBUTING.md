# Contributing to faman

از علاقه شما به مشارکت در faman متشکریم!

## راه‌های مشارکت

- افزودن یا بهبود صفحات راهنما (`pages/fa/*.md`)
- رفع باگ در CLI
- بهبود رندر و تجربه کاربری
- نوشتن تست
- ترجمه یا بهبود مستندات
- پیشنهاد ویژگی جدید (از طریق Issue)

## استاندارد صفحات راهنما

هر صفحه باید شامل front matter زیر باشد:

```yaml
---
title: command
aliases:
- alias1
category: category
difficulty: beginner|intermediate|advanced
keywords:
- kw1
- kw2
---
```

و بخش‌های زیر (به ترتیب ترجیحی):

1. `# Introduction`
2. `# Syntax`
3. `# Options`
4. `# Examples`
5. `# Common mistakes`
6. `# Tips`
7. `# Related commands`

- محتوای واقعی و مفید بنویسید (نه Lorem Ipsum).
- زبان ساده و روان فارسی.
- مثال‌ها باید قابل کپی و اجرا باشند.
- اشتباهات رایج واقعی را پوشش دهید.

## توسعه

```bash
git clone https://github.com/faman-project/faman.git
cd faman
go mod download
go test ./...
go run ./cmd/faman ls
```

## سبک کد

- Idiomatic Go
- Clean Architecture بدون over-engineering
- کامنت فقط جایی که لازم است
- تست برای منطق اصلی

## Pull Request

1. از شاخه `main` جدا شوید.
2. تغییرات مرتبط را در یک PR نگه دارید.
3. تست‌ها باید پاس شوند.
4. توضیح واضح در توضیحات PR بنویسید.

## Code of Conduct

با شرکت در این پروژه، [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) را می‌پذیرید.
