# نمایش صحیح فارسی در ترمینال

اگر حروف فارسی به‌هم‌ریخته، جدا از هم، مربع خالی (□)، یا «توهم‌رفته» دیده می‌شوند، معمولاً یکی از این‌هاست:

1. **locale غیر UTF-8**
2. **فونت ترمینال بدون گلیف عربی/فارسی**
3. **شکستن خط وسط کلمه** (اتصال حروف عربی را خراب می‌کند)
4. **ترمینال ضعیف در RTL**

---

## راه‌حل سریع در faman

```bash
# خروجی ساده بدون رنگ و wrap تهاجمی
FAMAN_PLAIN=1 faman ls

# بدون شکستن خط (هر خط همان‌طور که در فایل است)
FAMAN_WRAP=0 faman ls
```

اگر locale شما UTF-8 نباشد، faman یک هشدار روی stderr چاپ می‌کند.

---

## ۱) UTF-8 کردن locale

```bash
# موقت
export LANG=en_US.UTF-8
export LC_ALL=en_US.UTF-8

# یا فارسی
export LANG=fa_IR.UTF-8
export LC_ALL=fa_IR.UTF-8

# ماندگار: به ~/.bashrc یا ~/.zshrc اضافه کنید
```

بررسی:

```bash
locale
echo -e "سلام"
```

اگر `echo` هم خراب است، مشکل از faman نیست — از سیستم/ترمینال است.

---

## ۲) فونت مناسب

فونت‌هایی که فارسی را خوب پوشش می‌دهند:

| فونت | توضیح |
|------|--------|
| **Vazirmatn** | عالی برای UI و ترمینال |
| **Noto Sans Arabic** / Noto Naskh Arabic | بسته گوگل |
| **DejaVu Sans Mono** | monospace با پوشش خوب |
| **Fira Code** + fallback فارسی | بعضی ترمینال‌ها |

### تنظیم در ترمینال‌های رایج

- **GNOME Terminal / Tilix**: Preferences → Profile → Custom font
- **Konsole**: Profile → Appearance → Font
- **Windows Terminal**: `settings.json` → `font.face`
- **VS Code / Cursor terminal**: `terminal.integrated.fontFamily`

مثال VS Code:

```json
"terminal.integrated.fontFamily": "Vazirmatn, DejaVu Sans Mono, monospace"
```

---

## ۳) ترمینال‌های پیشنهادی

| ترمینال | یادداشت |
|---------|----------|
| Kitty | پشتیبانی خوب Unicode |
| Alacritty | سریع؛ فونت را درست انتخاب کنید |
| WezTerm | fallback فونت قوی |
| GNOME Terminal | پایدار روی دسکتاپ |
| Windows Terminal | روی ویندوز بهترین گزینه معمول |

از ترمینال‌های خیلی قدیمی یا `linux console` خام بدون فونت فارسی اجتناب کنید.

---

## ۴) کارهایی که faman انجام می‌دهد

- متن فارسی را **وسط کلمه نمی‌شکند** (soft-wrap فقط روی فاصله)
- جداول و بلوک‌های کد را دست‌نخورده چاپ می‌کند
- با `FAMAN_PLAIN=1` مسیر ساده UTF-8 بدون glamour
- هشدار locale غیر UTF-8

محدودیت: بیشتر ترمینال‌ها RTL کامل ندارند؛ مخلوط فارسی و انگلیسی ممکن است گاهی ترتیب بصری عجیبی داشته باشد — این محدودیت ترمینال است نه encoding.

---

## ۵) تست سریع

```bash
printf '%s\n' 'سلام دنیا — اتصال حروف: می‌شود'
faman echo
FAMAN_PLAIN=1 faman echo
```

اگر `printf` درست و `faman` خراب است، Issue باز کنید و خروجی `locale`، نام ترمینال و فونت را بفرستید.
