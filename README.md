### golang Rest API
using gin framework

# get todo
curl -X GET http://localhost:9090/todos

# add todo
curl -H "Content-Type: application/json" -d '{"id": "john", "title": "changeme", "completed": true}' -X POST http://localhost:9090/todos

# get specific todo
curl -X GET http://localhost:9090/todos/3

# patch todo status
curl -X PATCH http://localhost:9090/todos/3