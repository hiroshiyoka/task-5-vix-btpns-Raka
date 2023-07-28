# REST API

Project-Based Intern: Fullstack Developer Virtual Internship Experience BTPN Syariah

![Banner](./src/banner.png)

This is a simple Rest-API built using GO-Language. Users can either log in, register, and log out. Afterward, users can add a profile photo and change the photo. Only logged-in or signed-up users have the privilege to delete or add a profile photo. Different users cannot delete or modify photos created by other users. The API utilizes JSON Web Token and Mux Middleware for both Authentication and Authorization.

## Features

- User can log in, register, and log out.
- User can add, update, delete, and view profile photos.
- Authentication and authorization using JSON Web Token and Mux Middleware.

## Technology/Framework Used

- GO
- Gorilla Mux
- Gorm
- MySQL

## Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/hiroshiyoka/task-5-vix-btpns-Raka.git
```

Change into the project directory:

```bash
cd task-5-vix-btpns-Raka
```

Install the necessary dependencies:

```bash
go get -u github.com/gorilla/mux gorm.io/gorm gorm.io/driver/mysql golang.org/x/crypto
```

Install the JSON Web Token:

```bash
go get github.com/golang-jwt/jwt/v4
```

Run the project:

```bash
go run main.go
```