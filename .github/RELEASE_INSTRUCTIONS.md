# Release (immutable tags)

GitHub **immutable releases** mean:

- After a tag is used for a published release, **you cannot reuse that tag_name**.
- You cannot delete-and-recreate the same version to fix missing assets.
- Fix = **bump version** (v0.2.0 → **v0.2.1**).

## Correct procedure for v0.2.1

### 1) Clean junk (optional)

On GitHub → Releases, delete any **untagged** / broken drafts if listed.

### 2) Create and push the git tag (from your machine)

```bash
git checkout main && git pull
git tag -a v0.2.1 -m "faman v0.2.1"
git push origin v0.2.1
```

(Repo rules may block Actions from creating tags; local push is reliable.)

### 3) Run Release workflow

- Actions → Release → Run workflow
- Tag: **`v0.2.1`**
- prerelease: **false**

Or the tag push alone triggers the workflow (`on: push tags: v*`).

### 4) Verify

https://github.com/erfankasraie/Faman/releases/tag/v0.2.1

Assets must include tar.gz / zip / SHA256SUMS.

```bash
sha256sum -c SHA256SUMS --ignore-missing
```

## Do not

- Re-run workflow with tag `v0.2.0` after it was already published as immutable
- Expect `gh release delete` + recreate same tag to work
