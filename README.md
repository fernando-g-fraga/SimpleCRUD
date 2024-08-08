# simpleCRUD with Go and Postgres

## Overview

`simpleCRUD` is a simple CRUD application built with Go and PostgreSQL. This project demonstrates basic operations for creating, reading, updating, and deleting records in a PostgreSQL database using Go.

## Features

- **Create**: Add new records to the database.
- **Read**: Retrieve and display records from the database.
- ~~**Update**: Modify existing records in the database.~~ - UPCOMING!
- **Delete**: Remove records from the database.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- [Go](https://golang.org/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)

## Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/simpleCRUD.git
   cd simpleCRUD
   ```

2. **Initialize Go modules:**

   ```bash
   go mod init simpleCRUD
   go mod tidy
   ```

3. **Create the database and table:**

   - Start PostgreSQL and create a new database.

   ```sql
   CREATE DATABASE simpleCRUD;
   ```

   - Create a table in your database:

   ```sql
   CREATE TABLE contato (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100),
       email VARCHAR(100)
   );
   ```

4. **Configure your database connection:**

   Update the database connection details in `main.go`:

   ```go
   const (
       host     = "localhost"
       port     = 5432
       user     = "your_user"
       password = "your_password"
       dbname   = "simpleCRUD"
   )
   ```

5. **Run the application:**

   ```bash
   go run main.go
   ```

## Usage

Upon running the application, you’ll be presented with a menu in the terminal where you can choose to:

- Add a new contact
- View all contacts
- View details from a contact
- Delete a contact

Follow the on-screen prompts to perform these operations.

## Project Structure

```
simpleCRUD/
├── main.go          # Entry point of the application
├── db/
│   └── db.go        # Database interaction functions
├── models/
│   └── user.go      # Struct definitions for data models
└── go.mod           # Go module file
```

## Dependencies

This project uses the following dependencies:

- `github.com/lib/pq` - PostgreSQL driver for Go

## Contributing

Contributions are welcome! If you have any improvements, suggestions, or bug reports, feel free to open an issue or submit a pull request.
