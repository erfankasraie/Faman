# Cutting a pre-release

## Automated (recommended)

1. Open **Actions → Release → Run workflow**
2. Tag: `v0.1.2-pre`
3. prerelease: **true**
4. Run — builds Linux / Windows / macOS archives with `pages/fa` and publishes a GitHub Pre-release.

Or from git:

```bash
git checkout main && git pull
git tag -a v0.1.2-pre -m "v0.1.2-pre"
git push origin v0.1.2-pre
```

Pushing the tag triggers the same workflow.

## Assets produced

- `faman-*-linux-amd64.tar.gz`
- `faman-*-linux-arm64.tar.gz`
- `faman-*-windows-amd64.zip`
- `faman-*-darwin-amd64.tar.gz`
- `faman-*-darwin-arm64.tar.gz`
- `SHA256SUMS`

Do **not** use a long-lived branch named like a version tag; tags only.
