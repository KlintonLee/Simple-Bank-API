# Simple Bank API

## 🎯 Motivação
Recentemente iniciei meus estudos em Arquitetura Hexagonal e a linguagem Golang,
então, este repositório foi criado com o objetivo de praticar.

## 👨🏻‍💻 Desafio (a melhorar)
Um banco precisa realizar a gestão de contas de seus clientes, para isso, será
necessário desenvolver uma API-REST com as seguintes Entidades:
```
Customers
id         string (uuid)
name       string
cpf        string
birth      datetime
created_at datetime
updated_at datetime

Account
id                     string (uuid)
customer_id            string (uuid)
balance                float
withdrawal_daily_limit float
isActive               bool
account_type           int8
created_at             datetime
updated_at             datetime

Transactions
id               string (uuid)
account_id       string (uuid)
value            float
transaction_type string (deposit | withdrawal)
created_at       datetime
```
- Implementar path que realiza a criação de uma conta;
- Implementar path que realiza operação de depósito em uma conta;
- Implementar path que realiza operação de consulta de saldo em determinada conta;
- Implementar path que realiza operação de saque em uma conta;
- Implementar path que realiza o bloqueio de uma conta;
- Implementar path que recupera o extrato de transações de uma conta;
- Implementar extrato por período;
- Elaborar manual de execução;
- Elaborar documentação;
- Elaborar testes.

## 🛠 Ferramentas necessárias
Em construção... 🧱

## 💻 Como executar
Em construção... 🧱