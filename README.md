# Expense Tracker & Debt Settlement API (Go)

## Overview

This project implements a RESTful API for managing shared expenses among groups and computing optimal debt settlements. The system allows users to record expenses, split costs among participants, calculate individual balances, and determine the minimum set of transactions required to settle all outstanding debts.

The application models the core functionality of expense-sharing platforms such as Splitwise, with a focus on backend design, correctness, and efficiency.

---

## Objectives

* Provide a structured system for tracking shared expenses
* Automate fair cost distribution among group members
* Compute net balances for each participant
* Generate optimal settlement transactions
* Demonstrate backend development using Go

---

## System Architecture

The application follows a layered architecture:

Client → REST API (Gin) → Business Logic → Database (SQLite)

---

## Data Model

The system is based on the following entities:

* **User** — Represents an individual participant
* **Group** — Represents a collection of users sharing expenses
* **GroupMember** — Associates users with groups
* **Expense** — Records a payment made within a group
* **Split** — Represents each user's share of an expense

---

## Technology Stack

| Component            | Technology  |
| -------------------- | ----------- |
| Programming Language | Go (Golang) |
| Web Framework        | Gin         |
| ORM                  | GORM        |
| Database             | SQLite      |

---

## API Endpoints

### User Management

* `POST /users` — Create a new user

### Group Management

* `POST /groups` — Create a group
* `POST /groups/{id}/add` — Add a user to a group

### Expense Management

* `POST /expenses` — Record a shared expense

### Settlement

* `GET /settlement/{groupId}` — Retrieve optimal settlement transactions

---

## Settlement Algorithm

The system computes settlements using a greedy optimization approach:

1. Calculate the net balance of each participant
2. Classify users as creditors (positive balance) or debtors (negative balance)
3. Match the largest debtor with the largest creditor
4. Transfer the minimum required amount
5. Repeat until all balances are settled

This approach minimizes the number of transactions required to clear all debts.

**Time Complexity:** O(n log n)

---

## Example Scenario

Three users share an expense of ₹900 paid by one member.

Each user’s share = ₹300

Net balances:

* User A: +600
* User B: −300
* User C: −300

Optimal settlements:

* User B pays User A ₹300
* User C pays User A ₹300

---

## Setup Instructions

### Prerequisites

* Go (version 1.18 or higher)

### Steps to Run

```bash
go mod tidy
go run .
```

The server will start on:

http://localhost:8080

---

## Limitations

* Supports equal expense splitting only
* No authentication or authorization
* Single-instance deployment

---

## Future Enhancements

* Unequal split support
* Payment integration
* User authentication
* Web or mobile interface
* Distributed deployment

---

## Conclusion

This project demonstrates the design and implementation of a backend system for expense management, including data modeling, REST API development, persistence, and algorithmic optimization.
