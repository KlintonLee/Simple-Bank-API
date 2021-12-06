mocks:
	docker exec -it simple-bank-api mockgen -destination=internal/modules/customers/mocks/customer.go -package=mock_customer -source=internal/modules/customers/customer.go Customer
	docker exec -it simple-bank-api mockgen -destination=internal/modules/customers/mocks/customer_service.go -package=mock_customer -source=internal/modules/customers/customer_service.go Customer

clean:
	rm -rf internal/modules/customers/mocks

test:
	docker exec -it simple-bank-api go test ./...

PHONY : mocks clean test