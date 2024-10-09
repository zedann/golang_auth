run-postgres:
	docker exec -it 6db15fbab4a3 psql -U postgres

start-postgres:
	docker start golang_auth-db-1

stop-postgres:
	docker stop golang_auth-db-1

db-up:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5434/golang_auth?sslmode=disable" --verbose up

db-down:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5434/golang_auth?sslmode=disable" --verbose down

migrate-table:
	migrate create -ext sql -dir db/migrations add_$(NAME)_table

build:
	go build -o ./bin/app cmd/main.go

run: build
	./bin/app
