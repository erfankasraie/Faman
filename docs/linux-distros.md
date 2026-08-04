# faman روی توزیع‌های لینوکس

یک باینری استاتیک Go برای **linux-amd64** و **linux-arm64** روی تقریباً همهٔ توزیع‌های مدرن کار می‌کند. تفاوت اصلی در **نصب وابستگی build** و **مسیر بستهٔ بومی** است.

## روش مشترک (همهٔ توزیع‌ها)

### ۱) یک‌خطی (پیشنهادی)

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash
export PATH="$HOME/.local/bin:$PATH"
faman version
```

با فونت و UTF-8:

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash -s -- --rtl
```

### ۲) آرشیو Release

| CPU | فایل |
|-----|------|
| x86_64 | `faman-*-linux-amd64.tar.gz` |
| aarch64 / ARM64 | `faman-*-linux-arm64.tar.gz` |

```bash
tar -xzf faman-*-linux-amd64.tar.gz
cd faman-*-linux-amd64
export FAMAN_PAGES="$PWD/pages/fa"
./faman ls
```

### ۳) از سورس

```bash
git clone https://github.com/erfankasraie/Faman.git && cd Faman
go build -o faman ./cmd/faman
export FAMAN_PAGES="$PWD/pages/fa"
```

---

## Debian / Ubuntu / Mint / Pop!_OS

```bash
# اسکریپت
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash -s -- --rtl

# یا بسته deb (اگر در Release یا dist/ باشد)
sudo dpkg -i faman_*_amd64.deb
# صفحات: /usr/local/share/faman/pages/fa
```

وابستگی build (اختیاری):

```bash
sudo apt update
sudo apt install -y git golang-go make curl
# اگر go قدیمی بود: sudo snap install go --classic
```

---

## Fedora / RHEL / Rocky / AlmaLinux

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash

# وابستگی build
sudo dnf install -y git golang make curl
```

بستهٔ `.rpm` رسمی هنوز در Release نیست؛ از tar.gz یا `get.sh` استفاده کنید.  
ساخت rpm بعداً با nFPM/goreleaser در رودمپ است.

---

## Arch Linux / Manjaro / EndeavourOS

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash

# یا از سورس
sudo pacman -S --needed git go make
git clone https://github.com/erfankasraie/Faman.git && cd Faman
go build -o faman ./cmd/faman
```

بستهٔ **AUR** هنوز منتشر نشده؛ می‌توانید موقتاً با `get.sh` پیش بروید.

---

## openSUSE

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash
sudo zypper install -y git go make curl   # برای build
```

---

## Alpine

باینری Release با `CGO_ENABLED=0` معمولاً روی Alpine glibc/musl بسته به build متفاوت است. **پیشنهاد روی Alpine: build از سورس روی خود Alpine** یا استفاده از `get.sh` اگر لینک‌پذیر بود.

```bash
apk add --no-cache go git make
git clone https://github.com/erfankasraie/Faman.git && cd Faman
CGO_ENABLED=0 go build -o faman ./cmd/faman
export FAMAN_PAGES="$PWD/pages/fa"
./faman ls
```

داخل Docker:

```dockerfile
FROM golang:1.22-alpine AS build
WORKDIR /src
RUN apk add --no-cache git
RUN git clone --depth 1 https://github.com/erfankasraie/Faman.git .
RUN CGO_ENABLED=0 go build -o /faman ./cmd/faman

FROM alpine:3.20
COPY --from=build /faman /usr/local/bin/faman
COPY --from=build /src/pages /usr/local/share/faman/pages
ENV FAMAN_PAGES=/usr/local/share/faman/pages/fa
ENTRYPOINT ["faman"]
```

---

## NixOS / Guix

بستهٔ رسمی در nixpkgs نیست. گزینه‌ها:

- `get.sh` در محیط کاربر (اگر مجاز باشد)
- `nix-shell -p go git` سپس `go build`
- یا آرشیو `linux-amd64` + `FAMAN_PAGES`

---

## WSL (ویندوز + لینوکس)

مثل همان توزیع داخل WSL (معمولاً Ubuntu):

```bash
curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash -s -- --rtl
```

فونت را در **Windows Terminal** تنظیم کنید، نه فقط داخل WSL.  
راهنما: [windows.md](windows.md) + [terminal-persian.md](terminal-persian.md)

---

## مسیر صفحات بعد از نصب

| روش | مسیر معمول |
|------|------------|
| `get.sh` / `--user` | `~/.local/share/faman/pages/fa` |
| `deb` / prefix سیستمی | `/usr/local/share/faman/pages/fa` |
| آرشیو دستی | کنار باینری `pages/fa` یا `FAMAN_PAGES` |

```bash
faman update --pages   # تازه‌سازی محتوا
```

---

## جدول سریع

| توزیع | بهترین روش الان |
|--------|------------------|
| Ubuntu / Debian | `get.sh` یا `.deb` |
| Fedora / RHEL | `get.sh` یا tar.gz |
| Arch | `get.sh` یا سورس |
| openSUSE | `get.sh` یا tar.gz |
| Alpine | build از سورس |
| WSL | `get.sh` مثل Ubuntu |
| Raspberry Pi (ARM64) | `linux-arm64.tar.gz` یا `get.sh` |

بسته‌بندی برای maintainerها: [packaging.md](packaging.md)  
مک: [macos.md](macos.md)  
ویندوز: [windows.md](windows.md)
