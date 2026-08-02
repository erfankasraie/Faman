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

`direnv` با ورود به یک پوشه، متغیرهای محیطی را از فایل `.envrc` بار می‌کند و با خروج برمی‌گرداند — مناسب venv، کلید API و تنظیمات پروژه.

# Syntax

```
direnv allow
direnv deny
direnv reload
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `allow` | تأیید `.envrc` فعلی |
| `deny` | لغو |
| `reload` | بارگذاری مجدد |
| `status` | وضعیت |

# Examples

```bash
# نصب hook در شل (یک‌بار) — مثلاً در ~/.bashrc:
# eval "$(direnv hook bash)"

cat > .envrc <<'EOF'
export APP_ENV=dev
source_env_if_exists .env
layout python3   # یا: source .venv/bin/activate
EOF

direnv allow
cd .    # یا خروج و ورود مجدد به پوشه
echo "$APP_ENV"
```

# Common mistakes

- commit کردن `.envrc` با **اسرار**؛ اسرار را در `.env` بگذارید و gitignore کنید.
- فراموش کردن `direnv allow` بعد از هر تغییر `.envrc`.

# Tips

- `layout python` / `layout node` در مستندات direnv.
- با `asdf` و `nvm` خوب ترکیب می‌شود.

# Related commands

- `export` / `env`
- `venv`
- `asdf`
