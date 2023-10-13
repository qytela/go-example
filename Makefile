download:
	@go mod download

migration_up:
	@migrate -database "mysql://root:@tcp(localhost:3306)/go-example" -path migrations up

migration_down:
	@migrate -database "mysql://root:@tcp(localhost:3306)/go-example" -path migrations down
