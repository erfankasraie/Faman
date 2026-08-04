# Roadmap — faman

به‌روز: **۲۰۲۶-۰۸-۰۴** · نسخه هدف ریلیز: **`0.1.4-pre`**

---

## وضعیت فعلی

| محور | وضعیت |
|------|--------|
| نسخه | `0.1.4-pre` |
| صفحات | ~۱۳۵+ |
| `faman update` | چک GitHub + `--pages` |
| نصب | `get.sh` / `install.sh` / `install.ps1` |
| عمق man | find/grep/sed/awk/tar/rsync/curl |
| CI / Release workflow | فعال |

---

## همین ریلیز (v0.1.4-pre) — انجام‌شده / در حال بستن

- [x] update واقعی (check + pages)
- [x] مدیران بسته و محیط مجازی
- [x] عمق man دستهٔ اول
- [x] مستندات install / update / README / CHANGELOG
- [ ] تگ `v0.1.4-pre` + آرتیفکت لینوکس/ویندوز روی GitHub Releases
- [ ] عمق man دستهٔ دوم: ssh, git, docker, systemctl, journalctl, chmod, ip, tmux, jq, make

---

## مرحله بعد — کوتاه‌مدت (۱–۲ هفته)

1. **محتوا:** غنی‌سازی man سنگین باقی‌مانده + ۲۰ صفحه کمبود (dd, htop, ssh-keygen, xz, …)
2. **کیفیت:** `scripts/check-pages.sh` (front matter، حداقل طول)
3. **CLI:** `faman list` / `faman categories` / `faman doctor`
4. **Release:** goreleaser یا همین workflow روی هر تگ پایدار

## میان‌مدت — v0.2.0

- search با `--cat`
- completion نام صفحات
- update باینری از آرتیفکت Release (اختیاری)
- ≥۱۵۰ صفحه

## بلندمدت — v0.3 → v1.0

- AUR / deb
- checksum صفحات
- TUI بعد از قفل فرمت صفحه
- ≥۲۰۰ صفحه باکیفیت

---

## اصول

1. عمق man سنگین > تعداد خام
2. نصب زیر ۲ دقیقه
3. صفحات جدا از باینری قابل‌به‌روزرسانی (`update --pages`)
