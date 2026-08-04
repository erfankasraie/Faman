---
title: git
aliases:
category: development
difficulty: intermediate
keywords:
- vcs
- commit
- branch
- github
---

# Introduction

`git` سیستم کنترل نسخهٔ توزیع‌شده است: تاریخچهٔ تغییرات، شاخه، ادغام و همکاری از طریق remote (GitHub/GitLab).

این صفحه **روزمره** است؛ git خیلی بزرگ‌تر از یک man کوتاه است.

# Syntax

```
git <command> [<args>]
```

# Options

## تنظیم و کلون

| دستور | کار |
|--------|-----|
| `git config --global user.name "..."` | نام |
| `git config --global user.email "..."` | ایمیل |
| `git clone URL [DIR]` | کپی ریپو |
| `git init` | ریپو جدید |

## وضعیت و تاریخچه

| دستور | کار |
|--------|-----|
| `git status` | وضعیت working tree |
| `git log --oneline --graph` | تاریخچه |
| `git diff` | تغییرات stage‌نشده |
| `git diff --staged` | آمادهٔ commit |
| `git show` | یک commit |

## تغییر و ذخیره

| دستور | کار |
|--------|-----|
| `git add PATH` | stage |
| `git add -p` | تکه‌تکه |
| `git commit -m "msg"` | ثبت |
| `git commit --amend` | اصلاح آخرین commit |
| `git restore PATH` | دور انداختن تغییر کاری |
| `git restore --staged PATH` | unstage |

## شاخه و remote

| دستور | کار |
|--------|-----|
| `git branch` | لیست |
| `git switch NAME` | عوض کردن شاخه |
| `git switch -c NAME` | ساخت و رفتن |
| `git merge BRANCH` | ادغام |
| `git rebase BRANCH` | بازنویسی روی پایه |
| `git remote -v` | remoteها |
| `git fetch` | گرفتن بدون merge |
| `git pull` | fetch + integrate |
| `git push -u origin BRANCH` | ارسال |

# Examples

## چرخهٔ روزانه

```bash
git clone https://github.com/erfankasraie/Faman.git
cd Faman
git switch -c feature/docs
# ویرایش فایل‌ها
git status
git add docs/macos.md
git commit -m "docs: improve macOS install guide"
git push -u origin feature/docs
```

## تاریخچه و مقایسه

```bash
git log --oneline -20
git log -- path/to/file
git diff main...HEAD
```

## اصلاح قبل از push

```bash
git add forgotten.md
git commit --amend --no-edit
# اگر هنوز push نشده باشد
```

## undoهای امن‌تر

```bash
# تغییر فایل قبل از commit
git restore file.txt

# از stage خارج کن
git restore --staged file.txt

# آخرین commit را برگردان ولی تغییرات بمانند
git reset --soft HEAD~1
```

## همگام با upstream

```bash
git fetch origin
git switch main
git pull --ff-only
```

# Common mistakes

- `git push --force` روی `main` مشترک.
- commit کردن secrets (`.env`, کلید).
- `git clone` با HTTPS وقتی فقط SSH تنظیم شده (یا برعکس).
- پیام commit مبهم (`fix` alone).

# Tips

- `.gitignore` را زود بنویسید.
- `git config --global pull.rebase false` (یا true) را آگاهانه انتخاب کنید.
- برای فایل بزرگ: Git LFS.

# Related commands

- `gh` (GitHub CLI) · `diff` · `ssh` · `faman docker`
