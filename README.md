## MyBooks API Documentation

### Introduction
Welcome to the documentation for the MyBooks API. MyBooks is a SaaS project designed to help users organize their personal library. This API provides endpoints for managing libraries, books, user profiles, billing, loan status, and reading status.

### Authentication
Authentication is required for most endpoints. MyBooks supports magic link authentication and authentication via Google account.

#### Endpoints:
- `/auth`: Authenticate using magic link.
- `/auth/google`: Authenticate using Google account

### Libraries
Manage libraries where users can organize their books.

#### Endpoints:
- `GET /libraries`: Get all libraries.
- `POST /libraries`: Create a new library.
- `PUT /libraries/{libraryId}`: Update a library.
- `DELETE /libraries/{libraryId}`: Delete a library.
- `POST /libraries/{libraryId}/books`: Add books to a library.

### Books
Manage books within libraries.

#### Endpoints:
- `GET /books`: Get all books.
- `POST /books`: Create a new book.
- `PUT /books/{bookId}`: Update a book.
- `DELETE /books/{bookId}`: Delete a book.

### Profile
Manage user profiles.

#### Endpoints:
- `PUT /profile/photo`: Update profile photo.
- `PUT /profile`: Update name, email, and password.
- `DELETE /profile`: Delete the account.

### Billing
Manage billing details and subscription plans.

#### Endpoints:
- `GET /billing`: Get billing details for the account ($5 per account).
- `POST /subscribe`: Subscribe to a plan.

### Loan
Manage loan status of books.

#### Endpoints:
- `PUT /books/{bookId}/loan`: Mark a book as loaned and indicate to whom.
- `PUT /books/{bookId}/return`: Mark a book as returned.

### Reading Status
Manage reading status of books.

#### Endpoints:
- `PUT /books/{bookId}/read`: Mark a book as read.
- `DELETE /books/{bookId}/read`: Remove the reading status from a book.

### Running Locally
To run the service locally, you can use [Air](https://github.com/cosmtrek/air) for hot-reloading. Run the following command:
```
air
```

### Running in Production with Docker
To run the service in a production environment using Docker, follow these steps:
1. Build the Docker image:
```
docker build --tag mybooks .
```
2. Run the Docker container, mapping port 8080 on your local machine to port 8080 in the container:
```
docker run -p 8080:8080 mybooks
```

Now you should have the MyBooks API service up and running locally or in a production environment.

## Features roadmap

### Authentication

- [ ] Should be able to authenticate using magic link;
- [ ] Should be able to authenticate using Google account;

### Library

- [ ] Should be able to create a new library;
- [ ] Should be able to get all libraries;
- [ ] Should be able to update a library;
- [ ] Should be able to delete a library;
- [ ] Should be able to add books in a library;

### Books

- [x] Should be able to create a new book;
- [x] Should be able to get all books;
- [X] Should be able to get book by id;
- [ ] Should be able to update a book;
- [X] Should be able to delete a book;

### Profile

- [ ] Should be able to update the profile photo;
- [ ] Should be able to update name, email, and password if they exist;
- [ ] Should be able to delete the account;

### Billing

- [ ] Should be able to get billing details for the account ($5 per account);
- [ ] Should be able to subscribe to a plan;

### Loan

- [ ] Should be able to mark a book as loaned and indicate to whom;
- [ ] Should be able to mark a book as returned;

### Reading Status

- [ ] Should be able to mark a book as read;
- [ ] Should be able to remove the reading status from a book.
