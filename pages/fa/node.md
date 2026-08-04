---
title: node
aliases:
- nodejs
category: development
difficulty: beginner
keywords:
- javascript
- runtime
- server
- development
---

# Introduction

`node` (Node.js) محیط اجرای جاوااسکریپت خارج از مرورگر است؛ برای اجرای اسکریپت‌های سمت سرور، ابزارهای خط‌فرمان، و توسعه وب‌اپلیکیشن‌ها استفاده می‌شود.

# Syntax

```
node [OPTIONS] [SCRIPT.js] [ARGUMENTS]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-e CODE` | اجرای مستقیم یک قطعه کد بدون فایل |
| `-v` / `--version` | نمایش نسخه Node.js |
| `-i` | ورود به حالت تعاملی (REPL) |
| `--watch` | اجرای دوباره خودکار هنگام تغییر فایل |
| `--inspect` | فعال‌کردن دیباگر برای اتصال ابزارهایی مثل Chrome DevTools |

# Examples

```bash
# اجرای یک فایل جاوااسکریپت
node app.js

# اجرای مستقیم یک خط کد
node -e "console.log(1 + 1)"

# ورود به محیط تعاملی برای تست سریع کد
node

# اجرا با راه‌اندازی مجدد خودکار هنگام ذخیره فایل
node --watch server.js

# دیدن نسخه نصب‌شده
node -v

# اجرا با آرگومان‌های خط‌فرمان قابل‌دسترسی در process.argv
node script.js arg1 arg2
```

# Common mistakes

- استفاده از دستورات مرورگر (مثل `document`, `window`) در محیط Node.js که این اشیا فقط در مرورگر وجود دارند.
- عدم مدیریت نسخه Node.js بین پروژه‌های مختلف؛ استفاده از ابزارهایی مثل `nvm` برای سوییچ سریع بین نسخه‌ها توصیه می‌شود.
- فراموش‌کردن `require`/`import` درست ماژول‌ها که منجر به خطای «module not found» می‌شود.

# Tips

- برای مدیریت چند نسخه Node.js روی یک سیستم (مثلاً برای پروژه‌های مختلف)، از `nvm` استفاده کنید.
- محیط REPL تعاملی (`node` بدون آرگومان) برای تست سریع قطعه‌کد بسیار مفید است.

# Related commands

- `npm` — مدیر بسته پیش‌فرض Node.js
- `npx` — اجرای مستقیم اجرایی‌های بسته‌ها بدون نصب دائمی
- `nvm` — مدیریت نسخه‌های مختلف Node.js
- `yarn` — مدیر بسته جایگزین npm
