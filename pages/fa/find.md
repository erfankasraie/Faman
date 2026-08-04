---
title: find
aliases:
category: filesystem
difficulty: intermediate
keywords:
- search
- files
- recursive
- locate
- exec
---

# Introduction

`find` درخت دایرکتوری را پیمایش می‌کند و بر اساس **عبارت‌ها** (نام، نوع، اندازه، زمان، مجوز، مالک و …) فایل/پوشه پیدا می‌کند. برخلاف `locate` که از ایندکس استفاده می‌کند، `find` همان لحظه دیسک را می‌خواند — دقیق‌تر و کندتر.

# Syntax

```
find [مسیر...] [گزینه‌های سراسری] [عبارت]
```

ترتیب مهم است: اول مسیرها، بعد آزمون‌ها و عمل‌ها. پیش‌فرض مسیر `.` و عمل پیش‌فرض `-print` است.

# Options

## گزینه‌های سراسری (قبل از عبارت)

| گزینه | توضیح |
|-------|--------|
| `-maxdepth N` | حداکثر عمق از نقاط شروع |
| `-mindepth N` | حداقل عمق (مثلاً ۱ = خود ریشه را رد کن) |
| `-xdev` / `-mount` | از filesystem خارج نشو |
| `-L` | symlink را دنبال کن (مثل `-follow`) |
| `-P` | symlink را دنبال نکن (پیش‌فرض) |
| `-H` | فقط symlinkهای خط فرمان را دنبال کن |
| `-daystart` | زمان را از ابتدای امروز حساب کن |

## آزمون نام و مسیر

| عبارت | توضیح |
|-------|--------|
| `-name PATTERN` | فقط basename؛ حساس به حروف؛ glob شل (`* ? []`) |
| `-iname PATTERN` | مثل name بدون حساسیت حروف |
| `-path PATTERN` | کل مسیر نسبت به نقطهٔ شروع |
| `-ipath PATTERN` | path بدون حساسیت |
| `-regex PATTERN` | regex روی کل مسیر (سینتکس emacs پیش‌فرض) |
| `-iregex PATTERN` | regex بدون حساسیت |
| `-regextype TYPE` | `posix-extended` و … |

**همیشه** الگو را کوتیشن کنید تا شل `*` را زود expand نکند.

## نوع، اندازه، خالی

| عبارت | توضیح |
|-------|--------|
| `-type f` | فایل معمولی |
| `-type d` | پوشه |
| `-type l` | symlink |
| `-type b/c/p/s` | block / char / pipe / socket |
| `-xtype` | مثل type ولی بعد از resolve برای لینک |
| `-empty` | فایل خالی یا پوشهٔ بدون ورودی |
| `-size N` | دقیقاً N واحد |
| `-size +N` / `-size -N` | بزرگ‌تر / کوچک‌تر |
| واحد size | `c` بایت، `k` کیبی‌بایت، `M`، `G` (پیش‌فرض 512-byte blocks اگر واحد ننویسید) |

## زمان

| عبارت | معنی |
|-------|--------|
| `-mtime N` | دادهٔ فایل N*۲۴ ساعت پیش عوض شده |
| `-mtime -N` | کمتر از N روز |
| `-mtime +N` | بیشتر از N روز |
| `-mmin N` | بر حسب دقیقه |
| `-atime` / `-amin` | آخرین دسترسی |
| `-ctime` / `-cmin` | تغییر inode (مجوز/مالک و …) |
| `-newer FILE` | جدیدتر از FILE |
| `-newermt 'YYYY-MM-DD'` | جدیدتر از تاریخ (GNU) |

## مجوز و مالک

| عبارت | توضیح |
|-------|--------|
| `-user NAME` / `-uid N` | مالک |
| `-group NAME` / `-gid N` | گروه |
| `-perm 644` | مجوز دقیقاً |
| `-perm -644` | همهٔ بیت‌های 644 روشن باشند |
| `-perm /644` | هر یک از بیت‌ها |
| `-readable` / `-writable` / `-executable` | برای کاربر فعلی (GNU) |

