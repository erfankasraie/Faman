# faman روی ویندوز

پشتیبانی ویندوز روی **همان شاخه `main`** است (برنچ جدا لازم نیست).

## پیش‌نیاز

- [Windows Terminal](https://aka.ms/terminal) (پیشنهادی)
- فونت با گلیف فارسی: **Cascadia Mono**، **Noto Sans Mono**، یا **Vazirmatn**
- اختیاری: Go 1.22+ برای ساخت از سورس

## نصب از سورس

PowerShell:

```powershell
git clone https://github.com/erfankasraie/Faman.git
cd Faman
go build -o faman.exe ./cmd/faman

# صفحات کنار باینری یا در LocalAppData
$dest = Join-Path $env:LOCALAPPDATA "faman\pages"
New-Item -ItemType Directory -Force -Path $dest | Out-Null
Copy-Item -Recurse -Force .\pages\fa $dest\

# PATH کاربر
$dir = Join-Path $env:LOCALAPPDATA "faman\bin"
New-Item -ItemType Directory -Force -Path $dir | Out-Null
Copy-Item .\faman.exe $dir\
$env:Path = "$dir;$env:Path"
```

یا اسکریپت:

```powershell
Set-ExecutionPolicy -Scope CurrentUser RemoteSigned   # یک‌بار در صورت نیاز
irm https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.ps1 | iex
```

## مسیر صفحات

ترتیب جستجو:

1. `FAMAN_PAGES`
2. `pages\fa` کنار `faman.exe`
3. `%LOCALAPPDATA%\faman\pages\fa`
4. `%APPDATA%\faman\pages\fa`
5. پوشه جاری (توسعه)

```powershell
$env:FAMAN_PAGES = "C:\path\to\pages\fa"
faman.exe ls
```

## UTF-8 و فارسی

در **Windows Terminal** → Settings → Defaults → Appearance:

- Font face: `Cascadia Mono` یا فونت فارسی
- مطمئن شوید profile روی UTF-8 است

PowerShell 7:

```powershell
[Console]::OutputEncoding = [System.Text.UTF8Encoding]::new()
$OutputEncoding = [Console]::OutputEncoding
```

اگر حروف خراب بود:

```powershell
$env:FAMAN_PLAIN = "1"
faman.exe ls
```

راهنمای عمومی فارسی: [terminal-persian.md](terminal-persian.md)

## CMD کلاسیک

```bat
chcp 65001
set FAMAN_PLAIN=1
faman.exe ls
```

ترجیح با Windows Terminal است؛ `cmd.exe` قدیمی برای RTL/اتصال حروف ضعیف‌تر است.

## CI

هر push روی `main` آرتیفکت `faman-windows-amd64` می‌سازد (Actions → Artifacts).

## محدودیت‌ها

- محتوای صفحات دربارهٔ **دستورات لینوکس** است؛ روی ویندوز برای یادگیری/WSL مفید است
- RTL کامل به ترمینال میزبان بستگی دارد (مثل لینوکس)
