# چک‌لیست ریلیز v0.2.2 (فردا)

## قبل از تگ

```bash
git checkout main && git pull
bash scripts/check-pages.sh
go test ./...
go build -o /tmp/faman ./cmd/faman
/tmp/faman version   # باید 0.2.2 باشد
```

## تگ و انتشار

```bash
git tag -a v0.2.2 -m "faman v0.2.2"
git push origin v0.2.2
```

یا Actions → Release → tag `v0.2.2` · prerelease=false.

## بعد از سبز شدن workflow

- [ ] https://github.com/erfankasraie/Faman/releases/tag/v0.2.2
- [ ] فایل‌ها: tar.gz (linux/darwin) · windows zip · **deb** · SHA256SUMS
- [ ] `sha256sum -c SHA256SUMS --ignore-missing`
- [ ] `sudo dpkg -i faman_0.2.2_amd64.deb` روی یک VM تست

## یادآوری immutable

اگر چیزی جا ماند → **v0.2.3**؛ دوباره `v0.2.2` ساخته نمی‌شود.
