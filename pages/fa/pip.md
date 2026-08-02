---
title: pip
aliases:
- pip3
category: package
difficulty: beginner
keywords:
- python
- package
- pypi
- install
---

# Introduction

`pip` نصب‌کنندهٔ بسته‌های **Python** از PyPI است. برای پروژه‌ها بهتر است داخل venv استفاده شود، نه روی Python سیستم.

# Syntax

```
pip [COMMAND] [OPTIONS]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `install PKG` | نصب |
| `uninstall PKG` | حذف |
| `list` | لیست |
| `freeze` | نسخه‌های پین برای requirements |
| `show PKG` | اطلاعات |
| `install -r requirements.txt` | از فایل |

# Examples

```bash
python3 -m venv .venv
source .venv/bin/activate   # Windows: .venv\Scripts\activate
python -m pip install --upgrade pip
pip install requests
pip freeze > requirements.txt
pip install -r requirements.txt
```

# Common mistakes

- `sudo pip install` روی سیستم و خراب کردن بسته‌های distro.
- فراموش کردن فعال‌سازی venv.

# Tips

- ترجیح: `python -m pip` تا pip درست به همان مفسر وصل شود.
- روی دبیان: بسته‌های سیستمی با `apt`؛ کتابخانه‌های پروژه با pip+venv.

# Related commands

- `python` / `venv`
- `poetry` / `uv` — ابزارهای مدرن‌تر
- `npm` — معادل جاوااسکریپت
