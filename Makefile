DB_URL := "postgres://cource_user:cource_password@localhost:5432/cource_app?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_URL)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down
	
run:
	go run cmd/app/main.go