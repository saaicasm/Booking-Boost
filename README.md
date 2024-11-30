# Booking and Reservation System

This is a Go-based project for managing bookings and reservations. It leverages modern tools and libraries to provide secure and efficient features such as authentication, session management, and email notifications.

## Features
- **Booking and Reservation Management**
- **User Authentication**
- **Session Management** using `scs`
- **Routing** with `chi` router
- **Database Management** using PostgreSQL
- **Email Notifications**

## Tech Stack
- **Programming Language:** Go (version 1.22.5)
- **Router:** [Chi Router](https://github.com/go-chi/chi)
- **Session Management:** [SCS](https://github.com/alexedwards/scs)
- **Database:** PostgreSQL

## Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/iamlego/Booking-Boost.git
   cd Booking-Boost
   ```

2. **Install Dependencies**
   Make sure you have Go installed. Run the following command to install project dependencies:
   ```bash
   go mod tidy
   ```

3. **Set Up Environment Variables**
   Create a `.env` file and configure the following variables:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=booking_db
   SMTP_HOST=smtp.example.com
   SMTP_PORT=587
   SMTP_USER=your_email@example.com
   SMTP_PASSWORD=your_email_password
   ```

4. **Run Database Migrations**
   Use a tool like `migrate` to set up the database schema:
   ```bash
   migrate -path migrations -database "postgres://your_db_user:your_db_password@localhost:5432/booking_db?sslmode=disable" up
   ```

5. **Run the Application**
   Start the server with:
   ```bash
   go run cmd/web/main.go
   ```

## Usage
- Access the application at `http://localhost:8080`.
- Create a user account and log in to manage bookings and reservations.

## Folder Structure
```plaintext
Booking-Boost/
├── cmd/web/            # Web application entry point
│   ├── main.go         # Main application file
│   ├── middleware.go   # Middleware for the application
│   └── routes.go       # Application routes
├── pkg/config/         # Configuration package
│   └── config.go       # Configuration setup
├── pkg/handler/        # Handlers for routes
│   └── handler.go      # Route handlers
├── pkg/models/         # Models for data structures
│   └── templateData.go # Template data structures
├── pkg/render/         # Rendering utilities
│   └── render.go       # Template rendering
├── templates/          # HTML templates
│   ├── about.page.tmpl # About page template
│   ├── base.layout.tmpl# Base layout template
│   └── home.page.tmpl  # Home page template
├── .env.example        # Example environment variables file
├── go.mod              # Dependency management
└── migrations/         # Database migration files
```

## Libraries Used
- [Chi Router](https://github.com/go-chi/chi) - Lightweight, idiomatic Go HTTP router
- [SCS](https://github.com/alexedwards/scs) - Secure Cookie-based Session Management
- [pq](https://github.com/lib/pq) - PostgreSQL driver for Go
- [Godotenv](https://github.com/joho/godotenv) - Environment variable management

## Contribution
Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a new branch for your feature/bug fix.
3. Commit your changes and create a pull request.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

---

Feel free to reach out with any questions or suggestions!
