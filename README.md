# Book Management 📚
![image](https://github.com/user-attachments/assets/74cfad5d-4f4f-453b-b26d-7785e3f30bd1)

This project is built using Go with the Gin framework for handling HTTP requests, Gorm as the ORM for database operations, JWT for user authentication, and bcrypt for password hashing.

## API Endpoints
### Authentication
For authentication, I’m using JWT. When users log in, they receive a signed JWT token, which they use to authenticate their future requests. Passwords are hashed using bcrypt to make more secure.
- POST `/register`: Registers a new user.<br>

| Parameter | Type | Example|
| ------------- | ------------- | ------------- |
| username | string | "user" |
| password | string | "test" |
<br>

- POST `/login`: Logs a user in and returns a JWT token.<br>

| Parameter | Type | Example|
| ------------- | ------------- | ------------- |
| username | string | "user" |
| password | string | "test" |
<br>

If the user registers successfully and logs in correctly, they’ll receive a token that can be used to authenticate all requests related to managing books (Add bearer token in header).

### Book management
I’ve kept things simple and straightforward, but also ensured the API covers key functionality like creating, reading, updating, and deleting books.

- GET `/books/list`: List all books.<br>
- GET `/books/detail`: Get a book by ID.<br>

| Parameter | Type | Example|
| ------------- | ------------- | ------------- |
| ID | int | 1 |
<br>

- POST `/books/create`: Create a new book.<br>

| Parameter | Type | Example|
| ------------- | ------------- | ------------- |
| name | string | "Test Book" |
| author | string | "Test Author 1" |
| description | string | "Best Book!" |
| price | int | 1000 |
<br>

- PUT `/books/update`: Update a book by ID.<br>

| Parameter | Type | Example|
| ------------- | ------------- | ------------- |
| ID | int | 1 |
| name | string | "Test Book" |
| author | string | "Test Author 1" |
| description | string | "Best Book!" |
| price | int | 1000 |
<br>

- DELETE `/books/delete`: Delete a book by ID.<br>


| Parameter | Type | Example|
| ------------- | ------------- | ------------- |
| ID | int | 1 |
<br>



