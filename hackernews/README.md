https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

```shell
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xzv

mv migrate go/bin/

migrate -database mysql://root:password@/hackernews -path internal/pkg/db/migrations/mysql up

go get -u github.com/go-chi/chi/v5
```

## jwt

https://pkg.go.dev/github.com/golang-jwt/jwt#section-readme
