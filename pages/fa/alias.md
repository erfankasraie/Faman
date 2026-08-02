---
title: alias
aliases:
category: shell
difficulty: beginner
keywords:
- shortcut
- shell
- config
---

# Introduction

`alias` میانبر برای دستورات طولانی در شل می‌سازد. builtin است.

# Syntax

```
alias
alias name='command'
unalias name
```

# Options

بدون گزینه لیست aliasها را نشان می‌دهد.

# Examples

```bash
alias ll='ls -lah'
alias gs='git status'
alias grep='grep --color=auto'
unalias ll
```

# Common mistakes

- انتظار ماندگاری بدون نوشتن در `~/.bashrc` یا `~/.zshrc`.

# Tips

- در اسکریپت‌ها alias معمولاً غیرفعال است؛ از تابع استفاده کنید.

# Related commands

- `type`
- `which`
- `function` در bash
