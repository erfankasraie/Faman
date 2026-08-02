---
title: direnv
aliases:
category: environment
difficulty: intermediate
keywords:
- env
- shell
- autoenv
- dotenv
---

# Introduction

`direnv` با ورود به پوشه، متغیرها را از `.envrc` بار می‌کند و با خروج برمی‌گرداند.

# Syntax

```
direnv allow | deny | reload | status
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `allow` | تأیید `.envrc` |
| `deny` | لغو |
| `reload` | بارگذاری مجدد |
| `status` | وضعیت |

# Examples

## فعال‌سازی خودکار Python venv

```bash
# یک‌بار در ~/.bashrc یا ~/.zshrc:
# eval "$(direnv hook bash)"   # یا zsh

cd ~/projects/myapi
python3 -m venv .venv

cat > .envrc <<'EOF'
export APP_ENV=development
export DATABASE_URL=postgres://localhost:5432/myapi
source .venv/bin/activate
EOF

echo '.env' >> .gitignore
direnv allow

# الان با cd به این پوشه، venv و متغیرها فعال می‌شوند:
which python
echo "$APP_ENV"
```

## جدا کردن اسرار

```bash
cat > .env <<'EOF'
API_KEY=sk-secret-do-not-commit
EOF
echo '.env' >> .gitignore

cat > .envrc <<'EOF'
source_env_if_exists .env
export APP_ENV=development
EOF
direnv allow
```

## ترکیب با nvm

```bash
cat > .envrc <<'EOF'
# نیاز به nvm بارگذاری‌شده در شل
nvm use
EOF
direnv allow
```

# Common mistakes

- اسرار داخل `.envrc` و commit شدن.
- فراموش کردن `direnv allow` بعد از ویرایش.

# Tips

- `.envrc` را برای دستورات غیرحساس commit کنید؛ اسرار فقط در `.env`.

# Related commands

- `export` · `venv` · `nvm` · `asdf`
