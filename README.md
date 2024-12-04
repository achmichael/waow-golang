# API Documentation

## Table of Contents
- [Technologies Used](#technologies-used)
- [Libraries Used](#libraries-used)
- [Routes Documentation](#routes-documentation)
- [Route Security and Middlewares](#route-security-and-middlewares)
- [Getting Started](#getting-started)
- [License](#license)

## Technologies Used

This project uses the following technologies:

### 1. [Gin](https://github.com/gin-gonic/gin)
Gin is a high-performance web framework for Go (Golang). It provides a fast HTTP router and various features like middleware support, routing, and more.

- **Why Gin?**  
  Gin is known for its speed, small memory footprint, and ease of use. It is suitable for building REST APIs and web applications in Go.

- **Official Documentation**: [Gin Documentation](https://gin-gonic.com/docs/)

### 2. [JWT (JSON Web Tokens)](https://jwt.io/)
JWT is a compact and self-contained way to securely transmit information between parties as a JSON object. It is commonly used for authorization and information exchange.

- **Why JWT?**  
  JWT is used in this project for securely managing authentication tokens and ensuring that the user is authenticated before accessing protected routes.

- **Official Documentation**: [JWT Documentation](https://jwt.io/introduction/)

### 3. [GORM](https://gorm.io/)
GORM is an Object Relational Mapping (ORM) library for Golang. It provides a simple and powerful way to interact with SQL databases.

- **Why GORM?**  
  GORM allows developers to interact with the database using Go objects, making it easier to write and maintain database queries.

- **Official Documentation**: [GORM Documentation](https://gorm.io/docs/)

### 4. [MySQL](https://www.mysql.com/)
MySQL is an open-source relational database management system. It is widely used for storing structured data in the form of tables.

- **Why MySQL?**  
  MySQL is a reliable and fast RDBMS that is widely supported and commonly used in production environments.

- **Official Documentation**: [MySQL Documentation](https://dev.mysql.com/doc/)

### 5. [UUID](https://github.com/google/uuid)
UUID is a library for generating universally unique identifiers (UUIDs) in Go. It is often used to generate unique identifiers for resources in distributed systems.

- **Why UUID?**  
  UUID ensures that each generated ID is unique across systems and machines, which is useful for creating globally unique resources in APIs.

- **Official Documentation**: [UUID Documentation](https://pkg.go.dev/github.com/google/uuid)

### 6. [Godotenv](https://github.com/joho/godotenv)
Godotenv loads environment variables from a `.env` file into the Go environment, making it easy to configure applications.

- **Why Godotenv?**  
  Godotenv simplifies the process of managing environment variables, especially for configuration settings like database credentials, API keys, and other secrets.

- **Official Documentation**: [Godotenv Documentation](https://pkg.go.dev/github.com/joho/godotenv)

---

## Libraries Used

This project uses the following libraries:

- **[Gin](https://github.com/gin-gonic/gin)** for web routing and middleware.
- **[JWT](https://github.com/golang-jwt/jwt)** for handling JSON Web Tokens for authentication and authorization.
- **[GORM](https://gorm.io/)** for interacting with MySQL databases.
- **[UUID](https://github.com/google/uuid)** for generating unique identifiers.
- **[Godotenv](https://github.com/joho/godotenv)** for managing environment variables.

---

## Routes Documentation

### 1. **Authentication Routes**
- **POST** `/login`:  
  - **Description**: Login the user by verifying credentials.
  - **Request Body**: JSON with username and password.
  ```json
  {
    "username": "johndoe",
    "password": "password123"
  }
  ```
  - **Response**: JWT token on success.
  ```json
  {
    "status": true,
    "data": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzMzODM2MTksInJvbGUiOiJVU0VSIiwidXNlcm5hbWUiOiJqb2huZG9lMTIzIn0.UFUXHV1tKBzWAtZr6NUhic6CGBYQ_TJp6Zhzg7_vVwE"
    }
  }
  ```

- **POST** `/register`:  
  - **Description**: Register a new user.
  - **Request Body**: JSON with user details (username, password, email, displayName, bio, profilePictureUrl, role).
  ```json
  {
    "username": "michael243",
    "password": "password123",
    "email": "michael32@example.com",
    "displayName": "John Doe",
    "bio": "Software engineer",
    "profilePictureUrl": "#",
    "role": "ADMIN"
  }
  ```
  - **Response**: Success or error message.
  ```json
  {
    "status": true,
    "message": "User created successfully"
  }
  ```

### 2. **User Routes**
- **POST** `/user`:  
  - **Description**: Register a new user (requires authentication).
  - **Request Body**: JSON with user details.
  ```json
  {
    "username": "michael243",
    "password": "password123",
    "email": "michael32@example.com",
    "displayName": "John Doe",
    "bio": "Software engineer",
    "profilePictureUrl": "#",
    "role": "ADMIN"
  }
  ```
  - **Response**: Success or error message.
  ```json
  {
    "status": true,
    "message": "User created successfully"
  }
  ```

- **GET** `/user`:  
  - **Description**: Get all users (requires admin authentication).
  - **Response**: JSON array of users.
  ```json
  {
    "status": true,
    "data": []
  }
  ```

- **GET** `/user/:id`:  
  - **Description**: Get a specific user by ID (requires admin authentication).
  - **Response**: JSON with user details.
  ```json
  {
    "status": true,
    "data": {
      "id": 1,
      "username": "john_doe",
      "email": "john@example.com",
      "displayName": "John Doe",
      "bio": "Software engineer",
      "profilePictureUrl": "#",
      "role": "USER"
    }
  }
  ```

- **PUT** `/user/:id`:  
  - **Description**: Update user details by ID (requires admin authentication).
  - **Request Body**: JSON with updated user details.
  ```json
  {
    "username": "john_doe",
    "email": "john@example.com",
    "displayName": "John Doe",
    "bio": "Software engineer",
    "role": "USER"
  }
  ```
  - **Response**: Success or error message.
  ```json
  {
    "status": true,
    "message": "User updated successfully"
  }
  ```

- **DELETE** `/user/:id`:  
  - **Description**: Delete a user by ID (requires admin authentication).
  - **Response**: Success or error message.
  ```json
  {
    "status": true,
    "message": "User deleted successfully"
  }
  ```

### 3. **Article Routes**
- **POST** `/article`:  
  - **Description**: Create a new article (requires authentication).
  - **Request Body**: JSON with article details.
  ```json
  {
    "user_id": "user_id",
    "category_id": "category_id",
    "title": "Michael",
    "content": "Michael anak yang baik",
    "slug": "michael",
    "view_count": 0
  }
  ```
  - **Response**: Success or error message.
  ```json
  {
    "status": true,
    "message": "Article created successfully"
  }
  ```

- **GET** `/article`:  
  - **Description**: Get all articles.
  - **Response**: JSON array of articles.
  ```json
  {
    "status": true,
    "data": [
      {
        "id": 1,
        "user_id": "user_id",
        "category_id": "category_id",
        "title": "Michael",
        "content": "Michael anak yang baik",
        "slug": "michael",
        "view_count": 0
      },
      {
        "id": 2,
        "user_id": "user_id",
        "category_id": "category_id",
        "title": "Michael",
        "content": "Michael anak yang baik",
        "slug": "michael",
        "view_count": 0
      }
    ]
  }
  ```

- **GET** `/article/:id`:  
  - **Description**: Get a specific article by ID.
  - **Response**: JSON with article details.
  ```json
  {
    "status": true,
    "data": {
      "id": 1,
      "user_id": "user_id",
      "category_id": "category_id",
      "title": "Michael",
      "content": "Michael anak yang baik",
      "slug": "michael",
      "view_count": 0
    }
  }
  ```

- **PUT** `/article/:id`:  
  - **Description**: Update an article by ID (requires authentication).
  - **Request Body**: JSON with updated article details.
  ```json
  {
    "user_id": "user_id",
    "category_id": "category_id",
    "title": "Michael",
    "content": "Michael anak yang baik",
    "slug": "michael",
    "view_count": 0
  }
  ```
  - **Response**: Success or error message.
  ```json
  {
    "status": true,
    "message": "Article updated successfully"
  }
  ```

- **DELETE** `/article/:id`:  
  - **Description**: Delete an article by ID (requires authentication).
  - **Response**: Success or error message.
  ```json
  {
    "status": true,
    "message": "Article deleted successfully"
  }
  ```

### 4. **Category Routes**
- **POST** `/categories`:  
  - **Description**: Create a new category (requires authentication).
  - **Request Body**: JSON with category details.
  ```json
  {
    "name": "Michael",
    "description": "Michael anak yang baik"
  }
  ```
  - **Response**: Success or error message.
  ```json
  {
    "status": true,
    "data": {
      "id": "category_id",
      "updated_at": "0001-01-01T00:00:00Z",
      "created_at": "2024-12-04T14:56:50.124+07:00",
      "name": "Michael",
      "description": "haloooo",
      "slug": "michael"
    }
  }
  ```

- **GET** `/categories`:  
  - **Description**: Get all categories.
  - **Response**: JSON array of categories.
  ```json
  {
    "status": true,
    "data": [
      {
        "id": "category_id",
        "updated_at": "0001-01-01T00:00:00Z",
        "created_at": "2024-12-04T14:56:50.124+07:00",
        "name": "Michael",
        "description": "haloooo",
        "slug": "michael"
      }
    ]
  }
  ```

- **GET** `/categories/:id`:  
  - **Description**: Get a specific category by ID.
  - **Response**: JSON with category details.
  ```json
  {
    "status": true,
    "data": {
      "id": "category_id",
      "updated_at": "0001-01-01T00:00:00Z",
      "created_at": "2024-12-04T14:56:50.124+07:00",
      "name": "Michael",
      "description": "haloooo",
      "slug": "michael"
    }
  }
  ```

- **PUT** `/categories/:id`:  
  - **Description**: Update a category by ID (requires authentication).
  - **Request Body**: JSON with updated category details.
  ```json
  {
    "name": "Michael",
    "description": "Michael anak yang baik"
  }
  ```
  - **Response**: Success or error message.
  ```json
  {
    "status": true,
    "message": "Category updated successfully"
  }
  ```

- **DELETE** `/categories/:id`:  
  - **Description**: Delete a category by ID (requires authentication).
  - **Response**: Success or error message.
  ```json
  {
    "status": true,
    "message": "Category deleted successfully"
  }
  ```

---

## Route Security and Middlewares

### Middlewares Applied to Routes:
- **AuthJWT Middleware**: Ensures that the user is authenticated by verifying their JWT token.
- **MustUser Middleware**: Ensures that the authenticated user has basic user permissions.
- **MustAdmin Middleware**: Ensures that the authenticated user has admin permissions.

### Route Groups:
- **Article Routes**:  
  These routes require the user to be authenticated and have user-level access for creating, updating, and deleting articles.

- **Category Routes**:  
  These routes require the user to be authenticated and have user-level access for creating, updating, and deleting categories.

- **User Routes**:  
  These routes require the user to be authenticated and have admin-level access for creating, updating, and deleting users.

---

## Getting Started

1. Clone this repository.
2. Install dependencies using `go get`.
3. Create a `.env` file with necessary environment variables, such as database credentials and JWT secret.
4. Run the application using `go run main.go`.

---

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

