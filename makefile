migrate_dev_db:
	migrate -database "postgresql://mrmixr_dev_user:local_dev_pwd@localhost:5432/mrmixr_dev_db?sslmode=disable" -path migrations $d

create_migration:
	migrate create -ext sql -dir migrations -seq $d

