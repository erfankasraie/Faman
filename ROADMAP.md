# Roadmap — faman

نسخهٔ در حال آماده‌سازی: **`0.2.2`** (ریلیز برنامه‌ریزی‌شده: فردا)

---

## وضعیت

| محور | مقدار |
|------|--------|
| صفحات | ~۱۷۶ (پس از batch 0.2.2) |
| CLI | show · search(--cat) · list · categories · random · doctor · update(--verify) · completion |
| بسته‌ها | tar.gz · zip · **deb** · SHA256SUMS |
| zsh | alias + recomplete + docs |

---

## چک‌لیست ریلیز v0.2.2 (فردا)

- [x] version = 0.2.2 در کد
- [x] CHANGELOG بخش 0.2.2
- [x] deb در workflow
- [x] search --cat + completion
- [x] صفحات کمبود دسته فشرده‌سازی / دیسک / شبکه
- [ ] `bash scripts/check-pages.sh` سبز
- [ ] CI سبز روی main
- [ ] `git tag -a v0.2.2 -m "faman v0.2.2" && git push origin v0.2.2`
- [ ] تأیید Assets شامل `.deb` و `SHA256SUMS`

---

## بعد از 0.2.2

1. پوشش بیشتر pages (~۲۰۰)
2. deb arm64 اختیاری
3. self-update باینری اختیاری با verify
4. v0.3.0
