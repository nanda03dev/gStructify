build:
	go build -o ms-name .
run: build
	./ms-name
dev:
	go run main.go
run-sql-db:
	docker-compose -f docker-compose.sql.db.yml up -d