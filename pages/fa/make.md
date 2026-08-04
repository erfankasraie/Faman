---
title: make
aliases:
category: development
difficulty: intermediate
keywords:
- build
- makefile
- automation
- target
---

# Introduction

`make` فرمان‌ها را بر اساس **Makefile** اجرا می‌کند: targetها، وابستگی‌ها، و بازسازی فقط وقتی ورودی جدیدتر است. رایج در C/C++ و هر پروژه با کارهای تکراری.

# Syntax

```
make [OPTIONS] [TARGET...]
make -f FILE TARGET
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-f FILE` | نام Makefile |
| `-j N` | اجرای موازی N job |
| `-n` | dry-run (چاپ بدون اجرا) |
| `-B` | اجبار rebuild |
| `-s` | ساکت |
| `-C DIR` | تغییر پوشه قبل از اجرا |
| `-k` | با خطا ادامه بده |
| `VAR=value` | متغیر از خط فرمان |

## ساختار Makefile (خلاصه)

```makefile
target: dependencies
\tcommand
\tcommand
```

توجه: فرورفتگی دستورات باید **Tab** باشد نه فاصله.

| مفهوم | معنی |
|--------|------|
| `$@` | نام target |
| `$<` | اولین پیش‌نیاز |
| `$^` | همهٔ پیش‌نیازها |
| `.PHONY` | target غیر فایلی |
| `VAR := value` | متغیر |

# Examples

## Makefile ساده

```makefile
.PHONY: all clean test

all: app

app: main.o util.o
\tgcc -o $@ $^

%.o: %.c
\tgcc -c -o $@ $<

test:
\t./app --test

clean:
\trm -f app *.o
```

```bash
make
make -j4
make clean
make test
make CFLAGS='-O2 -Wall'
```

## فقط اسکریپت بدون compile

```makefile
.PHONY: run lint
run:
\tpython app.py
lint:
\truff check .
```

```bash
make run
make -n lint    # ببین چه اجرا می‌شود
```

## از پوشه دیگر

```bash
make -C build
make -f Makefile.release all
```

# Common mistakes

- فاصله به‌جای Tab → خطای confusing.
- target بدون `.PHONY` وقتی فایلی هم‌نام وجود دارد.
- وابستگی ناقص → rebuild نشدن بعد از تغییر هدر.

# Tips

- `make -j$(nproc)` روی لینوکس.
- برای پروژهٔ بزرگ گاهی `ninja` / سیستم build دیگر.
- `make help` را خودتان با یک target `help` مستند کنید.

# Related commands

- `gcc` · `cmake` · `ninja` · `cargo` · `npm`
