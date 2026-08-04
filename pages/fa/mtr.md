---
title: mtr
aliases:
category: network
difficulty: intermediate
keywords:
- traceroute
- ping
- latency
---

# Introduction

`mtr` ترکیب **traceroute + ping** است؛ هر hop را به‌صورت زنده با آمار loss/latency نشان می‌دهد.

# Syntax

```
mtr [OPTIONS] HOST
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-r` / `--report` | یک گزارش و خروج |
| `-c N` | تعداد ping در report |
| `-n` | بدون resolve نام |
| `-4` / `-6` | IPv4 / IPv6 |
| `-w` | wide report |

# Examples

```bash
mtr 1.1.1.1
mtr -r -c 20 example.com
mtr -n4 github.com
```

# Common mistakes

- loss روی hop میانی همیشه مشکل نیست (نرخ ICMP محدود).

# Tips

- برای تیکت به ISP خروجی `--report` بفرستید.

# Related commands

- `traceroute` · `ping` · `ip`
