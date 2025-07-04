# 🌟 Domain-Driven Golang Task Manager

Welcome to the **Domain-Driven Golang Task Manager**! 🚀 This is a robust, modern Go application built with **Domain-Driven Design (DDD)** principles to manage tasks and categories efficiently. Powered by **Gin**, **GORM**, and **PostgreSQL**, it offers a clean architecture, a feature-rich RESTful API, and cross-platform development support. Whether you're exploring DDD or building a scalable task management system, this project is your perfect starting point! 🎉

---

## ✨ Features

- 🗂 **DDD Architecture**: Organized into `domain`, `application`, `infrastructure`, `handlers`, `routes`, `dto`, and `common` packages for maintainability and scalability.
- 📡 **RESTful API**: Create, read, update (via PATCH), and delete tasks and categories with standardized JSON responses.
- 🔍 **Advanced Querying**: Supports **paging**, **search** by title/description for tasks and name/description for categories, **sorting** by multiple fields, and **filtering** by task status.
- 🎨 **Category Colors**: Categories have a randomly assigned color on creation, updatable via a predefined palette of 10 distinct colors.
- 🔗 **Task-Category Relationship**: Each task can optionally belong to one category, with category data included in task responses by default.
- 🗄 **PostgreSQL Integration**: Persistent storage with GORM for seamless database operations.
- ⚙️ **Cross-Platform Development**: `Makefile` supports Windows, macOS, and Linux with automated PostgreSQL container management.
- 🔒 **Environment Configuration**: Securely manage database credentials using `.env` and `.env.example`.
- 🛠 **Extensible**: Ready for adding authentication, validation, or additional features.

---

## 🧠 What is Domain-Driven Design (DDD)?

### Overview
Domain-Driven Design (DDD) is a software design approach introduced by Eric Evans in his book *"Domain-Driven Design: Tackling Complexity in the Heart of Software"*. It focuses on aligning software development with the **business domain** by creating models that reflect real-world processes and concepts. DDD emphasizes collaboration between developers and domain experts to ensure the software solves business problems effectively. 🧩

