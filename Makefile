.PHONY: run migrate_up migrate_down

migrate_up:
	migrate -path scripts/migration -database "mysql://root:109339Lam@@tcp(127.0.0.1:3306)/socialnetapp" up

migrate_down:
	migrate -path scripts/migration -database "mysql://root:109339Lam@@tcp(127.0.0.1:3306)/socialnetapp" down

migrate_force_db:
	migrate -path scripts/migration -database "mysql://root:109339Lam@@tcp(127.0.0.1:3306)/socialnetapp" force 1

gen_cal:
	protoc --go_out=pkg/proto/authen_and_post --go_opt=paths=source_relative --go-grpc_out=pkg/proto/authen_and_post --go-grpc_opt=paths=source_relative pkg\proto\authen_and_post.proto

run_server:
	go run .\cmd\authen_and_post_svc\main.go