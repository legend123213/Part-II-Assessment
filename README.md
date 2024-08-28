# Part-II-Assessment
# Library Management System

This project is a Library Management System built using Go (Golang) and MongoDB. It provides functionalities for user registration, user verification, password reset, and borrowing books.

## Table of Contents
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Installation
```bash
git clone https://github.com/legend123213/ass2.git
cd ass2
go mod tidy
```

## Configuration
Set up MongoDB:
- Ensure you have MongoDB installed and running.
- Update the MongoDB URI in the configuration file.

Create a configuration file `config.json` in the `Config` directory with the following structure:
```json
{
   "Database": {
      "Uri": "your_mongodb_uri",
      "Name": "your_database_name"
   },
   "Email": {
      "EmailKey": "your_email_key"
   },
   "Port": ":8000"
}
```

## Usage
...

## API Endpoints
...
The server will start on the port specified in the configuration file.

### User Endpoints
- POST /users: Register a new user.
- GET /users/: Verify user registration.
- POST /users/login: User login.
- GET /users/password-reset: Request password reset.
- PUT /users/password-update: Update password.
- DELETE /users/:id: Delete a user.

### Borrow Endpoints
- POST /books/borrow: Create a borrow request.
- GET /books/borrow: Get a specific borrow request.
- PATCH /books/borrow/:id/status: Update borrow request status.
- DELETE /books/borrow/:id: Delete a borrow request.
- GET /admin/borrows: Get all borrow requests (Admin only).

## Project Structure
....
├── Config
│   └── config.json
├── Delivery
│   └── Controller
│       ├── usercontroller.go
│       └── borrowcontroller.go
├── Domain
│   ├── user.go
│   └── borrow.go
├── Infrastructure
│   └── send_email.go
├── Repositery
│   ├── userrepo.go
│   └── borrowbook.go
├── Usecase
│   ├── userusecase.go
│   └── borrowusecase.go
├── main.go
└── routes.go


this my postman documentation
<https://documenter.getpostman.com/view/26863846/2sAXjJ6t5B>