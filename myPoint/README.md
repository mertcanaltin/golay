# Go User API

This is a Go project that generates random user data and provides a simple API to filter users based on their age.

## Usage

1. Clone or download the project to your local machine.

2. Make sure you have Go installed on your computer.

3. Open a command prompt and navigate to the project directory.

4. Run `go run main.go` to start the application.

5. The API will be available at `http://localhost:8080/users`.

6. To retrieve all users, send a GET request to `http://localhost:8080/users`.

7. To retrieve users of a specific age, send a GET request to `http://localhost:8080/users?age=<age>`, replacing `<age>` with the desired age value.

8. The API will return the user data in JSON format.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
