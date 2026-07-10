# CLI Todo App in Go

> A modular, repository-driven CLI Todo application built in Go to explore clean architecture, dependency injection, interchangeable storage backends, and production-oriented software engineering practices.

![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)
![Architecture](https://img.shields.io/badge/Architecture-Repository%20Pattern-blue)
![Storage](https://img.shields.io/badge/Storage-JSON%20%7C%20SQLite-success)
![Status](https://img.shields.io/badge/Status-Active-brightgreen)

---

# Table of Contents

- [Project Overview](#project-overview)
- [Features](#features)
- [Why This Project?](#why-this-project)
- [Project Structure](#project-structure)
- [Architecture](#architecture)
- [Storage Architecture](#storage-architecture)
- [Repository Pattern](#repository-pattern)
- [Configuration](#configuration)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Design Principles](#design-principles)
- [Roadmap](#roadmap)
- [Contributing](#contributing)

---

# Project Overview

This project began as a simple command-line Todo application and gradually evolved into a backend-oriented Go project focused on software architecture rather than CRUD functionality alone.

The primary goal is to learn and demonstrate production-grade backend concepts while keeping the application small enough to understand completely.

Current architecture includes:

- Repository Pattern
- Dependency Injection
- Runtime Storage Selection
- Configuration Management
- Factory Pattern
- Contract Testing
- SQLite Persistence
- JSON Persistence
- Modular Package Structure

The project is intentionally designed to support future migration to REST APIs, PostgreSQL, Docker, and CI/CD without major architectural changes.

---

# Features

- ✅ Add Todos
- ✅ List Todos
- ✅ Mark Todos as Completed
- ✅ Delete Todos
- ✅ Persistent Storage
- ✅ JSON Repository
- ✅ SQLite Repository
- ✅ Runtime Backend Selection
- ✅ Centralized Configuration
- ✅ Structured Logging
- ✅ Repository Contract Tests
- ✅ Clean Modular Architecture

---

# Why This Project?

Most Todo applications stop after implementing CRUD operations.

This project intentionally goes further.

Instead of focusing solely on features, it focuses on software engineering practices commonly used in production Go applications.

Topics explored include:

- Repository Pattern
- Dependency Injection
- Factory Pattern
- Interface-driven Design
- Configuration Management
- Contract Testing
- Modular Project Organization
- Clean Architecture Principles

The goal is to serve as both a learning resource and a foundation for future backend development.

---

# Project Structure

```text
internal/
├── cli/                  # CLI commands and routing
├── config/               # Application configuration
├── constants/            # Shared constants
├── logger/               # Structured logging
├── repository/
│   ├── contract/         # Repository interfaces and shared models
│   ├── json/             # JSON repository implementation
│   ├── sqlite/           # SQLite repository implementation
│   ├── test/             # Shared repository contract tests
│   └── factory.go        # Runtime repository selection
└── todo/                 # Business logic
```

---

# Architecture

The application follows a layered architecture.

```text
                    CLI
                     │
                     ▼
              Todo Service
                     │
                     ▼
        TodoRepository Interface
              ▲             ▲
              │             │
      JSON Repository   SQLite Repository
```

Each layer has a single responsibility.

- CLI handles user interaction.
- Todo Service contains business logic.
- Repository defines persistence behavior.
- Repository implementations handle storage details.

---

# Storage Architecture

Persistence is abstracted behind a repository interface.

Instead of coupling business logic directly to JSON files or SQLite, every storage implementation satisfies the same interface.

```text
CLI
 │
 ▼
Todo Service
 │
 ▼
Repository Interface
 │
 ├── JSON Repository
 └── SQLite Repository
```

Because the service depends only on the interface, storage implementations can be replaced without modifying business logic.

---

# Repository Pattern

Every storage backend implements the same contract.

```go
type TodoRepository interface {
	Create(title string) (Todo, error)
	GetAll() ([]Todo, error)
	Update(todo Todo) error
	Delete(id int) error
}
```

Benefits:

- Loose coupling
- Easier testing
- Cleaner architecture
- Storage independence
- Future extensibility

---

# Runtime Storage Selection

Storage backend selection happens during application startup.

No code modifications are required.

Simply configure the desired backend.

```env
STORAGE_TYPE=json
```

or

```env
STORAGE_TYPE=sqlite
```

A repository factory selects the correct implementation based on the configuration.

---

# Configuration

Application configuration is centralized inside the `internal/config` package.

Supported environment variables:

| Variable | Default | Description |
|----------|----------|-------------|
| STORAGE_TYPE | sqlite | Storage backend |
| DB_PATH | todos.db | SQLite database path |
| LOG_LEVEL | info | Logging level |

Example:

```env
STORAGE_TYPE=sqlite
DB_PATH=todos.db
LOG_LEVEL=info
```

---

# SQLite Initialization

When SQLite is selected:

1. Opens the configured database.
2. Creates the database file automatically if it does not exist.
3. Creates required tables during startup.
4. Starts the application.

No manual database setup is required.

---

# Installation

## Clone the repository

```bash
git clone https://github.com/Ashwani1405/cli-todo.git
cd cli-todo
```

## Install dependencies

```bash
go mod tidy
```

---

# Usage

## Add Todo

```bash
go run . add "Learn Go"
```

## List Todos

```bash
go run . list
```

## Mark Todo Complete

```bash
go run . done 1
```

## Delete Todo

```bash
go run . delete 1
```

---

# Testing

Run all tests:

```bash
go test ./...
```

Verbose mode:

```bash
go test -v ./...
```

The project uses reusable repository contract tests.

The same behavioral tests are executed against:

- JSON Repository
- SQLite Repository

This guarantees consistent behavior across all storage implementations.

---

# Design Principles

The project demonstrates several software engineering concepts commonly used in production Go applications.

- Repository Pattern
- Dependency Injection
- Factory Pattern
- Interface-driven Design
- Configuration Management
- Separation of Concerns
- Contract Testing
- Modular Architecture

---

# Roadmap

## Completed

- Modular CLI Architecture
- Repository Pattern
- JSON Repository
- SQLite Repository
- Runtime Repository Factory
- Centralized Configuration
- Repository Contract Tests

## Planned

- PostgreSQL Repository
- REST API
- Docker Support
- GitHub Actions CI
- Database Migrations
- Authentication
- Pagination & Filtering

---

# Contributing

Contributions are welcome.

1. Fork the repository.
2. Create a feature branch.

```bash
git checkout -b feature/my-feature
```

3. Commit your changes.

```bash
git commit -m "feat: add my feature"
```

4. Push your branch.

```bash
git push origin feature/my-feature
```

5. Open a Pull Request.

Please ensure:

- Code follows Go formatting conventions.
- Tests pass (`go test ./...`).
- New features include appropriate documentation.
- Architectural consistency is maintained.

---

# Learning Outcomes

This project explores practical backend engineering concepts including:

- Clean Architecture
- Repository Pattern
- Dependency Injection
- Factory Pattern
- Configuration Management
- SQL using SQLite
- Contract Testing
- Interface-based Design
- Modular Go Applications

It serves as a foundation for progressively building production-ready backend systems in Go.