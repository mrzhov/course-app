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

gen:
	go tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go