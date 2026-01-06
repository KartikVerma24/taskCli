# taskCli

**taskCli** is a simple, interactive **command-line task manager written in Go**. 
It lets you add, view, update, complete, and delete tasks directly from your terminal.

Built to practice clean Go project structure and CLI design.

---

## Features

- Add tasks with description and priority 
- List all tasks 
- Update task status or priority 
- Mark tasks as done 
- Delete tasks 
- Interactive shell-based CLI 

---

## Installation

```bash
git clone https://github.com/KartikVerma24/taskCli.git
cd taskCli
go build -o taskCli .
sudo cp taskCli /usr/local/bin/
sudo mv taskCli /usr/bin/
```

Run the CLI:

```bash
taskCli --version
```

---

## Usage

When you start the app, you’ll see:

```
Welcome to the taskCLI
Type 'help' to see commands, 'exit' to quit
task>
```

Type `help` to see all available commands.

---

## 📌 Available Commands

### 1. `add` :- Add a new task

**Tags**
- `--desc` *(mandatory)*: task description (must be in double quotes)
- `--priority` *(optional)*: `low | medium | high | critical`

**Examples**
```bash
add --desc "testing app"
add --desc "code review" --priority low
```

---

### 2. `list-all` :- List all tasks

```bash
list-all
```

---

### 3. `change` :- Update task status or priority

**Tags**
- `--id` *(mandatory)*: task ID
- `--status`: `todo | wip | done | cancelled`
- `--priority`: `low | medium | high | critical`

**Examples**
```bash
change --id 1 --status wip
change --id 2 --priority critical
```

---

### 4. `done` :- Mark a task as done

**Tags**
- `--id` *(mandatory)*

```bash
done --id 1
```

---

### 5. `delete` :- Delete a task

**Tags**
- `--id` *(mandatory)*

```bash
delete --id 1
```

---

### 6. `sort` :- Sort the tasks on the basis of priority or status

**Tags**
- `--by` *(mandatory)* : `status | priority`

```bash
sort --by status
sort --by priority
```

---

### 7. `clear` :- Clear the terminal screen

```bash
clear
```

---

## Purpose

This project is focused on :-
- Learning Go project structure 
- Building an interactive CLI 
- Separating concerns cleanly (CLI, domain, service, DB)

---
