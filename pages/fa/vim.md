---
title: vim
aliases:
- vi
- nvim
category: editor
difficulty: intermediate
keywords:
- editor
- text
- modal
---

# Introduction

`vim` ویرایشگر قدرتمند و modal است. یادگیری‌اش زمان‌بر است ولی برای کار روزمره روی سرور بسیار رایج است.

# Syntax

```
vim [OPTIONS] [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `+N` | رفتن به خط N |
| `-R` | فقط خواندنی |
| `-o` / `-O` | چند پنجره افقی/عمودی |
| `-p` | چند تب |

# Examples

```bash
vim file.txt
vim +20 file.txt
vim -R /etc/hosts
```

حالت‌های اصلی:

- **Normal** — پیمایش و فرمان (پیش‌فرض)
- **Insert** — با `i` / `a` / `o`
- **Visual** — با `v`
- **Command** — با `:`

فرمان‌های ضروری:

| کلید | کار |
|------|-----|
| `i` | درج متن |
| `Esc` | برگشت به Normal |
| `:w` | ذخیره |
| `:q` | خروج |
| `:wq` / `ZZ` | ذخیره و خروج |
| `:q!` | خروج بدون ذخیره |
| `/pattern` | جستجو |
| `dd` | حذف خط |
| `yy` / `p` | کپی / چسباندن |

# Common mistakes

- گیر کردن در vim بدون دانستن `:q!`.
- ویرایش با دسترسی اشتباه و ناتوانی در ذخیره.

# Tips

- `vimtutor` برای آموزش تعاملی.
- `nvim` جایگزین مدرن بسیاری از کاربران است.

# Related commands

- `nano` — ساده‌تر برای مبتدی
- `less` — فقط مشاهده
