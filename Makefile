migrateup:
	migrate -path repository/migration -database "postgresql://chris:password@localhost:5432/mockdb?sslmode=disable" -verbose up

migratedown:
	migrate -path repository/migration -database "postgresql://chris:password@localhost:5432/mockdb?sslmode=disable" down