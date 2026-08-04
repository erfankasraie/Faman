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
| عمق man (انجام‌شده) | find · grep · sed · awk · tar · rsync · curl |
| CI | GitHub Actions + golangci-lint |

**درصد تقریبی مسیر محتوا تا v1.0:** ~۵۰–۵۵٪ تعداد · عمق هنوز ناهمگن.

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
| 0.6 | CI سبز روی main | Actions success | [x] پس از fix lint |

**خروجی:** Pre-release عمومی قابل لینک‌دادن.

---

### فاز ۱ — عمق man دستهٔ دوم  ·  **← الان اینجا**  ·  ۳–۷ روز

هدف: همان سطح `find`/`curl` برای دستورات پربازدید باقی‌مانده.

| # | صفحه | تمرکز محتوا | وضعیت |
|---|------|-------------|--------|
| 1.1 | `ssh` | کلید، config، تونل، scp/sftp، ProxyJump | [ ] در حال |
| 1.2 | `scp` / ارتباط با ssh | مسیر remote، `-r`، محدودیت | [ ] |
| 1.3 | `git` | clone/commit/branch/remote/log روزمره | [ ] در حال |
| 1.4 | `docker` | run/build/ps/images/compose اشاره | [ ] در حال |
| 1.5 | `systemctl` | start/enable/status/unit types | [ ] در حال |
| 1.6 | `journalctl` | فیلتر unit/time/follow | [ ] در حال |
| 1.7 | `chmod` / اشاره ACL | نمادین و عددی، `-R` | [ ] در حال |
| 1.8 | `ip` | addr/link/route/neigh | [ ] |
| 1.9 | `tmux` | session/window/pane، کلیدها | [ ] |
| 1.10 | `jq` | filter پایه، pipe، از فایل | [ ] |
| 1.11 | `make` | target، متغیر، `-j` | [ ] |
| 1.12 | `crontab` | فرمت زمان، crontab -e | [ ] |

**معیار پذیرش فاز ۱:** هر صفحه ≥ ~۲.۵KB یا معادل جداول+مثال؛ کاربر با `faman ssh` سناریوی عملی ببیند.

---

### فاز ۲ — کیفیت صفحات و ابزار ممیزی  ·  ۳–۵ روز

| # | کار | جزئیات |
|---|-----|--------|
| 2.1 | `scripts/check-pages.sh` | front matter اجباری، بخش‌های ۷گانه، حداقل بایت |
| 2.2 | اجرای check در CI | fail روی صفحهٔ ناقص جدید |
| 2.3 | یکنواختی aliases/keywords | جستجوی بهتر |
| 2.4 | ممیزی صفحات < ۱KB | لیست و غنی‌سازی |

---

### فاز ۳ — CLI محصولی `v0.2.0`  ·  ۲–۴ هفته

| # | دستور / قابلیت | جزئیات |
|---|----------------|--------|
| 3.1 | `faman list` | فهرست صفحات، `--cat`، `--diff` |
| 3.2 | `faman categories` | شمارش هر دسته |
| 3.3 | `faman random` | یادگیری تصادفی |
| 3.4 | `faman doctor` | locale، مسیر pages، نسخه، نوشتنی بودن |
| 3.5 | search `--cat` | فیلتر دسته |
| 3.6 | completion | نام فایل‌های pages به‌عنوان arg |
| 3.7 | ≥ ۱۵۰ صفحه | فاز ۱ + ۲۰ صفحهٔ جدید |

**خروجی:** تگ `v0.2.0` با یادداشت مهاجرت.

---

### فاز ۴ — صفحات کمبود (محتوا)  ·  موازی با ۳

اولویت اضافه شدن:

| دسته | نمونه‌ها |
|------|----------|
| دیسک | `dd`, `fdisk`, `lsblk`✓, `mkfs`, `blkid`, `sync` |
| متن | `column`, `paste`, `nl`, `hexdump`, `strings` |
| perf | `htop`, `nice`, `strace`, `timeout`, `pidof` |
| شبکه | `ssh-keygen`, `nmap` پایه, `nft`/`ufw`✓, `resolvectl` |
| فشرده | `xz`, `zstd`, `bzip2` |
| امن | `setfacl`, `visudo`, `userdel` |

---

### فاز ۵ — توزیع و update قوی `v0.3.0`

| # | کار |
|---|-----|
| 5.1 | آرتیفکت deb روی Release workflow |
| 5.2 | nFPM/goreleaser برای deb+rpm |
| 5.3 | `faman update` باینری از Release (اختیاری) |
| 5.4 | SHA256 اجباری برای صفحات |
| 5.5 | اسکلت AUR و یادداشت Homebrew |

---

### فاز ۶ — TUI `v0.4.0`

فقط بعد از قفل نسبی فرمت صفحه (فاز ۲).

bubbletea: لیست، جستجو، viewport، احترام به `FAMAN_PLAIN`.

---

### فاز ۷ — v1.0

| معیار | هدف |
|--------|------|
| صفحات باکیفیت | ≥ ۲۰۰ |
| صفحات عمیق man سنگین | ≥ ۲۵ |
| نصب بدون Go | بله |
| schema صفحه | قفل‌شده در CONTRIBUTING |
| semver | سیاست پشتیبانی مشخص |

---

## ۴. ترتیب کار هفتگی پیشنهادی

| هفته | تمرکز |
|------|--------|
| **جاری** | فاز ۱: ssh…chmod (شروع در همین PR/commitها) |
| +۱ | اتمام فاز ۱ (ip…crontab) + check-pages.sh |
| +۲ | list/categories/doctor |
| +۳ | ۲۰ صفحهٔ جدید + release v0.2.0 |
| +۴ | deb در CI / update binary |

---

## ۵. تصمیم‌های باز

1. آیا `go:embed` برای pages؟ (آسان برای `go install`، سخت برای update جزئی)
2. نام رسمی: `faman list` در برابر `faman pages`
3. عمق git: یک صفحه در برابر `git-commit` / `git-branch` جدا

---

## ۶. خارج از محدوده (عمداً عقب)

- AI داخل باینری
- RTL کامل همهٔ ترمینال‌ها
- چندزبانه‌سازی قبل از عمق فارسی
"}, {