# Cutting release v0.2.0 (stable, not prerelease)

## بعد از merge شدن نسخه 0.2.0 روی main

### Actions (پیشنهادی)

1. [Actions → Release](https://github.com/erfankasraie/Faman/actions/workflows/release.yml)
2. **Run workflow**
3. Tag: **`v0.2.0`**
4. prerelease: **false** (خاموش)
5. Run

workflow اگر تگ شامل `pre|rc|alpha|beta` نباشد، `prerelease=false` می‌گذارد.

### Git

```bash
git checkout main && git pull
git tag -a v0.2.0 -m "faman v0.2.0"
git push origin v0.2.0
```

### تأیید

- https://github.com/erfankasraie/Faman/releases/tag/v0.2.0
- باید **Latest** باشد نه Pre-release
- آرتیفکت: linux/darwin tar.gz + windows zip
