# AuthoBookshelf

AuthoBookshelf, a robust online bookstore, excels in user authentication, authorization, and access management. It boasts secure registration and authentication, enabling account management with data retention adherence. The platform prioritizes system logging, offering resilience during failures, while users effortlessly explore, add to carts, and review books. Administrators efficiently handle inventory. Built on PostgreSQL, AuthoBookshelf ensures a seamless user experience with clear setup instructions.

## Getting Started

### Installing

<!-- ```sh
clone, cd into project directory,pull out main.go , go mod tidy,go run main.go, application statrted on 9010 can go on broser and test the api
``` -->

Step 1 : Clone the Repository

Step 2 : Go into Root of the application

```
cd online-book
```

Step 3 : Install the dependencies

```
go mod tidy
```

Step 4 : Run The application

```
go run main.go
```

To test the application, kindly paste the following URL into your web browser:

http://localhost:9010/

## Features

### API Documentation :

- API endpoints are documented using Postman, and the documentation is not hard-coded. You can access the documentation by visiting the following link: https://documenter.getpostman.com/view/18129767/2s9Y5cu1ao

### User Authentication and Authorization :

- JSON Web Tokens (JWT) are used to manage user authentication and authorization.

### Data Models :

- GORM is employed to establish data models for users, books, cart items, orders, and order details.

### Admin Capabilities :

- Administrators have the ability to manage the book inventory with the following actions:

  - Create new books
  - Retrieve and view books
  - Update book information
  - Delete books from the database

### User Actions :

- After registering and logging in, users can perform the following actions:

  - Add books to their cart and view cart details
  - Remove books from their cart
  - Decrement the quantity of books in the cart
  - Add books to their order, purchase them, and view order details
  - Delete their user account
  - Deactivate their user account
  - Add book reviews

### Database :

- The application uses a PostgreSQL-based database to store and manage data.

### Docker Integration :

- A Dockerfile is used to build and run the application.

## Directory Structure

```

online-bookstore
│───main.go
│───Dockerfile
└───package
    │───auth
    | └──auth.go
    │───config
    | └──app.go
    │───controllers
    │ │──book-controller.go
    | │──cart-controller.go
    | └──user-controller.go
    │───helper
    | └──helper.go
    │───middleware
    | └──middleware.go
    │───models
    │ │──book.go
    | │──cart.go
    | └──user.go
    │───routes
    | │──book-routes.go
    | └──user-routes.go
    │───utils
    | └──utils.go


```
