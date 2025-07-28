# Task Manager API Documentation

A simple RESTful API for managing tasks. This API uses MongoDB as the database to store task information. You can create, view, update, and delete tasks using standard HTTP methods. This API supports user authentication and authorization using JWT tokens. Only admin users can create, update, or delete tasks.

---

## Clean Architecture Update

This project now follows the principles of Clean Architecture, which separates concerns into distinct layers for better maintainability, testability, and scalability. The main layers are:

- **Domain:** Contains core business logic and entity definitions (e.g., `Task`, `User`, `Role`).
- **Usecases:** Implements application-specific business rules and orchestrates interactions between repositories and entities.
- **Repositories:** Handles data access and persistence, abstracting database operations for tasks and users.
- **Infrastructure:** Provides supporting services such as JWT authentication, password hashing, and middleware for authentication and authorization.
- **Delivery:** Manages HTTP routing and controllers, exposing the API endpoints and handling HTTP requests/responses.

**Key Benefits:**
- Clear separation of concerns.
- Easy to swap out infrastructure (e.g., database, authentication) without affecting business logic.
- Improved testability and code organization.

**Directory Structure:**
```
/Domain         # Entities and core business logic
/Usecases       # Application use cases
/Repositories   # Data access and persistence
/Infrastructure # Supporting services (auth, password, middleware)
/Delivery       # HTTP controllers, routers, and entry point
```

All API functionality remains the same, but the codebase is now organized according to Clean Architecture best practices.


## How to Run

1. **Clone the repository:**
    ```sh
    git clone https://github.com/yourusername/Task-manager-go.git
    cd Task-manager-go
    git checkout clean-architecture
    ```
    2. **Set up MongoDB:**
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

        func main() {
          godotenv.Load()
          mongoURI := os.Getenv("MONGODB_URI")
          clientOptions := options.Client().ApplyURI(mongoURI)
          // ...
        }
        ```
It is the same for jwt_secret
3. **Build and start the server:**
    ```sh
    cd Delivery
    go run main.go
    ```
    or:
    ```sh
    cd Delivery
    go run .
    ```

4. The API will be available at:
    ```
    http://localhost:8080
    ```

## Authentication & Authorization

- All task endpoints require authentication via JWT token.
- Only admin users can create, update, or delete tasks.
- Register and login endpoints are public.
- Use the `/register` endpoint to create a user. The first registered user is assigned the admin role; subsequent users are regular users.
- Use the `/login` endpoint to obtain a JWT token. Include this token in the `Authorization` header for all requests to protected endpoints:
  ```
  Authorization: Bearer <your-jwt-token>
  ```


## MongoDB Integration

- The API uses a MongoDB collection named `Tasks` in a database called `Task-Database`.
- Each task is stored as a document in the collection with the following fields:
  - `id` (string): Unique identifier for the task.
  - `title` (string): Title of the task.
  - `description` (string): Description of the task.
  - `due_date` (string, RFC3339 format): Due date. (Stored as `time.Time` in Go, serialized as string in JSON.)
  - `status` (string): Status of the task (e.g., Pending, Completed).

### Error Handling

- **Database Connection Errors:** If the API cannot connect to MongoDB, the server will log the error and terminate.
- **Task Not Found:** If a task with the specified ID does not exist, the API will return a `404 Not Found` response.
- **Duplicate Task ID:** When creating a task, if a task with the same ID already exists, the API will return a `409 Conflict` response.
- **Invalid Request Body:** If the request body is invalid (e.g., missing required fields), the API will return a `400 Bad Request` response.

## API Usage

This project includes a Postman collection and curl examples for testing and understanding the API.

### Getting Started
Open [Postman](https://web.postman.co/workspace/e85ab91d-850f-41f6-8c20-cb0459fbaf68/collection/42847133-eba53ce0-400d-49db-a4f0-6f87d274047a?action=share&source=copy-link&creator=42847133)

Or view the docs online: [API Documentation](https://documenter.getpostman.com/view/42847133/2sB34kCdiH)
1. Click **“Run in Postman”** (top-right of the page)
2. Postman will automatically open and load the collection for you
3. Start making requests using the examples provided in the docs

## Base URL

```
http://localhost:8080
```

---
## User Management

- The API uses a MongoDB collection named `Users` in the same database.
- Each user has:
  - `username` (string): Unique username.
  - `password` (string): Hashed password (not returned in responses).
  - `role` (string): Either `"Admin"` or `"user"`.

## Endpoints


#### 0. Register

- **POST** `/register`
- **Description:** Register a new user. The first user is assigned the admin role.
- **Body:**
  ```json
  {
    "username": "your_username",
    "password": "your_password"
  }
  ```
- **Response:**
  - `200 OK` with username and success message.
  - `409 Conflict` if username already exists.

#### 0. Login

- **POST** `/login`
- **Description:** Login and receive a JWT token.
- **Body:**
  ```json
  {
    "username": "your_username",
    "password": "your_password"
  }
  ```
- **Response:**
  - `200 OK` with access token.
  - `401 Unauthorized` if credentials are invalid.

#### 0. Promote User to Admin

- **PATCH** `/promote/{username}`
- **Description:** Promote a user to admin. Requires admin JWT.
- **Header:** `Authorization: Bearer <token>`
- **Response:**
  - `200 OK` on success.
  - `400 Bad Request` if user not found or error.

---

## Task Endpoints

> **All task endpoints require JWT authentication. Only admin users can create, update, or delete tasks.**

### 1. Get All Tasks

- **GET** `/tasks`
- **Header:** `Authorization: Bearer <token>`
- **Description:** Retrieve a list of all tasks.

#### Request

- No parameters required.

#### Response

- **Status:** 200 OK
- **Body:** Array of Task objects.

```json
[
  {
    "id": "1",
    "title": "Task 1",
    "description": "First task",
    "due_date": "2024-06-10T12:00:00Z",
    "status": "Pending"
  }
]
```

#### Curl Test

**Linux:**
```sh
curl -X GET http://localhost:8080/tasks \
  -H "Authorization: Bearer <your-jwt-token>"
