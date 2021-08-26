<img src="theironbankofbraavosby.jpeg" alt="The Iron Bank of Braavos" width="280" align="right" />

# 🏦 The Iron Bank of Braavos
Welcome to the Iron Bank of Braavos.<br />
Make your transactions and your credit requests but remember the Iron Bank will have its due.<br />
## 💻 About
The API was created to simulate a digital bank, through account management. It was built following the RESTful API Design and clean architecture concepts. <br />
Public image available on [Docker Hub](https://hub.docker.com/r/rafaelcmedrado/desafio)
## Table of Contents

* [About](#about)
* [Tabela de Conteudo](#table-of-contents)
* [Features](#features)
* [How It Works](#how-it-works)
    * [Application / Dependencies](#application--dependencies)
    * [Environment Variables](#environment-variables)
    * [How to Begin](#how-to-begin)
      * [Makefile](#makefile)
    * [Technologies](#technologies)
    * [Endpoints](#endpoints)
       * [Accounts](#accounts)
       * [Login](#login) 
       * [Transfer](#transfer)
* [Credits](#credits)

### ⚙️ Features
- [x] Create an Account
- [x] Getting Accounts
- [x] Authenticate users
- [x] Transfer between accounts
- [x] Get transfer list
### 🚀 How It Works

This project is just the back end part (server folder)
### 🔢 Environment Variables
| Name                      | Description                                                |
|---------------------------|------------------------------------------------------------|
| API_PORT                  | Port that will be listened on for the new request          |
| API_LOG_LEVEL             | Structured api log level                                   |
| DB_PROTOCOL               | DB instance protocol                                       |
| DB_USERNAME               | DB instance user                                           |
| DB_SECRET                 | DB instance password                                       |
| DB_HOST                   | DB instance host                                           |
| DB_PORT                   | DB instance port                                           |
| DB_DATABASE               | DB instance's default database name                        |
| DB_OPTIONS                | Options placed for the db                                  |
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
### Makefile
  There are three commands to be used as a shortcut
  - Makefile build-image:
    - This shortcut corresponds to "docker build -t rafaelcmedrado/desafio:latest -f build/Dockerfile ."
  - Makefile push-image:
    - This shortcut corresponds to "docker push rafaelcmedrado/desafio:latest"
  - Makefile run-local:
    - This shortcut corresponds to "docker-compose -f deploy/local/docker-compose.yml up"
### 🛠 Technologies

The following tools were used in the construction of the project:

- [Postman](https://www.postman.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [Golang](https://golang.org/)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Docker](https://www.docker.com/)
- [Logrus](https://github.com/sirupsen/logrus)
- [Envconfig](github.com/kelseyhightower/envconfig)

### 📎 Endpoints

#### Accounts
- `POST /accounts` - Create an Account 
  - To make the request, pass through json or xml the customer's name, the amount that will deposit, the customer's cpf and the password that he will use as a form of access.
  - Request Example: 
    ```bash
    {
    "name": "Leao",
    "cpf": "102.502.200-53",
    "secret": "hash",
    "balance": 1000
    }
    ```
  - Response Example:
    ```bash
    {
    "id": "61110291-7db2-4d6f-ac0e-d9eb7b269bc4"
    }
    ```
  - Erros Examples:
    ```bash
    {
    "errors": "given cpf is already used"                                             
    "errors": "unable to insert"
    "errors": "unable to select"
    "errors": "given the balance amount is invalid"
    "errors": "given cpf is invalid"
    }
    ```
- `GET /accounts` - Get the list of accounts
   - It is not necessary to request.
  - Response Example:
    ```bash
    {
    {
    "accounts": [
        {
            "id": "61110291-7db2-4d6f-ac0e-d9eb7b269bc4",
            "name": "Leao",
            "cpf": "10250220053",
            "secret": "0800fc577294c34e0b28ad2839435945",
            "balance": 1000,
            "created_at": "2021-08-26T15:16:49.052764Z"
        },
        {
            "id": "2c7cde32-d9c6-4317-bc3d-4eb073e3e391",
            "name": "Rafael",
            "cpf": "56250221053",
            "secret": "0800fc577294c34e0b28ad2839435945",
            "balance": 1000,
            "created_at": "2021-08-26T17:45:12.348536Z"
        }
    ]
    }
    }
    ```
  - Erros Examples:
    ```bash
    {                                        
    "errors": "unable to insert"
    }
    ```
- `GET /accounts/{id}/balance` - Get account balance
  - It is not necessary to request.
  - Response Example:
    ```bash
    {
    "balance": 1000
    }
    ```
  - Erros Examples:
    ```bash
    {                                        
    "errors": "given id is invalid"
    "errors": "given the UUID is incorrect"
    "errors": "unable to select"
    }
    ```

#### Login
- `POST /login` - authenticate the user
  - To make the request, pass through json or xml the user's cpf and the user's access password.
  - Request Example:
    ```bash
    {
    "cpf": "10250220053",
    "secret": "hash"
    }
    ```
  - Response Example:
    ```bash
    {
    "token": "MjYvMDgvMjAyMSAxNDo1NjoyMjo2MTExMDI5MS03ZGIyLTRkNmYtYWMwZS1kOWViN2IyNjliYzQ6YTNjYzQ1MjAtYzE2YS00YjViLTllYzMtYWU4NWVkZGFhOWJh"
    }
    ```
  - Erros Examples:
    ```bash
    {
    "errors": "given secret or CPF are incorrect"                                             
    "errors": "unable to insert"
    "errors": "unable to select"
    }
    ```
#### Transfer
- `GET /transfers` - get the list of transfers from the authenticated user.
    - Requires `Authorization` header entry.
    - It is not necessary to request.
    - Response Example:
      ```bash  
      {
      {
      "transfers": [
         {
              "id": "65d3e765-e8fc-4491-9633-ceaed8337479",
              "origin_account_id": "04b7c433-d054-439b-8149-d247a105ad98",
              "destination_account_id": "36d5eb62-d335-43d6-80b5-e7d64bf162ce",
              "amount": 400,
              "created_at": "2021-08-26T15:15:24.699843-03:00"
         }
      ]
      }
      }
      ```
    - Erros Examples:
      ```bash
      {
       "errors": "given token is invalid"                                             
       "errors": "unable to insert"
       "errors": "unable to select"
       "errors":  "given the UUID is incorrect"
       }
       ```
- `POST /transfers` - transfers from one Account to another.
    - Requires `Authorization` header entry.
    - To make the request, pass through json or xml the id of the user who will receive the transfer and the amount that will be transferred.
    - Request Example:
      ```bash
      {
      "account_destination_id": "36d5eb62-d335-43d6-80b5-e7d64bf162ce",
      "amount": 400
      }
      ```
  - Response Example:
    ```bash
    {
    "id": "593ab439-7866-4215-b608-27e610748da8"
    }
    ```
  - Erros Examples:
    ```bash
    {
    "errors": "given account without balance"                                           
    "errors": "given amount is invalid"
    "errors": "given account is the same as the account destination"
    "errors": "unable to insert"
    "errors": "unable to select"
    "errors": "given the UUID is incorrect"
    "errors": "given account destination id is invalid"
    "errors": "given token is invalid"
    "errors": "given id is invalid"
    }
    ```
## ✅ Credits
I thank all the authors and content available by the application developers and my mentor [Pedro](https://github.com/pedroyremolo) for guiding me in creating the project.
