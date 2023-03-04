# run db migartion
db-init:
	go run migration/main.go init
	go run migration/main.go up

db-up:
	go run migration/main.go up

db-down:
	go run migration/main.go down

db-reset:
	go run migration/main.go reset

# run redis
run-redis:
	docker run --name product-api-redis -p 6379:6379 -d redis

# run http server
run-http-server-local:
	go build -o "./cmd/product-http/product-http" ./cmd/product-http && ./cmd/product-http/product-http

# run seeder
run-seed:
	go run seed/main.go