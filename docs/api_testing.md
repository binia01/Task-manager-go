# 🧪 Task Management API - Unit Testing 

## 🔧 Prerequisites

- Go 1.20+
- `testify` library: `go get github.com/stretchr/testify`

**Set up MongoDB:**
    - This project uses MongoDB as the database. Ensure you have a MongoDB instance running.
    - You must provide your MongoDB Atlas username and password in the connect_db function inside Reposiories/task_repository_test.go and Reposiories/user_repository_test.go.
    - Example connection string (replace <username> and <password>):
    
      ```go
      clientOptions := options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.tj8um.mongodb.net/?retryWrites=true&w=majority")
      ```

## ▶️ Running Tests

Run all tests:

```bash
go test ./... -v
```

Check test coverage:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## ✅ Coverage

- Usecase Layer: Fully tested (`TaskUsecase`, `UserUsecase`)
- Infrastructure: PasswordService, JWTService, AuthMiddleware tested
- Middleware token/role scenarios included
- Reposiory Layer: Fully tested
(`TaskRepository`, `UserRepository`)

## 🔄 CI Integration

GitHub Actions is configured to automatically run tests on each push to `main`. See `.github/workflows/go.yml`.


