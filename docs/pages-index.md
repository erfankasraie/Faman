# فهرست صفحات faman

حدود **۱۳۴** صفحهٔ فارسی.

---

## محیط‌های مجازی و نسخه (Virtual environments)

| دستور | کاربرد | سطح |
|--------|--------|------|
| `venv` | محیط مجازی استاندارد Python | مبتدی |
| `virtualenv` | محیط مجازی کلاسیک Python | مبتدی |
| `pipenv` | Pipfile + env | متوسط |
| `poetry` | وابستگی + lock + env | متوسط |
| `conda` | Anaconda / Miniconda / mamba | متوسط |
| `pyenv` | چند نسخهٔ Python | متوسط |
| `nvm` | چند نسخهٔ Node.js | مبتدی |
| `asdf` | version manager چندزبانه | متوسط |
| `direnv` | بارگذاری خودکار `.envrc` | متوسط |
| `docker` | کانتینر (ایزولهٔ کامل) | — |
| `podman` | کانتینر rootless | متوسط |
| `nix` | بسته/محیط اعلانی | پیشرفته |

```bash
faman venv
faman poetry
faman conda
faman nvm
faman direnv
```

---

## مدیران بسته (Package managers)

| دستور | اکوسیستم | سطح |
|--------|-----------|------|
| `apt` | Debian / Ubuntu | مبتدی |
| `dpkg` | deb (سطح‌پایین) | متوسط |
| `snap` | Ubuntu Snap | مبتدی |
| `flatpak` | Flathub / دسکتاپ | مبتدی |
| `dnf` | Fedora / RHEL (yum) | مبتدی |
| `rpm` | RPM (سطح‌پایین) | متوسط |
| `zypper` | openSUSE | مبتدی |
| `pacman` | Arch | متوسط |
| `yay` | Arch AUR | متوسط |
| `apk` | Alpine / Docker | مبتدی |
| `brew` | macOS / Linuxbrew | مبتدی |
| `nix` | Nix / nix-shell | پیشرفته |
| `pip` | Python / PyPI | مبتدی |
| `npm` | Node.js | مبتدی |
| `cargo` | Rust / crates | مبتدی |

---

## فایل‌سیستم و مسیر

`ls` `cd` `pwd` `mkdir` `rmdir` `cp` `mv` `rm` `ln` `touch` `find` `du` `df` `file` `stat` `mount` `umount` `lsblk` `chmod` `chown` `chgrp` `basename` `dirname` `realpath` `locate` `umask`

## متن و فیلتر

`cat` `less` `head` `tail` `grep` `sed` `awk` `sort` `uniq` `cut` `tr` `wc` `tee` `echo` `diff` `patch` `base64` `sha256sum` `seq` `yes`

## شل و محیط

`alias` `which` `env` `export` `history` `clear` `xargs` `sudo` `source` `time` `watch`

## فرایند

`ps` `top` `kill` `pgrep` `pkill` `nohup` `jobs`

## کاربر و مجوز

`id` `whoami` `passwd` `useradd` `who` `last` `getent`

## شبکه

`curl` `wget` `ssh` `scp` `ping` `ip` `ss` `dig` `traceroute` `hostname` `nc`

## سیستم و سرویس

`systemctl` `journalctl` `uname` `uptime` `date` `free` `lsof` `reboot` `shutdown` `crontab` `ufw` `hostnamectl` `timedatectl`

## آرشیو و انتقال

`tar` `zip` `unzip` `gzip` `gunzip` `rsync`

## توسعه و ویرایش

`nano` `vim` `git` `docker` `make` `jq` `tmux` `screen`

---

برای صفحهٔ جدید: قالب `pages/fa/ls.md` و [CONTRIBUTING.md](../CONTRIBUTING.md).
