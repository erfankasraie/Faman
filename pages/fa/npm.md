---
title: npm
aliases:
category: package
difficulty: beginner
keywords:
- node
- javascript
- package
- install
---

# Introduction

`npm` مدیر بسته پیش‌فرض **Node.js** است: نصب ماژول‌ها، اجرای اسکریپت‌ها و انتشار پکیج.

# Syntax

```
npm <command> [args]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `install` / `i` | نصب از package.json |
| `install PKG` | افزودن بسته |
| `uninstall PKG` | حذف |
| `run SCRIPT` | اجرای اسکریپت |
| `init` | شروع پروژه |
| `update` | به‌روزرسانی |
| `list` | درخت وابستگی |

# Examples

```bash
npm init -y
npm install express
npm install -D typescript
npm uninstall express
npm run build
npm update
```

# Common mistakes

- commit کردن `node_modules` به‌جای `package-lock.json`.
- نصب سراسری بدون نیاز (`-g`) و شلوغ شدن سیستم.

# Tips

- `npx` برای اجرای موقت باینری پکیج.
- جایگزین‌ها: `yarn`، `pnpm`.

# Related commands

- `node`
- `npx`
- `pip` — پایتون
