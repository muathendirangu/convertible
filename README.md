![example workflow](https://github.com/muathendirangu/convertible/actions/workflows/go.yml/badge.svg)

# convertible - currency converter REST based web service
A REST based Go web service to convert currencies in Go. Currently supports conversions of following currencies(Nigerian Naira(NGN), Ghanaian Cedis(GHS), and Kenyan Shillings (KSH))

### Assumptions
- A simple REST based service that takes a json payload of Currency to Convert from, Currency to convert to and an Amount that returns a response json object containing the initial json payload with additional Message field to give more context on the conversion.
- That conversion rate are pre-selected(taken from google conversion rates on date 16th July 2021) and stored in a local slice of of structs containing From,To and conversion rate per object. Any request that is sent to the endpoint utilizes this hard-corded slice of structs as a point of reference for computation.

## Technologies used
  - [Golang](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
  - [GoDotEnv](https://pkg.go.dev/github.com/joho/godotenv) A Go (golang)  which loads env vars from a .env file. It is easier and straightforward to use and widely used by most open-source project


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


### 3. Running unit tests

To run tests use the following commands
`go test -v ./... --cover`

### 4. Run the application

Run the application: `go run main.go`.

Using curl:
`curl -d '{ "from": "STRINGVALUE", "to": "STRINGVALUE","amount": NUMBERVALUE}' -H "Content-Type: application/json" -X POST http://localhost:PORT/`

NB. Allowed currency choice representation
- KES,GHC and NGN
- sample request response:
`http://localhost:PORT/`

**method:** POST
Convert from Kenyan Shillings(KES) to Ghanian Cedi(GHC)
**payload:**

```
{
    "amount": "5",
    "from": "KES",
    "to": "GHC"
}
```
**response:**
```
{
    "From": "KES",
    "To": "GHC",
    "InitialAmount": 5,
    "ConvertedAmount": 0.28,
    "DefaultConversionRate": 0.055,
    "Message": " Amount 5 KES is equivalent to 0.28 GHC after conversion using the rate of 1 KES equals 0.055 GHC"
}
```
Convert from Kenyan Shillings(KES) to Nigerian Naira(NGN)
**payload:**

```
{
    "amount": "200",
    "from": "KES",
    "to": "NGN"
}
```
**response:**
```
{
    "From": "KES",
    "To": "NGN",
    "InitialAmount": 200,
    "ConvertedAmount": 760,
    "DefaultConversionRate": 3.8,
    "Message": " Amount 200 KES is equivalent to 760 NGN after conversion using the rate of 1 KES equals 3.8 NGN"
}
```
Convert from Ghanian Cedi(GHC) to Kenyan Shillings(KES) 
**payload:**

```
{
    "amount": "200",
    "from": "GHC",
    "to": "KES"
}
```
**response:**
```
{
    "From": "GHC",
    "To": "KES",
    "InitialAmount": 200,
    "ConvertedAmount": 3634,
    "DefaultConversionRate": 3.8,
    "Message": " Amount 200 KES is equivalent to 3634 KES after conversion using the rate of 1 GHC equals 18.17 KES"
}
```
## Author

- [Charles Ndirangu](https://twitter.com/muathendirangu)

