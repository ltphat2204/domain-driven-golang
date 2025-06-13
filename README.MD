# 🌟 Domain-Driven Golang Task Manager

Welcome to the **Domain-Driven Golang Task Manager**! 🚀 This is a robust, modern Go application built with **Domain-Driven Design (DDD)** principles to manage tasks efficiently. Powered by **Gin**, **GORM**, and **PostgreSQL**, it offers a clean architecture, RESTful API, and cross-platform development support. Whether you're a developer exploring DDD or building a scalable task management system, this project is your perfect starting point! 🎉

---

## ✨ Features

- 🗂 **DDD Architecture**: Organized into `domain`, `application`, `infrastructure`, `handlers`, `routes`, `dto`, and `common` packages for maintainability and scalability.
- 📡 **RESTful API**: Create, read, update (via PATCH), and delete tasks with standardized JSON responses.
- 🗄 **PostgreSQL Integration**: Persistent storage with GORM for seamless database operations.
- ⚙️ **Cross-Platform Development**: `Makefile` supports Windows, macOS, and Linux with automated PostgreSQL container management.
- 🔒 **Environment Configuration**: Securely manage database credentials using `.env` and `.env.example`.
- 📝 **Normalized Responses**: Consistent API responses with `common` package for success and error handling.
- 🛠 **Extensible**: Ready for adding authentication, validation, or additional features.

---

## 🛠 Tech Stack

- **Go** (1.24.3) 🐹
- **Gin** (HTTP framework) 🌐
- **GORM** (ORM for PostgreSQL) 🗃
- **PostgreSQL** (16-alpine via Docker) 🐘
- **godotenv** (Environment variable management) 🔑
- **Makefile** (Build automation) ⚒

---

## 📂 Project Structure

The project follows DDD principles with a modular structure:

```
domain-driven-golang/
├── .env              # Environment variables (ignored by Git)
├── .env.example      # Template for .env
├── .gitignore        # Git ignore rules
├── Makefile          # Build automation
├── go.mod            # Go module dependencies
├── common/           # Common response structures
├── config/           # Database configuration
├── domain/           # Domain entities and interfaces
├── application/      # Business logic
├── infrastructure/   # Database and external services
├── handlers/         # HTTP handlers
├── routes/           # API routes
├── dto/              # Data transfer objects
└── main.go           # Application entry point
```

---

## 📋 Prerequisites

Before you dive in, ensure you have the following installed:

- 🐹 **Go** (1.24.3 or later): [Download Go](https://go.dev/dl/)
- 🐳 **Docker**: [Install Docker Desktop](https://www.docker.com/products/docker-desktop/)
- ⚒ **Make**: 
  - **Windows**: Install via [Chocolatey](https://chocolatey.org/) (`choco install make`) or use Git Bash/WSL.
  - **macOS/Linux**: Usually pre-installed; verify with `make --version`.
- 📝 **Git**: [Install Git](https://git-scm.com/downloads)

---

## 🚀 Getting Started

Follow these steps to set up and run the project locally:

1. **Clone the Repository** 📥
   ```bash
   git clone https://github.com/ltphat2204/domain-driven-golang.git
   cd domain-driven-golang
   ```

2. **Initialize Go Module** 📦
   ```bash
   go mod init github.com/ltphat2204/domain-driven-golang
   go mod tidy
   ```

3. **Set Up Environment Variables** 🔑
   - Copy the example environment file:
     ```bash
     cp .env.example .env
     ```
   - The `.env` file is automatically updated by the `Makefile` when starting PostgreSQL, but you can manually edit it if needed:
     ```env
     DB_HOST=localhost
     DB_USER=postgres
     DB_PASSWORD=12345
     DB_NAME=tasks_db
     DB_PORT=5432
     ```

4. **Run the Application** 🚀
   - Use the `Makefile` to start PostgreSQL and the Go app:
     ```bash
     make run
     ```
   - This command:
     - Starts a PostgreSQL 16-alpine container (if not already running).
     - Updates `.env` with database credentials.
     - Runs the Go application on `http://localhost:8080`.

5. **Verify the API** 🌐
   - Test the `POST /tasks` endpoint:
     ```bash
     curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title":"Buy groceries","description":"Milk, eggs, bread"}'
     ```
   - Expected response:
     ```json
     {
         "success": true,
         "data": {
             "ID": 1,
             "Title": "Buy groceries",
             "Description": "Milk, eggs, bread",
             "Status": "pending",
             "CreatedAt": "2025-06-14T02:21:00Z"
         }
     }
     ```

---

## 📚 API Endpoints

The API follows RESTful conventions with normalized JSON responses. Below are the available endpoints:

| Method | Endpoint          | Description                     | Example Payload                                      |
|--------|-------------------|---------------------------------|-----------------------------------------------------|
| `POST` | `/tasks`          | Create a new task               | `{"title":"Task 1","description":"Details"}`        |
| `GET`  | `/tasks/:id`      | Get a task by ID                | -                                                   |
| `GET`  | `/tasks`          | Get all tasks                   | -                                                   |
| `PATCH`| `/tasks/:id`      | Partially update a task         | `{"status":"completed"}`                            |
| `DELETE` | `/tasks/:id`    | Delete a task                   | -                                                   |

### Example Requests

- **Create Task**:
  ```bash
  curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title":"Buy groceries","description":"Milk, eggs, bread"}'
  ```

- **Update Task (PATCH)**:
  ```bash
  curl -X PATCH http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"status":"completed"}'
  ```

- **Get All Tasks**:
  ```bash
  curl http://localhost:8080/tasks
  ```

---

## 🛠 Makefile Commands

The `Makefile` simplifies development tasks and is cross-platform (Windows, macOS, Linux). Run `make help` to see all commands:

```bash
make help
```

| Command          | Description                                      |
|------------------|--------------------------------------------------|
| `make run`       | Start PostgreSQL and run the Go app             |
| `make build`     | Build the executable (`bin/domain-driven-golang`) |
| `make test`      | Run Go tests                                    |
| `make postgres`  | Start the PostgreSQL container                  |
| `make postgres-stop` | Stop and remove the PostgreSQL container    |
| `make clean`     | Stop PostgreSQL and remove build artifacts      |
| `make fmt`       | Format Go code                                  |

---

## 🐳 PostgreSQL Setup

The project uses a PostgreSQL 16-alpine container managed via Docker. The `Makefile` automates its setup:

- **Container Name**: `dev-postgres`
- **Database**: `tasks_db`
- **User**: `postgres`
- **Password**: `12345`
- **Port**: `5432`

To interact with the database:
```bash
make postgres-shell
```

---

## 🤝 Contributing

Contributions are welcome! 🌟 Follow these steps:

1. Fork the repository 🍴
2. Create a feature branch (`git checkout -b feature/awesome-feature`) 🌿
3. Commit your changes (`git commit -m "Add awesome feature"`) 📝
4. Push to the branch (`git push origin feature/awesome-feature`) 🚀
5. Open a Pull Request 📬

Please ensure your code follows Go conventions and passes `make fmt` and `make test`.

---

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## 🎉 Acknowledgements

- **Gin**: For a fast and flexible HTTP framework.
- **GORM**: For seamless ORM integration.
- **PostgreSQL**: For reliable data storage.
- **Go Community**: For awesome tools and libraries! 🐹

---

Built with ❤️ by [ltphat2204](https://github.com/ltphat2204). Happy coding! 🚀