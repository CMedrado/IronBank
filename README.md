<img src="docs/imgs/theironbankofbraavosby.jpeg" alt="The Iron Bank of Braavos" width="280" align="right" />

# 🏦 The Iron Bank of Braavos

Welcome to the Iron Bank of Braavos.<br />
Make your transactions and your credit requests but remember the Iron Bank will have its due.

## 💻 About

The API was created to simulate a digital bank, through account management. It was built following the RESTful API.
Design and clean architecture concepts. <br />
Public image available on [Docker Hub](https://hub.docker.com/r/rafaelcmedrado/desafio).

## Table of Contents

* [About](#-about)
* [Table of Contents](#table-of-contents)
* [Features](#-features)
* [How It Works](#-how-it-works)
    * [Application / Dependencies](#-application--dependencies)
    * [How to Begin](#-how-to-begin)
       * [Makefile](#makefile)
    * [Environment Variables](#-environment-variables)
    * [Technologies](#-technologies)
    * [Endpoints](#-endpoints)
    * [Accounts](#accounts)
    * [Login](#login)
    * [Transfer](#transfer)
* [Acknowledgements](#-acknowledgements)

### ⚙️ Features

- [x] Create an Account.
- [x] Getting Accounts.
- [x] Authenticate users.
- [x] Transfer between accounts.
- [x] Get transfer list.

## 🚀 How It Works

### 🚧 Application / Dependencies

Before starting, you will need to have the following tools installed on your machine:
[Git](https://git-scm.com), [Golang](https://golang.org/dl/), [PostgreSQL](https://www.postgresql.org/).

I recommend having an editor to work with code like [VSCode](https://code.visualstudio.com/) and having an api client
like [Postman](https://www.postman.com/downloads/).

The Api is delivered in containers using [Docker](https://www.docker.com/).

### 🎲 How to Begin

```bash
# Clone this repository
$ git clone <https://github.com/CMedrado/DesafioStone>

# Access the project folder in the terminal/cmd
$ cd DesafioStone

# Run the application
$ Make run-local

# The server will start on port:5000 - go to <http://localhost:5000>
```

### Makefile

There are three commands to be used as a shortcut.

- make build-image:
    - This shortcut corresponds to "docker build -t rafaelcmedrado/desafio:latest -f build/Dockerfile .".
- make push-image:
    - This shortcut corresponds to "docker push rafaelcmedrado/desafio:latest".
- make run-local:
    - This shortcut corresponds to "docker-compose -f deploy/local/docker-compose.yml up".
  
### 🔢 Environment Variables

| Name                      | Description                                                | Examples             |
|---------------------------|------------------------------------------------------------|----------------------|
| API_PORT                  | Port that will be listened on for the new request          | 5000                 |    
| API_LOG_LEVEL             | Structured api log level                                   | INFO                 |
| DB_PROTOCOL               | DB instance protocol                                       | postgres             |
| DB_USERNAME               | DB instance user                                           | postgres             |
| DB_SECRET                 | DB instance password                                       | example              |
| DB_HOST                   | DB instance host                                           | db                   |
| DB_PORT                   | DB instance port                                           | 5432                 |
| DB_DATABASE               | DB instance's default database name                        | desafio              |
| DB_OPTIONS                | Options placed for the db                                  | sslmode=disable      |

### 🛠 Technologies

The following tools were used in the construction of the project:

- [Postman](https://www.postman.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [Golang](https://golang.org/)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Docker](https://www.docker.com/)
- [Logrus](https://github.com/sirupsen/logrus)
- [Envconfig](https://github.com/kelseyhightower/envconfig)

### 📎 Endpoints

#### Accounts

- `POST /accounts` - Account creation endpoint.

    - The request and response bodies will be in json.
  
    - Request Example:
      ```bash
      {
          "name": "Lion",
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
      
    - Possible Errors:

      | Error Code       | Body JSON                                       | Description                                              |
      |------------------|-------------------------------------------------|----------------------------------------------------------|
      | 400  Bad Request | "errors": "given cpf is already used"           | The `cpf` has already been used by another account.      |
      | 400  Bad Request | "errors": "unable to insert"                    | The database is unable to insert.                        |
      | 400 Bad Request  | "errors": "unable to select"                    | The database is unable to select.                        |
      | 400 Bad Request  | "errors": "given the balance amount is invalid" | The `balance` amount on account creation is less than 0. |
      | 400 Bad Request  | "errors": "given cpf is invalid"                | The `cpf` entered to create the account is wrong.        |

- `GET /accounts` - Account listing endpoint.

    - It is not necessary to send body request and the response body will be by json.
  
    - Response Example:
      ```bash
      {
      "accounts": [
          {
              "id": "61110291-7db2-4d6f-ac0e-d9eb7b269bc4",
              "name": "Lion",
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
      ```
    - Possible Errors:

      | Error Code      | Body JSON                    | Description                       |
      |---------------- |------------------------------|-----------------------------------|
      | 400 Bad Request | "errors": "unable to insert" | The database is unable to insert. |
      
- `GET /accounts/{id}/balance` - Account balance display endpoint.

    - It is not necessary to send body request only the id by url and the response body will be by json.
  
    - Response Example:
      ```bash
      {
          "balance": 1000
      }
      ```
    - Possible Errors:

      |      Error Code    |            Body JSON                    |    Description                                                      |
      |--------------------|-----------------------------------------|---------------------------------------------------------------------|
      | 400 Bad Request    | "errors": "given the UUID is incorrect" | The UUID cannot be converted because there is an error in the UUID. |
      | 400 Bad Request    | "errors": "unable to select"            | The database is unable to select.                                   |
      | 406 Not Acceptable | "errors": "given id is invalid"         | The entered `id` is not valid.                                      |

#### Login

- `POST /login` - Account authentication endpoint.

    - The request and response bodies will be in json.
  
    - This function returns a token of type base64 to be placed in headers as a form of authentication.
  
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
    - Possible Errors:

      |      Error Code  |            Body JSON                          |    Description                    |
      |------------------|-----------------------------------------------|-----------------------------------|
      | 401 Unauthorized | "errors": "given secret or CPF are incorrect" | THe `cpf` or `password` is wrong. |
      | 400 Bad Request  | "errors": "unable to insert"                  | The database is unable to insert. |
      | 400 Bad Request  | "errors": "unable to select"                  | The database is unable to select. |

#### Transfer

- `GET /transfers` - Account transfer listing endpoint.

    - Requires basic type `Authorization` credential header entry.
  
    - It is not necessary to send body request and the response body will be by json.
  
    - Response Example:
      ```bash  
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
      ```
  - Possible Errors:

     |   Error Code     |            Body JSON                    |    Description                                                      |
     |------------------|-----------------------------------------|---------------------------------------------------------------------|
     | 400 Bad Request  | "errors": "given id is invalid"         | The entered `id` is not valid.                                      |
     | 400 Bad Request  | "errors": "given the UUID is incorrect" | The UUID cannot be converted because there is an error in the UUID. |
     | 400 Bad Request  | "errors": "unable to select"            | The database is unable to select.                                   |
     | 401 Unauthorized | "errors": "given token is invalid"      | The passed token is not correct.                                    |
      
- `POST /transfers` - Transfer creation endpoint between accounts.

    - Requires basic type `Authorization` credential header entry.
  
    - The request of the body and the response will be in json.
  
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

    - Possible Errors:

      |   Error Code       |            Body JSON                                             |    Description                                                      |
      |--------------------|------------------------------------------------------------------|---------------------------------------------------------------------|
      | 406 Not Acceptable | "errors": "given id is invalid"                                  | The entered `id` is not valid.                                      |
      | 400 Bad Request    | "errors": "given the UUID is incorrect"                          | The UUID cannot be converted because there is an error in the UUID. |
      | 400 Bad Request    | "errors": "unable to select"                                     | The database is unable to select.                                   |
      | 400 Bad Request    | "errors": "unable to insert"                                     | The database is unable to insert.                                   |
      | 401 Unauthorized   | "errors": "given token is invalid"                               | The passed `token` is not correct.                                  |
      | 400 Bad Request    | "errors": "given account without balance"                        | The origin account has no `balance`.                                |
      | 400 Bad Request    | "errors": "given amount is invalid"                              | The two accounts are the same.                                      |
      | 400 Bad Request    | "errors": "given account is the same as the account destination" | The `amount` is less than zero.                                     |
      | 401 Unauthorized   | "errors": "given account destination id is invalid"              | The entered `id` is not valid.                                      |

## ✅ Acknowledgements

I thank all the authors and content available by the application developers and my
mentor [Pedro](https://github.com/pedroyremolo) for guiding me in creating the project.
