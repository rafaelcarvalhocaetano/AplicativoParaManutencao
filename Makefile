
server:
	go run ./cmd/server.go

migrate-up:
	bash ./scripts/migrate.sh migrate_up

migrate-down:
	bash ./scripts/migrate.sh migrate_down -y

migrate-create:
	bash ./scripts/migrate.sh migrate_create

graphl-generate:
	go run github.com/99designs/gqlgen generate