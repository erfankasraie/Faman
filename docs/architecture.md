# Architecture

faman follows a simple Clean Architecture style without unnecessary layers.

```
cmd/faman          → entrypoint (main)
internal/app       → CLI commands (cobra)
internal/parser    → load & parse Markdown pages
internal/renderer  → terminal rendering (lipgloss + glamour)
internal/search    → full-text-ish search over pages
internal/update    → future online updates (placeholder)
pages/fa           → content (Markdown)
```

## Design principles

- Small and fast
- Easy to contribute
- No over-engineering
- Content is king (pages live as plain Markdown)

## Page loading

The parser looks for `pages/fa` relative to the executable or current working directory. This keeps development and installed layouts simple.
