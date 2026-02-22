# Engineering Notes

This document records the reasoning behind key implementation choices made during development.

## Scope of the Project

The project focuses on backend functionality: data management, business logic, and settlement computation. A graphical user interface can be built separately to consume these APIs.

## Database Selection

SQLite was chosen because it is lightweight and requires no external setup. This allows the project to run consistently across different environments without configuration overhead.

## Expense Splitting Assumption

The current implementation assumes equal splits among all group members. Supporting custom ratios would require additional validation and a more complex data model.

## Settlement Strategy

A greedy method is used to match users who owe money with users who should receive money. This approach is simple, efficient, and sufficient for small group sizes.

## Reliability Considerations

The system prioritizes correctness of calculations over performance optimizations. For typical use cases, response times remain well within acceptable limits.

## Potential Improvements

Future versions could include:

* Custom split ratios
* Authentication and user sessions
* Persistent user management
* Integration with payment systems
* A web or mobile client interface

## Development Constraints

The implementation was designed to remain understandable and maintainable within a limited development timeframe while still demonstrating core backend concepts.
