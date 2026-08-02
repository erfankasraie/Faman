---
title: poetry
aliases:
category: environment
difficulty: intermediate
keywords:
- python
- dependency
- lock
- venv
---

# Introduction

`poetry` ابزار مدیریت وابستگی و بسته‌بندی **Python** است: محیط مجازی، قفل نسخه (`poetry.lock`) و انتشار را یکپارچه می‌کند.

# Syntax

```
poetry <command> [options]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `init` / `new` | شروع پروژه |
| `add PKG` | افزودن وابستگی |
| `install` | نصب از lock |
| `update` | به‌روزرسانی |
| `run CMD` | اجرا داخل env |
| `shell` | فعال‌سازی شل |
| `build` / `publish` | بسته و انتشار |

# Examples

```bash
poetry new myapp && cd myapp
poetry add requests
poetry add -D pytest
poetry install
poetry run python -m myapp
poetry shell
```

# Common mistakes

- مخلوط کردن `pip install` دستی با poetry در همان پروژه.
- commit نکردن `poetry.lock` برای اپلیکیشن‌ها.

# Tips

- `poetry config virtualenvs.in-project true` تا `.venv` داخل پروژه ساخته شود.
- جایگزین سبک: `uv` یا `pip-tools`.

# Related commands

- `pip` / `venv`
- `pipenv`
- `conda`
