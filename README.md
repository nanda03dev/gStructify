# gStructify 🚀

A Golang package designed to generate high-performance backend services using fasthttp. It follows the Clean Architecture pattern, automatically generating the domain, infrastructure, handlers, and routing layers for each service, streamlining the development process.


## Project Structure

```
project/
├── sql-migrations/                      # Migration scripts for PostgreSQL
│   └── user_migration.sql
├── main.go                              # Application entry point
├── go.mod
├── src/                                 # Source code folder
│   ├── app_module/                      # Central module initialization
│   │   └── app_module.go                # Initializes repositories and services
│   ├── common/                          # Common utility functions
│   │   ├── constants.go                 # UUID generator
│   ├── core/                            # Contains all the layers
│   │   ├── domain/
│   │   │   └── aggregates/
│   │   │       ├── user.go              # Business logic for User (Domain model)
│   │   ├── application/
│   │   │   └── service/
│   │   │       ├── user_service.go      # Business logic for User
│   │   ├── infrastructure/
│   │   │   ├── db/
│   │   │   │   ├── db.go                # General Database connection
│   │   │   │   └── postgres-db.go       # PostgreSQL connection logic
│   │   │   ├── entity/
│   │   │   │   ├── user.go              # MongoDB User entity
│   │   │   ├── repository/
│   │   │   │   ├── user_repository_impl.go  # MongoDB UserRepository implementation
│   │   ├── interfaces/
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

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd project
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the Application

To start the application, run:
```bash
go run main.go
```

### Testing

To run tests, execute:
```bash
go test ./...
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Clean Architecture by Robert C. Martin](https://www.oreilly.com/library/view/clean-architecture-a/9780134494166/)
- [Golang Documentation](https://golang.org/doc/)
```