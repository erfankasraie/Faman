# Cutting a release

## v0.1.4-pre (current)

### Option A — Actions (recommended)

1. Open [Actions → Release](https://github.com/erfankasraie/Faman/actions/workflows/release.yml)
2. **Run workflow**
3. Tag: `v0.1.4-pre`
4. prerelease: **true**
5. Run

This builds Linux / Windows / macOS archives (binary + `pages/fa`) and publishes a GitHub Pre-release.

### Option B — Git tag

```bash
git checkout main && git pull
git tag -a v0.1.4-pre -m "faman v0.1.4-pre"
git push origin v0.1.4-pre
```

### After publish

- Verify: https://github.com/erfankasraie/Faman/releases
- Smoke: download linux amd64 tarball → `./faman version` → `./faman update --check`

### Version in code

`internal/app/root.go` → `version = "0.1.4-pre"` (override with `-ldflags` in CI).
