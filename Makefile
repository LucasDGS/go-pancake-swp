generate-swagger:
	swag init --parseDependency --parseInternal	

migrate-up:
	migrate -database postgres://postgres:postgres@localhost:5432/evolution_fitness?sslmode=disable -path ./db/migrations up

migrate-down:
	migrate -database postgres://postgres:postgres@localhost:5432/evolution_fitness?sslmode=disable -path ./db/migrations down
