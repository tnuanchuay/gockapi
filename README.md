# gockapi

Mock API made by Golang.

### Get Started
Create workspace like this
MyMock
   |_ gockapi.exe
   |_api
      |_api-get.get
      |_api-post.post

api-get.get
api-post.get
```json
{
  api:"api"
}
```

run gockapi
browse to localhost:8080/api/get
use Postman request to localhost:8080/api/get with method post

### Install
```
go install github.com/tspn/gockapi
```
