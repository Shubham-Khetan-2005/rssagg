# RSS Aggregator

An RSS Aggregator built with Go and PostgreSQL that collects and manages RSS feeds from multiple sources, enabling users to efficiently stay updated with their favorite content.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [Contact](#contact)

## Features

- **User Authentication**: Create users with unique API keys for personalized usage.
- **Feed Management**: Add and view RSS feeds from various sources.
- **Post Management**: Fetch and display posts from subscribed feeds.
- **Feed Follows**: Follow/unfollow feeds and manage your subscriptions.
- **Health Check**: Verify the application's readiness with a simple endpoint.

## Tech Stack

- **Programming Language**: Go
- **Database**: PostgreSQL
- **Web Framework**: Chi
- **Environment Management**: GoDotEnv
- **CORS**: Chi CORS middleware

## Installation

### Prerequisites

- Go 1.20 or later
- PostgreSQL 15 or later
- Git

### Steps

1. Clone the repository:

    ```bash
    git clone https://github.com/Shubham-Khetan-2005/rssagg.git
    cd rssagg
    ```

2. Set up the PostgreSQL database:
    - Create a new database (e.g., `rss_aggregator`).
    - Update the database connection details in the `.env` file.
    - Example - 
    ```bash
     `PORT`: The port on which the server will run (e.g., `8080`).
     `DB_URL`: The connection string for the PostgreSQL database.
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Run database migrations (if applicable):

    ```bash
    go run cmd/migrate/main.go
    ```

5. Start the application:

    ```bash
    go build
    ./rssagg.exe
    ```

## Usage

- Authentication: Most endpoints require an Authorization header with an API key.
- Database: Ensure the PostgreSQL database is configured and accessible at the DB_URL specified in the environment variables.
- Scraping: The server periodically scrapes feeds for updates, controlled by the startScrapping function.


# API Documentation

This API provides endpoints for managing users, RSS feeds, posts, and feed subscriptions.

---

## Endpoints

### **Health Check**
- **GET** `/v1/healthz`
  - **Description**: Checks if the server is running and ready.
  - **Response**:
    ```json
    {}
    ```

---

### **Error Testing**
- **GET** `/v1/error`
  - **Description**: Simulates an intentional error for testing purposes.
  - **Response**:
    ```json
    {
      "error": "Intentional error for testing."
    }
    ```

---

## Users

### **Create User**
- **POST** `/v1/users`
  - **Description**: Creates a new user and returns their details, including an API key.
  - **Request Body**:
    ```json
    {
      "name": "Shubham Khetan"
    }
    ```
  - **Response**:
    ```json
    {
      "id": "uuid",
      "created_at": "2025-01-15T10:00:00Z",
      "updated_at": "2025-01-15T10:00:00Z",
      "name": "Shubham Khetan",
      "api_key": "randomly-generated-api-key"
    }
    ```

### **Get User**
- **GET** `/v1/users`
  - **Description**: Retrieves the authenticated user's details.
  - **Headers**:
    ```
    Authorization: ApiKey <api_key>
    ```
  - **Response**:
    ```json
    {
      "id": "uuid",
      "created_at": "2025-01-15T10:00:00Z",
      "updated_at": "2025-01-15T10:00:00Z",
      "name": "John Doe",
      "api_key": "randomly-generated-api-key"
    }
    ```

---

## Feeds

### **Create Feed**
- **POST** `/v1/feeds`
  - **Description**: Adds a new RSS feed.
  - **Headers**:
    ```
    Authorization: ApiKey <api_key>
    ```
  - **Request Body**:
    ```json
    {
      "name": "Tech News",
      "url": "https://example.com/rss"
    }
    ```
  - **Response**:
    ```json
    {
      "id": "uuid",
      "created_at": "2025-01-15T10:00:00Z",
      "updated_at": "2025-01-15T10:00:00Z",
      "name": "Tech News",
      "url": "https://example.com/rss",
      "user_id": "uuid"
    }
    ```

### **Get Feeds**
- **GET** `/v1/feeds`
  - **Description**: Retrieves all RSS feeds.
  - **Response**:
    ```json
    [
      {
        "id": "uuid",
        "created_at": "2025-01-15T10:00:00Z",
        "updated_at": "2025-01-15T10:00:00Z",
        "name": "Tech News",
        "url": "https://example.com/rss",
        "user_id": "uuid"
      }
    ]
    ```

---

## Posts

### **Get Posts**
- **GET** `/v1/posts`
  - **Description**: Retrieves posts from feeds followed by the authenticated user.
  - **Headers**:
    ```
    Authorization: ApiKey <api_key>
    ```
  - **Response**:
    ```json
    [
      {
        "id": "uuid",
        "created_at": "2025-01-15T10:00:00Z",
        "updated_at": "2025-01-15T10:00:00Z",
        "title": "Post Title",
        "description": "Optional Description",
        "published_at": "2025-01-15T10:00:00Z",
        "url": "https://example.com/post",
        "feed_id": "uuid"
      }
    ]
    ```

---

## Feed Follows

### **Create Feed Follow**
- **POST** `/v1/feed_follows`
  - **Description**: Subscribes the authenticated user to a specific RSS feed.
  - **Headers**:
    ```
    Authorization: ApiKey <api_key>
    ```
  - **Request Body**:
    ```json
    {
      "feed_id": "uuid"
    }
    ```
  - **Response**:
    ```json
    {
      "id": "uuid",
      "created_at": "2025-01-15T10:00:00Z",
      "updated_at": "2025-01-15T10:00:00Z",
      "user_id": "uuid",
      "feed_id": "uuid"
    }
    ```

### **Get Feed Follows**
- **GET** `/v1/feed_follows`
  - **Description**: Retrieves all feed follows for the authenticated user.
  - **Headers**:
    ```
    Authorization: ApiKey <api_key>
    ```
  - **Response**:
    ```json
    [
      {
        "id": "uuid",
        "created_at": "2025-01-15T10:00:00Z",
        "updated_at": "2025-01-15T10:00:00Z",
        "user_id": "uuid",
        "feed_id": "uuid"
      }
    ]
    ```

### **Delete Feed Follow**
- **DELETE** `/v1/feed_follows/{feedFollowID}`
  - **Description**: Deletes a feed follow for the authenticated user.
  - **Headers**:
    ```
    Authorization: ApiKey <api_key>
    ```
  - **Response**:
    ```json
    {}
    ```


## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Submit a pull request.

## Contact

For any queries or issues, please contact me at  [janukhetan2005@gmail.com](mailto:janukhetan2005@gmail.com).

---



