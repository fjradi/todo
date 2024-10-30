include .env

migrate-up:
	migrate -path migration -database "$(DATABASE_URL)" -verbose up

migrate-down:
	migrate -path migration -database "$(DATABASE_URL)" -verbose down

sqlc:
	sqlc generate