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
| `faman update` | `--check` + `--pages` (باینری جدا) |
| نصب | `get.sh` · `install.sh` · `install.ps1` |
| بسته‌بندی | `package.sh` → tar.gz / zip.exe / deb محلی |
| مستندات پلتفرم | linux-distros · macos · windows · packaging · update |
| عمق man (انجام‌شده) | find · grep · sed · awk · tar · rsync · curl · **ssh · git · docker · systemctl · journalctl · chmod** |
| CI | GitHub Actions + golangci-lint |

**درصد تقریبی مسیر محتوا تا v1.0:** ~۵۰–۵۵٪ تعداد · عمق در حال تکمیل فاز ۱.

---

## ۲. اصول اولویت

1. **عمق man سنگین > تعداد صفحهٔ خام**
2. قالب آموزشی ۷ بخشی حفظ شود (نه کپی man)
3. نصب کاربر جدید < ۲ دقیقه
4. صفحات با `update --pages` از باینری جدا بمانند
5. هر فاز خروجی قابل‌تست داشته باشد (چک‌باکس + معیار پذیرش)

### سیاست عمق

| نوع دستور | حداقل انتظار |
|-----------|----------------|
| کوتاه (`true`, `yes`) | مقدمه + مثال |
| متوسط | ۷ بخش کامل، ≥۳ مثال |
| سنگین (`ssh`, `git`, `find`, …) | جداول گزینهٔ گسترده + ≥۴ سناریو + اشتباهات |

---

## ۳. فازها

### فاز ۰ — بستن پیش‌انتشار `v0.1.4-pre`  ·  وضعیت: تقریباً تمام

| # | کار | پذیرش | وضعیت |
|---|-----|--------|--------|
| 0.1 | نسخه در کد = 0.1.4-pre | `faman version` | [x] |
| 0.2 | CHANGELOG / README / install / update docs | لینک‌های README درست | [x] |
| 0.3 | update واقعی | `--check` و `--pages` کار کنند | [x] |
| 0.4 | packaging script + docs مک/لینوکس | `docs/macos.md` و `linux-distros.md` | [x] |
| 0.5 | تگ GitHub + آرتیفکت Release | فایل linux/darwin/windows روی Releases | [ ] maintainer |
| 0.6 | CI سبز روی main | Actions success | [x] |

**خروجی:** Pre-release عمومی قابل لینک‌دادن.

---

### فاز ۱ — عمق man دستهٔ دوم  ·  **← الان اینجا**  ·  ۳–۷ روز

هدف: همان سطح `find`/`curl` برای دستورات پربازدید باقی‌مانده.

| # | صفحه | تمرکز محتوا | وضعیت |
|---|------|-------------|--------|
| 1.1 | `ssh` | کلید، config، تونل، ProxyJump | [x] |
| 1.2 | `scp` | مسیر remote، `-r` | [ ] بعدی |
| 1.3 | `git` | clone/commit/branch/remote/log | [x] |
| 1.4 | `docker` | run/build/ps/images/compose | [x] |
| 1.5 | `systemctl` | start/enable/status/unit | [x] |
| 1.6 | `journalctl` | unit/time/follow | [x] |
| 1.7 | `chmod` | نمادین و عددی، `-R`، ACL اشاره | [x] |
| 1.8 | `ip` | addr/link/route/neigh | [ ] بعدی |
| 1.9 | `tmux` | session/window/pane | [ ] بعدی |
| 1.10 | `jq` | filter پایه | [ ] بعدی |
| 1.11 | `make` | target، `-j` | [ ] بعدی |
| 1.12 | `crontab` | فرمت زمان | [ ] بعدی |

**معیار پذیرش فاز ۱:** هر صفحه جداول+مثال عملی؛ `faman ssh` سناریوی واقعی بدهد.

**پیشرفت:** ۶ از ۱۲ انجام شد.

---

### فاز ۲ — کیفیت صفحات و ابزار ممیزی  ·  ۳–۵ روز

| # | کار | جزئیات |
|---|-----|--------|
| 2.1 | `scripts/check-pages.sh` | front matter، ۷ بخش، حداقل بایت |
| 2.2 | check در CI | fail روی صفحهٔ ناقص |
| 2.3 | یکنواختی aliases/keywords | جستجوی بهتر |
| 2.4 | ممیزی صفحات کم‌حجم | غنی‌سازی |

---

### فاز ۳ — CLI محصولی `v0.2.0`  ·  ۲–۴ هفته

| # | قابلیت | جزئیات |
|---|--------|--------|
| 3.1 | `faman list` | `--cat`، `--diff` |
| 3.2 | `faman categories` | شمارش |
| 3.3 | `faman random` | یادگیری |
| 3.4 | `faman doctor` | locale، pages path |
| 3.5 | search `--cat` | فیلتر |
| 3.6 | completion نام صفحات | arg اول |
| 3.7 | ≥ ۱۵۰ صفحه | |

---

### فاز ۴ — صفحات کمبود (موازی)

`dd`, `fdisk`, `mkfs`, `htop`, `strace`, `ssh-keygen`, `xz`, `zstd`, `setfacl`, `nmap` پایه، …

---

### فاز ۵ — توزیع `v0.3.0`

deb در Release · nFPM/rpm · update باینری · SHA256 · AUR/Homebrew اسکلت

---

### فاز ۶ — TUI `v0.4.0`

بعد از فاز ۲ · bubbletea

---

### فاز ۷ — v1.0

≥۲۰۰ صفحه · schema قفل · semver پایدار

---

## ۴. ترتیب هفتگی

| هفته | تمرکز |
|------|--------|
| **جاری** | فاز ۱: اتمام scp · ip · tmux · jq · make · crontab |
| +۱ | فاز ۲ check-pages + CI |
| +۲ | list / categories / doctor |
| +۳ | +۲۰ صفحه · v0.2.0 |
| +۴ | deb در CI |

---

## ۵. تصمیم‌های باز

1. `go:embed` برای pages؟
2. `faman list` در برابر `faman pages`
3. git تک‌صفحه در برابر زیرصفحات

---

## ۶. خارج از محدوده

AI داخل باینری · RTL کامل همه ترمینال‌ها · i18n زودهنگام
