# System Design

## Architecture
Client → REST API (Gin) → Service Logic → Database (SQLite)

## Data Model
- User
- Group
- GroupMember
- Expense
- Split

## Workflow
1. Users join groups
2. Expenses are recorded
3. Amount is split among members
4. Net balances computed
5. Greedy algorithm generates minimal settlement transactions

## Design Decisions
- SQLite chosen for simplicity and portability
- RESTful endpoints for scalability
- GORM used for ORM abstraction