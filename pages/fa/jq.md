---
title: jq
aliases:
category: text
difficulty: intermediate
keywords:
- json
- filter
- parse
- api
---

# Introduction

`jq` پردازشگر **JSON** در خط فرمان است: انتخاب فیلد، فیلتر، map، و خروجی خوانا. جفت ایده‌آل `curl` برای APIها.

# Syntax

```
jq [OPTIONS] FILTER [FILE...]
... | jq FILTER
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-r` | خروجی raw (بدون کوتیشن رشته) |
| `-c` | فشرده یک‌خطی |
| `-S` | مرتب‌سازی کلیدها |
| `-n` | بدون ورودی؛ فقط FILTER |
| `-e` | کد خروج خطا اگر false/null |
| `-s` | کل ورودی به صورت یک آرایه |
| `-f FILE` | فیلتر از فایل |
| `--arg k v` | متغیر رشته از شل |
| `--argjson k v` | متغیر JSON |

## فیلترهای پایه

| فیلتر | معنی |
|--------|------|
| `.` | کل سند |
| `.key` / `."key-with-dash"` | فیلد |
| `.key.sub` | تو در تو |
| `.[0]` | عنصر آرایه |
| `.[]` | iterate آرایه/object |
| `.key?` | بدون خطا اگر نباشد |
| `select(cond)` | فیلتر شرطی |
| `map(f)` | اعمال روی آرایه |
| `keys` / `length` | کلیدها / طول |
| `+` / `//` | ادغام / پیش‌فرض |

# Examples

## خواندن و انتخاب

```bash
echo '{"name":"faman","stars":10}' | jq .
echo '{"name":"faman","stars":10}' | jq -r .name
curl -s https://api.github.com/repos/erfankasraie/Faman | jq '{name, stars: .stargazers_count}'
```

## آرایه

```bash
echo '{"items":[{"id":1},{"id":2}]}' | jq '.items[].id'
echo '[1,2,3]' | jq 'map(. * 2)'
echo '[{"n":1},{"n":5},{"n":3}]' | jq '[.[] | select(.n > 2)]'
```

## از فایل و شل

```bash
jq '.dependencies' package.json
jq --arg env prod '.config + {env: $env}' config.json
```

## با curl در اسکریپت

```bash
code=$(curl -s -o /tmp/body.json -w '%{http_code}' https://api.example.com/health)
jq -e '.ok == true' /tmp/body.json
```

# Common mistakes

- فراموش کردن `-r` وقتی خروجی را به شل می‌دهید (`"value"` با کوتیشن).
- JSON نامعتبر (ویرگول اضافه) → jq خطا می‌دهد.
- انتظار YAML؛ jq فقط JSON (برای YAML ابزار دیگر).

# Tips

- `jq empty` برای اعتبارسنجی خاموش.
- فیلترهای پیچیده را در فایل `.jq` نگه دارید.
- جایگزین: `python -m json.tool` فقط pretty-print.

# Related commands

- `curl` · `grep` · `awk` · `python` · `yq`
