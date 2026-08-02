---
title: pyenv
aliases:
category: environment
difficulty: intermediate
keywords:
- python
- version
- shims
---

# Introduction

`pyenv` چند نسخهٔ **Python** را نصب و بین آن‌ها جابه‌جا می‌کند.

# Syntax

```
pyenv <command> [args]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `install VER` | نصب نسخه |
| `versions` | لیست |
| `global VER` | پیش‌فرض کاربر |
| `local VER` | برای پوشه |
| `shell VER` | فقط این نشست |

# Examples

## دو پروژه، دو نسخه

```bash
pyenv install 3.11.8
pyenv install 3.12.2

mkdir -p ~/work/legacy ~/work/new

cd ~/work/legacy
pyenv local 3.11.8
python -m venv .venv && source .venv/bin/activate
pip install django==3.2

cd ~/work/new
pyenv local 3.12.2
python -m venv .venv && source .venv/bin/activate
pip install django

# فایل .python-version را commit کنید تا تیم هماهنگ شود
```

## تست سریع یک نسخه

```bash
pyenv shell 3.12.2
python -V
# خروج از shell override با بستن ترمینال یا:
pyenv shell --unset
```

# Common mistakes

- نبودن `eval "$(pyenv init -)"` در rc.
- انتظار مدیریت بسته‌ها بدون venv.

# Tips

- ترکیب طلایی: **pyenv** (نسخه) + **venv/poetry** (بسته).

# Related commands

- `venv` · `asdf` · `conda`
