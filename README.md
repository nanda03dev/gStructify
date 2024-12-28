# gStructify 🚀

A Golang package designed to generate high-performance backend services using fasthttp. It follows the Clean Architecture pattern, automatically generating the domain, infrastructure, handlers, and routing layers for each service, streamlining the development process.

## Features
- Automatically generates backend code structure for Go projects.
- Creates service folders with pre-configured Go modules.
- Simplifies repetitive tasks in setting up backend services.

## Project Structure

```
project/
├── sql-migrations/                      # Migration scripts for PostgreSQL
│   ├── seed/     
│   │   └── init-seed.sql   
│   ├── sql/     
│   │   └── init.sql
├── main.go                              # Application entry point
├── go.mod
├── src/                                 # Source code folder
│   ├── bootstrap/                       # Central module initialization
│   │   └── bootstrap.go               
│   ├── common/                          # Common utility functions
│   │   ├── constants.go                 # UUID generator
│   ├── core/                            # Contains all the layers
│   │   ├── domain/
│   │   │   └── aggregates/
│   │   │       ├── user.go              # Business logic for User (Domain model)
│   │   ├── application/
│   │   │   └── service/
│   │   │       ├── user_service.go      # Business logic for User
│   │   │   └── workers/
│   │   │       ├── worker.go            # Background worker
│   │   ├── infrastructure/
│   │   │   ├── db/
│   │   │   │   ├── db.go                # General Database connection
│   │   │   │   └── sql-db.go            # SQL-DB connection logic
│   │   │   ├── entity/
│   │   │   │   ├── user.go               # MongoDB User entity
│   │   │   ├── repository/
│   │   │   │   ├── user_repository_impl.go  # MongoDB UserRepository implementation
│   │   ├── interface/
│   │   │   └── dto/
│   │   │       ├── user_dto.go          # Request and Response DTO for User
│   │   │   └── handlers/
│   │   │       ├── user_handler.go      # HTTP handler for User
│   ├── helpers/                         # Additional utility functions 
│   │   ├── string_helpers.go            # String manipulation helpers

```

## Layers Explained

- **Domain Layer**: Contains the core business models and domain logic. It defines entities and their relationships, encapsulating the business rules.

- **Application Layer**: Holds the business logic and use cases. It coordinates the flow of data between the domain and infrastructure layers, processing requests and responses.

- **Infrastructure Layer**: Manages data access and external integrations (e.g., database interactions, third-party APIs). It provides concrete implementations for the repository interfaces defined in the domain layer.

- **Interface Layer**: Contains HTTP handlers for routing incoming requests and managing responses. It may also include Data Transfer Objects (DTOs) for structured data interchange between client and server.

## Getting Started

### Prerequisites

- Go 1.XX or higher
- Dependencies specified in `go.mod`

## Installation

To install `gStructify`, run the following command:

```bash
go install github.com/nanda03dev/gStructify@latest
```

## Usage

Follow these steps to generate entries for your Go microservice:

### Step 1: Create a Microservice Directory
Create a directory for your microservice:
```bash
mkdir user-ms
```

### Step 2: Initialize a Go Module
Navigate to the directory and initialize a `go.mod` file:
```bash
cd user-ms
go mod init user-ms
```

### Step 3: Run the Generator Command
Run the `gStructify` command with your desired entity name:
```bash
gStructify -entity=user

create file with name `gStructify.config.json` and entity details refer sample config file below
note "key and values should be snake-case senstive"
gStructify 
```

This will:
1. Generate a structured set of files for the `user, order` entities.
2. Populate the necessary Go files within the `user-ms` directory.

---

Happy coding with `gStructify`!

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Clean Architecture by Robert C. Martin](https://www.oreilly.com/library/view/clean-architecture-a/9780134494166/)
- [Golang Documentation](https://golang.org/doc/)
```