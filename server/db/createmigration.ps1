psql -h localhost -U postgres -c "create database goresizing;"

cd ./server/db
$POSTGRESQL_URL='postgres://postgres:123456@localhost:5432/goresizing?sslmode=disable'
.\migrate.exe create -ext sql -dir migrations -seq create_users_table
.\migrate.exe -database $POSTGRESQL_URL -path migrations up

cd ./server/db
$POSTGRESQL_URL='postgres://postgres:123456@localhost:5432/goresizing?sslmode=disable'
.\migrate.exe -database $POSTGRESQL_URL -path migrations down