# Moneway_Challenge
Challenge moneway

![Alt text](Misc/Topology.png?raw=true "Topology")


REST api :

GET /balance/{id} : return current balance of account specified by id

GET /transactions : return all transactions

POST /balance/createaccount params {accountname: name of the new account} : create an account on balance service

POST /transactions/creditaccount params {accountid: name of the account, 
                                        description: description, 
                                        amount: amout, notes: notes} : Debit the account 
                                        
POST /transactions/depositaccount params {accountid: name of the account, 
                                        description: description, 
                                        amount: amout, notes: notes} : Depose money to the account
