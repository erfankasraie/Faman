---
title: pipenv
aliases:
category: environment
difficulty: intermediate
keywords:
- python
- pipfile
- virtualenv
- package
- lock
---

# Introduction

`pipenv` مدیر **بسته + محیط مجازی** برای Python است: به‌جای `requirements.txt` خام، از `Pipfile` (اعلامی) و `Pipfile.lock` (قفل نسخه) استفاده می‌کند و خودش virtualenv می‌سازد.

نصب ابزار (یک‌بار):

```bash
pip install --user pipenv
# یا: apt install pipenv   (در بعضی توزیع‌ها)
```

# Syntax

```
pipenv <command> [options] [packages]
```

# Options

| دستور / گزینه | توضیح |
|----------------|--------|
| `install` | نصب همه از Pipfile / lock |
| `install PKG` | افزودن و نصب بستهٔ اصلی |
| `install PKG==1.2` | پین نسخه |
| `install -d PKG` / `--dev` | وابستگی توسعه (تست، linter) |
| `uninstall PKG` | حذف از محیط و Pipfile |
| `lock` | بازتولید `Pipfile.lock` |
| `lock --clear` | پاک کردن کش و lock دوباره |
| `sync` | نصب دقیق طبق lock |
| `update` / `update PKG` | به‌روزرسانی |
| `graph` | درخت وابستگی |
| `check` | هشدار امنیتی شناخته‌شده |
| `run CMD` | اجرا داخل محیط (بدون shell) |
| `shell` | فعال‌سازی شل در env |
| `--venv` | مسیر پوشهٔ محیط |
| `--rm` | حذف محیط مجازی |
| `--python 3.12` | انتخاب مفسر هنگام ساخت |

# Examples

## شروع پروژه و مدیریت پکیج

```bash
mkdir shop-api && cd shop-api
pipenv --python 3.12

# بسته‌های اصلی
pipenv install flask requests

# بسته‌های dev
pipenv install -d pytest ruff

# اجرای برنامه
cat > app.py <<'EOF'
from flask import Flask
app = Flask(__name__)

@app.get("/health")
def health():
    return {"ok": True}

if __name__ == "__main__":
    app.run(port=5000)
EOF

pipenv run python app.py
```

## افزودن، پین، حذف

```bash
pipenv install "requests>=2.31,<3"
pipenv install boto3==1.34.0
pipenv uninstall boto3
pipenv graph
cat Pipfile
```

نمونهٔ `Pipfile`:

```toml
[[source]]
url = "https://pypi.org/simple"
verify_ssl = true
name = "pypi"

[packages]
flask = "*"
requests = ">=2.31,<3"

[dev-packages]
pytest = "*"
ruff = "*"

[requires]
python_version = "3.12"
```

## قفل و نصب روی CI / سرور

```bash
# روی ماشین توسعه بعد از تغییر وابستگی:
pipenv lock

# روی سرور یا CI — فقط همان نسخه‌های lock:
pipenv sync
# یا فقط production:
pipenv sync --dev   # اگر تست لازم است

pipenv run pytest -q
```

## تولید requirements.txt (برای ابزارهای قدیمی)

```bash
pipenv requirements > requirements.txt
pipenv requirements --dev > requirements-dev.txt
```

## اسکریپت‌های تکرارشونده

```bash
pipenv run python -m pytest
pipenv run ruff check .
pipenv shell          # سپس دستورهای عادی python/pip
exit
```

## پاکسازی

```bash
pipenv --venv         # کجاست؟
pipenv --rm           # حذف env؛ Pipfile می‌ماند
rm -f Pipfile lock    # اگر پروژه را رها می‌کنید (با احتیاط)
```

# Common mistakes

- `pip install` مستقیم بیرون از `pipenv run` / `shell` → بسته در env پروژه نمی‌نشیند.
- commit نکردن **`Pipfile.lock`** برای اپلیکیشن (بازتولیدپذیری از بین می‌رود).
- دو سیستم موازی: هم `.venv` دستی هم pipenv.
- برای پروژهٔ جدید، اکوسیستم فعال‌تر اغلب **poetry** یا **uv** است؛ pipenv بیشتر در ریپوهای قدیمی دیده می‌شود.

# Tips

- متغیر `PIPENV_VENV_IN_PROJECT=1` → ساخت `.venv` داخل پروژه.
- `pipenv check` را گاه‌به‌گاه برای وابستگی‌های آسیب‌پذیر اجرا کنید.
- اگر فقط یک اسکریپت کوتاه است، `python -m venv` + `pip` ساده‌تر است.

# Related commands

- `pip` — نصب‌کنندهٔ سطح‌پایین
- `venv` — فقط محیط مجازی
- `poetry` — جایگزین مدرن‌تر
- `conda` — محیط‌های علمی
