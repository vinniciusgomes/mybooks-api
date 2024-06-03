## MyBooks API Documentation

### Introduction
Welcome to the documentation for the MyBooks API. MyBooks is a SaaS project designed to help users organize their personal library. This API provides endpoints for managing libraries, books, user profiles, billing, and loan status.

### Authentication
Authentication is required for most endpoints. MyBooks supports magic link authentication and authentication via Google account.

#### Endpoints:
- `POST v1/auth`: Authenticate using magic link.
- `POST v1/auth/google`: Authenticate using Google account

### Libraries
Manage libraries where users can organize their books.

#### Endpoints:
- `GET v1/libraries`: Get all libraries.
- `GET v1/libraries/{libraryId}`: Get library by ID.
- `POST v1/libraries`: Create a new library.
- `PUT v1/libraries/{libraryId}`: Update a library.
- `DELETE v1/libraries/{libraryId}`: Delete a library.
- `POST v1/libraries/{libraryId}/books`: Add books to a library.

### Books
Manage books within libraries.

#### Endpoints:
- `GET v1/books`: Get all books.
- `GET v1/books/{bookId}`: Get book by ID.
- `POST v1/books`: Create a new book.
- `PUT v1/books/{bookId}`: Update a book.
- `DELETE v1/books/{bookId}`: Delete a book.

### Profile
Manage user profiles.

#### Endpoints:
- `PUT v1/profile/photo`: Update profile photo.
- `PUT v1/profile`: Update name, email, and password.
- `DELETE v1/profile`: Delete the account.

### Billing
Manage billing details and subscription plans.

#### Endpoints:
- `GET v1/billing`: Get billing details for the account ($5 per account).
- `POST v1/subscribe`: Subscribe to a plan.

### Loan
Manage loan status of books.

#### Endpoints:
- `POST v1/loans`: Create a loan
- `GET v1/loans`: Get all loans
- `PUT v1/loans/:loanId/return`: Mark loan as returned

### Running Locally
To run the service locally, you can use [Air](https://github.com/cosmtrek/air) for hot-reloading. Run the following command:
```
air
```

### Running in Production with Docker
To run the service in a production environment using Docker, follow these steps:
1. Run Docker compose:
```
docker-compose up --build
```

Now you should have the MyBooks API service up and running locally or in a production environment.

## Features roadmap

### Authentication

- [ ] Should be able to authenticate using magic link;
- [ ] Should be able to authenticate using Google account;

### Library ✅

- [X] Should be able to create a new library;
- [X] Should be able to get all libraries;
- [X] Should be able to get a library by id;
- [X] Should be able to update a library;
- [X] Should be able to delete a library;
- [X] Should be able to add books in a library;

### Books ✅

- [x] Should be able to create a new book;
- [x] Should be able to get all books;
- [x] Should be able to get all books with params;
  - [X] Should be able to filter books by title
  - [X] Should be able to filter books by author
  - [X] Should be able to filter books by genre
  - [X] Should be able to filter books by ISBN
  - [X] Should be able to filter books by language
  - [X] Should be able to filter books by read
- [X] Should be able to get book by id;
- [X] Should be able to update a book;
- [X] Should be able to delete a book;

### Profile

- [ ] Should be able to update the profile photo;
- [ ] Should be able to update name, email, and password if they exist;
- [ ] Should be able to delete the account;

### Billing

- [ ] Should be able to get billing details for the account ($5 per account);
- [ ] Should be able to subscribe to a plan;

### Loan ✅

- [X] Should be able to create a loan and indicate to whom;
- [x] Should be able to get all loans;
- [X] Should be able to mark a loan as returned;
