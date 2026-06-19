# gox documentation website

Static documentation site for [gox](https://github.com/sahilkhaire/gox) — Node-friendly Go toolkit.

Built with [VitePress](https://vitepress.dev/). API reference pages are generated from gox source via a Go doc scanner.

## Prerequisites

- Node.js 20+
- Go 1.25+ (for `scripts/gendocs`)

## Development

```bash
cd gox-website
npm install
npm run docs:dev
```

Open http://localhost:5173

## Build

```bash
GOX_SRC=../gox npm run docs:build
```

Output: `docs/.vitepress/dist/`

## Commands

| Command | Description |
|---------|-------------|
| `npm run docs:dev` | Generate API docs + start dev server |
| `npm run docs:build` | Generate API docs + production build |
| `npm run docs:gen` | Run `scripts/gendocs` only |
| `npm run docs:check` | Fail if any gox export lacks a docs page |
| `npm run sync-guides` | Copy guides from `../gox/docs/` |

## Cloudflare Pages

| Setting | Value |
|---------|-------|
| Build command | `npm ci && npm run docs:build` |
| Build output | `docs/.vitepress/dist` |
| Root directory | `gox-website` |
| Environment | `GOX_SRC=../gox` |

## Layout

```
gox-website/
├── docs/                 # VitePress content
│   ├── guide/            # Getting started, architecture
│   ├── reference/        # Cheat sheet
│   ├── migration/        # Express, axios, Knex, …
│   └── packages/         # Generated API reference (267 pages)
├── scripts/
│   ├── gendocs/          # Go API scanner
│   └── sync-guides.mjs   # Copy guides from gox repo
└── public/               # Static assets
```

The `gox` library lives in the sibling `../gox` directory and is read-only at build time.
