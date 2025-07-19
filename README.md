# mrn

`mrn` is a lightweight and flexible CLI tool for Go developers that helps scaffold basic folder structures and boilerplate files for new feature modules.

## Features

- 🛠 Auto-generates a new folder with:
    - `handlers.go`
    - `models.go`
    - `services.go`
- 🔁 Simple and repeatable CLI usage
- 🧼 Clean and idiomatic file templates
- 🧱 Helps keep your project modular and consistent

## Installation

```bash
go install github.com/gnomedevreact/mrn@latest
```

# Usage

```bash
mrn generate <module_name> <target_path>
```

# Example
```
mrn generate users ./api
```
This will create:

 ```
./api/users/
├── handlers.go
├── models.go
└── services.go
```