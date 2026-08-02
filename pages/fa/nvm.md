---
title: nvm
aliases:
category: environment
difficulty: beginner
keywords:
- node
- javascript
- version
- npm
---

# Introduction

`nvm` چند نسخهٔ **Node.js** را برای کاربر نصب و بین آن‌ها سوییچ می‌کند.

# Syntax

```
nvm <command> [version]
```

# Options / Commands

| دستور | توضیح |
|-------|--------|
| `install VER` | نصب |
| `use VER` | فعال در این شل |
| `alias default VER` | پیش‌فرض |
| `ls` / `ls-remote` | لیست محلی/دور |
| `uninstall VER` | حذف |

# Examples

## پروژهٔ Node با نسخهٔ ثابت تیم

```bash
mkdir web && cd web
nvm install 20
nvm use 20
node -v

echo "20" > .nvmrc
npm init -y
npm install express

cat > server.js <<'EOF'
const express = require("express");
const app = express();
app.get("/", (_req, res) => res.send("ok"));
app.listen(3000, () => console.log("http://127.0.0.1:3000"));
EOF

node server.js
```

## سوییچ بین دو نسخه

```bash
nvm install 18
nvm install 22
nvm use 18 && node -v
nvm use 22 && node -v
nvm alias default 20
```

## ورود به پوشه با `.nvmrc`

```bash
cd web
nvm use          # از .nvmrc می‌خواند
npm test
```

# Common mistakes

- Node apt + nvm هم‌زمان روی PATH.
- نبودن `nvm` در shell rc.

# Tips

- CI: نسخه را از `.nvmrc` بخوانید.
- جایگزین سبک: `fnm`.

# Related commands

- `npm` · `npx` · `asdf` · `pyenv`
