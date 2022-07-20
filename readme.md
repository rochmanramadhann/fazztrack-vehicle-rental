<h1 align="center">Lectronic-Ku</h1>
<p align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/2560px-Go_Logo_Blue.svg.png" width="400px" alt="Golang.jpg" /></p>
<p align="center">
    <a href="https://golang.org/" target="blank">More about Golang</a>
</p>

## Built With

[![Golang](https://img.shields.io/badge/Golang-4.x-blue.svg?style=rounded-square)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-v.14.2-blue.svg?style=rounded-square)](https://www.postgresql.org/)

## Description

This backend application is used by the user to record incoming orders and manage products. In this application, users can display products, add products, delete products, edit products, and calculate the total price of an order. This application is built with Golang using the gorilla/mux package for routing. The databases used in this application are PostgreSQL and Redis for the caching process.

## Feature

- Authentication and Authorization
- JWT Web Token
- CRUD Product
- CRUD Category
- CRUD History
- Solid Principle
- Search Product Name
- Sort Product Name, Category, Date Update, and Price order by ASC or DESC

## Installation Steps

1. Install the Golang and GO Environment

```bash
https://golang.org/doc/install
```

2. Clone the repository

```bash
git clone git@github.com:wsaefulloh/cashier-backend.git (go get)
```

3. Add .env file at root folder project

```sh
DB_HOST=your_db_host
DB_USER=your_db_user
DB_PASS=your_db_pass
DB_NAME=your_dbname
```

4. Build the app

```bash
go build
```

5. Run the app

```bash
go run main.go serve
```

6. You are all set!

```bash
View the website at: http://localhost:8080
```

## End Point

You can see all the end point [here](https://documenter.getpostman.com/view/12136790/UzJPLExt#df1f63e0-c9bf-4db2-b0ca-67ec3a17e2a2)
