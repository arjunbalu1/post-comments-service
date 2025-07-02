# Post-Comments Service

A simple RESTful service for creating posts and comments, with Markdown support for comments (rendered as HTML).

---

## Database
- **Database name:** `post_comments_service`
- **Tables:** `posts`, `comments`

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

4. **Run the application:**
   ```sh
   go run main.go
   ```
   - The server will start on `http://localhost:8080`

---

## API Routes & Examples

### 1. Create a Post
- **POST** `/posts`
- **Body:**
  ```json
  {
    "title": "My First Post",
    "content": "Hello, world!"
  }
  ```
- **Response:**
  ```json
  {
    "ID": 1,
    "title": "My First Post",
    "content": "Hello, world!",
    "CreatedAt": "2025-07-02T22:18:51.351651+05:30",
    "comments": []
  }
  ```

---

### 2. List All Posts
- **GET** `/posts`
- **Response:**
  ```json
  [
    {
      "ID": 1,
      "title": "My First Post",
      "content": "Hello, world!",
      "CreatedAt": "2025-07-02T22:18:51.351651+05:30",
      "comments": []
    }
  ]
  ```

---

### 3. Get a Post with Comments (Markdown rendered as HTML)
- **GET** `/posts/{id}`
- **Response:**
  ```json
  {
    "ID": 1,
    "title": "My First Post",
    "content": "Hello, world!",
    "CreatedAt": "2025-07-02T22:18:51.351651+05:30",
    "comments": [
      {
        "ID": 1,
        "post_id": 1,
        "content": "\u003cp\u003eThis is a \u003cem\u003eMarkdown\u003c/em\u003e comment!\u003c/p\u003e\n",
        "CreatedAt": "2025-07-02T22:31:48.617642+05:30"
      }
    ]
  }
  ```

---

### 4. Add a Comment to a Post
- **POST** `/posts/{id}/comments`
- **Body:**
  ```json
  {
    "content": "This is a *Markdown* comment!"
  }
  ```
- **Response:**
  ```json
  {
    "ID": 1,
    "post_id": 1,
    "content": "This is a *Markdown* comment!",
    "CreatedAt": "2025-07-02T22:31:48.617642+05:30"
  }
  ```

---

## Notes
- Comments are stored as Markdown and rendered as HTML when fetching a post.
- HTML tags in the JSON response are escaped (e.g., `\u003c` for `<`) for security. This is normal and will render correctly in browsers/frontends.

---

## Architecture Overview
- **Go** with **Gin** for HTTP server
- **GORM** for ORM/database access
- **PostgreSQL** for data storage
- **gomarkdown/markdown** for Markdown rendering

---