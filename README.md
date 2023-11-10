# Flash-Cards API

**Description**: An API in Golang for creating, managing, and effectively studying flash-cards.

<p align="center">
  <img src="./assets/logo.svg" alt="Golang Flash-Cards API">
</p>

## Table of Contents

- [Requirements](#requirements)
- [Setup](#setup)
- [Database](#database)
- [Usage](#usage)
- [Examples](#examples)
- [Contribution](#contribution)

## Requirements

- [Golang](https://golang.org/) (version 1.19 or higher)
- Database (Mysql 8)
- Docker

## Database

1. A simple image description of our database:

<p align="center">
  <img src="https://user-images.githubusercontent.com/5201283/281923704-e931152f-0e81-47ce-99bf-aa7e6ecd1b32.png" alt="Golang Flash-Cards API">
</p>

## Setup

1. Clone the repository:

```bash
git clone https://github.com/andremartinsds/flash-cards-api
```

## Usage

1. Navigate to the project folder:

```bash
cd flash-cards-api
```

2. Configure environment variables:

Create a .env file in the project's root and set the necessary environment variables, such as database credentials and server settings.

```bash
  MYSQL_ROOT_PASSWORD=your_mysql_root_password
  MYSQL_DATABASE=your_mysql_database
  MYSQL_USER=your_mysql_user
  MYSQL_PASSWORD=your_mysql_password
```

3. Install dependencies:

```bash
go mod tyde
```

4. Execute docker-compose:

```bash
docker-compose up
```

5. Run api

In this case you can run the flash-cards-api with:

```bash
go run main.go
```

if you wnat execute with swagger you can run with make like:

```bash
make
```

## Examples

put the video gif here

## Contribution

fell free to send a pull request

## Author

André Martins
