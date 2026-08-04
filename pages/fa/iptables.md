---
title: iptables
aliases:
category: network
difficulty: advanced
keywords:
- firewall
- security
- network
- rules
---

# Introduction

`iptables` ابزار سنتی مدیریت فایروال در لینوکس است که قوانین فیلترکردن، هدایت (NAT) و مدیریت بسته‌های شبکه را در سطح کرنل (netfilter) تعریف می‌کند. در توزیع‌های جدید به‌تدریج جای خود را به `nftables` می‌دهد، اما هنوز بسیار رایج است.

⚠️ تغییر قوانین فایروال روی یک سرور راه دور می‌تواند دسترسی SSH شما را قطع کند؛ همیشه با احتیاط و ترجیحاً با یک قانون بازگشتی (rollback) کار کنید.

# Syntax

```
iptables [OPTIONS] -A CHAIN RULE
```

# Options

| گزینه | توضیح |
|-------|--------|
| `-A CHAIN` | افزودن قانون به انتهای یک زنجیره (مثل INPUT) |
| `-L` | نمایش قوانین فعلی |
| `-D` | حذف یک قانون |
| `-p PROTO` | تعیین پروتکل (tcp/udp/icmp) |
| `--dport PORT` | پورت مقصد |
| `-j TARGET` | اقدام نهایی (ACCEPT, DROP, REJECT) |
| `-F` | پاک‌کردن تمام قوانین یک زنجیره |

# Examples

```bash
# نمایش تمام قوانین فعلی با شماره خط
sudo iptables -L -n -v --line-numbers

# اجازه‌دادن به ترافیک ورودی روی پورت ۲۲ (SSH)
sudo iptables -A INPUT -p tcp --dport 22 -j ACCEPT

# مسدودکردن یک آدرس IP خاص
sudo iptables -A INPUT -s 203.0.113.5 -j DROP

# اجازه‌دادن به اتصالات موجود (established) که قبلاً برقرار شده‌اند
sudo iptables -A INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT

# حذف یک قانون خاص با شماره
sudo iptables -D INPUT 3

# ذخیره دائمی قوانین (بسته به توزیع متفاوت است)
sudo iptables-save > /etc/iptables/rules.v4
```

# Common mistakes

- ⚠️ تنظیم قانون `DROP` روی همه چیز بدون اجازه‌دادن به SSH که باعث قطع کامل دسترسی به سرور می‌شود.
- فراموش‌کردن این‌که قوانین `iptables` معمولاً پس از ریبوت پاک می‌شوند مگر با `iptables-save`/`netfilter-persistent` ذخیره شوند.
- ترتیب قوانین اهمیت دارد؛ اولین قانون منطبق اجرا می‌شود، نه لزوماً دقیق‌ترین.

# Tips

- قبل از اعمال قوانین محدودکننده روی سرور راه دور، یک قانون موقت با تایمر (`at now + 5 minutes 'iptables -F'`) تنظیم کنید تا در صورت قفل‌شدن دسترسی، خودکار بازگردد.
- برای مدیریت آسان‌تر روی سرورهای اوبونتو/دبیان، `ufw` رابط ساده‌تری روی iptables ارائه می‌دهد.

# Related commands

- `ufw` — رابط ساده‌تر مدیریت فایروال (روی iptables)
- `nftables` / `nft` — جانشین مدرن‌تر iptables
- `ss` — بررسی پورت‌های باز که نیاز به فیلترشدن دارند
