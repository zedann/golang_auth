run-postgres:
	docker exec -it 6db15fbab4a3 psql -U postgres
start-postgres:
	docker start golang_auth-db-1
stop-postgres:
	docker stop golang_auth-db-1
build:
	go build -o ./bin/app cmd/main.go
run: build
	./bin/app
