migrate-up:
	migrate -path migrations/ -database "postgres://postgres:1234@localhost:15432/mirlo_db?sslmode=disable" up $(n)

migrate-down:
	migrate -path migrations/ -database "postgres://postgres:1234@localhost:15432/mirlo_db?sslmode=disable" down $(n)

