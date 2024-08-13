# Task Manager Refactored Architecture Documentation

## 1. Overview

The Task Manager project is designed as a modular web application using Go. It utilizes the Gin framework for HTTP routing and middleware, and incorporates JWT for authentication and password hashing for security. The architecture follows a clean separation of concerns to ensure maintainability and scalability.

## 2. Folder Structure

```plaintext
task-manager/
├── api/
│   ├── middleware/
│   │   ├── admin.go
│   │   └── auth.go
│   ├── router/
│   │   └── router.go
│   └── controller/
│       ├── task_controller.go
│       └── user_controller.go
├── Domain/
│   └── task.go
│   └── user.go 
├── cmd/
│   └── main.go
├── config/
│   └── config.go
│   ├── auth_middleWare.go
│   ├── jwt_service.go
│   ├── password_service.go
├── Repository/
│   ├── task_repository.go
│   └── user_repository.go
├── usecase/
│   ├── task_usecases.go
│   └── user_usecases.go

```

## 3. Components

- **api/**: Contains the controllers, middleware, and router.
- **middleware/**: Contains middleware functions like authentication and admin check.
- **router/**: Contains route definitions.
- **controller/**: Contains handlers for HTTP requests.
- **Domain/**: Contains the domain models (entities) used throughout the application.
- **cmd/**: Entry point of the application.
- **config/**: Contains configuration, JWT, and password handling services.
- **Repository/**: Contains the data access layer, handling interactions with the database or in-memory storage.
- **usecase/**: Contains business logic, coordinating data flow between the controllers and repositories.
## API Endpoints

### User Endpoints

- **POST `/register`**: Register a new user.
  - Request Body:
    ```json
    {
      "username": "user123",
      "password": "password",
      "role": "user"
    }
    ```

- **POST `/login`**: Authenticate a user and receive a JWT token.
  - Request Body:
    ```json
    {
      "username": "user123",
      "password": "password"
    }
    ```
  - Response:
    ```json
    {
      "token": "jwt_token_here"
    }
    ```

### Task Endpoints

- **GET `/tasks`**: Retrieve all tasks (admin only) or tasks for the logged-in user.
  
- **GET `/tasks/:id`**: Retrieve a specific task by ID.
  
- **POST `/tasks`**: Create a new task.
  - Request Body:
    ```json
    {
      "title": "New Task",
      "description": "Task description",
      "duedate": "2023-09-01T00:00:00Z",
      "status": "in_progress"
    }
    ```

- **PUT `/tasks/:id`**: Update a specific task by ID.
  - Request Body:
    ```json
    {
      "title": "Updated Task",
      "description": "Updated task description",
      "duedate": "2023-09-01T00:00:00Z",
      "status": "complete"
    }
    ```

- **DELETE `/tasks/:id`**: Delete a specific task by ID.

## Middleware

### Authentication Middleware

- **Function**: Verifies the JWT token sent in the `Authorization` header.
- **Purpose**: Ensures that only authenticated users can access protected routes.
- **Implementation**:
  - The `AuthMiddleware` checks for the `Authorization` header, validates the JWT token, and extracts user details.
  - If the token is invalid or missing, the middleware returns a `401 Unauthorized` response.

### Admin Middleware

- **Function**: Restricts access to certain routes to users with the admin role.
- **Purpose**: Ensures that only admins can manage all tasks.
- **Implementation**:
  - The `AdminMiddleware` checks the user's role from the JWT token and allows access only if the role is `admin`.
  - If the user is not an admin, the middleware returns a `403 Forbidden` response.

## Models

### User Model

- **Structure**:
  ```go
  type User struct {
      ID       int    `json:"id"`
      Username string `json:"username"`
      Password string `json:"password"`
      Role     string `json:"role"`
  }
  ```

### Task Model

- **Structure**:
  ```go
  type Task struct {
      ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
      Title       string             `json:"title" bson:"title"`
      Description string             `json:"description" bson:"description"`
      DueDate     time.Time          `json:"duedate" bson:"duedate"`
      Status      TaskStatus         `json:"status" bson:"status"`
      UserID      int                `json:"userID" bson:"userID"`
  }


## 4. Design Decisions

1. **Separation of Concerns**:
   - **Reasoning**: By dividing the application into distinct layers (Delivery, Domain, Infrastructure, Repositories, Usecases), the architecture ensures that each layer has a single responsibility, which simplifies maintenance and testing.

2. **Use of Gin Framework**:
   - **Reasoning**: Gin provides a fast HTTP web framework with minimal overhead, making it suitable for high-performance applications.

3. **JWT for Authentication**:
   - **Reasoning**: JWT is used for stateless authentication, which simplifies scalability and security.

4. **Password Hashing**:
   - **Reasoning**: Secure password storage and comparison are implemented using hashing to protect user credentials.

5. **Unit Testing and Mocks**:
   - **Reasoning**: Mock implementations of use cases are used in tests to isolate and verify individual components without dependency on external systems.

## 5. Guidelines for Future Development

1. **Adding New Features**:
   - **Follow the Layered Architecture**: Ensure that new features are added in the appropriate layer (e.g., business logic in Usecases, routing in Controllers).
   - **Write Tests**: Add corresponding tests for new features to maintain code quality and functionality.

2. **Refactoring**:
   - **Maintain Consistency**: Refactor code while adhering to the architectural boundaries and naming conventions.
   - **Document Changes**: Update documentation to reflect any changes in the architecture or design.

3. **Error Handling**:
   - **Centralize Error Handling**: Use centralized error handling mechanisms in the Delivery layer to manage errors uniformly.

4. **Security**:
   - **Review Security Practices**: Regularly review and update security practices, especially in authentication and password management.

5. **Performance**:
   - **Profile and Optimize**: Use profiling tools to identify performance bottlenecks and optimize code as needed.

6. **Code Reviews**:
   - **Peer Reviews**: Implement a peer review process to ensure code quality and adherence to architectural guidelines.

## 6. Conclusion

This documentation provides a comprehensive overview of the refactored architecture of the Task Manager project. By adhering to the guidelines and design decisions outlined, future development can be streamlined, ensuring maintainability, scalability, and robustness of the application.
POStMAN DOCUMENTATION: [https://documenter.getpostman.com/view/37289771/2sA3rzKsPp]
