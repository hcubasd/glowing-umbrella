# glowing-umbrella

Glowing Umbrella is the CRM deals read API for the dashboard project of **[mlclogistica.app](https://mlclogistica.app)**. It assembles a fully-nested deals graph from the `sales.*` schema and serves it at `GET /deals`.

## Structure

```
src/
  schemas.ts     — Zod schemas for the full nested Deal graph (also the response validation/serialization source of truth)
  db.ts          — pg Pool, plus type-parser fixes for `numeric` (→ JS number) and `date` (→ raw string, avoids TZ drift)
  repository.ts  — all sales.* queries and in-memory graph assembly (getDeals())
  app.ts         — Fastify instance: CORS, Swagger, the /deals route (buildApp(), no listen — used directly by tests via .inject())
  main.ts        — entrypoint, binds buildApp() to 0.0.0.0:8080
  app.test.ts    — end-to-end test against a live migrated Postgres, via Fastify's .inject()
```

Compose: `compose.yaml` runs Postgres + `curly-spoon` migrations (image: `ghcr.io/hcubasd/curly-spoon:1.0.0` in dev). Ensure migrations have run before starting the server.

## Build & run (local)

```bash
npm ci
npm run dev      # tsx watch src/main.ts
```

Set Postgres env vars: `PGHOST`, `PGPORT`, `PGUSER`, `PGPASSWORD`, `PGDATABASE` (`pg` reads these automatically).

Production build:

```bash
npm run build    # tsc -> dist/
npm start        # node dist/main.js
```

## Testing

Tests are colocated with source (`src/*.test.ts`, matching bookish-lamp's convention) and run via Vitest using Fastify's `.inject()` — no server process or open port required, just a live migrated database. Use `scripts/integrate.sh` to install, test, and build together, or `docker compose up --wait` to bring up Postgres + migrations first.

```bash
npm test
```

## OpenAPI / Swagger

`@fastify/swagger` generates the OpenAPI spec directly from the same Zod schemas used for response validation — not from docstrings, so it can't drift from what the routes actually enforce. `@fastify/swagger-ui` serves it live at `/documentation` on the running server. This sits behind OAuth2 Proxy like the rest of the API (no public GitHub Pages docs — the source is public anyway, and the live spec never goes stale the way a separately-published static one could).

## Notes

- CORS in `app.ts` is restricted to `https://dashboard.mlclogistica.app`, `GET`/`OPTIONS` only, credentialed.
- The server binds `0.0.0.0:8080` explicitly — Fastify has no implicit default port or host.
- Response schemas are strict: any row shape Zod doesn't expect (an unexpected `null`, a `status` outside `won`/`lost`/`ongoing`) fails serialization rather than silently passing through, unlike the previous Go implementation.
