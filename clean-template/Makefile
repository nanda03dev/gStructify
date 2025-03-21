build:
	go build -o ms-name .
run: build
	./ms-name
dev:
	go run main.go
dev-docker:
	docker-compose up --build --remove-orphans
run-sql-db:
	docker-compose -f docker-compose.sql.db.yml up -d