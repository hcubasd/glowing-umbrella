# glowing-umbrella

Glowing Umbrella is a small Go HTTP service that exposes the CRM deals read API used by the dashboard. It assembles a fully‑nested deals graph from the `sales.*` schema and serves it at `GET /deals`.

Quick repo notes

- Source: `src/` (server, DB init, repository, models).
- Compose: `compose.yaml` runs Postgres + `curly-spoon` migrations (image: `ghcr.io/hcubasd/curly-spoon:1.0.0-rc.4` in dev). Ensure migrations have run before starting the server.

Build & run (local)

1. Ensure Go is installed and on PATH (tested with Go 1.26).
2. Set Postgres env vars: `PGHOST, PGPORT, PGUSER, PGPASSWORD, PGDATABASE`.
3. Build and run the server (or use `scripts/integrate.sh` to build, start, wait for readiness and run tests):

```bash
cd src
go mod download
go build -o app ./...
./app
```

Testing

- Integration tests live in `tests/`. They expect a running server and migrated database. Use `scripts/integrate.sh` to run the server and tests together in CI/local.

OpenAPI / Swagger

- The project uses `swaggo/swag` annotations for endpoint docs. To generate OpenAPI docs locally:

```bash
# install swag (requires Go)
go install github.com/swaggo/swag/cmd/swag@latest
# generate docs (outputs into ./docs)
swag init -g src/main.go -o docs
```

Notes

- CORS in `main.go` is restricted to the dashboard host.
- For production, set `GIN_MODE=release` and configure trusted proxies as documented by gin.

If you want, I can commit this README and add a small CI step to generate or validate the docs.