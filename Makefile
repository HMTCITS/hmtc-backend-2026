migrate:
	go run cmd/migrate/main.go

seed:
	go run cmd/migrate/main.go

build:
	docker build -t hmtc-backend .

run:
	docker container run \
	--name hmtc-backend-server \
	-p 8080:8080 \
	--env-file .env \
	hmtc-backend:latest