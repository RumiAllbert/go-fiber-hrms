# Human Resource Management System (HRMS) API written in Go

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

## Usage

1. Clone the repository
2. Run `go mod tidy` to install dependencies
3. Run `go run main.go` to start the server
4. Open `http://localhost:3000` in your browser

### NOTE

- There is currently no front-end for this project. You can use Postman or Paw to test the API. If you'd like to contribute, please feel free to do so.

## Endpoints and API

1. GET /employee
2. GET /employee/:id
3. POST /employee
4. PUT /employee/:id
5. DELETE /employee/:id

## Sample Request

```json
{
    "id": "2345lkjsdfgj3",
    "name": "John Doe",
    "email": "johndoe@gmail.com",
    "age": 45,
    "salary": 12353
}
```

## Requirements

1. Go 1.16+
2. MongoDB
