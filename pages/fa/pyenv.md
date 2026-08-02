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

`pyenv` چند نسخهٔ **Python** را کنار هم نصب و با shim بین آن‌ها جابه‌جا می‌کند (سراسری، هر پوشه، یا شل).

# Syntax

```
pyenv <command> [args]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `install -l` | نسخه‌های قابل نصب |
| `install 3.12.0` | نصب نسخه |
| `versions` | نصب‌شده‌ها |
| `global VER` | پیش‌فرض کاربر |
| `local VER` | برای این پوشه (`.python-version`) |
| `shell VER` | فقط این نشست |

# Examples

```bash
pyenv install 3.12.2
pyenv global 3.12.2
cd myproject
pyenv local 3.11.8
python -V
python -m venv .venv
```

# Common mistakes

- فراموش کردن init در shell rc (`pyenv init`).
- انتظار که pyenv به‌جای venv وابستگی پروژه را مدیریت کند — فقط نسخهٔ مفسر را عوض می‌کند.

# Tips

- ترکیب رایج: **pyenv** (نسخه) + **venv/poetry** (بسته‌ها).
- افزونه `pyenv-virtualenv` برای مدیریت envها.

# Related commands

- `venv`
- `asdf` — چندزبانه‌تر
- `conda`