```

**Windows CMD:**
```cmd
curl -X GET http://localhost:8080/tasks -H "Authorization: Bearer <your-jwt-token>"
```

---

### 2. Get Task By ID

- **GET** `/tasks/{id}`
- **Header:** `Authorization: Bearer <token>`
- **Description:** Retrieve a specific task by its ID.

#### Request

- **Path Parameter:** `id` (string) - Task ID

#### Response

- **Status:** 200 OK
- **Body:** Task object

```json
{
  "id": "1",
  "title": "Task 1",
  "description": "First task",
  "due_date": "2024-06-10T12:00:00Z",
  "status": "Pending"
}
```

- **Status:** 404 Not Found (if not found)
- **Body:**
```json
{ "message": "task not found!" }
```

#### Curl Test

**Linux:**
```sh
curl -X GET http://localhost:8080/tasks/1 \
  -H "Authorization: Bearer <your-jwt-token>"
```

**Windows CMD:**
```cmd
curl -X GET http://localhost:8080/tasks/1 -H "Authorization: Bearer <your-jwt-token>"
```

---

### 3. Create Task

- **POST** `/tasks`
- **Header:** `Authorization: Bearer <admin-token>`
- **Description:** Create a new task. only admin users.

#### Request

- **Body:** JSON Task object

```json
{
  "id": "4",
  "title": "Task 4",
  "description": "Fourth task",
  "due_date": "2024-06-13T12:00:00Z",
  "status": "Pending"
}
```

#### Response

- **Status:** 201 Created
- **Body:**
```json
{
  "message": "task created successfully!",
  "task": {
    "id": "4",
    "title": "Task 4",
    "description": "Fourth task",
    "due_date": "2024-06-13T12:00:00Z",
    "status": "Pending"
  }
}
```

- **Status:** 400 Bad Request (invalid body)
- **Body:**
```json
{ "message": "invalid request" }
```

- **Status:** 409 Conflict (duplicate ID)
- **Body:**
```json
{ "message": "Task with this ID already exists" }
```

#### Curl Test

**Linux:**
```sh
curl -X POST http://localhost:8080/tasks \
  -H "Authorization: Bearer <admin-token>" \
  -H "Content-Type: application/json" \
  -d '{"id":"4","title":"Task 4","description":"Fourth task","due_date":"2024-06-13T12:00:00Z","status":"Pending"}'
