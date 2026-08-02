---
title: jq
aliases:
category: text
difficulty: intermediate
keywords:
- json
- filter
- parse
---

# Introduction

`jq` پردازشگر سبک JSON در خط فرمان است؛ برای API و لاگ‌های JSON ضروری است.

# Syntax

```
jq [OPTIONS] FILTER [FILE...]
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-r` | خروجی raw (بدون کوتیشن رشته) |
| `-c` | فشرده یک خطی |
| `-e` | exit code بر اساس نتیجه |
| `-s` | slurp همه ورودی به آرایه |

# Examples

```bash
echo '{"name":"ali","age":30}' | jq .
echo '{"name":"ali"}' | jq -r .name
curl -s https://api.github.com/repos/erfankasraie/Faman | jq -r .full_name
jq '.items[].name' data.json
jq 'map(select(.ok==true))' list.json
```

# Common mistakes

- فراموش کردن `-r` وقتی رشته را به دستور بعدی می‌دهید.

# Tips

- فیلتر را در فایل نگه دارید: `jq -f filter.jq data.json`

# Related commands

- `curl` — دریافت JSON
- `python -m json.tool` — pretty-print ساده
- `yq` — برای YAML
