migrateup:
	migrate -path repository/db_migration -database "postgresql://chris:password@localhost:5432/mockdb?sslmode=disable" -verbose up

migratedown:
	migrate -path repository/db_migration -database "postgresql://chris:password@localhost:5432/mockdb?sslmode=disable" -verbose down

migratefix:
	migrate -path repository/db_migration -database "postgresql://chris:password@localhost:5432/mockdb?sslmode=disable" force 1

generate-proto:
	protoc --go_out=handler/campaign/grpc --go_opt=paths=source_relative \
    --go-grpc_out=handler/campaign/grpc --go-grpc_opt=paths=source_relative \
    campaign.proto

generate-proto-2:
	protoc --go_out=.  --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./handler/campaign/grpc/campaign.proto
