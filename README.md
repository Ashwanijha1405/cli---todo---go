# CLI Todo App in Go

## Project Overview

A persistent CRUD-based CLI todo application built using Go. It allows users to manage tasks directly from the terminal using commands like add, list, done, and delete. The project is designed with a modular structure to support future expansion into a backend-oriented system.

## Features

- Add todos
- List todos
- Mark todos as completed
- Delete todos
- Persistent JSON storage

## Installation

### 1. Clone the repository
```bash
git clone https://github.com/Ashwanijha1405/cli---todo---go
cd cli---todo---go
```

### 2. Install dependencies
```bash
go mod tidy
```

## Commands

### Add Todo

```bash
go run . add "Learn Go"
```
### List Todo

```bash
go run . list
```
### Mark Todo Completed

```bash
go run . done 1
```
### Delete Todo

```bash
go run . delete 1
```

## Project Structure

internal/
├── cli/        # Handles command parsing and routing
├── storage/    # Handles saving/loading todos (JSON persistence)
├── todo/       # Core logic (add, delete, update)

## Architecture Overview

The application follows a layered architecture:

- CLI Layer → Handles user input commands
- Business Layer → Core todo operations
- Storage Layer → Handles JSON persistence
- Logger Layer → Debugging and logging support

This separation improves scalability and maintainability.

## Contribution Guidelines

1. Fork the repository
2. Create a new branch:
   git checkout -b feature-name
3. Commit your changes:
   git commit -m "Add feature"
4. Push to branch:
   git push origin feature-name
5. Open a Pull Request