### Core Concepts
- **Domain**: The problem space or business area the software addresses (e.g., task and category management).
- **Entities**: Objects with unique identities (e.g., `Task`, `Category`).
- **Value Objects**: Immutable objects without identity (e.g., a task's `Status`).
- **Aggregates**: Clusters of related objects treated as a single unit (e.g., a `Task` with its `Category`).
- **Repositories**: Interfaces for data access, abstracting persistence (e.g., `TaskRepository`, `CategoryRepository`).
- **Services**: Business logic that doesn’t belong to entities or aggregates (e.g., `TaskService`, `CategoryService`).
- **Bounded Contexts**: Clear boundaries around specific domains to avoid ambiguity.

### Why Use DDD in This Project?
This project adopts DDD to create a scalable, maintainable, and business-focused task management system. Here’s why DDD is a great fit:
- **Clarity**: The `domain` package defines core business entities (`Task`, `Category`) and rules, making the codebase intuitive for developers and stakeholders. 📋
- **Separation of Concerns**: Layers like `application`, `infrastructure`, and `handlers` isolate business logic from technical details, improving modularity. 🛠
- **Scalability**: DDD’s structure supports adding new features (e.g., user authentication) without disrupting existing code. 📈
- **Maintainability**: Clear boundaries and abstractions make it easier to refactor or replace components (e.g., switching from PostgreSQL to another database). 🔧

### Visualization of DDD Architecture
Below is a diagram illustrating the DDD layered architecture, showing how requests flow through the layers to interact with the domain and persistence.

```mermaid
graph TD
    A[Client] -->|HTTP Requests| B[Handlers]
    B -->|Calls| C[Application]
    C -->|Uses| D[Domain]
    C -->|Accesses| E[Infrastructure]
    E -->|Persists| F[PostgreSQL]
    D -->|Defines| G[Entities & Rules]
    style A fill:#f9f,stroke:#333,stroke-width:2px
    style F fill:#bbf,stroke:#333,stroke-width:2px
```

- **Client**: Sends HTTP requests (e.g., via `curl` or a frontend).
- **Handlers**: Processes requests using Gin, validates input with `dto`.
- **Application**: Orchestrates business logic via `TaskService` and `CategoryService`.
- **Domain**: Defines core entities (`Task`, `Category`) and business rules.
- **Infrastructure**: Manages persistence with GORM and PostgreSQL.
- **PostgreSQL**: Stores task and category data.

---

## 🛠 Tech Stack

- **Go** (1.24.3) 🐹
- **Gin** (HTTP framework) 🌐
- **GORM** (ORM for PostgreSQL) 🗃
- **PostgreSQL** (16-alpine via Docker) 🐘
- **godotenv** (Environment variable management) 🔑
- **Makefile** (Build automation) ⚒

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
   - The `.env` file should contain your PostgreSQL database configuration. Update it with your credentials:
     ```env
     DB_HOST=localhost
     DB_USER=postgres
     DB_PASSWORD=12345
     DB_NAME=tasks_db
     DB_PORT=5432
     ```
   - Ensure the PostgreSQL database is running and accessible:
    ```bash
    make postgres
    ```

4. **Run the Application** 🚀
   - Use the `Makefile` to start PostgreSQL and the Go app:
     ```bash
     make run
     ```
   - This command runs the Go application on `http://localhost:8080`.

---

## 📚 API Endpoints

The API follows RESTful conventions with normalized JSON responses. Below are the available endpoints for both **tasks** and **categories**.

### Task Endpoints
| Method | Endpoint          | Description                     | Query Parameters / Payload                            |
|--------|-------------------|---------------------------------|------------------------------------------------------|
| `POST` | `/tasks`          | Create a new task               | `{"title":"Task 1","description":"Details","due_at":"2025-06-15T12:00:00Z","category_id":1}` |
| `GET`  | `/tasks/:id`      | Get a task by ID (includes category) | -                                                |
| `GET`  | `/tasks`          | Get tasks with pagination, search, sort, and filter (includes category) | `page`, `page_size`, `search`, `sort_by`, `sort_order`, `status` |
| `PATCH`| `/tasks/:id`      | Partially update a task         | `{"status":"Done","category_id":2}`                  |
| `DELETE` | `/tasks/:id`    | Delete a task                   | -                                                    |

#### GET /tasks Query Parameters
| Parameter    | Description                              | Example Values                     | Default          |
|--------------|------------------------------------------|------------------------------------|------------------|
| `page`       | Page number                              | `1`, `2`                          | `1`              |
| `page_size`  | Tasks per page                           | `10`, `20`                        | `10`             |
| `search`     | Keyword to search in title/description   | `groceries`                       | -                |
| `sort_by`    | Field to sort by                         | `title`, `due_at`, `created_at`   | `created_at`     |
| `sort_order` | Sort direction                           | `asc`, `desc`                     | `desc`           |
| `status`     | Filter by task status                    | `Pending`, `Doing`, `Done`        | -                |

#### Example Task Requests
- **Create Task**:
  ```bash
  curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title":"Buy groceries","description":"Milk, eggs, bread","due_at":"2025-06-15T12:00:00Z","category_id":1}'
  ```
  **Response**:
  ```json
  {
      "success": true,
      "data": {
          "ID": 1,
          "Title": "Buy groceries",
          "Description": "Milk, eggs, bread",
          "Status": "Pending",
          "CreatedAt": "2025-06-29T09:57:00Z",
          "UpdatedAt": "2025-06-29T09:57:00Z",
          "DueAt": "2025-06-15T12:00:00Z",
          "CategoryID": 1,
          "Category": {
              "ID": 1,
              "Name": "Work",
              "Description": "Work-related tasks",
              "Color": "#e6194b",
              "CreatedAt": "2025-06-29T09:56:00Z"
          }
      }
  }
  ```
- **Get Task by ID**:
  ```bash
  curl http://localhost:8080/tasks/1
  ```
  **Response**: Same as above.
- **Get Tasks with Query**:
  ```bash
  curl "http://localhost:8080/tasks?page=1&page_size=10&search=groceries&sort_by=title&sort_order=asc&status=Pending"
  ```
  **Response**:
  ```json
  {
      "success": true,
      "data": {
          "tasks": [
              {
                  "ID": 1,
                  "Title": "Buy groceries",
                  "Description": "Milk, eggs, bread",
                  "Status": "Pending",
                  "CreatedAt": "2025-06-29T09:57:00Z",
                  "UpdatedAt": "2025-06-29T09:57:00Z",
                  "DueAt": "2025-06-15T12:00:00Z",
                  "CategoryID": 1,
                  "Category": {
                      "ID": 1,
                      "Name": "Work",
                      "Description": "Work-related tasks",
                      "Color": "#e6194b",
                      "CreatedAt": "2025-06-29T09:56:00Z"
                  }
              }
          ],
          "meta": {
              "total": 1,
              "page": 1,
              "page_size": 10,
              "total_pages": 1
          }
      }
  }
  ```
- **Update Task (PATCH)**:
  ```bash
  curl -X PATCH http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"status":"Done","category_id":2}'
  ```
- **Delete Task**:
  ```bash
  curl -X DELETE http://localhost:8080/tasks/1
  ```

### Category Endpoints
| Method | Endpoint              | Description                     | Query Parameters / Payload                            |
|--------|-----------------------|---------------------------------|------------------------------------------------------|
| `POST` | `/categories`         | Create a new category           | `{"name":"Work","description":"Work tasks"}`         |
| `GET`  | `/categories/:id`     | Get a category by ID            | -                                                    |
| `GET`  | `/categories`         | Get categories with pagination, search, and sort | `page`, `page_size`, `search`, `sort_by`, `sort_order` |
| `PATCH`| `/categories/:id`     | Partially update a category     | `{"name":"Personal","color":"#3cb44b"}`             |
| `DELETE` | `/categories/:id`   | Delete a category               | -                                                    |

#### GET /categories Query Parameters
| Parameter    | Description                              | Example Values                     | Default          |
|--------------|------------------------------------------|------------------------------------|------------------|
| `page`       | Page number                              | `1`, `2`                          | `1`              |
| `page_size`  | Categories per page                      | `10`, `20`                        | `10`             |
| `search`     | Keyword to search in name/description    | `work`                            | -                |
| `sort_by`    | Field to sort by                         | `name`, `created_at`              | `created_at`     |
| `sort_order` | Sort direction                           | `asc`, `desc`                     | `desc`           |

#### Example Category Requests
- **Create Category** (color randomly assigned):
  ```bash
  curl -X POST http://localhost:8080/categories -H "Content-Type: application/json" -d '{"name":"Work","description":"Work-related tasks"}'
  ```
  **Response**:
  ```json
  {
      "success": true,
      "data": {
          "ID": 1,
          "Name": "Work",
          "Description": "Work-related tasks",
          "Color": "#e6194b",
          "CreatedAt": "2025-06-29T09:56:00Z"
      }
  }
  ```
- **Get Category by ID**:
  ```bash
  curl http://localhost:8080/categories/1
  ```
  **Response**: Same as above.
- **Get Categories with Query**:
  ```bash
  curl "http://localhost:8080/categories?page=1&page_size=10&search=work&sort_by=name&sort_order=asc"
  ```
  **Response**:
  ```json
  {
      "success": true,
      "data": {
          "categories": [
              {
                  "ID": 1,
                  "Name": "Work",
                  "Description": "Work-related tasks",
                  "Color": "#e6194b",
                  "CreatedAt": "2025-06-29T09:56:00Z"
              }
          ],
          "meta": {
              "total": 1,
              "page": 1,
              "page_size": 10,
              "total_pages": 1
          }
      }
  }
  ```
- **Update Category (PATCH)**:
  ```bash
  curl -X PATCH http://localhost:8080/categories/1 -H "Content-Type: application/json" -d '{"name":"Personal","color":"#3cb44b"}'
  ```
  **Response**:
  ```json
  {
      "success": true,
      "data": {
          "ID": 1,
          "Name": "Personal",
          "Description": "Work-related tasks",
          "Color": "#3cb44b",
          "CreatedAt": "2025-06-29T09:56:00Z"
      }
  }
  ```
- **Delete Category**:
  ```bash
  curl -X DELETE http://localhost:8080/categories/1
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
| `make stop-postgres` | Stop and remove the PostgreSQL container    |
| `make clean`     | Stop PostgreSQL and remove build artifacts      |
| `make fmt`       | Format Go code                                  |

---

## 📂 Project Structure

The project follows DDD principles with a modular structure:

```
domain-driven-golang/
├── .env              # Environment variables (ignored by Git)
├── .env.example      # Template for .env
├── .gitignore        # Git ignore rules
├── Makefile          # Build automation
├── README.md         # Project documentation
├── go.mod            # Go module dependencies
├── common/           # Common response structures
├── config/           # Database configuration
├── utils/            # Utility functions
├── modules/
│   ├── task/
│   │   ├── domain/
│   │   │   └── task.go
│   │   ├── application/
│   │   │   └── task_service.go
│   │   ├── infrastructure/
│   │   │   └── task_repository.go
│   │   ├── handlers/
│   │   │   └── task_handler.go
│   │   ├── routes/
│   │   │   └── routes.go
│   │   └── dto/
│   │       └── task.go
│   └── category/
│       ├── domain/
│       │   └── category.go
│       ├── application/
│       │   └── category_service.go
│       ├── infrastructure/
│       │   └── category_repository.go
│       ├── handlers/
│       │   └── category_handler.go
│       ├── routes/
│       │   └── routes.go
│       └── dto/
│           └── category.go
└── main.go           # Application entry point
```

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

Built with ❤️ by [ltphat2204](https://github.com/ltphat2204). Happy coding! 🚀# 🌟 Domain-Driven Golang Task Manager

Welcome to the **Domain-Driven Golang Task Manager**! 🚀 This is a robust, modern Go application built with **Domain-Driven Design (DDD)** principles to manage tasks and categories efficiently. Powered by **Gin**, **GORM**, and **PostgreSQL**, it offers a clean architecture, a feature-rich RESTful API, and cross-platform development support. Whether you're exploring DDD or building a scalable task management system, this project is your perfect starting point! 🎉

---

## ✨ Features

- 🗂 **DDD Architecture**: Organized into `domain`, `application`, `infrastructure`, `handlers`, `routes`, `dto`, and `common` packages for maintainability and scalability.
- 📡 **RESTful API**: Create, read, update (via PATCH), and delete tasks and categories with standardized JSON responses.
- 🔍 **Advanced Querying**: Supports **paging**, **search** by title/description for tasks and name/description for categories, **sorting** by multiple fields, and **filtering** by task status.
- 🎨 **Category Colors**: Categories have a randomly assigned color on creation, updatable via a predefined palette of 10 distinct colors.
- 🔗 **Task-Category Relationship**: Each task can optionally belong to one category, with category data included in task responses by default.
- 🗄 **PostgreSQL Integration**: Persistent storage with GORM for seamless database operations.
- ⚙️ **Cross-Platform Development**: `Makefile` supports Windows, macOS, and Linux with automated PostgreSQL container management.
- 🔒 **Environment Configuration**: Securely manage database credentials using `.env` and `.env.example`.
- 🛠 **Extensible**: Ready for adding authentication, validation, or additional features.

---

## 🧠 What is Domain-Driven Design (DDD)?

### Overview
Domain-Driven Design (DDD) is a software design approach introduced by Eric Evans in his book *"Domain-Driven Design: Tackling Complexity in the Heart of Software"*. It focuses on aligning software development with the **business domain** by creating models that reflect real-world processes and concepts. DDD emphasizes collaboration between developers and domain experts to ensure the software solves business problems effectively. 🧩

### Core Concepts
- **Domain**: The problem space or business area the software addresses (e.g., task and category management).
- **Entities**: Objects with unique identities (e.g., `Task`, `Category`).
- **Value Objects**: Immutable objects without identity (e.g., a task's `Status`).
- **Aggregates**: Clusters of related objects treated as a single unit (e.g., a `Task` with its `Category`).
- **Repositories**: Interfaces for data access, abstracting persistence (e.g., `TaskRepository`, `CategoryRepository`).
- **Services**: Business logic that doesn’t belong to entities or aggregates (e.g., `TaskService`, `CategoryService`).
- **Bounded Contexts**: Clear boundaries around specific domains to avoid ambiguity.

### Why Use DDD in This Project?
This project adopts DDD to create a scalable, maintainable, and business-focused task management system. Here’s why DDD is a great fit:
- **Clarity**: The `domain` package defines core business entities (`Task`, `Category`) and rules, making the codebase intuitive for developers and stakeholders. 📋
- **Separation of Concerns**: Layers like `application`, `infrastructure`, and `handlers` isolate business logic from technical details, improving modularity. 🛠
- **Scalability**: DDD’s structure supports adding new features (e.g., user authentication) without disrupting existing code. 📈
- **Maintainability**: Clear boundaries and abstractions make it easier to refactor or replace components (e.g., switching from PostgreSQL to another database). 🔧

### Visualization of DDD Architecture
Below is a diagram illustrating the DDD layered architecture, showing how requests flow through the layers to interact with the domain and persistence.

```mermaid
graph TD
    A[Client] -->|HTTP Requests| B[Handlers]
    B -->|Calls| C[Application]
    C -->|Uses| D[Domain]
    C -->|Accesses| E[Infrastructure]
    E -->|Persists| F[PostgreSQL]
    D -->|Defines| G[Entities & Rules]
    style A fill:#f9f,stroke:#333,stroke-width:2px
    style F fill:#bbf,stroke:#333,stroke-width:2px
```

- **Client**: Sends HTTP requests (e.g., via `curl` or a frontend).
- **Handlers**: Processes requests using Gin, validates input with `dto`.
- **Application**: Orchestrates business logic via `TaskService` and `CategoryService`.
- **Domain**: Defines core entities (`Task`, `Category`) and business rules.
- **Infrastructure**: Manages persistence with GORM and PostgreSQL.
- **PostgreSQL**: Stores task and category data.

---

## 🛠 Tech Stack

- **Go** (1.24.3) 🐹
- **Gin** (HTTP framework) 🌐
- **GORM** (ORM for PostgreSQL) 🗃
- **PostgreSQL** (16-alpine via Docker) 🐘
- **godotenv** (Environment variable management) 🔑
- **Makefile** (Build automation) ⚒

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
   - The `.env` file should contain your PostgreSQL database configuration. Update it with your credentials:
     ```env
     DB_HOST=localhost
     DB_USER=postgres
     DB_PASSWORD=12345
     DB_NAME=tasks_db
     DB_PORT=5432
     ```
   - Ensure the PostgreSQL database is running and accessible:
    ```bash
    make postgres
    ```

4. **Run the Application** 🚀
   - Use the `Makefile` to start PostgreSQL and the Go app:
     ```bash
     make run
     ```
   - This command runs the Go application on `http://localhost:8080`.

---

## 📚 API Endpoints

The API follows RESTful conventions with normalized JSON responses. Below are the available endpoints for both **tasks** and **categories**.

### Task Endpoints
| Method | Endpoint          | Description                     | Query Parameters / Payload                            |
|--------|-------------------|---------------------------------|------------------------------------------------------|
| `POST` | `/tasks`          | Create a new task               | `{"title":"Task 1","description":"Details","due_at":"2025-06-15T12:00:00Z","category_id":1}` |
| `GET`  | `/tasks/:id`      | Get a task by ID (includes category) | -                                                |
| `GET`  | `/tasks`          | Get tasks with pagination, search, sort, and filter (includes category) | `page`, `page_size`, `search`, `sort_by`, `sort_order`, `status` |
| `PATCH`| `/tasks/:id`      | Partially update a task         | `{"status":"Done","category_id":2}`                  |
| `DELETE` | `/tasks/:id`    | Delete a task                   | -                                                    |

#### GET /tasks Query Parameters
| Parameter    | Description                              | Example Values                     | Default          |
|--------------|------------------------------------------|------------------------------------|------------------|
| `page`       | Page number                              | `1`, `2`                          | `1`              |
| `page_size`  | Tasks per page                           | `10`, `20`                        | `10`             |
| `search`     | Keyword to search in title/description   | `groceries`                       | -                |
| `sort_by`    | Field to sort by                         | `title`, `due_at`, `created_at`   | `created_at`     |
| `sort_order` | Sort direction                           | `asc`, `desc`                     | `desc`           |
| `status`     | Filter by task status                    | `Pending`, `Doing`, `Done`        | -                |

#### Example Task Requests
- **Create Task**:
  ```bash
  curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title":"Buy groceries","description":"Milk, eggs, bread","due_at":"2025-06-15T12:00:00Z","category_id":1}'
  ```
  **Response**:
  ```json
  {
      "success": true,
      "data": {
          "ID": 1,
          "Title": "Buy groceries",
          "Description": "Milk, eggs, bread",
          "Status": "Pending",
          "CreatedAt": "2025-06-29T09:57:00Z",
          "UpdatedAt": "2025-06-29T09:57:00Z",
          "DueAt": "2025-06-15T12:00:00Z",
          "CategoryID": 1,
          "Category": {
              "ID": 1,
              "Name": "Work",
              "Description": "Work-related tasks",
              "Color": "#e6194b",
              "CreatedAt": "2025-06-29T09:56:00Z"
          }
      }
  }
  ```
- **Get Task by ID**:
  ```bash
  curl http://localhost:8080/tasks/1
  ```
  **Response**: Same as above.
- **Get Tasks with Query**:
  ```bash
  curl "http://localhost:8080/tasks?page=1&page_size=10&search=groceries&sort_by=title&sort_order=asc&status=Pending"
  ```
  **Response**:
  ```json
  {
      "success": true,
      "data": {
          "tasks": [
              {
                  "ID": 1,
                  "Title": "Buy groceries",
                  "Description": "Milk, eggs, bread",
                  "Status": "Pending",
                  "CreatedAt": "2025-06-29T09:57:00Z",
                  "UpdatedAt": "2025-06-29T09:57:00Z",
                  "DueAt": "2025-06-15T12:00:00Z",
                  "CategoryID": 1,
                  "Category": {
                      "ID": 1,
                      "Name": "Work",
                      "Description": "Work-related tasks",
                      "Color": "#e6194b",
                      "CreatedAt": "2025-06-29T09:56:00Z"
                  }
              }
          ],
          "meta": {
              "total": 1,
              "page": 1,
              "page_size": 10,
              "total_pages": 1
          }
      }
  }
  ```
- **Update Task (PATCH)**:
  ```bash
  curl -X PATCH http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"status":"Done","category_id":2}'
  ```
- **Delete Task**:
  ```bash
  curl -X DELETE http://localhost:8080/tasks/1
  ```

### Category Endpoints
| Method | Endpoint              | Description                     | Query Parameters / Payload                            |
|--------|-----------------------|---------------------------------|------------------------------------------------------|
| `POST` | `/categories`         | Create a new category           | `{"name":"Work","description":"Work tasks"}`         |
| `GET`  | `/categories/:id`     | Get a category by ID            | -                                                    |
| `GET`  | `/categories`         | Get categories with pagination, search, and sort | `page`, `page_size`, `search`, `sort_by`, `sort_order` |
| `PATCH`| `/categories/:id`     | Partially update a category     | `{"name":"Personal","color":"#3cb44b"}`             |
| `DELETE` | `/categories/:id`   | Delete a category               | -                                                    |

#### GET /categories Query Parameters
| Parameter    | Description                              | Example Values                     | Default          |
|--------------|------------------------------------------|------------------------------------|------------------|
| `page`       | Page number                              | `1`, `2`                          | `1`              |
| `page_size`  | Categories per page                      | `10`, `20`                        | `10`             |
| `search`     | Keyword to search in name/description    | `work`                            | -                |
| `sort_by`    | Field to sort by                         | `name`, `created_at`              | `created_at`     |
| `sort_order` | Sort direction                           | `asc`, `desc`                     | `desc`           |

#### Example Category Requests
- **Create Category** (color randomly assigned):
  ```bash
  curl -X POST http://localhost:8080/categories -H "Content-Type: application/json" -d '{"name":"Work","description":"Work-related tasks"}'
  ```
  **Response**:
  ```json
  {
      "success": true,
      "data": {
          "ID": 1,
          "Name": "Work",
          "Description": "Work-related tasks",
          "Color": "#e6194b",
          "CreatedAt": "2025-06-29T09:56:00Z"
      }
  }
  ```
- **Get Category by ID**:
  ```bash
  curl http://localhost:8080/categories/1
  ```
  **Response**: Same as above.
- **Get Categories with Query**:
  ```bash
  curl "http://localhost:8080/categories?page=1&page_size=10&search=work&sort_by=name&sort_order=asc"
  ```
  **Response**:
  ```json
  {
      "success": true,
      "data": {
          "categories": [
              {
                  "ID": 1,
                  "Name": "Work",
                  "Description": "Work-related tasks",
                  "Color": "#e6194b",
                  "CreatedAt": "2025-06-29T09:56:00Z"
              }
          ],
          "meta": {
              "total": 1,
              "page": 1,
              "page_size": 10,
              "total_pages": 1
          }
      }
  }
  ```
- **Update Category (PATCH)**:
  ```bash
  curl -X PATCH http://localhost:8080/categories/1 -H "Content-Type: application/json" -d '{"name":"Personal","color":"#3cb44b"}'
  ```
  **Response**:
  ```json
  {
      "success": true,
      "data": {
          "ID": 1,
          "Name": "Personal",
          "Description": "Work-related tasks",
          "Color": "#3cb44b",
          "CreatedAt": "2025-06-29T09:56:00Z"
      }
  }
  ```
- **Delete Category**:
  ```bash
  curl -X DELETE http://localhost:8080/categories/1
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
| `make stop-postgres` | Stop and remove the PostgreSQL container    |
| `make clean`     | Stop PostgreSQL and remove build artifacts      |
| `make fmt`       | Format Go code                                  |

---

## 📂 Project Structure

The project follows DDD principles with a modular structure:

```
domain-driven-golang/
├── .env              # Environment variables (ignored by Git)
├── .env.example      # Template for .env
├── .gitignore        # Git ignore rules
├── Makefile          # Build automation
├── README.md         # Project documentation
├── go.mod            # Go module dependencies
├── common/           # Common response structures
├── config/           # Database configuration
├── utils/            # Utility functions
├── modules/
│   ├── task/
│   │   ├── domain/
│   │   │   └── task.go
│   │   ├── application/
│   │   │   └── task_service.go
│   │   ├── infrastructure/
│   │   │   └── task_repository.go
│   │   ├── handlers/
│   │   │   └── task_handler.go
│   │   ├── routes/
│   │   │   └── routes.go
│   │   └── dto/
│   │       └── task.go
│   └── category/
│       ├── domain/
│       │   └── category.go
│       ├── application/
│       │   └── category_service.go
│       ├── infrastructure/
│       │   └── category_repository.go
│       ├── handlers/
│       │   └── category_handler.go
│       ├── routes/
│       │   └── routes.go
│       └── dto/
│           └── category.go
└── main.go           # Application entry point
```

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