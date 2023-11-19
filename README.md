# migrate-go 💨

This is a tool to migrate database.
Built upon [migrate](https://github.com/golang-migrate/migrate). 

Core feature is that it can read database dsn from file not only raw string. 🔥

## Usage 🧰

### Flags 🚩

- `-d` : database dsn raw string
- `-f` : file secret path to read database dsn (overwrite `-d`)
- `-dir` : migration directory absolute path

### Docker 🐳

1. Using DSN raw string
```bash
docker pull dekuyo/migrate-go:latest
docker run --rm -it -v $PWD:/app dekuyo/migrate-go -dir /app/migration -d "user:password@host:port/database"
```

2. Using DSN file
```bash
docker pull dekuyo/migrate-go:latest
docker run --rm -it -v $PWD:/app -v /path/to/dsn/file:/path/to/dsn/file dekuyo/migrate-go -dir /app/migration -f /path/to/dsn/file
```


### Docker Compose 🐳

1. Using DSN raw string
```yaml
db-migrate:
  image: dekuyo/migrate-go:latest
  volumes:
    - ./migrations:/migrations
  command: [ "-dir", "/migrations", "-dsn", "postgres://postgres:postgres@db:5432/order-shop?sslmode=disable" ]  
```

2. Using DSN file
```yaml
db-migrate:
  image: dekuyo/migrate-go:latest
  volumes:
    - ./migrations:/migrations
    - ./path/to/dsn/file:/path/to/dsn/file
  command: [ "-dir", "/migrations", "-file", "/path/to/dsn/file" ]
``
