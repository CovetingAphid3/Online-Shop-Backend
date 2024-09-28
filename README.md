
# Online Shop Backend 🛒

This is the backend for an online shop, built using **Go** as a challange. The project serves as a RESTful API to manage products, orders, and users, providing essential functionalities for an e-commerce platform.

## Features ✨

- **User Authentication**: Secure user registration and login processes.
- **Product Management**: Add, update, and delete products in the inventory.
- **Order Processing**: Create and manage orders placed by users.
- **Environment Configuration**: Loads environment variables for database connectivity and other configurations.

## Project Structure 📂

```
├── api
│   └── routes.go               # API route definitions
├── controllers
│   ├── order_controller.go      # Order-related logic
│   ├── product_controller.go    # Product-related logic
│   └── user_controller.go       # User-related logic
├── go.mod                       # Go module dependencies
├── go.sum                       # Go module checksum
├── initializers
│   ├── connectToDb.go          # Database connection logic
│   ├── loadEnvVariables.go      # Load environment variables
│   └── syncDatabase.go          # Synchronize the database schema
├── main.go                      # Entry point of the application
├── middleware
│   └── requireAuth.go           # Middleware for user authentication
└── models
    ├── order.go                 # Order model definition
    ├── product.go               # Product model definition
    └── user.go                  # User model definition
```

## Technologies Used 🛠️

- **Go**: Programming language used for backend development.
- **Gin**: Web framework for Go 
- **Gorm**: ORM for Go 
- **PostgreSQL/MySQL**: Database 

