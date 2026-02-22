# Expense Tracker & Bill Splitting API

## Overview
A REST API for managing shared expenses among groups and computing optimal debt settlements, similar to Splitwise.

## Key Features
- User and group management
- Add shared expenses
- Equal bill splitting
- Debt minimization using greedy algorithm
- Persistent storage using SQLite
- RESTful API design

## Tech Stack
- Go (Golang)
- Gin Web Framework
- GORM ORM
- SQLite Database

## API Endpoints

### Users
POST /users — Create user

### Groups
POST /groups — Create group  
POST /groups/{id}/add — Add user to group  

### Expenses
POST /expenses — Add expense  

### Settlement
GET /settlement/{groupId} — Calculate optimal settlements  

## Settlement Algorithm
Net balances are computed for each user. Creditors and debtors are matched greedily to minimize the number of transactions.

Time Complexity: O(n log n)

## How to Run

```bash
go mod tidy
go run .