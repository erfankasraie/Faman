# Roadmap — faman

نقشهٔ راه دقیق بر اساس وضعیت مخزن در **اوت ۲۰۲۶**.

---

## وضعیت فعلی (Snapshot)

| محور | وضعیت |
|------|--------|
| نسخهٔ کد | `0.1.0` (در `internal/app`) |
| صفحات فارسی | **~۱۱۲** فایل در `pages/fa/` |
| CLI | cobra: show / search / version / help / update (placeholder) / completion |
| رندر | glamour + lipgloss؛ `FAMAN_PLAIN` / `FAMAN_WRAP` |
| نصب | `scripts/install.sh` (یک‌خطی، RTL، user، uninstall، config) |
| شل | completion، zsh module، fzf، bat preview |
| تست / CI | parser / search / renderer + GitHub Actions |
| بسته‌بندی توزیع | هنوز نیست (AUR / deb / rpm) |
| `faman update` | placeholder |
| TUI | نیست |

**معیار پوشش «ابزار رایج»:** حدود ۲۰۰–۲۵۰ دستور هسته‌ای POSIX/GNU + ادمین روزمره.  
با ~۱۱۲ صفحه ≈ **۴۰–۵۰٪ مسیر محتوا تا v1.0** (بیش از هدف اولیه‌ٔ «یک‌سوم»).

---

## اصول اولویت

1. **کیفیت صفحه > تعداد خام** — قالب ۷ بخشی، مثال واقعی، اشتباه رایج.
2. **قابل‌نصب و قابل‌خواندن روی Ubuntu** — فارسی در ترمینال، یک‌خطی install.
3. **CLI پایدار قبل از TUI/AI** — فرمت صفحه و مسیر pages قفل شود.
4. **محتوا منبع حقیقت است** — ویژگی‌های فانتزی بدون پوشش محتوا ارزش کمی دارند.

---

## فاز ۰ — تثبیت قبل از ریلیز عمومی (v0.1.2)  ·  ۱–۲ هفته

هدف: یک تگ قابل معرفی به جامعه بدون شرمساری.

### محتوا
- [ ] ممیزی همهٔ صفحات: front matter کامل (`title`, `category`, `difficulty`, `aliases`, `keywords`)
- [ ] یکنواختی قالب ۷ بخشی روی صفحات کوتاه (<۸۰۰ بایت)
- [ ] غنی‌سازی ۲۰ صفحهٔ پربازدید: `ls`, `grep`, `find`, `chmod`, `systemctl`, `docker`, `git`, `ssh`, `curl`, `tar`, `jq`, `tmux`, `rsync`, `sed`, `awk`, `ps`, `ip`, `ufw`, `crontab`, `make`
- [ ] `docs/pages-index.md` دسته‌بندی‌شده و هم‌تراز با فایل‌های واقعی

### محصول
- [ ] هم‌راستا کردن نسخه: `version` در کد، `CHANGELOG`, تگ git `v0.1.2`
- [ ] اسکرین‌شات واقعی ترمینال در README (GNOME + FAMAN_PLAIN)
- [ ] لینک‌های README به `erfankasraie/Faman` (نه org فرضی)
- [ ] `go module` path یکدست با ریپوی واقعی

### کیفیت
- [ ] اسکریپت `scripts/check-pages.sh`: YAML front matter، عنوان، حداقل طول
- [ ] تست طلایی رندر برای ۲–۳ صفحهٔ نمونه (snapshot اختیاری)

**خروجی فاز ۰:** تگ `v0.1.2` + اعلام در README که «پیش‌نمایش عمومی محتوا».

---

## فاز ۱ — پوشش نیمه‌راه محتوا (v0.2.0-content)  ·  موازی با فاز ۰/۲

هدف: **~۱۵۰–۱۶۰ صفحه** (نیمهٔ مسیر تا ۲۰۰+).

### دسته‌های کمبود محتمل (اولویت اضافه شدن)

| دسته | نمونه‌های پیشنهادی |
|------|---------------------|
| فایل / دیسک | `dd`, `fdisk`, `parted`, `mkfs`, `blkid`, `sync`, `truncate`, `install` |
| متن / داده | `column`, `paste`, `join`, `comm`, `fold`, `fmt`, `nl`, `od`, `hexdump`, `strings` |
| فرایند / perf | `htop`, `nice`, `renice`, `strace`, `ltrace`, `time` (عمیق), `timeout`, `pidof` |
| شبکه | `nmap` (سطح پایه), `curl` عمیق‌تر, `wget` , `ip` route/neigh, `nft`, `firewall-cmd`, `resolvectl`, `sshd` |
| کاربر / امن | `userdel`, `usermod`, `groupadd`, `visudo`, `passwd` عمیق, `ssh-keygen`, `chmod` ACL/`setfacl` |
| systemd | `timedatectl` ✓, `hostnamectl` ✓, `loginctl`, `resolvectl`, `systemd-analyze` |
| بسته | `dpkg`, `apt-cache`, `snap`, `flatpak`, `pip`, `npm` (سطح CLI) |
| dev | `gcc` پایه, `python`/`python3` CLI, `git` زیرصفحه یا گسترش, `cargo`/`go` خلاصه |
| فشرده‌سازی | `xz`, `bzip2`, `zstd`, `7z` |

### قوانین صفحهٔ جدید
- حداقل: مقدمه، نحو، گزینه‌های پرکاربرد، ۳ مثال، اشتباه رایج، مرتبط‌ها
- بدون کپی ترجمه‌ٔ man خام؛ آموزش‌محور

