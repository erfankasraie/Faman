---
title: asdf
aliases:
category: environment
difficulty: intermediate
keywords:
- version
- plugin
- node
- python
- ruby
---

# Introduction

`asdf` یک version manager چندزبانه است: با پلاگین، Python، Node، Ruby، Go و … را یک‌جا مدیریت می‌کند.

# Syntax

```
asdf <command> [plugin] [version]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `plugin add NAME` | افزودن پلاگین |
| `install NAME VER` | نصب نسخه |
| `global NAME VER` | پیش‌فرض سراسری |
| `local NAME VER` | برای پوشه (`.tool-versions`) |
| `list` | نسخه‌های نصب‌شده |
| `current` | نسخهٔ فعال |

# Examples

```bash
asdf plugin add python
asdf plugin add nodejs
asdf install python 3.12.2
asdf install nodejs 20.11.0
asdf global python 3.12.2
asdf local nodejs 20.11.0
asdf current
```

# Common mistakes

- فراموش کردن `asdf reshim` بعد از نصب gem/npm سراسری.
- قاطی کردن با nvm/pyenv هم‌زمان روی PATH.

# Tips

- یک فایل `.tool-versions` برای کل تیم.
- اگر فقط Node می‌خواهید، `nvm`/`fnm` ساده‌ترند.

# Related commands

- `pyenv` / `nvm`
- `direnv`
