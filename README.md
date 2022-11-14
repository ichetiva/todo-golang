# Simple todo written in Go with net/http

Architecture:
```
│
├───cmd
│   └───app
│       └───main.go           - Main file which will be running
├───config
│   ├───config.go             - Loading config from config.yml
│   └───config.yml            - Config file
├───internal
│   ├───controllers
│   │   └───todo
│   │       ├───controller.go - Controller which contain all todo handlers
│   │       └───schemas.go    - Schemas for controller handlers
│   ├───routes
│   │   └───routes.go         - Include controllers
│   └───use_cases
│       └───server
│           └───server.go     - Setting server
└───pkg
    └───postgres
        ├───postgres.go       - Setting database for GORM
        └───models
            └───todo.go       - Database models
```

## Running:
1. Clone repository:
```
git clone https://github.com/ichetiva/todo-golang.git
```
2. Install modules:
```
go mod tidy
```
3. Rename config.yml.simple to config.yml and setting it
4. Run main file:
```
go run ./cmd/app
```
or\
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4. Build project:
```
Windows:
    go build -o ./bin/app.exe ./cmd/app
Linux:
    go build -o ./bin/app.exe ./cmd/app
```
5. Run built file:
```
Windows:
    ./bin/app.exe
Linux:
    ./bin/app
```
