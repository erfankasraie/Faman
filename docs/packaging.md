# بسته‌بندی و انتشار

راهنمای ساخت آرتیفکت برای **لینوکس (چند معماری/توزیع)**، **macOS** و **ویندوز**.

## خلاصه فرمت‌ها

| فرمت | پلتفرم | ابزار |
|------|--------|--------|
| `*-linux-amd64.tar.gz` | لینوکس x86_64 (Ubuntu, Fedora, Arch, …) | `package.sh` / Release |
| `*-linux-arm64.tar.gz` | لینوکس ARM64 (Pi, ARM server) | همان |
| `*-darwin-arm64.tar.gz` | macOS Apple Silicon | همان |
| `*-darwin-amd64.tar.gz` | macOS Intel | همان |
| `*-windows-amd64.zip` | ویندوز (`faman.exe`) | همان |
| `*_amd64.deb` | Debian/Ubuntu | `package.sh` + `dpkg-deb` |
| `.rpm` | Fedora/RHEL/openSUSE | هنوز نه (رودمپ: nFPM) |
| AUR / Homebrew | Arch / macOS | رودمپ |

یک باینری لینوکس برای **همهٔ distroهای glibc رایج** کافی است؛ نیازی به بستهٔ جدا per-distro نیست مگر بخواهید در مخزن رسمی آن‌ها باشید.

---

## یک‌خطی build

```bash
./scripts/package.sh
VERSION=0.1.4-pre ./scripts/package.sh
./scripts/package.sh --no-deb
./scripts/package.sh --deb-only
```

خروجی در `dist/`:

```text
faman-VERSION-linux-amd64.tar.gz
faman-VERSION-linux-arm64.tar.gz
faman-VERSION-darwin-amd64.tar.gz
faman-VERSION-darwin-arm64.tar.gz
faman-VERSION-windows-amd64.zip    # faman.exe + pages/
faman_VERSION_amd64.deb            # اگر dpkg-deb باشد
SHA256SUMS
```

هر آرشیو شامل **باینری + `pages/fa`** است.

---

## لینوکس — استفاده از آرتیفکت

```bash
tar -xzf faman-*-linux-amd64.tar.gz
cd faman-*-linux-amd64
export FAMAN_PAGES="$PWD/pages/fa"
./faman version
```

### deb (Debian/Ubuntu)

```bash
sudo apt install dpkg-dev
./scripts/package.sh --deb-only
sudo dpkg -i dist/faman_*_amd64.deb
```

### بقیهٔ توزیع‌ها

جدول نصب کاربر: [linux-distros.md](linux-distros.md)

---

## macOS — استفاده از آرتیفکت

```bash
# Apple Silicon
tar -xzf faman-*-darwin-arm64.tar.gz
cd faman-*-darwin-arm64
export FAMAN_PAGES="$PWD/pages/fa"
./faman ls

# Intel
tar -xzf faman-*-darwin-amd64.tar.gz
```

راهنمای کامل مک: [macos.md](macos.md)

Cross-compile از لینوکس:

```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o faman-darwin-arm64 ./cmd/faman
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o faman-darwin-amd64 ./cmd/faman
```

> امضای Apple (notarize) برای توزیع بیرون Gatekeeper فعلاً انجام نمی‌شود؛ برای استفادهٔ شخصی/توسعه‌دهنده معمولاً کافی است: System Settings → Privacy → Allow.

---

## ویندوز — `.exe`

داخل zip:

```powershell
Expand-Archive faman-*-windows-amd64.zip
cd faman-*-windows-amd64
$env:FAMAN_PAGES = "$PWD\pages\fa"
.\faman.exe ls
```

یا فقط باینری:

```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o faman.exe ./cmd/faman
```

[windows.md](windows.md)

---

## انتشار GitHub Release

```text
Actions → Release → Run workflow → tag v0.1.4-pre
```

یا تگ git. workflow می‌سازد: tar.gz لینوکس/داروین + zip ویندوز.  
`.deb` را با `package.sh` بسازید و در صورت نیاز دستی به Release ضمیمه کنید.

---

## نقشهٔ آینده

| هدف | روش پیشنهادی |
|------|----------------|
| deb/rpm یکدست | goreleaser + nFPM |
| AUR | `PKGBUILD` + `faman-bin` |
| Homebrew | tap جدا یا formula |
| Alpine | apk از سورس / build روی musl |
| امضای مک | Apple Developer + notarize |

---

## پیش‌نیاز build deb

```bash
sudo apt install dpkg-dev
```
