---
title: yarn
aliases:
category: development
difficulty: beginner
keywords:
- javascript
- node
- package
- manager
---

# Introduction

`yarn` یک مدیر بسته (package manager) برای جاوااسکریپت/Node.js است؛ جایگزینی برای `npm` با تمرکز بر سرعت نصب، قطعیت نسخه‌ها (`yarn.lock`) و کارایی بهتر در پروژه‌های بزرگ.

# Syntax

```
yarn [COMMAND] [OPTIONS]
```

# Options

زیردستورهای پرکاربرد:

| زیردستور | توضیح |
|-------|--------|
| `add PKG` | افزودن یک بسته به پروژه |
| `remove PKG` | حذف یک بسته |
| `install` (یا فقط `yarn`) | نصب تمام وابستگی‌های `package.json` |
| `run SCRIPT` | اجرای یک اسکریپت تعریف‌شده در `package.json` |
| `upgrade` | به‌روزرسانی بسته‌ها |
| `list` | نمایش درخت وابستگی‌ها |

# Examples

```bash
# نصب تمام وابستگی‌های پروژه (بر اساس package.json و yarn.lock)
yarn install

# افزودن یک بسته جدید
yarn add react

# افزودن یک بسته فقط برای محیط توسعه
yarn add --dev eslint

# حذف یک بسته
yarn remove lodash

# اجرای اسکریپت start تعریف‌شده در package.json
yarn start

# اجرای اسکریپت build
yarn build

# به‌روزرسانی همه بسته‌ها به آخرین نسخه مجاز
yarn upgrade
```

# Common mistakes

- ترکیب `yarn.lock` و `package-lock.json` (فایل قفل npm) در یک پروژه به‌صورت هم‌زمان که باعث ناسازگاری نسخه‌ها می‌شود؛ فقط از یکی استفاده کنید.
- فراموش‌کردن commit کردن فایل `yarn.lock` که باعث می‌شود نسخه‌های نصب‌شده بین اعضای تیم متفاوت باشد.

# Tips

- `yarn.lock` را همیشه در گیت commit کنید تا نسخه دقیق وابستگی‌ها بین همه اعضای تیم و سرورهای CI یکسان بماند.
- برای دیدن این‌که چرا یک بسته خاص نصب شده (کدام وابستگی به آن نیاز دارد): `yarn why PACKAGE_NAME`

# Related commands

- `npm` — مدیر بسته پیش‌فرض Node.js (جایگزین اصلی yarn)
- `npx` — اجرای مستقیم بسته‌های اجرایی بدون نصب دائمی
- `node` — موتور اجرای جاوااسکریپت
