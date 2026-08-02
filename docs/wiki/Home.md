# 🐐 faman

**Persian Manual Pages** — صفحات راهنمای فارسی لینوکس

faman یک پروژه متن‌باز است که معادل فارسی و مدرن صفحات man لینوکس را می‌سازد.  
ما ترجمه صرف نمی‌کنیم؛ **یاد می‌دهیم** — از زبان ساده برای مبتدی تا جزئیات حرفه‌ای.

مخزن: [erfankasraie/Faman](https://github.com/erfankasraie/Faman)

---

## ویژگی‌های اصلی

| دستور | کاربرد |
|-------|--------|
| `faman ls` | راهنمای فارسی دستور |
| `faman search docker` | جستجو در تمام صفحات |
| `faman version` | نسخه |
| `faman help` | راهنمای خود ابزار |
| `faman update` | به‌روزرسانی (placeholder تا v0.4) |

---

## نصب سریع

```bash
git clone https://github.com/erfankasraie/Faman.git
cd Faman
make build
sudo install -m 755 faman /usr/local/bin/faman
sudo mkdir -p /usr/local/share/faman
sudo cp -r pages /usr/local/share/faman/
```

جزئیات بیشتر در [README](https://github.com/erfankasraie/Faman#installation).

---

## صفحات ویکی

- [Roadmap](Roadmap) — نقشه راه نسخه‌ها
- [Architecture](https://github.com/erfankasraie/Faman/blob/main/docs/architecture.md) — معماری

---

ساخته شده با ♥ برای جامعه لینوکس فارسی
