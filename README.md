# A Simple CRUD Application using Golang

This is a simple RESTful API server for managing a collection of books, built with Golang and [Gorilla Mux](https://github.com/gorilla/mux).

## Features

- List all books
- Get a book by ID
- Create a new book (with or without specifying an ID)
- Update an existing book
- Delete a book

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) installed

### Install Dependencies

This project uses [Gorilla Mux](https://github.com/gorilla/mux) for routing. Install dependencies with:

```sh
go mod tidy
```

### Run the Server

Start the server with:

```sh
go run main.go
```

The server will start on `http://localhost:8080`.

## API Endpoints

### Get All Books

- **URL:** `/books`
- **Method:** `GET`
- **Description:** Returns a list of all books.

### Get Book by ID

- **URL:** `/books/{id}`
- **Method:** `GET`
- **Description:** Returns a single book by its ID.

### Create Book

- **URL:** `/books`
- **Method:** `POST`
- **Description:** Creates a new book. If `id` is omitted or empty, a random ID is generated.
- **Body Example:**
  ```json
  {
    "id": "",
    "title": "Harry Potter and the Sorcerer's Stone",
    "author": {
      "firstname": "J.K.",
      "lastname": "Rowling"
    },
    "price": 56.25
  }
  ```

### Update Book

- **URL:** `/books/{id}`
- **Method:** `PUT`
- **Description:** Updates the details of an existing book.
- **Body Example:**
  ```json
  {
    "title": "Harry Potter and the Chamber of Secrets",
    "author": {
      "firstname": "J.K.",
      "lastname": "Rowling"
    },
    "price": 1000
  }
  ```

### Delete Book

- **URL:** `/books/{id}`
- **Method:** `DELETE`
- **Description:** Deletes a book by its ID.

## Example Requests

You can use [Postman](https://www.postman.com/) to test the endpoints. A sample collection is provided in [`postman/go-crud-application.postman_collection.json`](postman/go-crud-application.postman_collection.json).

---

**Note:** All responses are in JSON format.  
If a book is not found, the server responds with a `404 Not Found` error.
