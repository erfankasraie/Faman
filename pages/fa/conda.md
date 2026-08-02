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

`conda` مدیر محیط و بسته (Anaconda/Miniconda) است: Python و کتابخانه‌های علمی را در محیط‌های جدا نگه می‌دارد.

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
| `install PKG` | نصب |
| `env list` | لیست |
| `remove -n NAME --all` | حذف |
| `env export` | خروجی YAML |

# Examples

## محیط داده/یادگیری ماشین

```bash
conda create -n ml python=3.11 -y
conda activate ml
conda install numpy pandas scikit-learn matplotlib -y

python - <<'EOF'
import numpy as np
print(np.random.rand(2, 2))
EOF

conda env export --from-history > environment.yml
conda deactivate
```

## بازسازی از YAML

```bash
conda env create -f environment.yml
conda activate ml
```

## فقط pip داخل conda env

```bash
conda activate ml
conda install pip -y
pip install requests   # بعد از بسته‌های conda
pip freeze > requirements-pip.txt
```

## پاکسازی

```bash
conda env list
conda env remove -n ml -y
conda clean -a -y      # کش
```

# Common mistakes

- شلوغ کردن محیط `base`.
- مخلوط بی‌نظم pip و conda.

# Tips

- **mamba** سریع‌تر است: `mamba install ...`.
- برای پروژهٔ فقط-pip، `venv` سبک‌تر است.

# Related commands

- `venv` · `pip` · `poetry` · `pyenv`
