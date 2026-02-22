# System Design

## Architecture Overview

The application follows a three-layer architecture:

1. Presentation Layer — REST API endpoints implemented using Gin
2. Application Layer — Business logic for expense processing and settlement
3. Data Layer — Persistent storage using SQLite via GORM

---

## Design Principles

* Simplicity and clarity
* Maintainability
* Correct handling of shared expenses
* Efficient computation of settlements

---

## Key Design Decisions

* SQLite selected for lightweight, zero-configuration deployment
* REST architecture for modularity and scalability
* GORM used for object–relational mapping
* Greedy algorithm chosen for efficient settlement computation

---

## Limitations

* Equal splitting only
* No user authentication
* Not optimized for distributed environments

---

## Possible Extensions

* Support for custom split ratios
* Role-based access control
* Integration with payment services
* Scalable database backend
* Frontend user interface
