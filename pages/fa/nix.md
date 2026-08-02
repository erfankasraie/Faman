---
title: nix
aliases:
- nix-env
- nix-shell
category: package
difficulty: advanced
keywords:
- nix
- declarative
- package
- reproducible
---

# Introduction

`nix` مدیر بسته و زبان پیکربندی Nix است: نصب ایزوله، نسخه‌های موازی و محیط‌های قابل‌تکرار (`nix-shell`).

# Syntax

```
nix [COMMAND]     # CLI جدید
nix-env [OPTIONS]
nix-shell [OPTIONS]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `nix profile install nixpkgs#PKG` | نصب در پروفایل کاربر |
| `nix-env -iA nixpkgs.PKG` | نصب (کلاسیک) |
| `nix search nixpkgs KEY` | جستجو |
| `nix-shell -p PKG` | شل موقت با بسته |
| `nix-collect-garbage` | پاکسازی store |

# Examples

```bash
nix search nixpkgs jq
nix-shell -p jq git
nix profile install nixpkgs#hello
nix-collect-garbage -d
```

# Common mistakes

- قاطی کردن کانال/flake بدون `experimental-features`.
- پر شدن دیسک store بدون garbage collection.

# Tips

- برای پروژه: `shell.nix` یا flake.
- روی NixOS با `configuration.nix` سیستم اعلانی است.

# Related commands

- `brew` — ساده‌تر برای macOS/Linux روزمره
- `guix` — رویکرد مشابه
