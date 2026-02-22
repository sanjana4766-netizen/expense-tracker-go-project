# Expense Tracker & Debt Settlement API (Go)

## Overview

This project implements a RESTful backend service for managing shared expenses within groups and computing optimal settlements among participants. The system records expenses, distributes costs across members, calculates net balances, and determines the minimal set of transactions required to settle outstanding debts.

The implementation focuses on correctness, simplicity, and clarity of design while demonstrating core backend development concepts using Go.

---

## Problem Context

When multiple individuals share expenses, determining who owes whom can quickly become complex, especially across multiple transactions. This system automates that process by:

* Tracking payments made within a group
* Calculating each member’s share
* Computing net balances
* Generating an efficient settlement plan

---

## System Architecture

The application follows a layered backend architecture:

Client → HTTP API (Gin) → Application Logic → Database (SQLite)

The API can be consumed by any frontend client, mobile application, or external service.

---

## Core Entities

The system models the domain using the following entities:

* **User** — An individual participant
* **Group** — A collection of users sharing expenses
* **GroupMember** — Association between users and groups
* **Expense** — A payment recorded within a group
* **Split** — Individual share of an expense

---

## Technology Stack

| Component     | Technology  |
| ------------- | ----------- |
| Language      | Go (Golang) |
| Web Framework | Gin         |
| ORM           | GORM        |
| Database      | SQLite      |

---

## API Endpoints

### User Management

* `POST /users` — Create a user

### Group Management

* `POST /groups` — Create a group
* `POST /groups/{id}/add` — Add a user to a group

### Expense Management

* `POST /expenses` — Record a shared expense

### Settlement

* `GET /settlement/{groupId}` — Retrieve optimal settlement transactions

---

## Settlement Methodology

The settlement computation is based on net balance analysis.

1. Determine how much each participant has paid versus their share
2. Classify participants as creditors or debtors
3. Match debtors to creditors iteratively
4. Transfer the minimum required amount per step
5. Continue until all balances reach zero

This approach reduces the total number of transactions required to clear debts.

**Time Complexity:** O(n log n)

---

## Example

Three users share a ₹900 expense paid entirely by one participant.

Each user’s share = ₹300

Net balances:

* User A: +600
* User B: −300
* User C: −300

Settlement:

* User B pays User A ₹300
* User C pays User A ₹300

---

## Setup and Execution

### Prerequisites

* Go 1.18 or higher

### Steps

```bash
go mod tidy
go run .
```

The server starts on:

http://localhost:8080

---

## Design Considerations

* SQLite chosen for portability and ease of setup
* RESTful architecture for extensibility
* Minimal dependencies for reliability
* Clear separation of data models and business logic

---

## Limitations

* Equal splitting only
* No authentication or authorization
* Single-node deployment
* No persistent user management beyond database storage

---

## Potential Extensions

* Custom split ratios
* Payment integration
* User authentication
* Real-time notifications
* Web or mobile client interface
* Scalable database backend

---

## Conclusion

This project demonstrates the design and implementation of a backend system for expense management, combining data modeling, API development, persistence, and algorithmic computation in a cohesive manner.
