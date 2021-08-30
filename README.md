<img src="docs/imgs/theironbankofbraavosby.jpeg" alt="The Iron Bank of Braavos" width="280" align="right" />

# 🏦 The Iron Bank of Braavos

Welcome to the Iron Bank of Braavos.<br />
Make your transactions and your credit requests but remember the Iron Bank will have its due.

## 💻 About

The API was created to simulate a digital bank, through account management. It was built following the RESTful API
Design and clean architecture concepts. <br />
Public image available on [Docker Hub](https://hub.docker.com/r/rafaelcmedrado/desafio)

## Table of Contents

* [About](#-about)
* [Table of Contents](#table-of-contents)
* [Features](#-features)
* [How It Works](#-how-it-works)
    * [Application / Dependencies](#-application--dependencies)
    * [Environment Variables](#-environment-variables)
    * [How to Begin](#-how-to-begin)
        * [Makefile](#makefile)
    * [Technologies](#-technologies)
    * [Endpoints](#-endpoints)
        * [Accounts](#accounts)
        * [Login](#login)
        * [Transfer](#transfer)
* [Credits](#-credits)

### ⚙️ Features

- [x] Create an Account
- [x] Getting Accounts
- [x] Authenticate users
- [x] Transfer between accounts
- [x] Get transfer list

### 🚀 How It Works

This project is just the back end part.

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

### 🚧 Application / Dependencies

Before starting, you will need to have the following tools installed on your machine:
[Git](https://git-scm.com), [Golang](https://golang.org/dl/), [PostgreSQL](https://www.postgresql.org/).

I recommend having an editor to work with code like [VSCode](https://code.visualstudio.com/) and having an api client
like [Postman](https://www.postman.com/downloads/)

The Api is delivered in containers using [Docker](https://www.docker.com/)

### 🎲 How to Begin

```bash
# Clone this repository
$ git clone <https://github.com/CMedrado/DesafioStone>

# Access the project folder in the terminal/cmd
$ cd DesafioStone

# Run the application
$ Make run-local

# The server will start on port:5000 - go to <http://localhost:3333>
```

### Makefile

There are three commands to be used as a shortcut

- Make build-image:
    - This shortcut corresponds to "docker build -t rafaelcmedrado/desafio:latest -f build/Dockerfile ."
- Make push-image:
    - This shortcut corresponds to "docker push rafaelcmedrado/desafio:latest"
- Make run-local:
    - This shortcut corresponds to "docker-compose -f deploy/local/docker-compose.yml up"

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

  - ID format is uuid
  - Token format is basic64

#### Accounts

- `POST /accounts` - Create an Account

    - To make the request, pass through json the customer's name, the amount that will deposit, the customer's cpf and
      the password that he will use as a form of access.
  
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
      
    - When the cpf has already been used by another account, the error will be returned:
       ```bash
      {
          "errors": "given cpf is already used" 
      }
      ```   
          
    - When the database is unable to insert, the error will be returned:
      ```bash
      {
          "errors": "unable to insert"
      }
      ```
      
    - When the database is unable to insert, the error will be returned:
      ```bash
      {
          "errors": "unable to select"
      }
      ```
      
    - When the balance amount on account creation is less than 0, the error will be returned:
      ```bash
      {
          "errors": "given the balance amount is invalid"
      }
      ```
      
    - When the cpf entered to create the account is wrong, the error will be returned:
      ```bash
      {
          "errors": "given cpf is invalid"
      }
      ```
- `GET /accounts` - Get the list of accounts

    - It is not necessary to request.
  
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
      
    - When the database is unable to insert, the error will be returned:
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
      
    - When the entered id is not valid, the error will be returned:
       ```bash
      {
          "errors": "given id is invalid"
      }
      ```   
          
    - When the uuid cannot be converted because there is an error in the uuid, the error will be returned:
      ```bash
      {
          "errors": "given the UUID is incorrect"
      }
      ```
      
    - When the database is unable to insert, the error will be returned:
      ```bash
      {
          "errors": "unable to select"
      }
      ```

#### Login

- `POST /login` - authenticate the user

    - To make the request, pass through json the user's cpf and the user's access password.
  
    - This function returns a token of type base64 to be placed in headers as a form of authentication
  
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
      
    - When the cpf or password is wrong , the error will be returned:
       ```bash
      {
          "errors": "given secret or CPF are incorrect"               
      }
      ```    
         
    - When the database is unable to insert, the error will be returned:
      ```bash
      {
          "errors": "unable to insert"
      }
      ```
      
    - When the database is unable to insert, the error will be returned:
      ```bash
      {
          "errors": "unable to select"
      }
      ```

#### Transfer

- `GET /transfers` - get the list of transfers from the authenticated user.

    - Requires basic type `Authorization` credential header entry.
  
    - It is not necessary to request.
  
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
      
    - When the entered id is not valid, the error will be returned:
       ```bash
      {
          "errors": "given id is invalid"
      }
      ```     
        
    - When the uuid cannot be converted because there is an error in the uuid, the error will be returned:
      ```bash
      {
          "errors": "given the UUID is incorrect"
      }
      ```
      
    - When the database is unable to insert, the error will be returned:
      ```bash
      {
          "errors": "unable to select"
      }
      ```
      
    - When the passed token is not correct, the error will be returned:
      ```bash
      {
          "errors": "given token is invalid"          
      }
      ```
      
- `POST /transfers` - transfers from one Account to another.

    - Requires basic type `Authorization` credential header entry.
  
    - To make the request, pass through json the id of the user who will receive the transfer and the amount that will
      be transferred.
  
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
      
    - When the entered id is not valid, the error will be returned:
       ```bash
      {
          "errors": "given id is invalid"
      }
      ```   
    
    - When the uuid cannot be converted because there is an error in the uuid, the error will be returned:
      ```bash
      {
          "errors": "given the UUID is incorrect"
      }
      ```
      
    - When the database is unable to insert, the error will be returned:
      ```bash
      {
          "errors": "unable to select"
      }
      ```
      
    - When the passed token is not correct, the error will be returned:
      ```bash
      {
          "errors": "given token is invalid"          
      }
      ```
      
    - When the origin account has no balance, the error will be returned:
       ```bash
      {
          "errors": "given account without balance"    
      }
      ```     
        
    - When past amount is less than zero, the error will be returned:
      ```bash
      {
          "errors": "given amount is invalid"
      }
      ```
      
    - When the two accounts are the same, the error will be returned:
      ```bash
      {
          "errors": "given account is the same as the account destination"
      }
      ```
      
    - When the destination account id does not exist, the error will be returned:
      ```bash
      {
          "errors": "given account destination id is invalid"      
      }
      ```
      
    - When the entered id is not valid, the error will be returned:
       ```bash
      {
          "errors": "given id is invalid"
      }
      ```  

## ✅ Credits

I thank all the authors and content available by the application developers and my
mentor [Pedro](https://github.com/pedroyremolo) for guiding me in creating the project.
