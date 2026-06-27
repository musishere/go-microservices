DB_URL=mysql://root:secret@tcp(127.0.0.1:3306)/ecom_db?tls=skip-verify

mysql:
	docker run --name ecom-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=ecom_db -d mysql:8.4

migrateup:
	migrate -path migrations -database "$(DB_URL)" up

migratedown:
	migrate -path migrations -database "$(DB_URL)" down

.PHONY: mysql migrateup migratedown


DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS products;