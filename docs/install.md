# نصب faman

## سریع

```bash
# فقط faman
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash

# faman + فونت + UTF-8 + کمک RTL
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash -s -- --with-rtl

# فقط فونت / RTL
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/setup-rtl.sh | bash
```

---

## تنظیمات پیشرفته اسکریپت

```bash
bash scripts/install.sh --help
```

### فلگ‌ها

| فلگ | معنی |
|------|------|
| `--with-rtl` | faman + فونت + locale + کمک RTL |
| `--rtl-only` | فقط فونت/locale/RTL |
| `--skip-deps` | نصب git/go سیستمی را رد کن |
| `--user` | نصب در `~/.local` بدون sudo برای باینری |
| `--prefix=DIR` | ریشه نصب (پیش‌فرض `/usr/local`) |
| `--branch=NAME` | شاخه یا تگ git |
| `--repo=URL` | آدرس مخزن |
| `--from-dir=PATH` | ساخت از کلون موجود |
| `--no-mlterm` | mlterm نصب نشود |
| `--no-shell-rc` | `.bashrc` / `.zshrc` دست نخورد |
| `--locale=en\|fa` | locale ترجیحی UTF-8 |
| `--plain-default` | `FAMAN_PLAIN=1` پیش‌فرض در shell |
| `--uninstall` | حذف باینری و pages از PREFIX |
| `--dry-run` | فقط چاپ کارها |
| `--verbose` | لاگ بیشتر |

### مثال‌ها

```bash
# نصب کاربر بدون root برای باینری
bash scripts/install.sh --user --with-rtl

# خروجی ساده فارسی به‌صورت پیش‌فرض
bash scripts/install.sh --with-rtl --plain-default --locale=fa

# از کلون فعلی
cd Faman && bash scripts/install.sh --from-dir=. --skip-deps --with-rtl

# پیش‌نمایش
bash scripts/install.sh --with-rtl --dry-run --verbose

# حذف
bash scripts/install.sh --uninstall
bash scripts/install.sh --user --uninstall
```

### فایل کانفیگ دائمی

```bash
mkdir -p ~/.config/faman
cp ~/.config/faman/install.env.example ~/.config/faman/install.env   # بعد از یک‌بار اجرای --with-rtl
# یا دستی:
cat > ~/.config/faman/install.env <<'EOF'
PREFIX=$HOME/.local
USER_INSTALL=1
WITH_RTL=1
PLAIN_DEFAULT=1
LOCALE_PREF=fa
NO_MLTERM=0
NO_SHELL_RC=0
BRANCH=main
EOF

bash scripts/install.sh   # مقادیر از فایل خوانده می‌شوند؛ CLI اولویت دارد
```

مسیر کانفیگ با `FAMAN_INSTALL_CONFIG` قابل تغییر است.

### متغیرهای محیطی

| متغیر | نقش |
|--------|------|
| `FAMAN_REPO_URL` | URL کلون |
| `FAMAN_BRANCH` | شاخه |
| `PREFIX` | مسیر نصب |
| `FAMAN_INSTALL_CONFIG` | مسیر فایل env |

---

## بعد از نصب (یک‌بار دستی)

### GNOME Terminal

Preferences → Custom font → `Vazirmatn` / `DejaVu Sans Mono`

### VS Code

```json
"terminal.integrated.fontFamily": "Vazirmatn, DejaVu Sans Mono, monospace"
```

### اگر حروف خراب است

```bash
export LANG=en_US.UTF-8
FAMAN_PLAIN=1 faman ls
```

بیشتر: [terminal-persian.md](terminal-persian.md)

## حذف دستی

```bash
sudo rm -f /usr/local/bin/faman
sudo rm -rf /usr/local/share/faman
# یا برای --user:
rm -f ~/.local/bin/faman
rm -rf ~/.local/share/faman
```
