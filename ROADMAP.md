# Roadmap — faman

آخرین به‌روزرسانی: **۲۰۲۶-۰۸-۰۴**  
نسخهٔ کد: **`0.2.0-pre`**

---

## وضعیت فعلی

| محور | مقدار |
|------|--------|
| صفحات | ~۱۶۳ |
| CLI | show · search · **list · categories · random · doctor** · update · version |
| کیفیت | `check-pages.sh` در CI |
| عمق man | فاز ۱ کامل |
| نصب | get.sh / install / package / docs پلتفرم |

---

## فازها

### فاز ۰ — pre-release · تقریباً تمام
تگ GitHub Releases هنوز با maintainer است.

### فاز ۱ — عمق man · **[x]**
### فاز ۲ — check-pages + CI · **[x]** (aliases یکنواخت اختیاری مانده)

### فاز ۳ — CLI محصولی · **بخش عمده [x]**

| # | کار | وضعیت |
|---|-----|--------|
| 3.1 | `faman list` (`--cat` `--diff` `--names`) | [x] |
| 3.2 | `faman categories` | [x] |
| 3.3 | `faman random` (`--open`) | [x] |
| 3.4 | `faman doctor` | [x] |
| 3.5 | search `--cat` | [ ] |
| 3.6 | completion نام صفحات | [ ] |
| 3.7 | ≥۱۵۰ صفحه | [x] (~۱۶۳) |
| 3.8 | تگ `v0.2.0` / pre | [ ] |

### فاز ۴+ — صفحات کمبود · deb در Release · TUI · v1.0

---

## بعدی

1. completion برای نام صفحات  
2. `search --cat`  
3. تگ Release `v0.2.0-pre`  
4. deb در workflow  
"}, {