# 💻 About
The API was created to simulate a digital bank, through account management. It was built following the clean architecture concepts created by Robert C. Martin.

Table of Contents
=================
* [About](#about)
* [Tabela de Conteudo](#table-of-contents)
* [Features](#features)
* [How It Works](#how-it-works)
    * [Application / Dependencies](#application--dependencies)
    * [How to Begin](#how-to-begin)
    * [Technologies](#technologies)
* [Endpoints](#endpoints)
    * [Accounts](#accounts)
    * [Login](#login)
    * [Transfer](#transfer)
* [Credits](#credits)
* [License](#license)
### ⚙️ Features
- [x] Create an Account
- [x] Getting Accounts
- [x] Authenticate users
- [x] Transfer between accounts
- [x] Get transfer list
### 🚀 How It Works

This project is just the back end part (server folder)
### 🚧 Application / Dependencies

Before starting, you will need to have the following tools installed on your machine:
[Git](https://git-scm.com), [Golang](https://golang.org/dl/), [PostgreSQL](https://www.postgresql.org/).

I recommend having an editor to work with code like [VSCode](https://code.visualstudio.com/) and having an api client like [Postman](https://www.postman.com/downloads/)


It was stored in containers using [Docker](https://www.docker.com/)

### 🎲 How to Begin

```bash
# Clone this repository
$ git clone <https://github.com/CMedrado/DesafioStone>

# Access the project folder in the terminal/cmd
$ cd DesafioStone

# Go to local folder within deploy
$ cd deploy/local

# Run the application
$ docker-compose up --build -d

# The server will start on port:5000 - go to <http://localhost:3333>
```

### 🛠 Technologies

The following tools were used in the construction of the project:

- [Postman](https://www.postman.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [Golang](https://golang.org/)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Docker](https://www.docker.com/)
- [Logrus](https://github.com/sirupsen/logrus)
- [Envconfig](github.com/kelseyhightower/envconfig)

## 📎 Endpoints

### Accounts
- `POST /accounts` - Create an Account
- `GET /accounts` - Get the list of accounts
- `GET /accounts/{id}/balance` - Get account balance

### Login
- `POST /login` - authenticate the user

### Transfer
- `GET /transfers` - get the list of transfers from the authenticated user.
    - Requires `Authorization` header entry.
- `POST /transfers` - transfers from one Account to another.
    - Requires `Authorization` header entry.


## ✅ Credits
I thank all the authors and content available by the application developers and my mentor [Pedro](https://github.com/pedroyremolo) for guiding me in creating the project.
## ✅ License
[MIT](https://choosealicense.com/licenses/mit/)