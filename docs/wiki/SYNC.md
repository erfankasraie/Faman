# همگام‌سازی با GitHub Wiki

محتوای ویکی در این پوشه نگهداری می‌شود تا در مخزن اصلی هم versioned باشد.

## صفحات

| فایل | صفحه ویکی |
|------|-----------|
| `Home.md` | Home |
| `Roadmap.md` | Roadmap |

## انتشار روی Wiki رسمی

```bash
git clone https://github.com/erfankasraie/Faman.wiki.git
cp docs/wiki/Home.md Faman.wiki/Home.md
cp docs/wiki/Roadmap.md Faman.wiki/Roadmap.md
cd Faman.wiki
git add Home.md Roadmap.md
git commit -m "docs: sync Roadmap and Home from main repo"
git push origin master
```

یا از رابط وب GitHub → Wiki → New Page.

لینک ویکی: https://github.com/erfankasraie/Faman/wiki
