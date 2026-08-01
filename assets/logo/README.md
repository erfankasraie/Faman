# Faman Logo

A geometric mark inspired by the stylized Persian wild goat (Capra
aegagrus) on the Marlik gold cup and Burnt City (Shahr-e Sukhteh)
pottery — reduced to its most essential visual idea: two large,
tapered, upward-curling horns, rendered as a single flat black
silhouette.

The horn curves are true logarithmic spirals (the same mathematical
curve real ram and ibex horns grow in), not freehand arcs — which is
why they read as "horn" rather than generic swoosh even at tiny
sizes. Symmetry is exact (verified pixel-for-pixel). No face, no
eyes, no texture, no gradient, no badge — just the silhouette,
in the spirit of GitHub / Rust / Docker / Bun-style marks.

## Files

| File | Purpose |
|---|---|
| `faman-logo-full.svg` | Full logo, tight (non-square) crop. Use for READMEs, docs headers, letterhead-style placement. |
| `faman-icon-24.svg` | Small icon, square canvas, padded and centered. Use for favicons, GitHub org/repo avatars, app icons, toolbars. |
| `faman-icon-inverse.svg` | Same icon, white fill, for dark backgrounds. |
| `favicon.ico` | Multi-size (16/32/48) favicon, white background. |
| `faman-icon-*.png` | Transparent-background PNG raster at 16/32/64/128/256/512/1024px. |
| `app-icon-dark.png` / `app-icon-light.png` | 512×512 rounded-square app-icon mockups (context preview, not required deliverables). |
| `terminal-glyph.txt` | Unicode/ASCII terminal representations. |

## Usage

- **Black on white** (default): use `faman-logo-full.svg` or
  `faman-icon-24.svg` as-is.
- **Dark backgrounds**: use `faman-icon-inverse.svg`, or apply
  `fill="currentColor"` in your own CSS-controlled context.
- **Minimum size**: legible down to 16×16; below that, prefer the
  terminal glyph instead of the vector mark.
- **Do not**: add gradients, drop shadows, outlines, or place inside
  a circular badge — the silhouette is the whole identity.

## Three official reductions

1. **Full logo** — `faman-logo-full.svg`
2. **Small icon (24×24)** — `faman-icon-24.svg`
3. **Terminal glyph** — see `terminal-glyph.txt`

### Terminal glyph (preview)

```
    ╭─╮ ╭─╮
   ╱   ╲╱   ╲
       ▾
```

Compact:

```
   ╭╮  ╭╮
    ╲▾╱
```

ASCII-only:

```
   /‾\ /‾\
  /   X   \
      V
```

Single character: `∧` or `Λ`
