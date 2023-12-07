migrateup:
	migrate -path repository/db_migration -database "postgresql://chris:password@localhost:5432/mockdb?sslmode=disable" -verbose up 2

migratedown:
	migrate -path repository/db_migration -database "postgresql://chris:password@localhost:5432/mockdb?sslmode=disable" -verbose down

migratefix:
	migrate -path repository/db_migration -database "postgresql://chris:password@localhost:5432/mockdb?sslmode=disable" force 1

generate-proto:
	protoc --go_out=handler/campaign/grpc --go_opt=paths=source_relative \
    --go-grpc_out=handler/campaign/grpc --go-grpc_opt=paths=source_relative \
    campaign.proto

generate-proto-2:
	protoc --go_out=./handler/campaign --go-grpc_out=./handler/campaign handler/campaign/grpc/*.proto

check-grpc-service:
	grpcurl -plaintext localhost:8080 list

check-grpc-service-method:
	grpcurl -plaintext localhost:8080 list campaign.CampaignHandler

test-grpc-getbyid:
	grpcurl -plaintext -format text -d 'id: 17' \
	localhost:8080 campaign.CampaignHandler.GetCampaignByID