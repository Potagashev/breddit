Чтобы запустить, введите команду 
`docker-compose up`

Чтобы применить миграции, введите команду
(Креды смотреть в .env, а порты в docker-compose.yml)
`migrate -database postgresql://postgres:postgrespw@localhost:49153/postgres?sslmode=disable -path ./db/migrations up`

Чтобы создавать миграции, пишем
`migrate create -ext sql -dir db/migrations -seq create_threads_table`
