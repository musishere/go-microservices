DB_URL=mysql://root:secret@tcp(127.0.0.1:3306)/ecom_db?tls=skip-verify

mysql:
	docker run --name ecom-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=ecom_db -d mysql:8.4

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" down

generate:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./ecom-grpc/pb/api.proto

.PHONY: mysql migrateup migratedown