**خروجی:** `CHANGELOG` با «+N pages» و به‌روز `pages-index`.

---

## فاز ۲ — CLI محصولی (v0.2.0)  ·  ۲–۴ هفته

هدف: ابزار «تمام‌روز» بدون TUI.

- [ ] `faman list` / `faman ls-pages` — فهرست با فیلتر category و difficulty
- [ ] `faman random` — صفحه تصادفی (یادگیری)
- [ ] `faman categories` — درخت دسته‌ها
- [ ] بهبود search: امتیاز شفاف‌تر، فیلتر `--cat`
- [ ] `faman doctor` — locale، مسیر pages، فونت؟، نسخه Go runtime نه؛ فقط محیط اجرا
- [ ] embed اختیاری صفحات در باینری (`go:embed`) برای `go install` بدون کپی pages
- [ ] completion غنی‌تر: نام صفحات به‌عنوان arg اول
- [ ] `faman update` حداقل: چک release GitHub + پیام «صفحات را با install دوباره بگیر»

**خروجی:** نسخهٔ semver `v0.2.0` با یادداشت مهاجرت.

---

## فاز ۳ — به‌روزرسانی و توزیع (v0.3.0)  ·  ۳–۵ هفته

- [ ] `faman update` واقعی: دانلود tarball صفحات یا clone depth=1 به `~/.local/share/faman`
- [ ] امضا / checksum ساده (SHA256 مجموعه‌صفحات)
- [ ] GitHub Release با باینری linux-amd64/arm64 (goreleaser)
- [ ] بستهٔ **AUR** (`faman-git` / `faman`)
- [ ] `.deb` ساده (nfpm یا goreleaser nFPM)
- [ ] Homebrew optional (linuxbrew) — اولویت پایین

**خروجی:** کاربر غیرتوسعه‌دهنده بدون `go build` نصب کند و محتوا را به‌روز کند.

---

## فاز ۴ — TUI (v0.4.0)  ·  بعد از پایدار شدن فرمت صفحه

- [ ] bubbletea: لیست صفحات، جستجوی زنده، نمایش با viewport
- [ ] علاقه‌مندی‌ها و تاریخچه در `~/.config/faman/`
- [ ] حالت `faman tui` بدون شکستن CLI کلاسیک
- [ ] احترام به `FAMAN_PLAIN` در TUI

**وابستگی:** فاز ۰–۲ (فرمت صفحه و list API).

---

## فاز ۵ — گسترش و اکوسیستم (v0.5 → v1.0)

- [ ] ۲۰۰+ صفحهٔ باکیفیت
- [ ] مشارکت‌کنندهٔ بیرونی: قالب Issue/PR، `good first page`
- [ ] صفحات کاربر: `~/.config/faman/pages/fa` با اولویت merge
- [ ] (اختیاری) لایهٔ AI جدا — **نه** در هسته؛ فقط اگر API و حریم مشخص باشد
- [ ] قفل schema فرمت صفحه برای v1.0
- [ ] اعلام v1.0 با سیاست پشتیبانی semver

---

## چیزهایی که فعلاً عمداً عقب می‌مانند

| موضوع | دلیل |
|--------|------|
| AI داخل باینری | پیچیدگی، هزینه، حواس‌پرتی از محتوا |
| RTL کامل در همه ترمینال‌ها | محدودیت امولاتور؛ تمرکز روی plain + فونت + docs |
| چندزبانه | اول فارسی را عمیق کنیم |
| پلاگین باینری پویا | زود است؛ صفحات کاربری کافی است |

---

## متریک‌های موفقیت

| متریک | هدف نزدیک (v0.2) | هدف v1.0 |
|--------|------------------|----------|
| تعداد صفحات fa | ≥ ۱۵۰ | ≥ ۲۰۰ |
| صفحات «کامل» (قالب + ≥۳ مثال) | ≥ ۸۰٪ | ≥ ۹۰٪ |
| نصب بدون Go (release binary) | بله | بله |
| `faman update` محتوا | حداقل پیام / ideal دانلود | پایدار |
| زمان تا اولین `faman ls` برای کاربر جدید | < ۲ دقیقه | < ۱ دقیقه |

---

## ترتیب پیشنهادی کار هفتگی

| هفته | تمرکز |
|------|--------|
| ۱ | ممیزی front matter + ۲۰ صفحهٔ پربازدید + check-pages + نسخه ۰.۱.۲ |
| ۲ | ۲۰–۳۰ صفحهٔ جدید از جدول کمبود + list/categories CLI |
| ۳ | goreleaser + update اسکلت + README اسکرین‌شات |
| ۴ | embed pages یا مسیر user share پایدار + doctor |
| ۵+ | TUI یا محتوای ۱۵۰+ بر اساس بازخورد |

---

## تصمیم‌های باز (باید با maintainer بسته شود)

1. **Module path نهایی:** `github.com/erfankasraie/Faman` در برابر `faman-project/faman`
2. **آیا صفحات داخل باینری embed شوند؟** (ساده برای `go install`، سخت برای update جزئی)
3. **سطح عمق man:** آموزشی کوتاه در برابر مرجع کامل
4. **نام دستور list:** `faman list` در برابر `faman pages`

---

آخرین به‌روزرسانی رودمپ: ۲۰۲۶-۰۸-۰۲  
منبع وضعیت: درخت `main`، ~۱۱۲ صفحه، installer/zsh/fzf/bat، CLI ۰.۱.۰