## عمل‌ها (Actions)

| عبارت | توضیح |
|-------|--------|
| `-print` | چاپ مسیر + newline (پیش‌فرض) |
| `-print0` | جدا با `\0` (امن برای `xargs -0`) |
| `-printf FORMAT` | خروجی قالب‌دار |
| `-ls` | شبیه `ls -dils` |
| `-delete` | حذف؛ بهتر است آخر عبارت باشد |
| `-exec CMD {} \;` | برای هر فایل یک بار CMD |
| `-exec CMD {} +` | چند فایل در یک فراخوانی (سریع‌تر) |
| `-execdir CMD {} \;` | اجرا از داخل پوشهٔ فایل |
| `-ok CMD {} \;` | مثل exec با تأیید تعاملی |
| `-quit` | بعد از اولین match متوقف شو (GNU) |

## عملگرهای منطقی

| عملگر | معنی |
|--------|------|
| `expr1 expr2` | AND ضمنی |
| `-a` / `-o` | AND / OR صریح |
| `!` یا `-not` | نقیض |
| `( )` | گروه‌بندی (در شل escape: `\( \)`) |
| `,` | هر دو را اجرا کن؛ نتیجهٔ سمت راست |

# Examples

## جستجوی روزمره

```bash
# همهٔ .log زیر /var/log
find /var/log -type f -name '*.log'

# نام بدون حساسیت
find . -iname 'readme*'

# فقط پوشه‌ها تا عمق ۲
find . -maxdepth 2 -type d

# فایل‌های > 100MB در home
find /home -xdev -type f -size +100M -ls
```

## زمان و پاکسازی

```bash
# تغییر در ۷ روز اخیر
find /var/log -type f -mtime -7

# قدیمی‌تر از ۳۰ روز — اول ببین، بعد حذف
find /tmp -type f -mtime +30 -print
find /tmp -type f -mtime +30 -delete

# خالی
find . -type d -empty
find . -type f -empty -delete
```

## exec و pipeline امن

```bash
# chmod دسته‌ای
find . -type f -name '*.sh' -exec chmod 755 {} +

# حذف امن با فاصله در نام
find . -type f -name '*.tmp' -print0 | xargs -0 rm -f

# grep داخل نتایج find
find . -type f -name '*.py' -print0 | xargs -0 grep -n 'TODO'
# یا:
find . -type f -name '*.py' -exec grep -Hn 'TODO' {} +
```

## مجوز و مالک

```bash
find /var/www -type f -perm 777 -ls
find /home -type f -user alice -group staff
find . -type f ! -perm -644
```

## OR و نقیض

```bash
find . -type f \( -name '*.jpg' -o -name '*.png' \)
find . -type f ! -name '*.git*' -a ! -path './node_modules/*'
```

# Common mistakes

- `-name *.log` بدون کوتیشن → شل زودتر expand می‌کند.
- `-delete` یا `-exec rm` بدون یک‌بار `-print` آزمایشی.
- `find /` بدون `-xdev` / `-maxdepth` → کند و خطرناک.
- `-size 100M` یعنی دقیقاً آن اندازه؛ برای «حداقل» از `+100M` استفاده کنید.
- گذاشتن `-delete` قبل از آزمون‌ها می‌تواند رفتار غیرمنتظره بدهد (آخر بگذارید).

# Tips

- برای فقط نام در کل سیستم: `locate` / `plocate` بعد از `updatedb`.
- جایگزین مدرن با سینتکس ساده‌تر: `fd`.
- `-exec ... {} +` معمولاً از `\;` و از حلقهٔ شل سریع‌تر است.
- خروجی برای انسان: `-printf '%p %u %s\n'`.

# Related commands

- `locate` / `updatedb` — جستجوی ایندکسی
- `fd` — جایگزین کاربرپسند
- `grep -r` — جستجوی **محتوا**
- `xargs` — آرگومان از stdin
- `stat` — جزئیات یک inode
