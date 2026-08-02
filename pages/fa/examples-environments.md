---
title: examples-environments
aliases:
- محیط-مجازی
- virtual-environments
category: environment
difficulty: beginner
keywords:
- cookbook
- examples
- venv
- docker
---

# Introduction

این صفحه **دستورهای آماده** برای سناریوهای رایج محیط مجازی و ایزوله است. جزئیات هر ابزار را با `faman venv`، `faman poetry` و … ببینید.

# Syntax

سناریوها مستقل‌اند؛ فقط بلوک مورد نیاز را کپی کنید.

# Options

انتخاب سریع:

| نیاز | ابزار پیشنهادی |
|------|----------------|
| پروژهٔ Python ساده | `venv` + `pip` |
| تیم + قفل نسخه | `poetry` |
| داده / علمی | `conda` |
| چند نسخهٔ Node | `nvm` |
| فعال‌سازی خودکار پوشه | `direnv` |
| ایزولهٔ کامل سیستم | `docker` / `podman` |

# Examples

## ۱) Python از صفر تا requirements

```bash
mkdir demo-py && cd demo-py
python3 -m venv .venv
source .venv/bin/activate
pip install httpx
python -c "import httpx; print(httpx.get('https://example.com').status_code)"
pip freeze > requirements.txt
deactivate
```

## ۲) Node LTS + اسکریپت npm

```bash
mkdir demo-node && cd demo-node
nvm install --lts && nvm use --lts
echo "lts/*" > .nvmrc
npm init -y
npm pkg set scripts.start="node -e \"console.log('hi')\""
npm start
```

## ۳) یک سرویس در کانتینر (بدون کثیف کردن میزبان)

```bash
mkdir demo-ctr && cd demo-ctr
cat > Dockerfile <<'EOF'
FROM python:3.12-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
CMD ["python", "-c", "print('container ok')"]
EOF
echo "requests==2.32.3" > requirements.txt
docker build -t demo-ctr .
docker run --rm demo-ctr
```

## ۴) direnv + venv در یک پروژه

```bash
python3 -m venv .venv
printf '%s\n' 'source .venv/bin/activate' 'export APP_ENV=dev' > .envrc
direnv allow
# با هر cd به پروژه، محیط آماده است
```

## ۵) مقایسهٔ سریع «کجا هستم؟»

```bash
echo "Python: $(command -v python) $(python -V 2>/dev/null)"
echo "Node:   $(command -v node) $(node -v 2>/dev/null)"
echo "CONDA:  ${CONDA_DEFAULT_ENV:-}"
echo "VIRTUAL_ENV=${VIRTUAL_ENV:-}"
```

# Common mistakes

- چند manager هم‌زمان روی PATH بدون نظم (nvm + node apt).
- اسرار در git.

# Tips

- یک پروژه = یک روش اصلی (مثلاً فقط poetry، نه poetry+pipenv با هم).

# Related commands

- `venv` · `poetry` · `conda` · `nvm` · `direnv` · `docker` · `podman`
