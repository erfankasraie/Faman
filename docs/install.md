# نصب آسان faman

## یک خط — فقط faman

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash
```

## یک خط — faman + فونت فارسی + locale + کمک RTL

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.sh | bash -s -- --with-rtl
```

این حالت علاوه بر faman:

- فونت‌های Noto / DejaVu (و Vazirmatn اگر در مخزن باشد) را نصب می‌کند
- `en_US.UTF-8` و در صورت امکان `fa_IR.UTF-8` را generate می‌کند
- در صورت امکان **mlterm** را نصب می‌کند (ترمینال قوی‌تر برای RTL)
- snippet مربوط به UTF-8 را به `~/.bashrc` / `~/.zshrc` اضافه می‌کند
- راهنمای فونت را در `~/.config/faman/terminal-font-hint.txt` می‌نویسد

## فقط RTL / فونت (اگر faman را قبلاً نصب کرده‌اید)

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/setup-rtl.sh | bash
```

## از داخل مخزن کلون‌شده

```bash
git clone https://github.com/erfankasraie/Faman.git
cd Faman
bash scripts/install.sh --with-rtl
```

## بعد از نصب — کارهایی که یک‌بار دستی می‌مانند

ترمینال نمی‌تواند فونت پروفایل GUI را همیشه عوض کند. یک‌بار:

### GNOME Terminal (Ubuntu پیش‌فرض)

1. منوی ☰ → **Preferences**
2. پروفایل → **Custom font**
3. `Vazirmatn` یا `Noto Sans Mono` یا `DejaVu Sans Mono`

### VS Code / Cursor

در `settings.json`:

```json
"terminal.integrated.fontFamily": "Vazirmatn, DejaVu Sans Mono, monospace"
```

### WezTerm

```lua
return {
  bidi_enabled = true,
  font = require("wezterm").font_with_fallback({ "Vazirmatn", "DejaVu Sans Mono" }),
}
```

## اگر حروف هنوز خراب است

```bash
export LANG=en_US.UTF-8
export LC_ALL=en_US.UTF-8
FAMAN_PLAIN=1 faman ls
```

جزئیات: [terminal-persian.md](terminal-persian.md)

## حذف

```bash
sudo rm -f /usr/local/bin/faman
sudo rm -rf /usr/local/share/faman
```
