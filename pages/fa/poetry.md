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

## از صفر تا اجرای تست

```bash
poetry new blog && cd blog
poetry add requests
poetry add -D pytest

# کد نمونه
cat > blog/fetch.py <<'EOF'
import requests

def status(url: str) -> int:
    return requests.get(url, timeout=5).status_code
EOF

cat > tests/test_fetch.py <<'EOF'
from blog.fetch import status

def test_example_com():
    assert status("https://example.com") == 200
EOF

poetry install
poetry run pytest -q
```

## env داخل خود پروژه

```bash
poetry config virtualenvs.in-project true
poetry install
ls -la .venv
poetry run python -c "import sys; print(sys.executable)"
```

## اسکریپت در pyproject.toml

```toml
[tool.poetry.scripts]
hello = "blog.fetch:status"
```

```bash
poetry run hello https://example.com
```

## به‌روزرسانی امن

```bash
poetry show --outdated
poetry update requests
# poetry.lock را commit کنید
```

# Common mistakes

- `pip install` دستی کنار poetry.
- commit نکردن `poetry.lock` برای اپ‌ها.

# Tips

- CI: `poetry install --no-root` یا `--only main`.
- جایگزین سبک: `uv`.

# Related commands

- `pip` · `venv` · `pipenv` · `conda`
