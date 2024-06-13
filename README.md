## MyBooks API Documentation

### Introduction
Welcome to the documentation for the MyBooks API. MyBooks is a SaaS project designed to help users organize their personal library. This API provides endpoints for managing libraries, books, user profiles, billing, and loan status.

## File structure
```
/myapp
|-- /cmd
|   |-- /api
|   |   |-- main.go
|   |
|-- /docs
|   |-- main.go
|   |-- swagger.json
|   |-- swagger.yaml
|   |
|-- /internal
|   |-- /domain
|   |   |-- /models
|   |   |   |-- auth_model.go
|   |   |   |-- book_model.go
|   |   |   |-- library_model.go
|   |   |   |-- loan_model.go
|   |   |   |-- user_model.go
|   |   |-- /repositories
|   |   |   |-- auth_repository.go
|   |   |   |-- book_repository.go
|   |   |   |-- library_repository.go
|   |   |   |-- loan_repository.go
|   |   |-- /services
|   |   |   |-- auth_service.go
|   |   |   |-- book_service.go
|   |   |   |-- library_service.go
|   |   |   |-- loan_service.go
|   |
|   |-- /infrastructure
|   |   |-- /api
|   |   |   |-- /handlers
|   |   |   |   |-- auth_handler.go
|   |   |   |   |-- books_handler.go
|   |   |   |   |-- library_handler.go
|   |   |   |   |-- loans_handler.go
|   |   |   |-- /middlewares
|   |   |   |   |-- auth_middleware.go
|   |   |   |   |-- cors_middleware.go
|   |   |   |-- server.go
|   |   |-- /config
|   |   |   |-- database.go
|   |   |-- /constants
|   |   |   |-- constants.go
|   |   |-- /helpers
|   |   |   |-- get_user_from_context.go
|   |   |   |-- handle_error.go
|   |
|-- /pkg
|   |-- generate_random_id.go
|   |-- send_email.go
|   |-- validate_model_struct.go
|
|-- Dockerfile
|-- docker-compose.yml
|-- go.mod
|-- go.sum
|-- makefile
```

## Features
### Authentication
Authentication is required for most endpoints. MyBooks supports credentials authentication and authentication via Google account.

#### Endpoints:
- `POST v1/auth/signup/credentials`: Create user with credentials.
- `POST v1/auth/signin/credentials`: Authentication user with credentials.
- `POST v1/auth/signout`: Logout user. 
- `POST v1/auth/forgot-password`: Forgot password.
- `POST v1/auth/reset-password/{token}`: Create a new password.
- `GET v1/auth/validate`: Validate access_token cookie.

### Libraries
Manage libraries where users can organize their books.

#### Endpoints:
- `GET v1/libraries`: Get all libraries.
- `GET v1/libraries/{libraryId}`: Get library by ID.
- `POST v1/libraries`: Create a new library.
- `PUT v1/libraries/{libraryId}`: Update a library.
- `DELETE v1/libraries/{libraryId}`: Delete a library.
- `POST v1/libraries/{libraryId}/books/{bookId}`: Add book to a library.
- `DELETE v1/libraries/{libraryId}/books/{bookId}`: Remove book from a library.

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

## Roadmap

### Authentication

- [X] Should be able to authenticate using credentials;
- [X] Should be able to logout;
- [X] Should be able to reset password;
- [ ] Should be able to refresh token;
- [ ] Should be able to authenticate using Google account;

### Library ✅

- [X] Should be able to create a new library;
- [X] Should be able to get all libraries;
- [X] Should be able to get a library by id;
- [X] Should be able to update a library;
- [X] Should be able to delete a library;
- [X] Should be able to add book in a library;
- [X] Should be able to remove book from a library;

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

## Installation

To use this project, you need to follow these steps:

1. Clone the repository: `git clone https://github.com/vinniciusgomes/mybooks-api`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./cmd/api/main.go`

## Running local with Air
To run the service locally, you can use [Air](https://github.com/cosmtrek/air) for hot-reloading. Run the following command:
```
air
```

## Makefile Commands

The project includes a Makefile to help you manage common tasks more easily. Here's a list of the available commands and a brief description of what they do:

- `make run`: Run the application without generating API documentation.
- `make run-with-docs`: Generate the API documentation using Swag, then run the application.
- `make build`: Build the application and create an executable file named `gopportunities`.
- `make test`: Run tests for all packages in the project.
- `make docs`: Generate the API documentation using Swag.
- `make clean`: Remove the `gopportunities` executable and delete the `./docs` directory.

To use these commands, simply type `make` followed by the desired command in your terminal. For example:

```sh
make run
```

## Docker and Docker Compose

This project includes a `Dockerfile` and `docker-compose.yml` file for easy containerization and deployment. Here are the most common Docker and Docker Compose commands you may want to use:

- `docker build -t your-image-name .`: Build a Docker image for the project. Replace `your-image-name` with a name for your image.
- `docker run -p 8080:8080 -e PORT=8080 your-image-name`: Run a container based on the built image. Replace `your-image-name` with the name you used when building the image. You can change the port number if necessary.

If you want to use Docker Compose, follow these commands:

- `docker compose build`: Build the services defined in the `docker-compose.yml` file.
- `docker compose up`: Run the services defined in the `docker-compose.yml` file.

To stop and remove containers, networks, and volumes defined in the `docker-compose.yml` file, run:

```sh
docker-compose down
```

For more information on Docker and Docker Compose, refer to the official documentation:

- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Used Tools

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [GoORM](https://gorm.io/) for database communication
- [Swagger](https://swagger.io/) for API documentation and testing

## Contributing

To contribute to this project, please follow these guidelines:

1. Fork the repository
2. Create a new branch: `git checkout -b feature/your-feature-name`
3. Make your changes and commit them using Conventional Commits
4. Push to the branch: `git push origin feature/your-feature-name`
5. Submit a pull request

---

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.

## Credits

This project was created by [vinniciusgomes](https://github.com/vinniciusgomes).
