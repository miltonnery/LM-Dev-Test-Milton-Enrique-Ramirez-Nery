# Life Bank

This project was made for a test which tries to accomplish the following requirements: 
1.  Clients authentication -> lb_authentication_svc
2. Client´s owned products -> lb_account_svc
3. Client´s transactions of owned products -> lb_account_svc
4. Client transfers for its own account
5. Client´s cc payment 
6. Client´s Loan Payment
7. Beneficiary affiliation -> lb_account_svc
8. Beneficiary email update -> lb_account_svc
9. Beneficiary deletion -> lb_account_svc
10. Beneficiary transfer
11. Beneficiary cc payment
12. Beneficiary loan payment

# Features!

LifeBank is intended to be a Web-App made of the following main elements:
  - SPA WebApp.
  - Secured API: This can work for Web and Mobile clients.
  - Information storaged in database: The Database Schema and scripts are provided.

# Tech
Gopher Guard uses a number of technologies to work properly:

* [Gin Gonic](https://gin-gonic.com/) - The fastest full-featured web framework for Go. Crystal clear!
* [GORM](http://gorm.io/) - The fantastic ORM library for Golang
* [Viper](https://github.com/spf13/viper) - GO configuration with fangs!

# INSTALLATION
### Database
Before running the services is necessary to execute the following Postgre SQL scripts:
1. lb_db.sql : Contains all the DDL and Schema definition
2. lb_db_data.sql : Contains the inicialization data (for endpoints testing and some configurations)

### Microservices
The Back-End microservices are located inside of:
```sh
/Backend/lb_name_of_the_microservice_svc
```
Before running them is necessary to compile them. Due to the files configuration, you have to move this projects to your GO-PATH. Normally (after installing Go in your computer) your GO-PATH should be:
```sh
C:\Users\username\go\src
```
The microservices are: 
1. lb_authorization_svc: PORT:90; Contains the Authorization / Gateway support for the service, this service encapsulates other microservices providing a security layer based on JWT.
2. lb_authentication_svc: PORT: 91; Microservice for login. It provides the JWT for services consumption.
3. lb_account_svc: PORT 92; Contains all the information related with account services contracted for a client, transactions, beneficiaries management.
4. lb_payment_svc: Payment handling (NOT DONE)
5. lb_transfer_svc: Transfers handling (NOT DONE)

### To-Dos

Here is a list of the remainng elements to develop
- lb_payment_svc:
- lb_transfer_svc:

License
----
Testing software belongs to LifeMiles
