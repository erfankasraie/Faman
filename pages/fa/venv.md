---
title: venv
aliases:
- python -m venv
category: environment
difficulty: beginner
keywords:
- python
- virtualenv
- isolation
- pip
---

# Introduction

`venv` ماژول استاندارد **Python** برای ساخت محیط مجازی است: وابستگی‌های پروژه از Python سیستم جدا می‌مانند.

# Syntax

```
python3 -m venv [OPTIONS] ENV_DIR
```

# Options

| گزینه | توضیح |
|-------|--------|
| `ENV_DIR` | مسیر پوشهٔ محیط (مثلاً `.venv`) |
| `--system-site-packages` | دسترسی به بسته‌های سیستم |
| `--clear` | پاک و بازسازی |
| `--upgrade` | ارتقای pip/setuptools داخل محیط |

# Examples

```bash
python3 -m venv .venv
source .venv/bin/activate          # Linux/macOS
# Windows: .venv\Scripts\activate

python -m pip install --upgrade pip
pip install requests
deactivate

# حذف محیط: فقط پوشه را پاک کنید
rm -rf .venv
```

# Common mistakes

- `sudo pip install` بیرون از venv و خراب کردن Python سیستم.
- فراموش کردن `activate` قبل از نصب بسته.
- commit کردن `.venv` در git (باید در `.gitignore` باشد).

# Tips

- همیشه ترجیح: `python -m pip` داخل محیط فعال.
- در IDE مسیر interpreter را روی `.venv` بگذارید.

# Related commands

- `pip` — نصب بسته داخل محیط
- `virtualenv` — ابزار قدیمی‌تر/انعطاف‌پذیرتر
- `poetry` / `pipenv` — مدیریت پروژه + قفل نسخه
