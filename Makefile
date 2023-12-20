setup:
	go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go mod tidy

migrate-up:
	migrate -path db/migrations -database "mysql://habi:habi123@tcp(localhost:3306)/ecommerse" -verbose up

migrate-down:
	migrate -path db/migrations -database "mysql://habi:habi123@tcp(localhost:3306)/ecommerse" -verbose down

migrate-create:
	migrate create -ext sql -dir db/migrations $(mg)

build:
	go build -o ecommerse