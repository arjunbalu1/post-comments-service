# Post-Comments Service

A simple RESTful service for creating posts and comments, with Markdown support for comments (rendered as HTML). Now includes user authentication (username + password, JWT), and input validation.

---

## Database
- **Database name:** `post_comments_service`
- **Tables:** `users`, `posts`, `comments`
- **Why PostgreSQL?** I picked PostgreSQL because I like my data consistent and my queries snappy. GORM lets me move fast without breaking things.

---

## Setup Instructions

1. **Install dependencies:**
   - Go
   - PostgreSQL (running on `localhost:5432`)

2. **Create the database:**
   - Use pgAdmin or psql:
     ```sql
     CREATE DATABASE post_comments_service;
     ```

3. **Configure environment (optional):**
   - By default, the app connects with:
     - user: `postgres`
     - password: `postgres`
     - host: `localhost`
     - db: `post_comments_service`
   - To override, set the `DATABASE_URL` environment variable:
     ```sh
     export DATABASE_URL="host=localhost user=postgres dbname=post_comments_service sslmode=disable password=YOUR_PASSWORD"
     ```
   - To set a custom JWT secret:
     ```sh
     export JWT_SECRET="your_secret_key"
     ```

4. **Run the application:**
   ```sh
   go run main.go
   ```
   - The server will start on `http://localhost:8080`

---

## API Routes & Examples

### 1. Register a User
- **POST** `/register`
- **Body:**
  ```json
  {
    "username": "alice",
    "password": "password123"
  }
  ```
- **Response:**
  ```json
  { "message": "User registered successfully" }
  ```

---

### 2. Login
- **POST** `/login`
- **Body:**
  ```json
  {
    "username": "alice",
    "password": "password123"
  }
  ```
- **Response:**
  ```json
  { "token": "<JWT token>" }
  ```

---

### 3. Create a Post (Authenticated)
- **POST** `/posts`
- **Headers:**
  - `Authorization: Bearer <JWT token>`
- **Body:**
  ```json
  {
    "title": "My First Post",
    "content": "Hello, world!"
  }
  ```
- **Validation:**
  - `title` and `content` are required and must be non-empty.
- **Response:**
  ```json
  {
    "ID": 1,
    "title": "My First Post",
    "content": "Hello, world!",
    "CreatedAt": "...",
    "comments": [],
    "username": "alice"
  }
  ```

---

### 4. List All Posts
- **GET** `/posts`
- **Response:**
  ```json
  [
    {
      "ID": 1,
      "title": "My First Post",
      "content": "Hello, world!",
      "CreatedAt": "...",
      "comments": [],
      "username": "alice"
    }
  ]
  ```

---

### 5. Get a Post with Comments (Markdown rendered as HTML)
- **GET** `/posts/{id}`
- **Response:**
  ```json
  {
    "ID": 1,
    "title": "My First Post",
    "content": "Hello, world!",
    "CreatedAt": "...",
    "comments": [
      {
        "ID": 1,
        "post_id": 1,
        "content": "<p>This is a <em>Markdown</em> comment!</p>\n",
        "CreatedAt": "...",
        "username": "alice"
      }
    ],
    "username": "alice"
  }
  ```

---

### 6. Add a Comment to a Post (Authenticated)
- **POST** `/posts/{id}/comments`
- **Headers:**
  - `Authorization: Bearer <JWT token>`
- **Body:**
  ```json
  {
    "content": "This is a *Markdown* comment!"
  }
  ```
- **Validation:**
  - `content` is required and must be non-empty.
- **Response:**
  ```json
  {
    "ID": 1,
    "post_id": 1,
    "content": "This is a *Markdown* comment!",
    "CreatedAt": "...",
    "username": "alice"
  }
  ```

---

## Authentication Flow
- Register a user with `/register`.
- Login with `/login` to receive a JWT token.
- Use the JWT token in the `Authorization` header (`Bearer <token>`) for all protected endpoints (creating posts/comments).

---

## Input Validation
- **Register:** `username` and `password` are required.
- **Create Post:** `title` and `content` are required and must be non-empty.
- **Add Comment:** `content` is required and must be non-empty.

---

## Architecture Overview
- **Go** with **Gin** for HTTP server
- **GORM** for ORM/database access
- **PostgreSQL** for data storage
- **gomarkdown/markdown** for Markdown rendering
- **JWT** for authentication
- **bcrypt** for password hashing

---

## Notes
- Comments are stored as Markdown and rendered as HTML when fetching a post.
- HTML tags in the JSON response are escaped (e.g., `\u003c` for `<`) for security. This is normal and will render correctly in browsers/frontends.
- All post and comment creation requires authentication.
- Username is the only user identifier (no email).

---