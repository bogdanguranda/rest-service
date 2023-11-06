# Rest Service

# Overview

A simple HTTP REST server that can search on a sorted dataset using the binary search algorithm.

# Project structure

- **db**: logic used to read the input data
- **service**: the search algorithm
- **util**: helpful logic (i.e. configuration, logging)
- **web**: logic to set up and create the web server with the request handler

# Prerequisites

- Go 1.16 or higher (tested on Go 1.21.3)

# Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/bogdanguranda/rest-service.git
    ```

2. Navigate to the project directory:

    ```bash
    cd rest-service
    ```

# Running the Application

Run the main.go file to start the server:

```bash
go run main.go
```


## Testing

## Manual testing

### Request

Send a GET request (e.g. with Postman or curl) to `localhost:8080/value/<value>`, for example:

```
GET localhost:8080/value/10
GET localhost:8080/value/1800
GET localhost:8080/value/0
```

### Response

Example of response format:

```json
{
   "index": 3,
   "value": 100
}
```

Or with message in case of errors (which is optional):
```json
{
   "index": 0,
   "value": 0,
   "message": "Invalid 'value' query parameter, must be a number."
}
```

To receive the `index` at which the value is located.

## Unit tests
Run all the tests in the project:

```bash
make test
```