migrateup:
	migrate --path db/migration --database "postgresql://name:pass@localhost:5432/language-learning?sslmode=disable" --verbose up

migratedown:
	migrate --path db/migration --database "postgresql://name:pass@localhost:5432/language-learning?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

testsum:
	gotestsum --watch --format testname -- -v ./... 


.PHONY: migrateup migratedown sqlc test