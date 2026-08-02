---
title: pipenv
aliases:
category: environment
difficulty: intermediate
keywords:
- python
- pipfile
- virtualenv
---

# Introduction

`pipenv` ترکیب `pip` + محیط مجازی با فایل‌های `Pipfile` و `Pipfile.lock` است.

# Syntax

```
pipenv <command> [options]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `install` | نصب از Pipfile |
| `install PKG` | افزودن بسته |
| `uninstall PKG` | حذف |
| `shell` | شل داخل env |
| `run CMD` | اجرای فرمان |
| `lock` | بازسازی lock |
| `--venv` | مسیر محیط |

# Examples

```bash
cd myproject
pipenv install requests
pipenv install -d pytest
pipenv shell
pipenv run python app.py
pipenv --rm    # حذف محیط
```

# Common mistakes

- پروژه‌های جدید اغلب به **poetry** یا **uv** مهاجرت کرده‌اند؛ pipenv کمتر نگهداری می‌شود.
- دو محیط موازی (venv دستی + pipenv).

# Tips

- برای پروژهٔ تازه، `poetry` یا `python -m venv` + `pip` را در نظر بگیرید.

# Related commands

- `poetry`
- `venv`
- `pip`
