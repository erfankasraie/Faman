# faman روی macOS

پشتیبانی مک روی همان شاخهٔ `main` است (باینری `darwin-amd64` و `darwin-arm64`).

## کدام آرشیو؟

| سخت‌افزار | آرشیو Release |
|-----------|----------------|
| Apple Silicon (M1/M2/M3/…) | `faman-*-darwin-arm64.tar.gz` |
| Intel Mac | `faman-*-darwin-amd64.tar.gz` |

تشخیص معماری:

```bash
uname -m
# arm64 → Apple Silicon
# x86_64 → Intel
```

## نصب سریع (پیشنهادی)

نیاز: Homebrew یا Xcode CLT برای `git`/`go` اگر از سورس بسازید.

### با اسکریپت (مثل لینوکس)

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash
export PATH="$HOME/.local/bin:$PATH"
faman version
faman ls
```

`get.sh` روی مک هم `--user` را پیش‌فرض می‌گذارد (`~/.local`).

> فلگ `--rtl` بیشتر برای فونت/locale لینوکس است؛ روی مک فونت ترمینال را دستی تنظیم کنید.

### از GitHub Releases

```bash
# مثال Apple Silicon
curl -sL -o faman.tgz \
  "https://github.com/erfankasraie/Faman/releases/latest/download/faman-0.1.4-pre-darwin-arm64.tar.gz"
# یا فایل را از صفحه Releases دانلود کنید (نام نسخه را چک کنید)

tar -xzf faman-*-darwin-arm64.tar.gz
cd faman-*-darwin-arm64
export FAMAN_PAGES="$PWD/pages/fa"
./faman version
./faman ls

# اختیاری: کپی دائمی
mkdir -p "$HOME/.local/bin" "$HOME/.local/share/faman"
cp faman "$HOME/.local/bin/"
cp -a pages "$HOME/.local/share/faman/"
export PATH="$HOME/.local/bin:$PATH"
# pages از مسیر share پیدا می‌شود یا:
# export FAMAN_PAGES="$HOME/.local/share/faman/pages/fa"
```

### از سورس

```bash
brew install go git   # در صورت نیاز
git clone https://github.com/erfankasraie/Faman.git
cd Faman
go build -o faman ./cmd/faman
export FAMAN_PAGES="$PWD/pages/fa"
./faman ls
```

## ترمینال و فارسی

- **Terminal.app** یا **iTerm2** یا **Warp**
- فونت با گلیف فارسی: مثلاً *SF Mono* + fallback، یا *Vazirmatn* / *Noto Sans Mono*
iTerm2 → Profiles → Text → Font

اگر حروف خراب بود:

```bash
export LANG=en_US.UTF-8
FAMAN_PLAIN=1 faman ls
```

جزئیات کلی: [terminal-persian.md](terminal-persian.md)

## به‌روزرسانی

```bash
faman update --check
faman update --pages
# باینری: دوباره get.sh یا دانلود آرشیو darwin مناسب
```

## محدودیت

- محتوای صفحات دربارهٔ **دستورات لینوکس** است؛ روی مک برای یادگیری و کار با Linux VM/Docker/SSH مفید است، نه جایگزین `man` اختصاصی macOS.
- بستهٔ Homebrew formula هنوز رسمی نیست (در رودمپ).

## بسته‌بندی برای توسعه‌دهنده

```bash
./scripts/package.sh --no-deb
# → dist/faman-*-darwin-amd64.tar.gz
# → dist/faman-*-darwin-arm64.tar.gz
```

راهنمای همهٔ فرمت‌ها: [packaging.md](packaging.md)
