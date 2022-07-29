### golang Rest API
using gin framework

### go commands
### get dependencies of the program
go get -d ./...
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

#### get books
curl -X GET http://localhost:9090/books

#### get single book
curl -X GET http://localhost:9090/books/1

#### add books
curl -d '{"title": "Star Worlds","author": "David Blod"}' -X POST http://localhost:9090/books

#### patch book record
curl -d '{"title": "The Infinite Game"}' -X PATCH http://localhost:9090/books/1

#### delete a book record
curl -X DELETE http://localhost:9090/books/1

#### login and get Jwt
curl -d '{"username": "bod"}' -X POST http://localhost:9090/login

#### use the Jwt to access private Rest API
curl -H 'Accept: application/json' -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE1MTQ4MDgwMDAsInVzZXIiOiJkZWQifQ.xYsQfBdWNtUlfMSMatGDLstQRgnhz3DU3rwv1sVKXQg" -X GET http://localhost:9090/private/test/1