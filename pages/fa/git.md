---
title: git
aliases:
category: development
difficulty: intermediate
keywords:
- version-control
- source
- commit
- repository
---

# Introduction

`git` سیستم کنترل نسخه توزیع‌شده و استاندارد رایج توسعه نرم‌افزار است. این صفحه پرکاربردترین فرمان‌ها را جمع می‌کند.

# Syntax

```
git <command> [<args>]
```

# Options

| فرمان | توضیح |
|-------|--------|
| `init` / `clone` | شروع یا کلون |
| `status` / `log` / `diff` | وضعیت و تاریخچه |
| `add` / `commit` | ثبت تغییرات |
| `branch` / `switch` / `merge` | شاخه |
| `fetch` / `pull` / `push` | ریموت |
| `restore` / `stash` | بازگردانی موقت |

# Examples

```bash
git clone https://github.com/erfankasraie/Faman.git
git status
git add .
git commit -m "Fix login bug"
git push

git switch -c feature/login
git log --oneline --graph
git restore file.txt
git stash && git stash pop
```

# Common mistakes

- `git add .` بدون دیدن `status`.
- کامیت مستقیم روی main به‌جای feature branch.
- force push روی شاخه اشتراکی.
- نبودن `.gitignore` مناسب.

# Tips

- پیام کامیت کوتاه و واضح.
- قبل از push: `git log` و `git diff`.
- ابزار کمکی: `lazygit`, `delta`, `gh`.

# Related commands

- `gh` — GitHub CLI
- `diff` / `patch`
- `make` — build
