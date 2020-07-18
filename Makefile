lint:
	docker-compose exec app golangci-lint run
migrate-up:
	docker-compose exec app migrate -database mysql://random:random@tcp\(db:3306\)/go-random?multiStatements=true -path db/migrations up
migrate-down:
	docker-compose exec app migrate -database mysql://random:random@tcp\(db:3306\)/go-random?multiStatements=true -path db/migrations down