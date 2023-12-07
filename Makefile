migrateup:
	migrate -path repository/db_migration -database "postgresql://chris:password@localhost:5432/mockdb?sslmode=disable" -verbose up

migratedown:
	migrate -path repository/db_migration -database "postgresql://chris:password@localhost:5432/mockdb?sslmode=disable" -verbose down

migratefix:
	migrate -path repository/db_migration -database "postgresql://chris:password@localhost:5432/mockdb?sslmode=disable" force 1
