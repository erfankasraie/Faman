# Contributing to faman

از علاقه شما به مشارکت در faman متشکریم!

## راه‌های مشارکت

- افزودن یا بهبود صفحات راهنما (`pages/fa/*.md`)
- رفع باگ در CLI
- بهبود رندر و تجربه کاربری (مخصوصاً یونیکد/فارسی)
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

و **فقط** این بخش‌های سطح ۱ (به همین نام انگلیسی):

1. `# Introduction`
2. `# Syntax`
3. `# Options`
4. `# Examples`
5. `# Common mistakes`
6. `# Tips`
7. `# Related commands`

قواعد:

- محتوای واقعی و مفید (نه متن پرکننده).
- فارسی روان؛ نام دستورات و فلگ‌ها انگلیسی داخل `` `code` ``.
- مثال‌ها قابل کپی/اجرا باشند.
- جدول گزینه‌ها با `| گزینه | توضیح |`.
- از شکستن مصنوعی کلمات فارسی خودداری کنید.

قالب مرجع: [`pages/fa/ls.md`](pages/fa/ls.md)

## نمایش فارسی در ترمینال

اگر حروف خراب دیده می‌شوند، ابتدا `docs/terminal-persian.md` را بخوانید.
برای تست:

```bash
FAMAN_PLAIN=1 go run ./cmd/faman ls
```

## توسعه

```bash
git clone https://github.com/erfankasraie/Faman.git
cd Faman
go mod download
go test ./...
go run ./cmd/faman ls
```

## سبک کد

- Idiomatic Go
- بدون over-engineering
- تست برای منطق اصلی

## Pull Request

1. از `main` شاخه بسازید.
2. تغییرات مرتبط در یک PR.
3. تست‌ها سبز باشند.
4. توضیح واضح بنویسید.

## Code of Conduct

با مشارکت، [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) را می‌پذیرید.
