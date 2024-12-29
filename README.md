# gStructify 🚀

A Golang package designed to generate high-performance backend services using fasthttp. It follows the Clean Architecture pattern, automatically generating the domain, infrastructure, handlers, and routing layers for each service, streamlining the development process.

## Features
- Automatically generates backend code structure for Go projects.
- Creates service folders with pre-configured Go modules.
- Simplifies repetitive tasks in setting up backend service.

## Project Structure

```
project/
├── sql-migrations/                          # Migration scripts for PostgreSQL
│   ├── seed/     
│   │   └── init-seed.sql   
│   ├── sql/     
│   │   └── init.sql
├── main.go                                  # Application entry point
├── go.mod
├── src/                                     # Source code folder
│   ├── bootstrap/                           # Central module initialization
│   │   └── bootstrap.go               
│   ├── common/                              # Common utility functions
│   │   ├── constants.go                     # UUID generator
│   ├── core/                                # Contains all the layers
│   │   ├── domain/
│   │   │   └── aggregates/
│   │   │       ├── user.go                  # Business logic for User (Domain model)
│   │   ├── application/
│   │   │   └── service/
│   │   │       ├── user_service.go          # Business logic for User
│   │   │   └── workers/
│   │   │       ├── worker.go                # Background worker
│   │   ├── infrastructure/
│   │   │   ├── db/
│   │   │   │   ├── db.go                    # General Database connection
│   │   │   │   └── sql-db.go                # SQL-DB connection logic
│   │   │   ├── entity/
│   │   │   │   ├── user.go                  # User entity
│   │   │   ├── repository/
│   │   │   │   ├── user_repository_impl.go  # UserRepository implementation
│   │   ├── interface/
│   │   │   └── dto/
│   │   │       ├── user_dto.go              # Request and Response DTO for User
│   │   │   └── handlers/
│   │   │       ├── user_handler.go          # HTTP handler for User
│   ├── helpers/                             # Additional utility functions 
│   │   ├── string_helpers.go                # String manipulation helpers

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

## Step 1: Install gStructify

Run the following command to install gStructify:

```bash
go install github.com/nanda03dev/gStructify@latest
```

## Step 2: Create Your Project Directory

Create and navigate to your project directory:

```bash
mkdir my-microservice
cd my-microservice
```

## Step 3: Initialize Your Go Module

Initialize a new Go module for your project:

```bash
go mod init my-microservice
```

## Step 4: Generate Service Code

Run `gStructify` to automatically create your service structure:

```bash
gStructify -entity=user
```

By default, gStructify will generate the entire structure for one entity with only an `id` field. For more advanced setups, you can define multiple entities and their fields in a configuration file, as explained in Step 5.

## Step 5: Configure Your Entities

To customize the generated code, create a `gStructify.config.json` file in your project directory. Define your entities and their fields as shown below:

```json
{
    "entities": [
        {
            "entity_name": "user",
            "fields": [
                { "field_name": "id", "type": "string" },
                { "field_name": "email", "type": "string" }
            ]
        },
        {
            "entity_name": "order",
            "fields": [
                { "field_name": "id", "type": "string" },
                { "field_name": "user_id", "type": "string" },
                { "field_name": "order_amount", "type": "int" }
            ]
        }
    ]
}

```

### Notes:
1. The fields in the configuration file should be defined using `snake_case`.
2. Both keys and values in the configuration file are case-sensitive.
3. You can directly use Go data types (e.g., `string`, `int`, `float64`, etc.) in the field definitions.

Once you’ve added fields in the configuration file, Run gStructify. 

Run `gStructify` to automatically create your service structure:

```bash
gStructify
```

it will generate files with all the specified fields included and create CRUD APIs for those entities.

This will:
1. Generate a structured set of files for the `user, order` entities.
2. Populate the necessary Go files within the `user-ms` directory.

## Step 6: Run the Application

If you have a URL for the SQL database, you can add it in the `.env` file under the `SQL_DB_URI` variable. Alternatively, if you want to run the database locally using Docker, follow these steps:

1. Start the PostgreSQL database:

   ```bash
   make run-sql-db
   ```

   This command will start the PostgreSQL database in detached mode.

2. Run the application locally:

   ```bash
   make run-dev
   ```

   This command will start the application on your local machine.
   
Once the application is running, you can start executing the generated APIs for your entities.

Now you're ready to build efficient and scalable Go backend services with gStructify. 🚀

---

Happy coding with `gStructify`!

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Clean Architecture by Robert C. Martin](https://www.oreilly.com/library/view/clean-architecture-a/9780134494166/)
- [Golang Documentation](https://golang.org/doc/)