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

## پروژهٔ وب ساده با Flask

```bash
mkdir myapi && cd myapi
python3 -m venv .venv
source .venv/bin/activate
python -m pip install --upgrade pip
pip install flask

cat > app.py <<'EOF'
from flask import Flask
app = Flask(__name__)

@app.get("/")
def hello():
    return {"msg": "سلام از venv"}

if __name__ == "__main__":
    app.run(debug=True)
EOF

python app.py
# مرورگر: http://127.0.0.1:5000/

pip freeze > requirements.txt
deactivate
```

## بازسازی روی ماشین دیگر

```bash
git clone <repo> && cd <repo>
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
python app.py
```

## ویندوز (PowerShell)

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install requests
deactivate
```

## چند نکتهٔ روزمره

```bash
which python          # باید .../.venv/bin/python باشد
python -V
pip list
rm -rf .venv          # حذف کامل محیط
```

# Common mistakes

- `sudo pip install` بیرون از venv.
- فراموش کردن `activate`.
- commit کردن `.venv` (در `.gitignore` بگذارید).

# Tips

- همیشه: `python -m pip` داخل محیط فعال.
- نام رایج پوشه: `.venv` یا `venv`.

# Related commands

- `pip` · `virtualenv` · `poetry` · `direnv`
