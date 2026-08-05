# Project Generator

A fast, interactive CLI tool to scaffold Go projects with boilerplate code and best practices. Spin up new HTTP backends, CLI applications, or empty Go projects in seconds.

## Overview

Project Generator is a Go-based project scaffolding tool that removes the friction of starting new projects. Answer a few interactive questions, and get a fully structured project with:

- Pre-configured project layouts
- Optional database integrations (PostgreSQL, MySQL)
- Test case templates for HTTP backends
- Multiple CLI frameworks (Cobra, Bubble Tea, Huh)
- Cross-platform support (Windows, macOS, Linux)

Perfect for rapidly prototyping, learning Go, or maintaining consistent project structure across your team.

## Features

**Interactive Setup**
- Beautiful terminal UI with form validation
- Step-by-step project configuration
- Smart conditional prompts (database options only for HTTP backends)

**Project Templates**
- HTTP Backend - RESTful API with optional database and tests
- CLI Application - Command-line tools with multiple framework options
- Empty Project - Minimal Go project structure

**Database Support**
- PostgreSQL (PGX driver)
- MySQL
- None (for non-database projects)

**Testing**
- Optional test case templates for HTTP backends
- Pre-configured testing structure

**CLI Frameworks**
- Cobra - Standard flag-based CLI applications
- Bubble Tea - Interactive TUI (Text User Interface) applications
- Huh - Beautiful form-based CLI tools

## Installation

### Prerequisites
- Go 1.26 or higher

### From Source

```bash
git clone https://github.com/bran7230/projectGenerator.git
cd projectGenerator
go build -o project-gen main.go
```

Move the binary to your PATH:
```bash
# macOS/Linux
mv project-gen /usr/local/bin/

# Or add current directory to PATH
```

## Usage

Run the generator:
```bash
./project-gen
# or if installed globally
project-gen
```

### Interactive Prompts

1. **Project Name** - Enter a name for your new project
2. **Project Type** - Choose between:
   - Http-Backend
   - Cli-Application
   - Empty-Project
3. **Database** (HTTP Backend only) - Select database framework
4. **Test Cases** (HTTP Backend only) - Include test templates
5. **CLI Framework** (CLI Application only) - Select your CLI tool

### Example Workflow

```
$ project-gen

     ____             _           __          ______
    / __ \_________  (_)__  _____/ /_        / ____/___  ____
   / /_/ / ___/ __ \/ / _ \/ ___/ __/_______/ / __/ __ \/ __ \
  / ____/ /  / /_/ / /  __/ /__/ /_/_______/ /_/ /  __/ / / /
 /_/   /_/   \____/ /\___/\___/\__/        \____/\___/_/ /_/
              /___/

Enter project name: my-awesome-api
Confirm project name? Yes
Choose a project: Http-Backend
Confirm project type? Yes
Choose a database framework: PostgreSQL(PGX driver)
Confirm database? Yes
Include test cases? Yes

Project scaffolded successfully!
Project location: /path/to/my-awesome-api
```

Then start developing:
```bash
cd my-awesome-api
go run main.go
```

## Project Structure

```
project-generator/
├── main.go              # CLI entry point
├── go.mod              # Go module definition
├── go.sum              # Dependency checksums
├── makefile            # Build automation
├── project_generator/  # Core generation logic
└── .gitignore          # Git configuration
```

## Technologies Used

- Charm Huh - Beautiful terminal forms
- Charm Lipgloss - Terminal styling
- Bubble Tea - TUI framework (included in templates)
- Cobra - CLI framework (included in templates)

## Cross-Platform Support

- Windows
- macOS
- Linux

The tool automatically detects your operating system and adjusts behavior accordingly.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development

```bash
# Build the project
go build -o project-gen main.go

# Test it
./project-gen
```

## License

This project is open source and available under the MIT License.

## Support

Have questions or found a bug? Please open an issue on GitHub.

---

Happy scaffolding!
