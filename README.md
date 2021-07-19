![example workflow](https://github.com/muathendirangu/convertible/actions/workflows/go.yml/badge.svg)

# convertible - currency converter
A REST based Go web service to convert currencies in Go. Currently supports conversions of following currencies(Nigerian Naira(NGN), Ghanaian Cedis(GHS), and Kenyan Shillings (KSH))

## Technologies used
  - [Golang](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
  - [GoDotEnv](https://pkg.go.dev/github.com/joho/godotenv) A Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file).


## Build from scratch

### 1. Project setup
clone the code
`git clone git@github.com:muathendirangu/convertible.git`

change directory
`cd convertible`

### 2. Setup project
 
rename the file from `.env.example` to `.env` 

```
PORT=
```

Populate this file with your preferred port variable.

### 3. Run the application

Run the application: `go run main.go`.

Using curl:
`curl -d '{ "from": "STRINGVALUE", "to": "STRINGVALUE","amount": NUMBERVALUE}' -H "Content-Type: application/json" -X POST http://localhost:PORT/`

### 4. Running unit tests
To run tests use the following commands
`go test -v ./... --cover`

## Author

- [Charles Ndirangu](https://twitter.com/muathendirangu)

