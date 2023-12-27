## E-Commerse API

### Setup
```shell
make setup
# OR
go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go mod tidy
```

### Database
```shell
### Create database
CREATE DATABASE ecommerse
### Create table
source db/schema/Product.sql

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

### Start development
```shell
go run main.go web
# OR use Air - live reload
air web
```

### API Contract
#### 1. List of Products

##### Request
```http
GET /products
Content-Type: application/json
```

```http
GET /products?sortBy=title
Content-Type: application/json
```

```http
GET /products?sortBy=rating
Content-Type: application/json
```

##### Response
```http
GET /products
Content-Type: application/json

{
  "code": 200,
  "status": "OK",
  "errors": "",
  "data": [
    {
      "id": "product-id",
      "title": "product title",
      "description": "product-description",
      "rating": 4.0,
      "image": "/path/to/image",
      "created_at": "2023-12-20T14:10:45Z",
      "updated_at": "2023-12-20T14:10:45Z",
      "deleted_at": {
        "Time": "0001-01-01T00:00:00Z",
        "Valid": false
      }
    }
  ]
}
```

#### 2. Detail of Product

##### Request
```http
GET /products/:id
Content-Type: application/json
```

##### Response
```http
GET /products/:id
Content-Type: application/json

{
  "code": 200,
  "status": "OK",
  "errors": "",
  "data": {
    "id": "product-id",
    "title": "product title",
    "description": "product-description",
    "rating": 4.0,
    "image": "/path/to/image",
    "created_at": "2023-12-20T14:10:45Z",
    "updated_at": "2023-12-20T14:10:45Z",
    "deleted_at": {
      "Time": "0001-01-01T00:00:00Z",
      "Valid": false
    }
  }
}
```

### 3. Create a New Product

##### Request
```http
POST /products
Content-Type: application/json

{
  "id": "product-id",
  "title": "product title",
  "description": "product-description",
  "rating": 4.0,
  "image": "/path/to/image",
}
```
##### Response
```http
POST /products
Content-Type: application/json

{
  "code": 201,
  "status": "OK",
  "errors": "",
  "data": "Product created"
}
```

### 4. Update the product

##### Request
```http
PUT /products/:id
Content-Type: application/json

{
  "title": "product title",
  "description": "product-description",
  "rating": 4.0,
  "image": "/path/to/image",
}
```

##### Response
```http
PUT /products/:id
Content-Type: application/json

{
  "code": 200,
  "status": "OK",
  "errors": "",
  "data": "Product updated"
}
```

### 5. Delete the Product

##### Request
```http
DELETE /products/:id
Content-Type: application/json
```

##### Response
```http
DELETE /products/:id
Content-Type: application/json

{
  "code": 201,
  "status": "OK",
  "errors": "",
  "data": "Product deleted"
}