```

**Windows CMD:**
```cmd
curl -X POST http://localhost:8080/tasks -H "Authorization: Bearer <admin-token>" -H "Content-Type: application/json" -d "{\"id\":\"4\",\"title\":\"Task 4\",\"description\":\"Fourth task\",\"due_date\":\"2024-06-13T12:00:00Z\",\"status\":\"Pending\"}"
```

---

### 4. Update Task

- **PUT** `/tasks/{id}`
- **Header:** `Authorization: Bearer <admin-token>`
- **Description:** Update an existing task by ID. only admin users.

#### Request

- **Path Parameter:** `id` (string) - Task ID
- **Body:** JSON Task object (fields to update)

```json
{
  "title": "Updated Task 1",
  "description": "Updated description",
  "due_date": "2024-06-15T12:00:00Z",
  "status": "Completed"
}
```

#### Response

- **Status:** 200 OK
- **Body:**
```json
{
  "message": "task updated successfully!",
  "task": {
    "id": "1",
    "title": "Updated Task 1",
    "description": "Updated description",
    "due_date": "2024-06-15T12:00:00Z",
    "status": "Completed"
  }
}
```

- **Status:** 400 Bad Request (invalid body)
- **Body:**
```json
{ "message": "invalid request" }
```

- **Status:** 404 Not Found (if not found)
- **Body:**
```json
{ "message": "task not found!" }
```

#### Curl Test

**Linux:**
```sh
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Authorization: Bearer <admin-token>" \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Task 1","description":"Updated description","due_date":"2024-06-15T12:00:00Z","status":"Completed"}'
```

**Windows CMD:**
```cmd
curl -X PUT http://localhost:8080/tasks/1 -H "Authorization: Bearer <admin-token>" -H "Content-Type: application/json" -d "{\"title\":\"Updated Task 1\",\"description\":\"Updated description\",\"due_date\":\"2024-06-15T12:00:00Z\",\"status\":\"Completed\"}"
```

---

### 5. Delete Task

- **DELETE** `/tasks/{id}`
- **Header:** `Authorization: Bearer <admin-token>`
- **Description:** Delete a task by its ID. Only admin users.

#### Request

- **Path Parameter:** `id` (string) - Task ID

#### Response

- **Status:** 200 OK
- **Body:**
```json
{
  "message": "task deleted successfully!",
  "task": {
    "id": "1",
    "title": "Task 1",
    "description": "First task",
    "due_date": "2024-06-10T12:00:00Z",
    "status": "Pending"
  }
}
```

- **Status:** 404 Not Found (if not found)
- **Body:**
```json
{ "message": "task not found!" }
```

#### Curl Test

**Linux:**
```sh
curl -X DELETE http://localhost:8080/tasks/1 \
  -H "Authorization: Bearer <admin-token>"
```

**Windows CMD:**
```cmd
curl -X DELETE http://localhost:8080/tasks/1 -H "Authorization: Bearer <admin-token>"
```

---

## Task Object

| Field       | Type      | Description         |
|-------------|-----------|---------------------|
| id          | string    | Unique identifier   |
| title       | string    | Task title          |
| description | string    | Task description    |
| due_date    | string    | Due date (RFC3339)  |
| status      | string    | Task status         |

## User Object

| Field    | Type   | Description         |
|----------|--------|---------------------|
| username | string | Unique username     |
| password | string | Hashed, not returned|
| role     | string | "Admin" or "user"   |

---

## Error Handling

- **401 Unauthorized:** Missing or invalid JWT token.
- **403 Forbidden:** Insufficient permissions (non-admin for admin-only endpoints).
- **404 Not Found:** Resource not found.
- **409 Conflict:** Duplicate resource (e.g., username or task ID).
- **400 Bad Request:** Invalid request body.

---

## Notes

- All date fields (`due_date`) must be in RFC3339 format (e.g., `"2024-06-10T12:00:00Z"`).
- Passwords are hashed before storage.
- JWT tokens expire after 12 hours.
- Only admin users can create, update, or delete tasks, or promote other users.