---
title: conda
aliases:
- mamba
category: environment
difficulty: intermediate
keywords:
- python
- anaconda
- miniconda
- environment
---

# Introduction

`conda` مدیر محیط و بسته (Anaconda/Miniconda) است: علاوه بر Python می‌تواند کتابخانه‌های غیرpip و حتی R را هم در محیط‌های جدا نگه دارد.

# Syntax

```
conda <command> [options]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `create -n NAME` | محیط جدید |
| `activate NAME` | فعال‌سازی |
| `deactivate` | خروج |
| `install PKG` | نصب در محیط فعلی |
| `env list` | لیست محیط‌ها |
| `remove -n NAME --all` | حذف محیط |
| `env export > env.yml` | خروجی قابل‌بازسازی |

# Examples

```bash
conda create -n ml python=3.11
conda activate ml
conda install numpy pandas
conda env export > environment.yml
conda deactivate
conda env remove -n ml

# از فایل:
conda env create -f environment.yml
```

# Common mistakes

- قاطی کردن `pip` و `conda install` بدون نظم (ترجیح: اول conda، بعد pip در همان env).
- فعال بودن base برای همهٔ کارها و شلوغ شدن محیط پایه.

# Tips

- **Mamba** حل‌کنندهٔ سریع‌تر سازگار با conda است.
- برای فقط Python سبک: Miniconda یا حتی فقط `venv`.

# Related commands

- `venv` / `pip`
- `mamba`
- `poetry`
