---
title: basename
aliases:
category: shell
difficulty: beginner
keywords:
- path
- filename
---

# Introduction

`basename` نام فایل را از مسیر کامل جدا می‌کند.

# Syntax

```
basename PATH [SUFFIX]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-s SUFFIX` | حذف پسوند |
| `-a` | چند آرگومان |

# Examples

```bash
basename /var/log/syslog          # syslog
basename /home/a/report.tar.gz .gz
basename -s .md /docs/page.md
```

# Common mistakes

- انتظار داشتن مسیر باقی‌مانده — برای آن `dirname` است.

# Tips

- در bash می‌توان از `${file##*/}` هم استفاده کرد.

# Related commands

- `dirname`
- `realpath`
- `readlink`
