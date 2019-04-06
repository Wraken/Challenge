# Moneway_Challenge
Challenge moneway

![Alt text](Misc/Topology.png?raw=true "Topology")


REST api :

GET /balance/{id} : return current balance of account specified by id

GET /transactions : return all transactions

POST /balance/createaccount params {accountname: name of the new account} : create an account on balance service

POST /transactions/debitaccount params {accountid: name of the account, amount: amout, notes: notes} : Debit money the account

POST /transactions/creditaccount params {accountid: name of the account, amount: amout, notes: notes} : Credit money to the account
