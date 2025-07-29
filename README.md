- URL swagger docs: /swagger/index.html



```bash
migrate -path ./migrations -database "postgres://postgres:123456@localhost:5432/mydb?sslmode=disable" up

migrate create -ext sql -dir ./migrations create_db

migrate -path ./migrations -database "postgres://postgres:123456@localhost:5432/mydb?sslmode=disable" force 20250727135835 #force
```


- TODO
    - Create full database: 27/07
    - Implement API: 28 - 30/07