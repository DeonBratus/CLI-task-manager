# Task Management CLI Application

## Overview

This Go-based command-line application allows users to manage tasks within a task board. The program provides functions to create, update, delete, and view tasks organized by their status, helping users streamline their task management workflow.

## Features

- **Task Creation**: Add tasks with unique names to the task board.
- **Task Editing**: Update task name, description, or status.
- **Task Deletion**: Remove tasks by name.
- **Task Viewing**: View details of specific tasks or filter tasks by status (To Do, In Progress, Done).
- **Task Listing**: Display all tasks in the task board.

## Getting Started

### Prerequisites

Ensure that **Go** is installed on your machine to compile and run the application.

### Running the Application

1.**Compile** the application using the Go compiler:
```bash
go build -o .
```
And run application\

2.**Run** the application:
```bash
go run main.go
```
## Manage task
Commands:
### Working with tasks
```add <task name>:``` Adds a new task.

```del <task name>```: Deletes a specified task.

```edit <task name>```: Prompts for task attribute changes (name, description, or status).
### Showing info about tasks
```show <task name>```: Shows detailed task information.

```ldn```: Lists tasks with Done status.

```ltd```: Lists tasks with ToDo status.

```lip```: Lists tasks with In Progress status.

```lall```: Lists all tasks.