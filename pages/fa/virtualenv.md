---
title: virtualenv
aliases:
category: environment
difficulty: beginner
keywords:
- python
- venv
- isolation
---

# Introduction

`virtualenv` ابزار کلاسیک ساخت محیط مجازی Python است. امروز بیشتر پروژه‌ها `python -m venv` را کافی می‌دانند؛ virtualenv هنوز برای نسخه‌های قدیمی یا گزینه‌های بیشتر مفید است.

# Syntax

```
virtualenv [OPTIONS] ENV_DIR
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-p PYTHON` | مسیر مفسر (مثلاً `python3.11`) |
| `--system-site-packages` | دیدن بسته‌های سیستم |
| `--clear` | بازسازی |

# Examples

```bash
pip install --user virtualenv
virtualenv -p python3 .venv
source .venv/bin/activate
pip install flask
deactivate
```

# Common mistakes

- ساخت محیط با Python اشتباه (`-p` را چک کنید).
- فعال نکردن محیط قبل از `pip install`.

# Tips

- برای Python 3.3+ معمولاً `venv` کافی است.
- با `direnv` می‌توان activate را خودکار کرد.

# Related commands

- `venv`
- `pip`
- `pyenv` — انتخاب نسخهٔ Python
