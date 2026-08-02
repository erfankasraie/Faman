---
title: nvm
aliases:
category: environment
difficulty: beginner
keywords:
- node
- javascript
- version
- npm
---

# Introduction

`nvm` (Node Version Manager) چند نسخهٔ **Node.js** را برای کاربر نصب و بین آن‌ها سوییچ می‌کند.

# Syntax

```
nvm <command> [version]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `install VER` | نصب (مثلاً `20` یا `lts`) |
| `use VER` | فعال در این شل |
| `alias default VER` | پیش‌فرض |
| `ls` | نسخه‌های محلی |
| `ls-remote` | نسخه‌های دور |
| `uninstall VER` | حذف |

# Examples

```bash
nvm install --lts
nvm use 20
node -v && npm -v
nvm alias default 20

# در پروژه با .nvmrc:
echo "20" > .nvmrc
nvm use
```

# Common mistakes

- نصب Node با apt و nvm هم‌زمان و قاطی شدن PATH.
- فراموش کردن بارگذاری nvm در `~/.bashrc` / `~/.zshrc`.

# Tips

- برای تیم: فایل `.nvmrc` در ریشهٔ ریپو.
- جایگزین‌ها: `fnm`، `asdf`، `n`.

# Related commands

- `npm` / `npx`
- `asdf`
- `pyenv` — معادل پایتون
