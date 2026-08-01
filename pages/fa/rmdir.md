---
title: rmdir
aliases:
- rd
category: filesystem
difficulty: beginner
keywords:
- directory
- remove
- empty
---

# Introduction

دستور `rmdir` فقط پوشه‌های **خالی** را حذف می‌کند. اگر پوشه چیزی داخلش داشته باشد، خطا می‌دهد.

برای حذف پوشه‌های غیرخالی از `rm -r` استفاده کنید.

# Syntax

```
rmdir [OPTIONS] DIRECTORY...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-p` | حذف پوشه‌های والد خالی هم |
| `-v` | نمایش پوشه‌های حذف‌شده |

# Examples

```bash
# حذف یک پوشه خالی
rmdir empty_folder

# حذف زنجیره‌ای پوشه‌های خالی
rmdir -p projects/web/frontend/src
```

# Common mistakes

- تلاش برای حذف پوشه غیرخالی → خطا. از `rm -r` استفاده کنید.
- فراموش کردن اینکه `rmdir` امن‌تر از `rm -r` است.

# Tips

- اگر مطمئن نیستید پوشه خالی است، اول `ls -a` بزنید.
- `rmdir` امن‌تر است چون تصادفاً محتویات را پاک نمی‌کند.

# Related commands

- `rm -r` — حذف پوشه و محتویات
- `mkdir` — ساخت پوشه
- `rm` — حذف فایل
