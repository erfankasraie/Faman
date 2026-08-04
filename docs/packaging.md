# بسته‌بندی و انتشار (`.exe` · `.deb` · آرشیو)

## خلاصه فرمت‌ها

| فرمت | پلتفرم | چطور ساخته می‌شود |
|------|--------|-------------------|
| `faman.exe` داخل `.zip` | ویندوز x64 | `GOOS=windows GOARCH=amd64 go build` |
| `.tar.gz` | لینوکس / macOS | همان build + `tar` + پوشه `pages` |
| `.deb` | دبیان/اوبونتو amd64 | `scripts/package.sh` + `dpkg-deb` |
| GitHub Release | همه | workflow `release.yml` روی تگ `v*` |

---

## یک‌خطی محلی

```bash
# همه آرشیوها + .deb (اگر dpkg-deb باشد)
./scripts/package.sh

VERSION=0.1.4-pre ./scripts/package.sh
./scripts/package.sh --no-deb      # فقط tar/zip
./scripts/package.sh --deb-only    # فقط deb
```

خروجی در `dist/`:

- `faman-VERSION-linux-amd64.tar.gz`
- `faman-VERSION-linux-arm64.tar.gz`
- `faman-VERSION-windows-amd64.zip`  ← شامل **`faman.exe`** و `pages/`
- `faman-VERSION-darwin-*.tar.gz`
- `faman_VERSION_amd64.deb`
- `SHA256SUMS`

### نصب deb

```bash
sudo dpkg -i dist/faman_*_amd64.deb
faman version
# صفحات: /usr/local/share/faman/pages/fa
```

### ویندوز

```powershell
Expand-Archive faman-*-windows-amd64.zip
cd faman-*-windows-amd64
$env:FAMAN_PAGES = "$PWD\pages\fa"
.\faman.exe ls
```

---

## فقط `.exe` (بدون اسکریپت)

```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build \
  -ldflags "-s -w -X github.com/faman-project/faman/internal/app.version=0.1.4-pre" \
  -o faman.exe ./cmd/faman
```

صفحات را جداگانه کنار exe بگذارید یا `FAMAN_PAGES` را ست کنید.

---

## انتشار روی GitHub

1. Actions → **Release** → Run workflow با تگ مثلاً `v0.1.4-pre`
2. یا: `git tag -a v0.1.4-pre -m "..." && git push origin v0.1.4-pre`

workflow فعلی می‌سازد: **tar.gz + zip ویندوز** (هنوز deb را آپلود نمی‌کند).  
برای افزودن deb به Release، یا خروجی `package.sh` را دستی به Release ضمیمه کنید، یا workflow را گسترش دهید.

---

## بعداً (رودمپ)

| بسته | ابزار پیشنهادی |
|------|----------------|
| deb/rpm بهتر | [goreleaser](https://goreleaser.com) + nFPM |
| AUR | `PKGBUILD` جدا |
| Homebrew | formula جدا |
| MSI ویندوز | اختیاری؛ zip+exe کافی است |

---

## پیش‌نیاز .deb روی ماشین build

```bash
sudo apt install dpkg-dev   # dpkg-deb
```
