---
title: cargo
aliases:
category: package
difficulty: beginner
keywords:
- rust
- package
- crates
- build
---

# Introduction

`cargo` ابزار رسمی build و مدیریت بسته **Rust** است (crates.io).

# Syntax

```
cargo <command> [OPTIONS]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `new NAME` | پروژه جدید |
| `build` | کامپایل |
| `run` | اجرا |
| `test` | تست |
| `add DEP` | افزودن وابستگی |
| `install BIN` | نصب باینری سراسری |
| `search KEY` | جستجوی crate |

# Examples

```bash
cargo new hello && cd hello
cargo build
cargo run
cargo add serde --features derive
cargo test
cargo install ripgrep
```

# Common mistakes

- فراموش کردن `Cargo.lock` در باینری‌ها (برای reproducibility).
- `cargo install` بدون به‌روزرسانی toolchain.

# Tips

- `rustup` برای مدیریت نسخهٔ Rust.
- `cargo clippy` و `cargo fmt` برای کیفیت کد.

# Related commands

- `rustc`
- `rustup`
- `go` — ماژول‌های گو
