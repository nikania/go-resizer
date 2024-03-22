psql -h localhost -U postgres -c "create database goresizing;"

$POSTGRESQL_URL='postgres://postgres:123456@localhost:5432/goresizing?sslmode=disable'
.\migrate.exe create -ext sql -dir db/migrations -seq create_users_table
.\migrate.exe -database $POSTGRESQL_URL -path db/migrations up
