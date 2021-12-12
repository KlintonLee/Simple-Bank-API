createdb:
	touch db.sqlite
	docker exec simple-bank-api sqlite3 db.sqlite "CREATE TABLE customers (id string, name string, cpf string, birth string, created_at timestamp, updated_at timestamp)";

dropdb:
	rm db.sqlite

mocks:
	docker exec -it simple-bank-api mockgen -destination=internal/modules/customers/mocks/customer.go -package=mock_customer -source=internal/modules/customers/customer.go Customer
	docker exec -it simple-bank-api mockgen -destination=internal/modules/customers/mocks/customer_service.go -package=mock_customer -source=internal/modules/customers/customer_service.go Customer

clean:
	rm -rf internal/modules/customers/mocks

test:
	docker exec -it simple-bank-api go test ./...

http:
	docker exec -it simple-bank-api go run main.go http

initialize_cobra:
	docker exec -it simple-bank-api cobra init --pkg-name github.com/klintonlee/simple-bank-api

PHONY : createdb dropdb mocks clean test http initialize_cobra