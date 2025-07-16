# Task Manager API Documentation

A simple RESTful API for managing tasks. You can create, view, update, and delete tasks using standard HTTP methods. This API is suitable for learning, prototyping, or integrating basic task management features into your applications.

## How to Run

1. **Clone the repository:**
    ```sh
    git clone https://github.com/yourusername/Task-manager-go.git
    cd Task-manager-go
    ```

2. **Build and start the server:**
    ```sh
    go run main.go
    ```

3. The API will be available at:
    ```
    http://localhost:8080
    ```

## API Usage
This project includes a Postman collection for testing and understanding the API.

### Getting Started
1. Open [Postman](https://www.postman.com/)
2. Import the collection from `docs/postman_collection.json`
3. Use the provided endpoints with appropriate payloads

Or view the docs online: [API Documentation](https://documenter.getpostman.com/view/42847133/2sB34iiefB)
## Base URL

```
http://localhost:8080
```

---

## Endpoints

### 1. Get All Tasks

- **GET** `/tasks`
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
curl -X GET http://localhost:8080/tasks
```

**Windows CMD:**
```cmd
curl -X GET http://localhost:8080/tasks
```

---

### 2. Get Task By ID

- **GET** `/tasks/{id}`
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
curl -X GET http://localhost:8080/tasks/1
```

**Windows CMD:**
```cmd
curl -X GET http://localhost:8080/tasks/1
```

---

### 3. Create Task

- **POST** `/tasks`
- **Description:** Create a new task.

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
  -H "Content-Type: application/json" \
  -d '{"id":"4","title":"Task 4","description":"Fourth task","due_date":"2024-06-13T12:00:00Z","status":"Pending"}'
```

**Windows CMD:**
```cmd
curl -X POST http://localhost:8080/tasks ^
  -H "Content-Type: application/json" ^
  -d "{\"id\":\"4\",\"title\":\"Task 4\",\"description\":\"Fourth task\",\"due_date\":\"2024-06-13T12:00:00Z\",\"status\":\"Pending\"}"
```

---

### 4. Update Task

- **PUT** `/tasks/{id}`
- **Description:** Update an existing task by ID.

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
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Task 1","description":"Updated description","due_date":"2024-06-15T12:00:00Z","status":"Completed"}'
```

**Windows CMD:**
```cmd
curl -X PUT http://localhost:8080/tasks/1 ^
  -H "Content-Type: application/json" ^
  -d "{\"title\":\"Updated Task 1\",\"description\":\"Updated description\",\"due_date\":\"2024-06-15T12:00:00Z\",\"status\":\"Completed\"}"
```

---

### 5. Delete Task

- **DELETE** `/tasks/{id}`
- **Description:** Delete a task by its ID.

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
curl -X DELETE http://localhost:8080/tasks/1
```

**Windows CMD:**
```cmd
curl -X DELETE http://localhost:8080/tasks/1
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

