migrate-up:
	migrate -path migrations/ -database "postgres://postgres:1234@localhost:15432/mirlo_db?sslmode=disable" up $(n)

migrate-down:
	migrate -path migrations/ -database "postgres://postgres:1234@localhost:15432/mirlo_db?sslmode=disable" down $(n)

unit-tests:
	go test ./... -tags=unit_test

db-tests:
	go test ./... -tags=db_test

test-all:
	go test ./... -tags="unit_test,db_test"