# gox documentation website

Static documentation site for [gox](https://github.com/sahilkhaire/gox) — Node-friendly Go toolkit.

**Production:** [https://gox.varcore.dev](https://gox.varcore.dev)

Built with [VitePress](https://vitepress.dev/). API reference pages are generated from gox source via a Go doc scanner.

## Prerequisites

- Node.js 24+ (matches GitHub Actions runners)
- Go 1.25+ (for `scripts/gendocs`)
- Clone the [gox](https://github.com/sahilkhaire/gox) library beside this repo (or set `GOX_SRC`)


## Development

```bash
npm install
GOX_SRC=../gox npm run docs:dev
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

## Deploy (Cloudflare Pages + GitHub Actions)

Pushes to `main` / `master` build the site and deploy to Cloudflare Pages at **https://gox.varcore.dev**.

### One-time Cloudflare setup

1. **Create a Pages project** (optional — first deploy can create `gox-docs` automatically):
   - Cloudflare dashboard → Workers & Pages → Create → Pages → Connect to Git *or* leave empty for GitHub Actions–only deploys.
   - Project name: `gox-docs`

2. **API token** (My Profile → API Tokens → Create Token → **Custom token**):
   - **Permissions:** Account → **Cloudflare Pages** → **Edit**
   - **Account resources:** include the account that owns `gox-docs`
   - Do **not** use the Global API Key — use a scoped API token only
   - Save the token once; it cannot be viewed again

3. **Account ID** — Workers & Pages → Overview → **Account ID** (32-character hex, not the zone ID)

4. **GitHub secrets** on `sahilkhaire/gox-documentation`:
   | Secret | Value |
   |--------|-------|
   | `CLOUDFLARE_API_TOKEN` | API token from step 2 |
   | `CLOUDFLARE_ACCOUNT_ID` | Account ID from step 3 |

   If the workflow uses the `production` environment, add the same secrets under **Settings → Environments → production** as well (or remove environment protection so repo secrets apply).

   **401 Authentication error?** Regenerate the token with **Pages → Edit**, confirm the account ID matches the dashboard, and verify secret names are exact (no extra spaces).

5. **Custom domain** (Cloudflare Pages → `gox-docs` project → Custom domains):
   - Add `gox.varcore.dev`
   - If `varcore.dev` is on Cloudflare, confirm the CNAME `gox` → `gox-docs.pages.dev` (often auto-created).

### CI workflow

`.github/workflows/docs.yml`:

- Checks out this repo and `sahilkhaire/gox` into `./gox`
- Runs `gendocs`, coverage check, and VitePress build
- On push to `main`/`master`, deploys `docs/.vitepress/dist` via `cloudflare/wrangler-action@v3` to project **`gox-docs`**

Pull requests run build + coverage only (no deploy).

### Manual deploy (optional)

```bash
GOX_SRC=../gox npm run docs:build
npx wrangler pages deploy docs/.vitepress/dist --project-name=gox-docs
```

Requires `wrangler login` or `CLOUDFLARE_API_TOKEN` in the environment.

## Layout

```
gox-website/
├── docs/                 # VitePress content
│   ├── public/           # Static assets (hero.svg, logo.svg)
│   ├── guide/            # Getting started, architecture
│   ├── reference/        # Cheat sheet
│   ├── migration/        # Express, axios, Knex, …
│   └── packages/         # Generated API reference
├── scripts/
│   ├── gendocs/          # Go API scanner
│   └── sync-guides.mjs   # Copy guides from gox repo
└── wrangler.toml         # Cloudflare Pages project config
```

The `gox` library is read-only at build time (`GOX_SRC`, default `../gox` locally, `gox` in CI).
