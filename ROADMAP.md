# Roadmap — faman

آخرین به‌روزرسانی: **۲۰۲۶-۰۸-۰۴**  
نسخهٔ کد فعلی: **`0.1.4-pre`**  
معیار پوشش کامل محتوا: حدود **۲۰۰–۲۵۰** دستور رایج

---

## ۱. وضعیت فعلی (Snapshot)

| محور | مقدار |
|------|--------|
| صفحات `pages/fa` | ~۱۳۵ |
| CLI | show · search · version · help · update · completion |
| `faman update` | `--check` + `--pages` |
| نصب / بسته‌بندی | get.sh · install · package.sh · docs پلتفرم |
| عمق man | دستهٔ ۱ + **کل دستهٔ ۲ (فاز ۱ تمام)** |
| CI | Actions + golangci-lint |

---

## ۲. اصول اولویت

1. عمق man سنگین > تعداد خام  
2. قالب ۷ بخشی آموزشی  
3. نصب < ۲ دقیقه  
4. صفحات جدا با `update --pages`  
5. هر فاز معیار پذیرش دارد  

---

## ۳. فازها

### فاز ۰ — `v0.1.4-pre` · تقریباً تمام

| # | کار | وضعیت |
|---|-----|--------|
| 0.1–0.4, 0.6 | نسخه، docs، update، packaging، CI | [x] |
| 0.5 | تگ + آرتیفکت GitHub Releases | [ ] maintainer |

### فاز ۱ — عمق man دستهٔ دوم · **[x] تمام**

| # | صفحه | وضعیت |
|---|------|--------|
| 1.1 | `ssh` | [x] |
| 1.2 | `scp` | [x] |
| 1.3 | `git` | [x] |
| 1.4 | `docker` | [x] |
| 1.5 | `systemctl` | [x] |
| 1.6 | `journalctl` | [x] |
| 1.7 | `chmod` | [x] |
| 1.8 | `ip` | [x] |
| 1.9 | `tmux` | [x] |
| 1.10 | `jq` | [x] |
| 1.11 | `make` | [x] |
| 1.12 | `crontab` | [x] |

به‌علاوه از قبل: find · grep · sed · awk · tar · rsync · curl

### فاز ۲ — کیفیت و ممیزی · **← بعدی**

| # | کار |
|---|-----|
| 2.1 | `scripts/check-pages.sh` |
| 2.2 | اتصال به CI |
| 2.3 | یکنواختی aliases/keywords |
| 2.4 | ممیزی صفحات کم‌حجم |

### فاز ۳ — CLI `v0.2.0`

`list` · `categories` · `random` · `doctor` · search `--cat` · completion · ≥۱۵۰ صفحه

### فاز ۴ — صفحات کمبود

dd · fdisk · htop · strace · ssh-keygen · xz · zstd · setfacl · nmap پایه · …

### فاز ۵–۷

توزیع deb/rpm · TUI · v1.0 (≥۲۰۰ صفحه)

---

## ۴. ترتیب هفتگی

| هفته | تمرکز |
|------|--------|
| **بعدی** | فاز ۲: check-pages.sh + CI |
| +۱ | list / doctor / categories |
| +۲ | +۲۰ صفحه · آماده‌سازی v0.2.0 |
| +۳ | deb در Release |

---

## ۵. خارج از محدوده

AI داخل باینری · RTL کامل همه‌جا · i18n زودهنگام
