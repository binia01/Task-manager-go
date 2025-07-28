# 🧪 Task Management API - Unit Testing 

## 🔧 Prerequisites

- Go 1.20+
- `testify` library: `go get github.com/stretchr/testify`

 **Set up MongoDB:**
      - This project uses MongoDB as the database. Ensure you have a MongoDB instance running.
      - Store your MongoDB Atlas connection string in a `.env` file at the project root. Example `.env` content:
        ```
        MONGODB_URI=mongodb+srv://<username>:<password>@cluster0.tj8um.mongodb.net/?retryWrites=true&w=majority
        ```
      - The application will read the connection string from the `.env` file using an environment variable loader (e.g., `github.com/joho/godotenv`).
      - In `Delivery/main.go`, load the environment variable and use it in your connection code:
        ```go
        import (
          "os"
          "github.com/joho/godotenv"
          // other imports...
        )

        func setupTaskTestDB() {
          godotenv.Load()
          mongoURI := os.Getenv("MONGODB_URI")
          clientOptions := options.Client().ApplyURI(mongoURI)
          // ...
        }
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


