# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.2] — 2026-08-02

### Added

- پوشش گستردهٔ دستورات رایج برای نزدیک شدن به حدود یک‌سوم ابزارهای هسته‌ای CLI
- متن و شل: `sed`, `awk`, `sort`, `uniq`, `cut`, `tr`, `wc`, `tee`, `xargs`, `which`, `env`, `history`, `clear`
- فایل‌سیستم: `ln`, `touch`, `du`, `file`, `stat`, `mount`, `umount`, `lsblk`
- فرایند: `pgrep`, `pkill`, `nohup`, `jobs`
- کاربر: `id`, `whoami`, `passwd`, `useradd`
- شبکه: `ip`, `ss`, `dig`, `traceroute`, `hostname`
- سیستم: `uname`, `uptime`, `date`, `sudo`, `journalctl`, `lsof`, `vim`
- بسته و فشرده‌سازی: `dnf`, `pacman`, `gzip`, `gunzip`
- به‌روزرسانی `ROADMAP.md` و فهرست صفحات

## [0.1.1] — 2026-08-02

### Added

- نقطه ورود `cmd/faman/main.go` (پروژه اکنون به‌درستی build می‌شود)
- صفحات: `echo`, `ps`, `top`, `df`, `free`, `kill`, `ping`, `nano`, `apt`, `systemctl`
- محتوای ویکی در `docs/wiki/`
- گسترش `ROADMAP.md`

### Fixed

- `internal/update/update.go` اکنون از `strings.Repeat` استاندارد استفاده می‌کند

## [0.1.0] — 2026-08-01

### Added

- اولین انتشار عمومی
- دستورات: `faman <cmd>`, `search`, `version`, `help`, `update`
- پارسر Markdown با پشتیبانی از front matter
- رندر رنگی و واکنش‌گرا
- جستجوی ساده
- صفحات اولیه راهنمای فارسی
- تست‌ها و CI
- مستندات پروژه
