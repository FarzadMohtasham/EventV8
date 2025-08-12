# EventV8

EventV8 is a lightweight and efficient event management REST API built with Go (Golang). It provides endpoints for creating, managing, and retrieving events, making it suitable for event scheduling applications, calendars, or booking systems.

---

## Features

* **RESTful API** design
* Create, read, update, and delete events
* Search and filter events by criteria
* MySQL database integration
* Built with performance and scalability in mind using Golang
* Modular and maintainable codebase

---

## Tech Stack

* **Language:** Go (Golang)
* **Database:** MySQL
* **Framework:** Gin (for HTTP routing)
* **ORM:** sqlx

---

## Installation

### Prerequisites

Ensure you have the following installed on your system:

* Go (>= 1.20)
* MySQL (>= 8.0)
* Git

### Clone the Repository

```bash
git clone https://github.com/FarzadMohtasham/EventV8.git
cd EventV8
```

### Install Dependencies

```bash
go mod tidy
```

### Configure Database

1. Create a MySQL database.
2. Update the `config.yaml` file (or `.env` if used) with your database credentials.

Example `config.yaml`:

```yaml
database:
  host: localhost
  port: 3306
  user: root
  password: your_password
  name: eventv8
```

### Run the Application

```bash
go run main.go
```

The server will start (default: `http://localhost:8080`).

---

## API Endpoints

| Method | Endpoint      | Description        |
| ------ | ------------- | ------------------ |
| POST   | `/events`     | Create a new event |
| GET    | `/events`     | Get all events     |
| GET    | `/events/:id` | Get event by ID    |
| PUT    | `/events/:id` | Update event by ID |
| DELETE | `/events/:id` | Delete event by ID |

---

## Example Request

### Create Event

```bash
curl -X POST http://localhost:8080/events \
-H "Content-Type: application/json" \
-d '{
  "name": "Tech Conference 2025",
  "description": "A conference about the latest in technology.",
  "location": "Berlin, Germany",
  "dateTime": "2025-10-12T10:00:00Z",
  "user_id": 1
}'
```

---

## Project Structure

```
EventV8/
├── main.go
├── config/
├── controllers/
├── models/
├── routes/
├── utils/
└── go.mod
```

---

## Contributing

Contributions are welcome! Please fork the repository and create a pull request.

Steps:

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Author

**Farzad Mohtasham**
[GitHub Profile](https://github.com/FarzadMohtasham)
