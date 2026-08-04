# Roadmap — faman

آخرین به‌روزرسانی: **۲۰۲۶-۰۸-۰۴**  
نسخهٔ کد فعلی: **`0.1.4-pre`**  
معیار پوشش کامل محتوا: حدود **۲۰۰–۲۵۰** دستور رایج

> هماهنگی چند contributor/AI: قبل از edit گسترده روی `pages/fa` یا README، `git pull` و commitهای اخیر را چک کنید.

---

## ۱. وضعیت فعلی (Snapshot)

| محور | مقدار |
|------|--------|
| صفحات `pages/fa` | ~۱۶۰+ (فاز محتوا موازی) |
| CLI | show · search · version · help · update · completion |
| `faman update` | `--check` + `--pages` |
| نصب / بسته‌بندی | get.sh · install · package.sh · docs پلتفرم |
| عمق man | دستهٔ ۱ + دستهٔ ۲ (فاز ۱) |
| کیفیت صفحات | `scripts/check-pages.sh` + گام CI |
| CI | Actions + golangci-lint + check-pages |

---

## ۲. اصول اولویت

1. عمق man سنگین > تعداد خام  
2. قالب ۷ بخشی آموزشی  
3. نصب < ۲ دقیقه  
4. صفحات جدا با `update --pages`  
5. هر فاز معیار پذیرش دارد  
6. **از تداخل فایل اجتناب:** تغییرات اسکریپت/CI جدا از bulk content

---

## ۳. فازها

### فاز ۰ — `v0.1.4-pre` · تقریباً تمام

| # | کار | وضعیت |
|---|-----|--------|
| 0.1–0.4, 0.6 | نسخه، docs، update، packaging، CI lint | [x] |
| 0.5 | تگ + آرتیفکت GitHub Releases | [ ] maintainer |

### فاز ۱ — عمق man دستهٔ دوم · **[x] تمام**

ssh · scp · git · docker · systemctl · journalctl · chmod · ip · tmux · jq · make · crontab  
(+ قبلی: find · grep · sed · awk · tar · rsync · curl)

### فاز ۲ — کیفیت و ممیزی · **در حال انجام**

| # | کار | وضعیت |
|---|-----|--------|
| 2.1 | `scripts/check-pages.sh` | [x] |
| 2.2 | اتصال به CI (`Check pages schema`) | [x] |
| 2.3 | یکنواختی aliases/keywords | [ ] |
| 2.4 | ممیزی صفحات کم‌حجم / ناقص | [x] baseline (nslookup/whois) |

### فاز ۳ — CLI `v0.2.0` · بعدی بعد از تثبیت ۲.۳

`list` · `categories` · `random` · `doctor` · search `--cat` · completion · ≥۱۵۰ صفحه (تعداد الان قبلاً عبور کرده)

### فاز ۴ — صفحات کمبود

ادامهٔ پوشش موازی؛ هماهنگ با contributor دیگر تا duplicate title نشود.

### فاز ۵–۷

توزیع deb/rpm · TUI · v1.0

---

## ۴. ترتیب هفتگی

| هفته | تمرکز |
|------|--------|
| جاری | تثبیت فاز ۲ + عدم تداخل با صفحات جدید |
| +۱ | `faman list` / `doctor` |
| +۲ | آماده‌سازی v0.2.0 + Release |

---

## ۵. خارج از محدوده

AI داخل باینری · RTL کامل همه‌جا · i18n زودهنگام
