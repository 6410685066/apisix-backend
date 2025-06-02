# API Documentation

## Authentication

### POST /api/login
- **Description:** Login and receive a JWT token.
- **Request Body:**
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
- **Response:**
    - `200 OK`
        ```json
        {
          "success": true,
          "token": "jwt-token",
          "id": 1,
          "username": "user"
        }
        ```
    - `401 Unauthorized`
        ```json
        { "error": "Invalid username or password" }
        ```

---

## Product

> **All endpoints below require a JWT token in the `Authorization` header.**

### GET /api/data
- **Description:** Retrieve a list of products.
- **Query Parameters:** (As supported by the system)
- **Response:**
    - `200 OK`
        ```json
        [
          {
            "id": 1,
            "name": "Product A",
            "price": 100
          }
        ]
        ```

### POST /api/data
- **Description:** Create a new product.
- **Request Body:**
    ```json
    {
      "name": "Product A",
      "price": 100
    }
    ```
- **Response:**
    - `200 OK`
        ```json
        { "success": true }
        ```
    - `400 Bad Request`
        ```json
        { "error": "..." }
        ```

### PUT /api/data/:id
- **Description:** Update a product by ID (replace all fields).
- **Request Body:**
    ```json
    {
      "name": "Product A",
      "price": 120
    }
    ```
- **Response:**
    - `200 OK`
        ```json
        { "success": true }
        ```
    - `400 Bad Request`
        ```json
        { "error": "Invalid product ID" }
        ```

### PATCH /api/data/:id
- **Description:** Update specific fields of a product by ID.
- **Request Body:** (Only the fields to be updated)
    ```json
    {
      "price": 130
    }
    ```
- **Response:**
    - `200 OK`
        ```json
        { "success": true }
        ```
    - `400 Bad Request`
        ```json
        { "error": "Invalid product ID" }
        ```

### DELETE /api/data/:id
- **Description:** Delete a product by ID.
- **Response:**
    - `200 OK`
        ```json
        { "success": true }
        ```
    - `400 Bad Request`
        ```json
        { "error": "Invalid product ID" }
        ```

---

**Note:**  
- All endpoints except `/login` require a JWT token in the header:  
  `Authorization: Bearer <token>`
- Replace `/api` with the prefix from your config (`APIPrefix`) if different.