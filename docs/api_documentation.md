# Task Manager API Documentation

A simple RESTful API for managing tasks. This API uses MongoDB as the database to store task information. You can create, view, update, and delete tasks using standard HTTP methods. This API is suitable for learning, prototyping, or integrating basic task management features into your applications.

## How to Run

1. **Clone the repository:**
    ```sh
    git clone https://github.com/yourusername/Task-manager-go.git
    cd Task-manager-go
    git checkout persist
    ```

2. **Set up MongoDB:**
    - This project uses MongoDB as the database. Ensure you have a MongoDB instance running.
    - You must provide your MongoDB Atlas username and password in the connect_db function inside data/task_service.go.
    - Example connection string (replace <username> and <password>):
    
      ```go
      clientOptions := options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.tj8um.mongodb.net/?retryWrites=true&w=majority")
      ```

3. **Build and start the server:**
    ```sh
    go run main.go
    ```
    or:
    ```sh
    go run .
    ```

4. The API will be available at:
    ```
    http://localhost:8080
    ```

## MongoDB Integration

- The API uses a MongoDB collection named `Tasks` in a database called `Task-Database`.
- Each task is stored as a document in the collection with the following fields:
  - `id` (string): Unique identifier for the task.
  - `title` (string): Title of the task.
  - `description` (string): Description of the task.
  - `due_date` (string): Due date in RFC3339 format.
  - `status` (string): Status of the task (e.g., Pending, Completed).

### Error Handling

- **Database Connection Errors:** If the API cannot connect to MongoDB, the server will log the error and terminate.
- **Task Not Found:** If a task with the specified ID does not exist, the API will return a `404 Not Found` response.
- **Duplicate Task ID:** When creating a task, if a task with the same ID already exists, the API will return a `409 Conflict` response.
- **Invalid Request Body:** If the request body is invalid (e.g., missing required fields), the API will return a `400 Bad Request` response.

## API Usage

This project includes a Postman collection and curl examples for testing and understanding the API.

### Getting Started
Open [Postman](https://.postman.co/workspace/My-Workspace~e85ab91d-850f-41f6-8c20-cb0459fbaf68/collection/42847133-a206412b-02ed-4c30-8eec-53822038a224?action=share&creator=42847133)

Or view the docs online: [API Documentation](https://documenter.getpostman.com/view/42847133/2sB34iiefB)
1. Click **“Run in Postman”** (top-right of the page)
2. Postman will automatically open and load the collection for you
3. Start making requests using the examples provided in the docs

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

