# Go API and HTTP Request

This repository contains two Go programs:

## Go User API

The `myPoint/myPoint.go` file sets up a simple API that generates random user data and allows filtering users based on age. The API runs on `localhost:8080/users` and supports the following endpoints:

- `/users` (GET): Retrieves all users.
- `/users?age=<age>` (GET): Retrieves users of a specific age.

To run the API, make sure you have Go installed on your computer. Open a command prompt, navigate to the project directory, and run `go run main.go`. The API will be available at `http://localhost:8080/users`.

## Go HTTP Request

The `users/users.go` file sends an HTTP GET request to a specified URL and prints the response body. To use the program, modify the `targetURL` variable to the desired URL. Run the program by executing `go run http_request.go` in a command prompt.

## License

This repository is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
