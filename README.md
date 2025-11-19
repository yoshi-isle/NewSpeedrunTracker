# New Speedrun Tracker with Go REST API & PostgreSQL

This is a Go REST API project that performs CRUD operations on `Submission`, `Category`, and `Activity` resources using the Gin framework and PostgreSQL. The project also includes middleware for logging the IP address of incoming requests.

## Features

- Create, read, update, and delete submissions
- Create and assign new categories of PBs
- Create new activities and assign them to a category
- Log IP addresses of incoming requests

### Middleware

The project includes a middleware that logs the IP address of incoming requests. The middleware is defined in `middleware/ip_logger.go` and is registered in the router setup in `routes/routes.go`.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
