run:
	docker-compose exec app go run cmd/srv/main.go

lint:
	docker-compose exec app golangci-lint run

migrate-up:
	docker-compose exec app migrate -database mysql://random:random@tcp\(db:3306\)/go-random?multiStatements=true -path db/migrations up

migrate-down:
	docker-compose exec app migrate -database mysql://random:random@tcp\(db:3306\)/go-random?multiStatements=true -path db/migrations down

test:
	docker-compose exec app go test ./...

mod:
	docker-compose exec app go mod tidy
	docker-compose exec app go mod vendor

bash:
	@docker-compose exec app bash

mysql:
	docker-compose exec db mysql -urandom -prandom go-random
