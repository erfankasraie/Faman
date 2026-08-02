# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.3] — 2026-08-02

### Added

- ده‌ها صفحه جدید: `rsync`, `tmux`, `screen`, `crontab`, `ufw`, `jq`, `make`, `diff`, `patch`, `watch`, `time`, `basename`, `dirname`, `realpath`, `sha256sum`, `base64`, `alias`, `export`, `source`, `umask`, `chgrp`, `reboot`, `shutdown`, `who`, `last`, `getent`, `hostnamectl`, `timedatectl`, `nc`, `locate`, `yes`, `seq`, ...
- راهنمای نمایش فارسی در ترمینال: `docs/terminal-persian.md`
- حالت‌های `FAMAN_PLAIN` و `FAMAN_WRAP` برای کنترل خروجی

### Changed

- رندرر: soft-wrap امن برای فارسی (بدون شکستن وسط کلمه)، هشدار locale غیر UTF-8، حذف عرض ثابت باکس عنوان
- یکدست‌سازی صفحات قدیمی (`cat`, `grep`, `chmod`, `git`, `head`, `tail`, `less`, `ssh`, `scp`, `find`, `docker`, `curl`) با قالب استاندارد بخش‌ها

## [0.1.2] — 2026-08-02

### Added

- گسترش پوشش به حدود یک‌سوم دستورات رایج (~۸۰+ صفحه)
- فهرست موضوعی `docs/pages-index.md`

## [0.1.1] — 2026-08-02

### Added

- `cmd/faman/main.go`، صفحات اولیه بیشتر، `docs/wiki/`

### Fixed

- `internal/update/update.go` — `strings.Repeat`

## [0.1.0] — 2026-08-01

### Added

- انتشار اولیه CLI، پارسر، رندر، جستجو، CI، صفحات پایه
