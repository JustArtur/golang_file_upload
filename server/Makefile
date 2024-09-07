include .env.dev

DB_URL ?= postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)
MIGRATIONS_PATH = db/migrate

migrate_up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate_down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down

migrate_version:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version
