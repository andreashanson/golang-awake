dsn := "host=localhost port=5432 dbname=agent user=postgres password=abc"

docker run -e POSTGRES_PASSWORD=abc -p 5432:5432 --rm postgres

dbmate --url "postgres://postgres:abc@localhost:5432/agent?sslmode=disable" up

