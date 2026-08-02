---
title: chgrp
aliases:
category: permissions
difficulty: intermediate
keywords:
- group
- ownership
- permissions
---

# Introduction

`chgrp` گروه مالک فایل را تغییر می‌دهد.

# Syntax

```
chgrp [OPTIONS] GROUP FILE...
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-R` | بازگشتی |
| `-v` | verbose |
| `-h` | تغییر symlink نه مقصد |

# Examples

```bash
sudo chgrp www-data /var/www/html
sudo chgrp -R dev /home/project
```

# Common mistakes

- نداشتن مجوز برای تغییر گروه.

# Tips

- `chown user:group` هر دو را با هم عوض می‌کند.

# Related commands

- `chown`
- `chmod`
- `id` / `groups`
