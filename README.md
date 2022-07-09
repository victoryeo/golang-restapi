### golang Rest API
using gin framework

### go commands
#### to run the application
go run main.go 
#### to test the application
go test -v

### to call the Rest API using curl
#### get todo
curl -X GET http://localhost:9090/todos

#### add todo
curl -H "Content-Type: application/json" -d '{"id": "john", "title": "changeme", "completed": true}' -X POST http://localhost:9090/todos

#### get specific todo
curl -X GET http://localhost:9090/todos/3

#### patch todo status
curl -X PATCH http://localhost:9090/todos/3