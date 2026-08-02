---
title: grep
aliases:
- search-text
category: text
difficulty: intermediate
keywords:
- search
- text
- pattern
- regex
---

# Introduction

`grep` الگو را داخل فایل‌ها جستجو می‌کند. نامش از Global Regular Expression Print گرفته شده است.

# Syntax

```
grep [OPTIONS] PATTERN [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-i` | بدون حساسیت حروف |
| `-r` / `-R` | بازگشتی |
| `-n` | شماره خط |
| `-v` | خطوط غیرمطابق |
| `-l` | فقط نام فایل |
| `-c` | تعداد |
| `-A N` / `-B N` / `-C N` | context |
| `-E` | regex توسعه‌یافته |
| `-w` | کلمه کامل |
| `--color=auto` | رنگ |
| `--exclude-dir=DIR` | نادیده گرفتن پوشه |

# Examples

```bash
grep 'error' logfile.txt
grep -in 'error' logfile.txt
grep -r 'TODO' src/
grep -rl 'password' /etc/
grep -C 3 'exception' app.log
grep -E 'error|warning|fail' app.log
grep -r --exclude-dir=.git --exclude-dir=node_modules 'FIXME' .
```

# Common mistakes

- pattern پیچیده بدون کوتیشن.
- `grep -r` روی درخت عظیم بدون exclude.
- انتظار regex پیشرفته بدون `-E` یا `-P`.

# Tips

- جایگزین سریع: `rg` (ripgrep).
- ترکیب: `grep ... | less`
- با find: `find . -name '*.py' -print0 | xargs -0 grep -n 'def '`. 

# Related commands

- `rg` / `ag`
- `find`
- `sed` / `awk`
