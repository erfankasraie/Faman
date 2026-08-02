# faman installer for Windows (PowerShell)
#   irm https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.ps1 | iex
#   or:  powershell -File scripts/install.ps1
$ErrorActionPreference = "Stop"

$Repo = if ($env:FAMAN_REPO_URL) { $env:FAMAN_REPO_URL } else { "https://github.com/erfankasraie/Faman.git" }
$Prefix = Join-Path $env:LOCALAPPDATA "faman"
$BinDir = Join-Path $Prefix "bin"
$PagesDir = Join-Path $Prefix "pages"

Write-Host "==> Installing faman to $Prefix" -ForegroundColor Cyan

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
  Write-Host "Go not found. Install from https://go.dev/dl/ then re-run." -ForegroundColor Red
  exit 1
}
if (-not (Get-Command git -ErrorAction SilentlyContinue)) {
  Write-Host "git not found. Install Git for Windows." -ForegroundColor Red
  exit 1
}

$tmp = Join-Path $env:TEMP ("faman-build-" + [guid]::NewGuid().ToString("n"))
New-Item -ItemType Directory -Path $tmp | Out-Null
try {
  Write-Host "==> Cloning..."
  git clone --depth 1 $Repo (Join-Path $tmp "src")
  Set-Location (Join-Path $tmp "src")

  Write-Host "==> Building..."
  go build -ldflags "-s -w" -o faman.exe ./cmd/faman

  New-Item -ItemType Directory -Force -Path $BinDir | Out-Null
  New-Item -ItemType Directory -Force -Path $PagesDir | Out-Null
  Copy-Item -Force .\faman.exe (Join-Path $BinDir "faman.exe")
  if (Test-Path (Join-Path $PagesDir "fa")) {
    Remove-Item -Recurse -Force (Join-Path $PagesDir "fa")
  }
  Copy-Item -Recurse -Force .\pages\fa (Join-Path $PagesDir "fa")

  Write-Host "==> Done: $(Join-Path $BinDir 'faman.exe')" -ForegroundColor Green
  Write-Host "Pages: $(Join-Path $PagesDir 'fa')"
  Write-Host ""
  Write-Host "Add to PATH for this session:"
  Write-Host "  `$env:Path = '$BinDir;' + `$env:Path"
  Write-Host "Permanent user PATH: System Properties → Environment Variables, or:"
  Write-Host "  [Environment]::SetEnvironmentVariable('Path', '$BinDir;' + [Environment]::GetEnvironmentVariable('Path','User'), 'User')"
  Write-Host ""
  Write-Host "Docs: https://github.com/erfankasraie/Faman/blob/main/docs/windows.md"
}
finally {
  Set-Location $env:USERPROFILE
  Remove-Item -Recurse -Force $tmp -ErrorAction SilentlyContinue
}
