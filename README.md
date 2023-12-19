## E-Commerse API

```shell
### Create Migration
migrate create -ext sql -dir db/migrations migration_state
# OR
make migrate-create mg=migration_state

### Migrate Up
make migrate-up
# OR 
go run main.go migrate

### Migrate Down
make migrate-down
```