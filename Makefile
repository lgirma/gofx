SHELL := /bin/bash

dump_schema:
	sqlite3 ./data/s3.db .schema > ./data/s3.schema.sql
	sqlite3 ./data/s3_egziabher_ab_bmb.db .schema > ./data/s3_egziabher_ab_bmb.schema.sql

migrate_header_init:
	goose sqlite3 ./data/s3.db -dir ./data/migrations create 0001_header_init sql

migrate_header:
	goose sqlite3 ./data/s3.db -dir ./data/migrations up

migrate_tenant_init:
	goose sqlite3 ./data/s3_tenant.db -dir ./data/migrations_tenant create 0001_tenant_init sql

migrate_tenant:
	goose sqlite3 ./data/s3_tenant.db -dir ./data/migrations_tenant up

sqlc_gen:
	(cd data && sqlc generate)
	(cd data && sqlc generate --file sqlc_tenant.yaml)

minor_release:
	@curver=$$(cat ./api/version); vermaj=$$(echo $$curver | cut -d. -f1); vermin=$$(echo $$curver | cut -d. -f2); newver="$${vermaj}.$$((vermin+1)).0-$$(date -u '+%Y%m%d')"; echo $$newver > ./api/version

major_release:
	@curver=$$(cat ./api/version); vermaj=$$(echo $$curver | cut -d. -f1); vermin=$$(echo $$curver | cut -d. -f2); newver="$$((vermaj+1)).0.0-$$(date -u '+%Y%m%d')"; echo $$newver > ./api/version

dev:
	go run -ldflags "-s -w -X main.Environment=production -X main.Version=1.0.0" ./api

test:
	go test ./...

release:
	GOOS=linux go build -ldflags "-s -w -X main.Environment=production -X main.Version=1.0.0" -o ./out/ ./api

release_windows:
	GOOS=windows go build -ldflags "-s -w -X main.Environment=production -X main.Version=1.0.0" -o ./out/ ./api