# Changelog

## [Unreleased]

### Security
- ارتقای وابستگی `golang.org/x/net` برای رفع ۳ از ۴ آسیب‌پذیری گزارش‌شده توسط Dependabot (از `v0.17.0` به `v0.50.0` از طریق replace به mirror گیت‌هاب)
- آسیب‌پذیری چهارم (`CVE-2026-25680`, شدت متوسط) نیاز به `golang.org/x/net@v0.55.0` دارد که خودش نیازمند Go `>= 1.25.0` است؛ رفع کامل آن نیازمند ارتقای toolchain در محیطی با دسترسی کامل به `golang.org` است

### Content
- ۱۵ صفحه جدید دیگر (به‌جز `sftp` که هم‌زمان در نسخه ۰.۲.۲ اضافه شده بود):
  - متن: `nl`, `rev`, `shuf`, `fold`, `csplit`
  - فایل‌سیستم/دسترسی: `install`, `chattr`, `lsattr`
  - شبکه/امنیت: `arp`, `route`, `iptables`
  - توسعه: `go`, `yarn`, `node`, `python3`
- ۲۸ صفحه پیشین (دور اول): سیستم/سخت‌افزار (`dmesg`, `lscpu`, `lsusb`, `lspci`, `cal`)، پردازش (`htop`, `killall`, `nice`, `renice`, `strace`, `at`)، متن (`printf`, `paste`, `join`, `comm`, `column`, `split`, `bc`)، فایل‌سیستم (`dd`, `mktemp`, `readlink`, `shred`)، شبکه/امنیت (`openssl`, `ssh-keygen`, `gpg`, `whois`, `nslookup`, `netstat`)
- مجموع صفحات به ۱۹۱+ رسید

## [0.2.2] — 2026-08-05 (planned release)

### Added
- `search --cat` و completion نام صفحات / دسته‌ها
- `.deb` (amd64) در workflow ریلیز + bash completion داخل پکیج
- `faman update --pages --verify` با تأیید SHA256SUMS از Release
- صفحات جدید: xz · zstd · more · whereis · tree · ncdu · timeout · blkid · fdisk · parted · sftp · host · mtr
- ماژول zsh: aliasهای list/doctor/update، `faman-recomplete`، تشخیص completion
- مستندات: docs/completion.md، به‌روزرسانی zsh.md و update.md

### Fixed
- workflow ریلیز سازگار با immutable releases
- فیلد JSON `assets` در تأیید ریلیز

## [0.2.1] — 2026-08-04

ریلیز پایدار پس از محدودیت immutable روی v0.2.0؛ آرتیفکت کامل + SHA256SUMS.

## [0.2.0] — 2026-08-04

اولین انتشار پایدار: list · categories · random · doctor · update · ~۱۶۳ صفحه.

## [0.1.4-pre] — 2026-08-04

پیش‌نمایش: packaging، docs، عمق man.

## [0.1.0] — earlier

CLI اولیه.
