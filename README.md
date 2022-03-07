## database
```
docker-compose up database -d
```

## migrations:
```
migrate -path migrations -database "postgres://login:password@localhost:5432/dbname?sslmode=disable" -verbose up
```

## start project
```
docker-compose up -d
```