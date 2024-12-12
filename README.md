# MyGit

MyGit is a simple, custom implementation of a Git-like version control system (VCS). The goal is to understand and build the core features of Git from scratch, starting with the initialization (init) functionality. The project is written in Go and uses the Cobra framework for building CLI commands.

---

## Features (So Far)

### Initialize Repository:
Run `mygit init` to initialize a `.mygit` directory in your project folder, similar to `git init`.

---

## File Structure
```
mygit/
├── cmd/
│   ├── root.go      # Root command of the CLI tool
│   └── init.go      # Implementation of the 'init' command
├── internal/
│   └── core/
│       └── init.go  # Core logic for initializing the repository
├── pkg/
│   └── utils/
│       └── helpers.go  # Helper utilities for future features
├── main.go          # Entry point for the application
├── go.mod           # Dependency and module management
└── go.sum           # Dependency checksums
```

---

## Installation

### Clone the Repository:
```bash
git clone https://github.com/abhiii71/mygit.git
cd mygit
```

### Install Dependencies:
Ensure you have Go installed (1.19 or later) and run:
```bash
go mod tidy
```

### Run the CLI:
```bash
go run main.go 
```

---

## Usage

### Initialize a Repository
Navigate to the directory where you want to initialize the repository:
```bash
cd /path/to/your/project
```
Run the `init` command:
```bash
mygit init
```

#### What Happens:
- A `.mygit` directory is created in your project folder.
- A `HEAD` file is created inside `.mygit`, pointing to the `refs/heads/main` branch.

#### Output:
- If successful:
```bash
Initialized empty MyGit repository in /path/to/your/project/.mygit
```
- If `.mygit` already exists:
```bash
Error: repository already initialized
```

### Check Repository Status
Run the `status` command to view the current state of the repository:
```bash
mygit status

---

### Tracked and Untracked Files
The `status` command displays:
- **Tracked files**: Files that have been added to the `.mygit/index` file.
- **Untracked files**: Files in the directory that are not in `.mygit/index` and not ignored by `.mygitignore`.

#### Example:
```bash
Tracked files:
  example_tracked_file.txt

Untracked files:
  new_file.txt
  another_new_file.txt
Adding Files


------------------------------------------------

## Dependencies
- **Cobra**: Framework for building CLI commands.

---

## Future Enhancements

### Core Git Features:
- Add, Commit, Log, Branch, and Merge.

### Advanced Commands:
- Status, Revert, and Remote interactions.

### Tests:
- Add unit and integration tests for each command and feature.

---

